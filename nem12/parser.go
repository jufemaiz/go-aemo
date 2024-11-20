package nem12

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/jufemaiz/go-aemo/nmi"
)

const (
	// base64 for base 64 numbers.
	base64 = 64
)

const (
	// ParseModeStrict defaults to strict mode.
	ParseModeStrict ParseMode = iota
	// ParseModeIgnoreX ignores X when parsing.
	ParseModeIgnoreX
)

const (
	parseStateNeedHeader         parseState = iota // A header record is required.
	parseStateNeedNMIDataDetails                   // A NMI data details record is required.
	parseStateNextNMIDataDetails                   // Move to the next interval data record.
	parseStateNeedIntervalData                     // An interval data record is required.
	parseStateNextIntervalData                     // Move to the next interval data record.
	parseStateNeedIntervalEvent                    // An interval event record is required.
	parseStateNextIntervalEvent                    // Move to the next interval event record.
	parseStateNeedB2BDetails                       // A B2B details record is required.
	parseStateNextB2BDetails                       // Move to the next B2B details record.
	parseStateCheckEnd                             // A valid end record has been encountered.
	parseStateCheckClosed                          // Checks the file is closed properly.
	parseStateClosed                               // Closed. No more action.
)

// Parser provides the ability to parse nem12 files.
type Parser interface {
	// ReadDataStream() (*IntervalSet, err error)
	ReadDay() (*IntervalSet, error)
	// ReadIntervals(n int) (*IntervalSet, err error)

	// StreamDataStreams() (<-chan *IntervalSet, <-chan error)
	// StreamDays() (<-chan *IntervalSet, <-chan error)
	// StreamIntervals() (<-chan *Interval, <-chan error)
}

// ParseMode controls the parsing behaviour.
type ParseMode int

// parser implements the Parser interface. See <https://golang.org/src/text/template/parse/parse.go>
// for inspiration.
type parser struct {
	// mode ParseMode      // Supports various levels of strictness.
	r   *csv.Reader    // Source of nem12 file to be parsed.
	loc *time.Location // Configured for NEMTime.
	// Cached data.
	header         *Record           // Contents of the Header record of the MDFF file.
	nmiDataDetails *Record           // Contents of the most recently read NMI data details record.
	intervalLength *time.Duration    // Interval length from the NMI data details record.
	metadata       *IntervalMetadata // Formatted metadata.
	// Parsing only; cleared/updated after step of the parse.
	recCount   int            // Record count for use in reporting errors.
	record     [1]Record      // A lookahead buffer used when peeking.
	peekCount  int            // Number of records that have been placed in the buffer.
	state      parseState     // Present value of the Parser's finite state machine.
	valueCount int            // Number of values in an interval record, based on interval length.
	eventCount int            // Number of contiguous interval records covered by events.
	unit       UnitOfMeasure  // Unit of measure of the values.
	values     Intervals      // The interval for the day currently being read.
	events     intervalEvents // The slice of interval events.
}

// parseState provides the states for parsing steps.
type parseState int

// NewParser returns a new Parser that parses text read from r.
func NewParser(r io.Reader) Parser {
	p := &parser{loc: NEMTime()}

	if r != nil {
		p.r = csv.NewReader(r)
		p.r.ReuseRecord = true
	}

	return p
}

// ReadDay reads a day of data and returns it.
func (p *parser) ReadDay() (*IntervalSet, error) {
	if p.r == nil {
		return nil, ErrReaderNil
	}

	return p.readDay()
}

// readDay implements the actual read day capability, making it commonly
// available to the interface functions.
//
//nolint:gocognit,dupl,funlen
func (p *parser) readDay() (set *IntervalSet, err error) { //nolint:cyclop,gocyclo
	defer func() {
		if err != nil && !errors.Is(err, io.EOF) {
			err = p.error(err)
		}
	}()

	var (
		fin bool
		rec Record
	)

	for {
		rec, err = p.next()
		if err != nil {
			if p.state == parseStateCheckClosed && errors.Is(err, io.EOF) {
				p.state = parseStateClosed

				return nil, io.EOF
			}

			if errors.Is(err, io.EOF) {
				return nil, fmt.Errorf("unexpected end of file: %w", ErrParseUnexpectedEOF)
			}

			switch p.state { //nolint:exhaustive
			case parseStateNeedIntervalData, parseStateNextIntervalData, parseStateNeedIntervalEvent, parseStateNeedB2BDetails:
				// If there's a bad record, we should fast forward to the next nmi data details.
				p.state = parseStateNextNMIDataDetails
			}

			return nil, fmt.Errorf("invalid record: %w", err)
		}

		ri := rec.Indicator()

		switch p.state { //nolint:exhaustive
		// We start off needing a header.
		case parseStateNeedHeader:
			if ri == RecordHeader {
				p.header = &rec
				p.state = parseStateNeedNMIDataDetails

				continue
			}

			p.state = parseStateCheckEnd

			return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordHeaderMissing)

		case parseStateNeedNMIDataDetails:
			p.state = parseStateNextIntervalData

			if ri != RecordNMIDataDetails {
				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordNMIDataDetailsMissing)
			}

			fallthrough

		case parseStateNextNMIDataDetails:
			switch ri { //nolint:exhaustive
			case RecordNMIDataDetails:
				if err = p.cacheNMIDataDetails(rec); err != nil {
					// Recover onto the next set of nmi data details and interval data.
					p.state = parseStateNextNMIDataDetails

					return nil, err
				}

				p.state = parseStateNeedIntervalData

				continue

			case RecordEnd:
				p.state = parseStateCheckEnd
				p.backup()

				continue
			}

			continue

		case parseStateNeedIntervalData:
			p.state = parseStateNextIntervalData

			if ri != RecordIntervalData {
				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordIntervalDataMissing)
			}

			fallthrough

		case parseStateNextIntervalData:
			switch ri { //nolint:exhaustive
			case RecordB2BDetails:
				p.state = parseStateNextB2BDetails
				p.backup()

				continue

			case RecordIntervalEvent:
				p.state = parseStateNeedIntervalEvent
				p.backup()

				continue

			case RecordIntervalData:
				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.state = parseStateNeedIntervalData
					p.resetDay()
					p.backup()

					return set, nil
				}

				if err = p.cacheValues(rec); err != nil {
					// Recover onto the next set of nmi data details and interval data.
					p.state = parseStateNextNMIDataDetails

					return nil, err
				}

			case RecordNMIDataDetails:
				p.state = parseStateNeedNMIDataDetails

				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.resetDay()
					p.backup()

					return set, nil
				}

				p.backup()

				continue

			case RecordEnd:
				p.state = parseStateCheckEnd

				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.resetDay()
					p.backup()

					return set, nil
				}

				p.backup()

				continue

			default:
				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordNMIDataDetailsMissing)
			}

		case parseStateNeedIntervalEvent:
			if ri != RecordIntervalEvent {
				p.state = parseStateNextIntervalEvent

				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordIntervalEventMissing)
			}

			if err := p.cacheEvent(rec); err != nil {
				return nil, fmt.Errorf("cache event: %w", err)
			}

			p.state = parseStateNextIntervalEvent

			continue

		case parseStateNextIntervalEvent:
			switch ri { //nolint:exhaustive
			case RecordB2BDetails:
				p.state = parseStateNextB2BDetails
				p.backup()

				continue

			case RecordIntervalEvent:
				p.state = parseStateNeedIntervalEvent
				p.backup()

				continue

			case RecordIntervalData:
				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.state = parseStateNeedIntervalData
					p.resetDay()
					p.backup()

					return set, nil
				}

				if err = p.cacheValues(rec); err != nil {
					// Recover onto the next set of nmi data details and interval data.
					p.state = parseStateNextNMIDataDetails

					return nil, err
				}

			case RecordNMIDataDetails:
				p.state = parseStateNeedNMIDataDetails

				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.resetDay()
					p.backup()

					return set, nil
				}

				p.backup()

				continue

			case RecordEnd:
				p.state = parseStateCheckEnd

				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.resetDay()
					p.backup()

					return set, nil
				}

				p.backup()

				continue

			default:
				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordNMIDataDetailsMissing)
			}

		case parseStateNeedB2BDetails:
			p.state = parseStateNextB2BDetails

			continue

		case parseStateNextB2BDetails:
			switch ri { //nolint:exhaustive
			case RecordB2BDetails:
				p.state = parseStateNeedB2BDetails
				p.backup()

				continue

			case RecordIntervalData:
				p.state = parseStateNextIntervalData
				p.backup()

				continue

			case RecordNMIDataDetails:
				p.state = parseStateNextNMIDataDetails
				p.backup()

				continue

			case RecordEnd:
				p.state = parseStateCheckEnd

				// If we have any data, we return it, and set the current record as a lookahead value instead.
				if len(p.values) > 0 {
					set, err = p.intervalSet()
					if err != nil {
						return nil, err
					}

					p.resetDay()
					p.backup()

					return set, nil
				}

				p.backup()

				continue

			default:
				return nil, fmt.Errorf("record indicator transition from '%s' to '%s' %w", RecordB2BDetails.String(), ri.String(), ErrIsInvalid)
			}

		case parseStateCheckEnd:
			if ri != RecordEnd {
				return nil, fmt.Errorf("record indicator '%s': %w", ri.String(), ErrParseRecordEndMissing)
			}

			p.state = parseStateCheckClosed
			fin = true

		case parseStateCheckClosed:
			fin = true

		default:
			return nil, ErrParseInvalidState
		}

		if fin {
			break
		}
	}

	return set, nil
}

// error returns a ParseError, with the current record count for reference.
func (p *parser) error(err error) error {
	return &ParseError{Err: err, Line: p.line()}
}

// next returns the next record in the file without validating it, either
// consuming from the peeked record, or via a direct read.
func (p *parser) next() (Record, error) {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		r, err := p.readNextRecord()
		if err != nil {
			return nil, err
		}

		p.record[0] = r
	}

	return p.record[p.peekCount], nil
}

// backup backs the records being viewed up one record.
func (p *parser) backup() {
	p.peekCount++
}

// // peek returns but does not consume the next record.
// func (p *parser) peek() (Record, error) {
// 	var err error

// 	if p.peekCount > 0 {
// 		return p.record[p.peekCount-1], nil
// 	}

// 	p.peekCount = 1

// 	p.record[0], err = p.readNextRecord()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return p.record[0], nil
// }

// readNewRecord reads a new record and returns the record.
func (p *parser) readNextRecord() (Record, error) {
	var e *csv.ParseError

	cols, err := p.r.Read()
	rec, recErr := NewRecord(cols)

	isParseError := errors.As(err, &e)

	// Ignore issues of differing field count - we are operating in the non-canonical
	// world that is NEM12.
	if err == nil || (isParseError && errors.Is(e, csv.ErrFieldCount)) {
		p.recCount++

		return rec, recErr
	}

	// End of the file. Doh.
	if errors.Is(err, io.EOF) {
		return rec, fmt.Errorf("%w", io.EOF)
	}

	if isParseError {
		// If it's just a ParseError, we can assume a record was actually read.
		p.recCount++
	}

	return rec, err //nolint:wrapcheck
}

// line returns the current line, which is the rec count offset by the peek count.
func (p *parser) line() int {
	return p.recCount - p.peekCount
}

// cacheReset resets the cached values.
func (p *parser) cacheReset() {
	// reset nmi detail cache.
	p.nmiDataDetails = nil
	p.intervalLength = nil
	p.metadata = nil
	p.unit = UnitUndefined
	p.valueCount = 0
	p.eventCount = 0

	// reset the daily transient data.
	p.resetDay()
}

// resetDay resets the transient data for a nmi detail.
func (p *parser) resetDay() {
	p.values = Intervals{}
	p.events = intervalEvents{}
	p.eventCount = 0
}

// cacheNMIDataDetails caches the nmi data details, and sets up the additional data points.
func (p *parser) cacheNMIDataDetails(rec Record) (err error) {
	p.cacheReset()

	p.nmiDataDetails = &rec

	err = p.setMetadata(rec)
	if err != nil {
		return err
	}

	err = p.setIntervalLength(rec)
	if err != nil {
		return err
	}

	err = p.setUnit(rec)
	if err != nil {
		return err
	}

	p.valueCount = int((hoursInDay * time.Hour) / *p.intervalLength)

	return nil
}

// setMetadata takes the 200 record and sets the metadata.
func (p *parser) setMetadata(rec Record) (err error) {
	n, err := nmi.NewNmi(rec[1].Value)
	if err != nil {
		return fmt.Errorf("nmi: %w", err)
	}

	s, err := NewSuffix(rec[4].Value)
	if err != nil {
		return err
	}

	u, err := NewUnit(rec[7].Value)
	if err != nil {
		return err
	}

	m := &nmi.Meter{
		Nmi:          n.String(),
		Identifier:   string(s.Meter),
		Registers:    []*nmi.MeterRegister{{RegisterID: rec[3].Value, MeasurementStream: s.Type.Identifier()}},
		SerialNumber: &(rec[6].Value),
	}

	p.metadata = &IntervalMetadata{
		Nmi:           n,
		Meter:         m,
		Suffix:        &s,
		UnitOfMeasure: &u,
	}

	return nil
}

// setIntervalLength sets the interval length from the record.
func (p *parser) setIntervalLength(rec Record) (err error) {
	il, err := time.ParseDuration(fmt.Sprintf("%sm", rec[8].Value))
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}

	p.intervalLength = &il

	return nil
}

// setUnit sets the unit of measurement from the record.
func (p *parser) setUnit(rec Record) (err error) {
	u, err := NewUnit(rec[7].Value)
	if err != nil {
		return err
	}

	p.unit = u

	return nil
}

// cacheValues caches a record's values, to ensure any additional data can be
// added if needed.
func (p *parser) cacheValues(rec Record) (err error) { //nolint:cyclop,funlen,gocognit,gocyclo
	// Reset the values first.
	p.resetDay()

	var (
		date        time.Time
		qm          QualityMethod
		reason      *Reason
		reasonDesc  *string
		updatedAt   *time.Time
		msatsLoadAt *time.Time
	)

	vals := []float64{}

	for _, field := range rec {
		switch field.Type { //nolint:exhaustive
		case FieldIntervalDate:
			date, err = time.ParseInLocation(Date8Format, field.Value, p.loc)
			if err != nil {
				return fmt.Errorf("parse time: %w", err)
			}

		case FieldIntervalValue:
			v, err := strconv.ParseFloat(field.Value, base64)
			if err != nil {
				return fmt.Errorf("parse float: %w", err)
			}

			vals = append(vals, v)

		case FieldQualityMethod:
			qm, err = NewQualityMethod(field.Value)
			if err != nil {
				return err
			}

		case FieldReasonCode:
			if field.Value != "" {
				rc, err := NewReason(field.Value)
				if err != nil {
					return err
				}

				reason = &rc
			}

		case FieldReasonDescription:
			if field.Value != "" {
				str := field.Value

				reasonDesc = &str
			}

		case FieldUpdateDateTime:
			dt, err := time.ParseInLocation(DateTime14Format, field.Value, p.loc)
			if err != nil {
				return fmt.Errorf("parse time: %w", err)
			}

			updatedAt = &dt

		case FieldMSATSLoadDateTime:
			if field.Value == "" {
				continue
			}

			dt, err := time.ParseInLocation(DateTime14Format, field.Value, p.loc)
			if err != nil {
				return fmt.Errorf("parse time: %w", err)
			}

			msatsLoadAt = &dt
		}
	}

	il := p.intervalLength
	if il == nil {
		return ErrParseIntervalLengthInvalid
	}

	if len(vals) != p.valueCount {
		return ErrParseIntervalDataLengthInvalid
	}

	q, err := qm.Quality()
	if err != nil {
		return err
	}

	m, err := qm.Method()
	if err != nil {
		return err
	}

	if reason != nil && *reason != ReasonUndefined && reason.RequiresDescription() && reasonDesc == nil {
		return ErrFieldReasonDescriptionLengthInvalid
	}

	for i, val := range vals {
		interval := &Interval{
			Time:           date.Add(time.Duration((i + 1) * int(il.Nanoseconds()))),
			IntervalLength: *il,
			Value: IntervalValue{
				Value:        val,
				DecimalValue: decimal.NewFromFloat(val),
				QualityFlag:  q,
			},
		}

		if m != MethodUndefined {
			interval.Value.MethodFlag = &m
		}

		if reason != nil {
			interval.Value.ReasonCode = reason
		}

		if reasonDesc != nil {
			interval.Value.ReasonDescription = reasonDesc
		}

		if updatedAt != nil {
			interval.Value.UpdateDateTime = updatedAt
		}

		if msatsLoadAt != nil {
			interval.Value.MSATSLoadDateTime = msatsLoadAt
		}

		p.values = append(p.values, interval)
	}

	return nil
}

// cacheEvent cashes the interval event record.
func (p *parser) cacheEvent(rec Record) (err error) {
	ev, err := newIntervalEvent(rec)
	if err != nil {
		return err
	}

	p.events = append(p.events, ev)

	return nil
}

// intervalSet creates an intervalSet given the current interval dataset and the
// metadata.
func (p *parser) intervalSet() (*IntervalSet, error) {
	if err := p.events.validate(p.valueCount); err != nil {
		return nil, fmt.Errorf("interval events: %w", err)
	}

	is := &IntervalSet{Data: []*Interval{}, Metadata: p.metadata}
	im := p.events.intervalMap()

	for i, v := range p.values {
		if v == nil {
			continue
		}

		val := *v

		ev, ok := im[i+1]
		if ok { //nolint:nestif
			q, err := ev.QualityMethod.Quality()
			if err != nil {
				return nil, fmt.Errorf("interval %d: quality: %w", i+1, err)
			}

			val.Value.QualityFlag = q

			m, err := ev.QualityMethod.Method()
			if err != nil {
				return nil, fmt.Errorf("interval %d: method: %w", i+1, err)
			}

			if m != MethodUndefined {
				val.Value.MethodFlag = &m
			}

			if ev.Reason != nil {
				val.Value.ReasonCode = ev.Reason
			}

			if ev.ReasonDescription != nil {
				val.Value.ReasonDescription = ev.ReasonDescription
			}
		}

		is.Data = append(is.Data, &val)
	}

	return is, nil
}
