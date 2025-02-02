package nem12

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/jufemaiz/go-aemo/nmi"
)

// IntervalSet is a set of intervals, with reference to metadata.
type IntervalSet struct {
	// Data
	Data Intervals `json:"data,omitempty"`

	// Metadata, shared.
	Metadata *IntervalMetadata `json:"metadata,omitempty"`
}

// Normalize returns the interval set in SI units.
func (is *IntervalSet) Normalize() (*IntervalSet, error) {
	if is == nil {
		return nil, nil
	}

	if is.Metadata == nil {
		return nil, ErrIntervalMetadataNil
	}

	if is.Metadata.UnitOfMeasure == nil {
		return nil, ErrUnitOfMeasureNil
	}

	uomBase := is.Metadata.UnitOfMeasure.Base()

	norm := &IntervalSet{
		Metadata: &IntervalMetadata{
			Nmi:           is.Metadata.Nmi,
			Meter:         is.Metadata.Meter,
			Suffix:        is.Metadata.Suffix,
			UnitOfMeasure: &uomBase,
		},
		Data: Intervals{},
	}

	uom := is.Metadata.UnitOfMeasure

	for _, v := range is.Data {
		if v == nil {
			continue
		}

		nv, err := v.Normalize(uom)
		if err != nil {
			return nil, err
		}

		norm.Data = append(norm.Data, nv)
	}

	return norm, nil
}

// Interval is duration of time from a start to a finish, with a value.
type Interval struct {
	// Data.
	Time           time.Time     `json:"datetime"` //nolint:tagliatelle
	IntervalLength time.Duration `json:"intervalLength"`
	Value          IntervalValue `json:"intervalValue"` //nolint:tagliatelle

	// Metadata, shared.
	Metadata *IntervalMetadata `json:"metadata,omitempty"`
}

// Normalize returns the interval in SI units.
func (i *Interval) Normalize(uom *UnitOfMeasure) (*Interval, error) {
	if i == nil {
		return nil, ErrIntervalNil
	}

	norm := &Interval{
		Time:           i.Time,
		IntervalLength: i.IntervalLength,
		Value: IntervalValue{
			Value:             i.Value.Value,
			DecimalValue:      i.Value.DecimalValue,
			QualityFlag:       i.Value.QualityFlag,
			MethodFlag:        i.Value.MethodFlag,
			ReasonCode:        i.Value.ReasonCode,
			ReasonDescription: i.Value.ReasonDescription,
			UpdateDateTime:    i.Value.UpdateDateTime,
			MSATSLoadDateTime: i.Value.MSATSLoadDateTime,
		},
	}

	if uom != nil {
		if err := uom.Validate(); err != nil {
			return nil, fmt.Errorf("unit of measurement: %w", err)
		}

		norm.Value.Value *= uom.Multiplier()
		norm.Value.DecimalValue = norm.Value.DecimalValue.Mul(uom.DecimalMultiplier())
	}

	return norm, nil
}

// Intervals is a slice of Interval.
type Intervals []*Interval

// IntervalLength for custom json marshalling of a duration.
type IntervalLength time.Duration

// An IntervalValue represents a single meter interval value as presented by an
// NEM12 file.
type IntervalValue struct {
	Value        float64         `json:"value"`        // Value of the interval in the SI unit of measure.
	DecimalValue decimal.Decimal `json:"decimalValue"` // Value of the interval in SI unit of measure as a decimal.
	// Quality flag applicable to this value.
	QualityFlag Quality `json:"quality"` //nolint:tagliatelle
	// Method flag applicable to this value.
	MethodFlag *Method `json:"method,omitempty"` //nolint:tagliatelle
	// ReasonCode applicable to this value.
	ReasonCode        *Reason    `json:"reason,omitempty"`            //nolint:tagliatelle
	ReasonDescription *string    `json:"reasonDescription,omitempty"` // Text describing the reason for when ReasonCode equals FreeTextDescription.
	UpdateDateTime    *time.Time `json:"updateDateTime,omitempty"`    // Time that the metering data was created or updated as reported by the MDP.
	MSATSLoadDateTime *time.Time `json:"msatsLoadDateTime,omitempty"` // Time that the metering data was loaded into MSATS.
}

// IntervalMetadata holds the metadata for an interval.
type IntervalMetadata struct {
	Nmi           *nmi.Nmi       `json:"nmi,omitempty"`
	Meter         *nmi.Meter     `json:"meter,omitempty"`
	Suffix        *Suffix        `json:"suffix,omitempty"`
	UnitOfMeasure *UnitOfMeasure `json:"unit,omitempty"` //nolint:tagliatelle
}

// intervalEvent captures an interval event record.
type intervalEvent struct {
	Start             int           `json:"start"`
	Finish            int           `json:"finish"`
	QualityMethod     QualityMethod `json:"qualityMethod"`
	Reason            *Reason       `json:"reason"`
	ReasonDescription *string       `json:"reasonDescription"`
}

// newIntervalEvent returns a new interval event for a record.
func newIntervalEvent(rec Record) (*intervalEvent, error) {
	var (
		start      int
		finish     int
		qm         QualityMethod
		reason     *Reason
		reasonDesc *string
		err        error
	)

	for _, field := range rec {
		switch field.Type {
		case FieldRecordIndicator:
			ri, err := NewRecordIndicator(field.Value)
			if err != nil {
				return nil, err
			}

			if ri != RecordIntervalEvent {
				return nil, ErrRecordIndicatorInvalid
			}

		case FieldStartInterval:
			start, err = strconv.Atoi(field.Value)
			if err != nil {
				return nil, fmt.Errorf("field Start '%s' %w", field.Value, ErrIsInvalid)
			}

		case FieldFinishInterval:
			finish, err = strconv.Atoi(field.Value)
			if err != nil {
				return nil, fmt.Errorf("field Finish '%s' %w", field.Value, ErrIsInvalid)
			}

		case FieldQualityMethod:
			qm, err = NewQualityMethod(field.Value)
			if err != nil {
				return nil, err
			}

		case FieldReasonCode:
			if field.Value != "" {
				rc, err := NewReason(field.Value)
				if err != nil {
					return nil, err
				}

				reason = &rc
			}

		case FieldReasonDescription:
			if field.Value != "" {
				str := field.Value

				reasonDesc = &str
			}
		default:
			return nil, fmt.Errorf("field '%s' %w", field.GoString(), ErrIsInvalid)
		}
	}

	return &intervalEvent{
		Start:             start,
		Finish:            finish,
		QualityMethod:     qm,
		Reason:            reason,
		ReasonDescription: reasonDesc,
	}, nil
}

// validate ensures the interval event is valid.
func (i *intervalEvent) validate(intervalCount int) error {
	// a nil event is invalid.
	if i == nil {
		return fmt.Errorf("interval event %w", ErrIsNil)
	}

	if i.Start < 1 || i.Start > intervalCount {
		return fmt.Errorf("start '%d' not in range [1,%d] %w", i.Start, intervalCount, ErrIsInvalid)
	}

	if i.Finish < i.Start || i.Finish > intervalCount {
		return fmt.Errorf("finish '%d' not in range [%d,%d] %w", i.Finish, i.Start, intervalCount, ErrIsInvalid)
	}

	return nil
}

// intervalEvents is a slice of events.
type intervalEvents []*intervalEvent

// validate ensures the collection of interval events are valid.
func (i intervalEvents) validate(intervalCount int) error {
	// An empty set of events is valid.
	if len(i) == 0 {
		return nil
	}

	// Validate the individual interval events.
	for j, ev := range i {
		if err := ev.validate(intervalCount); err != nil {
			return fmt.Errorf("interval event %d: %w", j, err)
		}
	}

	// Check continuity of events.
	im := i.intervalMap()

	for i := 1; i <= intervalCount; i++ {
		_, ok := im[i]
		if !ok {
			return fmt.Errorf("interval %d event information %w", i, ErrIsMissing)
		}
	}

	return nil
}

// intervalMap returns a map if interval events to the interval number.
func (i intervalEvents) intervalMap() map[int]*intervalEvent {
	im := map[int]*intervalEvent{}

	for _, ev := range i {
		if ev == nil {
			continue
		}

		for j := ev.Start; j <= ev.Finish; j++ {
			im[j] = ev
		}
	}

	return im
}
