package nem12

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"log"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/pkg/errors"
// )

// // A ELRecordIndicator describes the first field of a record according to the MDFF Specification.
// type ELRecordIndicator string

// const (
// 	// Header is the first field of NEM12 Header (100) record.
// 	Header = ELRecordIndicator("100")

// 	// NMIDataDetails is the first field of NEM12 NMIDataDetails (200) record.
// 	NMIDataDetails = ELRecordIndicator("200")

// 	// IntervalData is the first field of NEM12 IntervalData (300) record.
// 	IntervalData = ELRecordIndicator("300")

// 	// IntervalEvent is the first field of NEM12 IntervalEvent (400) record.
// 	IntervalEvent = ELRecordIndicator("400")

// 	// B2BDetails is the first field of NEM12 B2BDetails (500) record.
// 	B2BDetails = ELRecordIndicator("500")

// 	// End is the first field of NEM12 End (900) record.
// 	End = ELRecordIndicator("900")
// )

// // A ELRecord holds all fields of a record according to the MDFF Specification.
// type ELRecord []string

// // Copy returns a duplicate of the ELRecord that does not share the same backing array.
// // This method is provided so ELParser can use the ReuseELRecord option of a csv.Reader.
// func (r *ELRecord) Copy() ELRecord {
// 	c := make(ELRecord, len(*r))
// 	copy(c, *r)
// 	return c
// }

// // Field returns field n of the ELRecord or the empty string if n is outside the range 0 to len(r)-1.
// func (r ELRecord) Field(n int) string {
// 	if n < 0 || n >= len(r) {
// 		return ""
// 	}
// 	return r[n]
// }

// // Indicator returns the ELRecordIndicator field of the ELRecord.
// func (r ELRecord) Indicator() ELRecordIndicator {
// 	return ELRecordIndicator(r.Field(0))
// }

// // A parseState holds the state of the ELParser's finite state machine.
// type parseState int

// const (
// 	needHeader    parseState = iota // A Header record is required.
// 	findHeader                      // Recover from an error by skipping to the next Header record.
// 	needDetails                     // An NMIDataDetails record is required.
// 	findDetails                     // Recover from an error by skipping to the next NMIDataDetails record.
// 	needIntervals                   // An IntervalData record is required.
// 	findIntervals                   // Recover from an error by skipping to the next IntervalData record.
// 	needEvent                       // An IntervalEvent record is required.
// 	needB2B                         // A B2BDetails record is required.
// 	checkEnd                        // A valid End record has been encountered.
// )

// // A ReadError wraps errors that were triggered by a failure to read new input.
// type ReadError struct {
// 	error
// }

// // A ELParser parses a file conforming to the MDFF specification.
// type ELParser struct {
// 	Header          ELRecord        // Contents of the Header record of the MDFF file.
// 	NMIDataDetails  ELRecord        // Contents of the most recently read NMIDataDetails record.
// 	Strict          bool            // If set to true, any departure from NEM12 triggers an error rather than a warning.
// 	loc             *time.Location  // The time zone with respect to which dates are stated. (Defined to be +1000.)
// 	saved           ELRecord        // Lookahead buffer used by readELRecord and unreadELRecord.
// 	r               *csv.Reader     // Source of CSV records to be parsed.
// 	recNum          int             // ELRecord number for use in reporting errors.
// 	state           parseState      // Present value of the ELParser's finite state machine.
// 	valueCount      int             // Number of values in an interval record.
// 	valueMultiplier float64         // Scale factor to apply to values in an interval record.
// 	values          []IntervalValue // The interval values for the day currently being read.
// }

// // NewELParser returns a new ELParser that parses text read from r.
// func NewELParser(r io.Reader) *ELParser {
// 	p := &ELParser{r: csv.NewReader(r), loc: time.FixedZone("AEST", 10*60*60)}
// 	p.r.ReuseRecord = true
// 	return p
// }

// // readELRecord reads and returns the next ELRecord in the file without validating it.
// // When there are no more ELRecords to read, it returns io.EOF.
// func (p *ELParser) readELRecord() (ELRecord, error) {
// 	if len(p.saved) > 0 {
// 		r := p.saved
// 		p.saved = ELRecord{}
// 		return r, nil
// 	}

// 	r, err := p.r.Read()
// 	e, isParseError := err.(*csv.ParseError)

// 	// It's OK for each record in the CSV to have a different number of fields
// 	// so ignore csv.ErrFieldCount errors returned by the CSV Reader.
// 	if err != nil && !(isParseError && e.Err == csv.ErrFieldCount) {
// 		if err == io.EOF {
// 			return ELRecord(r), io.EOF
// 		}

// 		if isParseError {
// 			// If it's just a ParseError, we can assume a record was actually read.
// 			p.recNum++
// 		} else {
// 			// It was not possible to read a new line of input.
// 			err = ReadError{err}
// 		}
// 		return ELRecord(r), errors.Wrap(err, "read record")
// 	}

// 	p.recNum++
// 	return ELRecord(r), nil
// }

// // unreadELRecord pushes back a record to be read on the next call to readELRecord.
// func (p *ELParser) unreadELRecord(r ELRecord) {
// 	p.saved = r
// }

// // loadDetails validates and extracts the contents of the NMI data details record r.
// func (p *ELParser) loadDetails(r ELRecord) (err error) {
// 	p.NMIDataDetails = r.Copy()
// 	defer func() {
// 		if err != nil {
// 			p.NMIDataDetails = ELRecord{}
// 			err = errors.Wrap(err, "NMI data details")
// 		}
// 	}()

// 	dt, err := p.IntervalLength()
// 	if err != nil {
// 		return err
// 	}

// 	m := dt.Round(time.Minute)
// 	day := 1440 * time.Minute

// 	switch {
// 	case m < time.Minute:
// 		return errors.Errorf("interval less than 1 minute: %s", dt)
// 	case m > day:
// 		return errors.Errorf("interval greater than 1 day: %s", dt)
// 	}

// 	u := p.UOM()
// 	_, err = u.Description()
// 	if err != nil {
// 		return err
// 	}
// 	p.valueMultiplier = u.Multiplier()
// 	p.valueCount = int(day / m)
// 	return nil
// }

// // Normalise applies the unit of measure scale factor to a value.
// func (p *ELParser) Normalise(v float64) float64 {
// 	return v * p.valueMultiplier
// }

// // loadValues extracts the contents of the interval data record, rec.
// // If the Quality Flag is 'V', vFlag is returned with the value true.
// func (p *ELParser) loadValues(rec ELRecord) (date time.Time, vFlag bool, err error) {
// 	defer func() {
// 		if err != nil {
// 			err = errors.Wrap(err, "interval data record")
// 		}
// 	}()

// 	const (
// 		fixedFieldCount        = 7
// 		dateIndex              = 1
// 		firstValueIndex        = 2
// 		qualityFieldOffset     = firstValueIndex
// 		reasonFieldOffset      = qualityFieldOffset + 1
// 		descriptionFieldOffset = reasonFieldOffset + 1
// 		updateDTOffset         = descriptionFieldOffset + 1
// 		msatsLoadDTOffset      = updateDTOffset + 1
// 	)

// 	if n := p.valueCount + fixedFieldCount; len(rec) != n {
// 		return date, false, errors.Errorf("should have %d fields, not %d", n, len(rec))
// 	}

// 	q, m, err := ParseQualityMethod(rec.Field(p.valueCount+qualityFieldOffset), p.Strict)
// 	if err != nil {
// 		return date, false, err
// 	}

// 	if q == Final && len(m) == 0 {
// 		log.Print("WARNING: nem12 record ", p.recNum, ": quality flag F with no method flag")
// 	}

// 	var rCode Reason
// 	rField := rec.Field(p.valueCount + reasonFieldOffset)
// 	rDesc := rec.Field(p.valueCount + descriptionFieldOffset)

// 	switch q {
// 	case Variable:
// 		vFlag = true
// 		if rField != "" {
// 			return date, false, errors.Errorf(`ReasonCode must be empty for quality V: "%s"`, rField)
// 		}
// 	case Actual, Null, Estimated:
// 		// ReasonCode is optional for these quality flag values.
// 		if rField == "" {
// 			break
// 		}
// 		fallthrough
// 	default:
// 		// ReasonCode is mandatory for everything else.
// 		rCode, err = ParseReasonCode(rField)
// 		if err != nil {
// 			return date, false, err
// 		}
// 		if rCode == FreeTextDescription && rDesc == "" {
// 			err = errors.Errorf(`empty ReasonDescription field with ReasonCode "free text description"`)
// 			if p.Strict {
// 				return date, false, err
// 			}

// 			//TODO: Inject the logger as an interface argument to gateway.New().
// 			// (see https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern)
// 			log.Print("WARNING: nem12 record ", p.recNum, ": ", err)
// 		}
// 	}

// 	date, err = time.ParseInLocation(Date8Format, rec.Field(dateIndex), p.loc)
// 	if err != nil {
// 		return date, false, err
// 	}

// 	// Parse UpdateDateTime
// 	var updateDT time.Time
// 	updateDTField := rec.Field(p.valueCount + updateDTOffset)
// 	// UpdateDateTime is not required if QualityMethod is "N"
// 	if m != "N" && updateDTField == "" {
// 		err = errors.Errorf("empty field with MethodFlag: %s", m)
// 	} else {
// 		updateDT, err = time.ParseInLocation(DateTime14Format, updateDTField, p.loc)
// 	}
// 	if err != nil {
// 		err = errors.Wrap(err, "UpdateDateTime")
// 		if p.Strict {
// 			return date, false, err
// 		}
// 		log.Print("WARNING: nem12 record ", p.recNum, ": ", err)
// 	}

// 	// Parse MSATSLoadDateTime
// 	var msatsLoadDT time.Time
// 	msatsLoadDTField := rec.Field(p.valueCount + msatsLoadDTOffset)
// 	// The MSATSLoadDateTime field may not be present (it is "Required if available" according to the NEM12 spec)
// 	if msatsLoadDTField != "" {
// 		msatsLoadDT, err = time.ParseInLocation(DateTime14Format, msatsLoadDTField, p.loc)
// 		if err != nil {
// 			err = errors.Wrap(err, "MSATSLoadDateTime")
// 			if p.Strict {
// 				return date, false, err
// 			}
// 			log.Print("WARNING: nem12 record ", p.recNum, ": ", err)
// 		}
// 	}

// 	p.values = make([]IntervalValue, p.valueCount)

// 	for i := range p.values {
// 		v, e := strconv.ParseFloat(rec.Field(firstValueIndex+i), 64)
// 		if e != nil {
// 			return date, vFlag, errors.Wrap(e, fmt.Sprintf("value no. %d", i+1))
// 		}
// 		p.values[i] = IntervalValue{
// 			Value:             v,
// 			QualityFlag:       q,
// 			MethodFlag:        m,
// 			ReasonCode:        rCode,
// 			ReasonDescription: rDesc,
// 			UpdateDateTime:    updateDT,
// 			MSATSLoadDateTime: msatsLoadDT,
// 		}
// 	}

// 	return date, vFlag, nil
// }

// // applyEvent validates and applies the contents of the interval event record, rec, to the currently loaded interval values.
// func (p *ELParser) applyEvent(rec ELRecord) (err error) {
// 	defer func() {
// 		if err != nil {
// 			err = errors.Wrap(err, "interval event record")
// 		}
// 	}()

// 	const (
// 		indicatorIndex = iota
// 		startIndex
// 		endIndex
// 		qualityIndex
// 		reasonIndex
// 		descriptionIndex
// 		fieldCount
// 	)

// 	if len(rec) != fieldCount {
// 		return errors.Errorf("interval record must have %d fields, not %d", fieldCount, len(rec))
// 	}

// 	start, err := strconv.ParseInt(rec.Field(startIndex), 10, 0)
// 	if err != nil {
// 		return errors.Wrap(err, "StartInterval")
// 	}
// 	if start < 1 || start > int64(p.valueCount) {
// 		return errors.Errorf("StartInterval must be in range 1 to %d: %d", p.valueCount, start)
// 	}

// 	end, err := strconv.ParseInt(rec.Field(endIndex), 10, 0)
// 	if err != nil {
// 		return errors.Wrap(err, "EndInterval")
// 	}
// 	if end < start || end > int64(p.valueCount) {
// 		return errors.Errorf("EndInterval must be in range %d to %d: %d", start, p.valueCount, end)
// 	}

// 	q, m, err := ParseQualityMethod(rec.Field(qualityIndex), p.Strict)
// 	if err != nil {
// 		return err
// 	}

// 	if q == Final && len(m) == 0 {
// 		log.Print("WARNING: nem12 record ", p.recNum, ": quality flag F with no method flag")
// 	}

// 	var rCode Reason
// 	rField := rec.Field(reasonIndex)
// 	rDesc := rec.Field(descriptionIndex)

// 	switch q {
// 	case Variable:
// 		return errors.New(`quality flag must be not be V`)
// 	case Actual, Null, Estimated:
// 		// ReasonCode is optional for these quality flag values.
// 		if rField == "" {
// 			break
// 		}
// 		fallthrough
// 	default:
// 		// ReasonCode is mandatory for everything else.
// 		rCode, err = ParseReasonCode(rField)
// 		if err != nil {
// 			return err
// 		}
// 		if rCode == FreeTextDescription && rDesc == "" {
// 			err = errors.Errorf(`empty ReasonDescription field with ReasonCode "free text description"`)
// 			if p.Strict {
// 				return err
// 			}

// 			//TODO: Inject the logger as an interface argument to gateway.New().
// 			// (see https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern)
// 			log.Print("WARNING: nem12 record ", p.recNum, ": ", err)
// 		}
// 	}

// 	// Update the values specified by the interval event record.
// 	for i := start - 1; i < end; i++ {
// 		v := p.values[i].Value
// 		p.values[i] = IntervalValue{
// 			Value:             v,
// 			QualityFlag:       q,
// 			MethodFlag:        m,
// 			ReasonCode:        rCode,
// 			ReasonDescription: rDesc,
// 		}
// 	}

// 	return nil
// }

// // applyB2B validates and applies the contents of the interval data record, rec, to the currently loaded interval values.
// // TODO: Validate and make available the record fields that are not currently used.
// func (p *ELParser) applyB2B(r ELRecord) (err error) {
// 	return
// }

// // ReadDay returns the a day's worth of interval data from the NEM12 file being parsed.
// func (p *ELParser) ReadDay() (date time.Time, values []IntervalValue, err error) {
// 	defer func() {
// 		if err != nil && err != io.EOF {
// 			err = errors.Wrap(err, fmt.Sprintf("nem12 record %d", p.recNum))
// 		}
// 	}()
// 	var rec ELRecord

// Loop:
// 	for {
// 		rec, err = p.readELRecord()
// 		if err != nil && p.state != checkEnd {
// 			if err == io.EOF {
// 				p.state = checkEnd
// 				return date, values, errors.Wrap(err, "unexpected end of file")
// 			}
// 			switch p.state {
// 			case needIntervals, findIntervals, needEvent, needB2B:
// 				// The bad record may have been an NMIDataDetails so resync on the next one.
// 				p.state = findDetails
// 			}
// 			return date, values, errors.Wrap(err, "bad record")
// 		}

// 		switch p.state {
// 		case needHeader, findHeader:
// 			switch rec.Indicator() {
// 			case Header:
// 				p.Header = rec.Copy()
// 				p.state = needDetails
// 				continue
// 			case End:
// 				p.state = checkEnd
// 			default:
// 				if p.state == findHeader {
// 					continue
// 				}
// 				p.state = findHeader
// 			}
// 			return date, values, errors.New("expected header record")

// 		case needDetails, findDetails:
// 			switch rec.Indicator() {
// 			case NMIDataDetails:
// 				err = p.loadDetails(rec)
// 				if err != nil {
// 					p.state = findDetails
// 					return date, values, err
// 				}
// 				p.state = needIntervals
// 				continue
// 			case Header:
// 				p.unreadELRecord(rec)
// 				p.state = needHeader
// 			case End:
// 				p.state = checkEnd
// 			default:
// 				if p.state == findDetails {
// 					// Don't report an error during recovery.
// 					continue
// 				}
// 				p.state = findDetails
// 			}
// 			return date, values, errors.New("expected NMI data details record")

// 		case needIntervals, findIntervals:
// 			switch rec.Indicator() {
// 			case IntervalData:
// 				d, vFlag, e := p.loadValues(rec)
// 				if e != nil {
// 					p.state = findIntervals
// 					return date, values, e
// 				}
// 				date = d
// 				if vFlag {
// 					p.state = needEvent
// 					continue
// 				}
// 				p.state = needIntervals
// 				break Loop
// 			case NMIDataDetails:
// 				//TODO: Check that at least one IntervalData was read since last NMIDataDetails.
// 				p.unreadELRecord(rec)
// 				p.state = needDetails
// 				continue
// 			case B2BDetails:
// 				//TODO: Check that at least one IntervalData was read since last NMIDataDetails.
// 				p.unreadELRecord(rec)
// 				p.state = needB2B
// 				continue
// 			case End:
// 				p.state = checkEnd
// 				continue
// 			case Header:
// 				p.unreadELRecord(rec)
// 				p.state = needHeader
// 			default:
// 				if p.state == findIntervals {
// 					continue
// 				}
// 				p.state = findIntervals
// 			}
// 			return date, values, errors.New("expected interval data record")

// 		case needEvent, needB2B:
// 			switch rec.Indicator() {
// 			case IntervalEvent:
// 				if p.state == needB2B {
// 					// Events must precede B2Bs.
// 					p.state = findIntervals
// 					break
// 				}
// 				err = p.applyEvent(rec)
// 				if err != nil {
// 					p.state = findIntervals
// 					return date, values, err
// 				}
// 				continue
// 			case B2BDetails:
// 				err = p.applyB2B(rec)
// 				if err != nil {
// 					p.state = findIntervals
// 					return date, values, err
// 				}
// 				p.state = needB2B
// 				continue
// 			case IntervalData:
// 				p.unreadELRecord(rec)
// 				p.state = needIntervals
// 				break Loop
// 			case NMIDataDetails:
// 				p.unreadELRecord(rec)
// 				p.state = needDetails
// 				break Loop
// 			case End:
// 				p.state = checkEnd
// 				break Loop
// 			case Header:
// 				p.unreadELRecord(rec)
// 				p.state = needHeader
// 			default:
// 				p.state = findIntervals
// 			}
// 			return date, values, errors.New("expected interval event record")

// 		case checkEnd:
// 			if err == nil {
// 				p.unreadELRecord(rec)
// 			}
// 			if err != io.EOF {
// 				// Reset the state machine so the caller has the option of continuing to parse records.
// 				p.state = needHeader
// 				err = errors.New("unexpected text after end record")
// 			}
// 			return date, values, err
// 		}
// 	}

// 	// Validate the interval values.
// 	for i, v := range p.values {
// 		if v.QualityFlag == Variable || v.QualityFlag == Quality("") {
// 			return date, p.values, errors.Errorf("interval value #%d has no QualityMethod", i+1)
// 		}
// 	}

// 	return date, p.values, nil
// }

// // VersionHeader returns the field of the same name from the Header record.
// func (p *ELParser) VersionHeader() string {
// 	return p.Header.Field(1)
// }

// // DateTime returns the field of the same name from the Header record.
// func (p *ELParser) DateTime() (time.Time, error) {
// 	t, err := time.ParseInLocation(DateTime12Format, p.Header.Field(2), p.loc)
// 	return t, errors.Wrap(err, "DateTime")
// }

// // FromParticipant returns the field of the same name from the Header record.
// func (p *ELParser) FromParticipant() string {
// 	return p.Header.Field(3)
// }

// // ToParticipant returns the field of the same name from the Header record.
// func (p *ELParser) ToParticipant() string {
// 	return p.Header.Field(4)
// }

// // NMI returns the field of the same name from the NMI data details record.
// func (p *ELParser) NMI() string {
// 	return p.NMIDataDetails.Field(1)
// }

// // NMIConfiguration returns a list of the NMISuffixes applicable to the NMI from the current NMI data details record.
// // Suffixes are assumed to be case-insensistive and are returned in upper case.
// func (p *ELParser) NMIConfiguration() ([]string, error) {
// 	c := strings.ToUpper(p.NMIDataDetails.Field(2))
// 	ss := []string{}
// 	for i := 0; i < len(c); i += 2 {
// 		if i > len(c)-2 {
// 			return ss, errors.New("NMIConfiguration: unexpected character after last NMISuffix")
// 		}
// 		ss = append(ss, c[i:i+2])
// 		//TODO: validate suffix
// 	}
// 	return ss, nil
// }

// // RegisterID returns the field of the same name from the NMI data details record.
// func (p *ELParser) RegisterID() string {
// 	return p.NMIDataDetails.Field(3)
// }

// // NMISuffix returns the field of the same name from the NMI data details record.
// // Suffixes are assumed to be case-insensistive and are returned in upper case.
// func (p *ELParser) NMISuffix() string {
// 	return strings.ToUpper(p.NMIDataDetails.Field(4))
// }

// // MDMDataStreamIdentifier returns the field of the same name from the NMI data details record.
// func (p *ELParser) MDMDataStreamIdentifier() string {
// 	return p.NMIDataDetails.Field(5)
// }

// // MeterSerialNumber returns the field of the same name from the NMI data details record.
// func (p *ELParser) MeterSerialNumber() string {
// 	return p.NMIDataDetails.Field(6)
// }

// // UOM returns the unit of measure field from the NMI data details record.
// func (p *ELParser) UOM() UnitOfMeasure {
// 	return UnitOfMeasure(p.NMIDataDetails.Field(7))
// }

// // IntervalLength returns the time in each interval period.
// func (p *ELParser) IntervalLength() (time.Duration, error) {
// 	t, err := time.ParseDuration(p.NMIDataDetails.Field(8) + "m")
// 	return t, errors.Wrap(err, "IntervalLength")
// }

// // NextScheduledReadDate returns the time in each interval period.
// // If the field is not present, a zero time.Time value is returned with a nil error.
// func (p *ELParser) NextScheduledReadDate() (t time.Time, err error) {
// 	s := p.NMIDataDetails.Field(9)
// 	if s == "" {
// 		return
// 	}
// 	t, err = time.ParseInLocation(Date8Format, s, p.loc)
// 	return t, errors.Wrap(err, "NextScheduledReadDate")
// }
