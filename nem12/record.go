package nem12

import (
	"fmt"
	"strings"
	"time"
)

const (
	// IntervalLengthPT5M is the duration of an interval for 5 minutes.
	IntervalLengthPT5M = time.Duration(5 * secondsInMinute * nanosInSecond)
	// IntervalLengthPT15M is the duration of an interval for 15 minutes.
	IntervalLengthPT15M = time.Duration(15 * secondsInMinute * nanosInSecond)
	// IntervalLengthPT30M is the duration of an interval for 30 minutes.
	IntervalLengthPT30M = time.Duration(30 * secondsInMinute * nanosInSecond)
	// IntervalLengthPT24H is the duration of an interval for 24 hours.
	IntervalLengthPT24H = time.Duration(hoursInDay * minutesInHour * secondsInMinute * nanosInSecond)
)

const (
	// RecordUndefined is the record that is undefined.
	RecordUndefined RecordIndicator = iota
	// RecordHeader is the header record.
	RecordHeader
	// RecordNMIDataDetails is the NMI data details record.
	//
	// Multiple 300-500 record blocks are allowed within a single 200 record.
	RecordNMIDataDetails
	// RecordIntervalData is the interval data record.
	//
	// 300 records must be presented in date sequential order. For example, with a
	// series of Meter Readings for a period, the current record is the next
	// incremental IntervalDate after the previous record. Or, where data for
	// individual, non-consecutive days is sent, the IntervalDate for each 300
	// record is later than the previous one.
	//
	// Where the same QualityMethod and ReasonCode apply to all IntervalValues in
	// the 300 record, the QualityMethod, ReasonCode and ReasonDescription in the
	// 300 Record must be used. If either of these fields contains multiple values
	// for the IntervalValues, the QualityMethod in the 300 record must be set to
	// “V” and the 400 record must be provided.
	//
	// The use of ‘V’ as the quality method in this example indicates the
	// QualityMethod, ReasonCode or ReasonDescription vary across the day and will
	// be provided, for each Interval, in the 400 records that would immediately
	// follow this record. Refer 4.5 for details on the use of the 400 records.
	RecordIntervalData
	// RecordIntervalEvent is the interval event record.
	//
	// This record is mandatory where the QualityFlag is ‘V’ in the 300 record or
	// where the quality flag is‘A’ and reason codes 79, 89, and 61 are used.
	//
	// The StartInterval/EndInterval pairs must be presented in ascending record
	// order. The StartInterval/EndInterval period must cover an entire day
	// without gaps or overlaps. For example, (based on a 30-minute Interval):
	//
	//  400,1,26,A,,
	//  400,27,31,S53,9,
	//  400,32,48,E52,,
	//
	// Refer to section 2 (c) of AEMO's MDFF specification for further rules
	// regarding the use of this record.
	RecordIntervalEvent
	// RecordB2BDetails is the business-to-business details record.
	//
	// This record is mandatory where a manual Meter Reading has been performed or
	// attempted. Only valid "500" records associated with the current Meter
	// Reading period must be provided. For example, a 500 record associated with
	// a Substitute will become invalid if Actual Metering Data subsequently
	// replace the Substitutes.
	//
	// This record must be repeated where multiple TransCodes or RetServiceOrders
	// apply to the day.
	RecordB2BDetails
	// RecordEnd is the end record.
	RecordEnd
)

var (
	// IntervalLengths is a slice of all valid interval lengths.
	IntervalLengths = []time.Duration{ //nolint:gochecknoglobals
		IntervalLengthPT5M,
		IntervalLengthPT15M,
		IntervalLengthPT30M,
	}

	// recordIndicators is a slice of all record indicators.
	recordIndicators = []RecordIndicator{ //nolint:gochecknoglobals
		RecordHeader,
		RecordNMIDataDetails,
		RecordIntervalData,
		RecordIntervalEvent,
		RecordB2BDetails,
		RecordEnd,
	}

	// RecordIndicatorName maps a record indicator to a name.
	RecordIndicatorName = map[RecordIndicator]string{ //nolint:gochecknoglobals
		RecordHeader:         "100",
		RecordNMIDataDetails: "200",
		RecordIntervalData:   "300",
		RecordIntervalEvent:  "400",
		RecordB2BDetails:     "500",
		RecordEnd:            "900",
	}

	// RecordIndicatorValue maps a record indicator from a name.
	RecordIndicatorValue = map[string]RecordIndicator{ //nolint:gochecknoglobals
		"100": RecordHeader,
		"200": RecordNMIDataDetails,
		"300": RecordIntervalData,
		"400": RecordIntervalEvent,
		"500": RecordB2BDetails,
		"900": RecordEnd,
	}

	// recordIndicatorDescriptions maps a record indicator to its description.
	recordIndicatorDescriptions = map[RecordIndicator]string{ //nolint:gochecknoglobals
		RecordHeader:         "header",
		RecordNMIDataDetails: "NMI data details",
		RecordIntervalData:   "interval data",
		RecordIntervalEvent:  "interval event",
		RecordB2BDetails:     "B2B details",
		RecordEnd:            "end of data",
	}

	// recordFields maps a record with the fields.
	recordFields = map[RecordIndicator][]FieldType{ //nolint:gochecknoglobals
		RecordHeader: {
			FieldRecordIndicator, FieldVersionHeader, FieldDateTime, FieldFromParticipant,
			FieldToParticipant,
		},
		RecordNMIDataDetails: {
			FieldRecordIndicator, FieldNMI, FieldNMIConfiguration, FieldRegisterID, FieldNMISuffix,
			FieldMDMDataStreamIdentifier, FieldMeterSerialNumber, FieldUnitOfMeasurement,
			FieldIntervalLength, FieldNextScheduledReadDate,
		},
		RecordIntervalData: {
			FieldRecordIndicator, FieldIntervalDate, FieldIntervalValue, FieldQualityMethod,
			FieldReasonCode, FieldReasonDescription, FieldUpdateDateTime, FieldMSATSLoadDateTime,
		},
		RecordIntervalEvent: {
			FieldRecordIndicator, FieldStartInterval, FieldFinishInterval, FieldQualityMethod,
			FieldReasonCode, FieldReasonDescription,
		},
		RecordB2BDetails: {
			FieldRecordIndicator, FieldTransactionCode, FieldRetServiceOrder, FieldReadDateTime,
			FieldIndexRead,
		},
		RecordEnd: {FieldRecordIndicator},
	}
)

// Record holds a record (row) of NEM12 data, separated by new line.
type Record []Field

// NewRecord returns a new record, given the field values.
func NewRecord(vals []string) (r Record, err error) {
	if len(vals) == 0 {
		return nil, ErrRecordNil
	}

	ri, err := NewRecordIndicator(vals[0])
	if err != nil {
		return nil, err
	}

	// Break out here to a NewRecordIntervalData, using the known non-interval data
	// count to indicate number of values.
	if ri == RecordIntervalData {
		n := len(vals) - 7

		return NewRecordIntervalData(n, vals)
	}

	r, err = newRecordFor(ri, vals)
	if err != nil {
		return nil, err
	}

	switch ri {
	case RecordHeader:
		err = validateRecordHeader(r)
	case RecordNMIDataDetails:
		err = validateRecordNMIDataDetails(r)
	case RecordIntervalData:
		return nil, ErrRecordIntervalDataWithoutIntervalCountInvalid
	case RecordIntervalEvent:
		err = validateRecordIntervalEvent(r)
	case RecordB2BDetails:
		err = validateRecordB2BDetails(r)
	case RecordEnd:
		err = validateRecordEnd(r)
	case RecordUndefined:
		fallthrough //nolint:gocritic
	default:
		return r, ErrRecordIndicatorInvalid
	}

	return r, err
}

// NewRecordIntervalData returns a new record, given the number of expected intervals
// and the field values.
func NewRecordIntervalData(n int, vals []string) (Record, error) {
	if len(vals) == 0 {
		return nil, ErrRecordNil
	}

	ri, err := NewRecordIndicator(vals[0])
	if err != nil {
		return nil, err
	}

	f, err := ri.IntervalDataFields(n)
	if err != nil {
		return nil, err
	}

	r, err := newRecord(f, vals)
	if err != nil {
		return nil, err
	}

	if err := validateRecordIntervalData(r); err != nil {
		return nil, err
	}

	return r, err
}

// newRecordFor returns a new record, given the indicator and values.
func newRecordFor(ri RecordIndicator, vals []string) (Record, error) {
	f := ri.Fields()

	return newRecord(f, vals)
}

// newRecordFor returns a new record, given the fields and values.
func newRecord(f []FieldType, vals []string) (Record, error) {
	if len(vals) != len(f) {
		return nil, fmt.Errorf("value count %d, expected %d: %w", len(vals), len(f), ErrRecordFieldLengthInvalid)
	}

	r := Record{}

	for i, v := range vals {
		field, err := NewField(f[i], v)
		if err != nil {
			return r, fmt.Errorf("field %d (%s) '%s': %w", i, f[i].Identifier(), v, err)
		}

		if err := field.Validate(); err != nil {
			return r, fmt.Errorf("field %d (%s) '%s': %w", i, f[i].Identifier(), v, err)
		}

		r = append(r, field)
	}

	return r, nil
}

// validateRecordHeader validates cross reference field dependencies within a row.
func validateRecordHeader(r Record) error {
	return nil
}

// validateRecordNMIDataDetails validates cross reference field dependencies within a row.
func validateRecordNMIDataDetails(r Record) error {
	pairCount := 2
	pairs := chunkString(r[2].Value, pairCount)
	found := false

	for _, pair := range pairs {
		if pair == r[4].Value {
			found = true

			break
		}
	}

	if !found {
		return ErrFieldNMIConfigurationMissingNMISuffix
	}

	return nil
}

// validateRecordIntervalData validates cross reference field dependencies within a row.
func validateRecordIntervalData(r Record) error {
	return nil
}

// validateRecordIntervalEvent validates cross reference field dependencies within a row.
func validateRecordIntervalEvent(r Record) error {
	return nil
}

// validateRecordB2BDetails validates cross reference field dependencies within a row.
func validateRecordB2BDetails(r Record) error {
	return nil
}

// validateRecordEnd validates cross reference field dependencies within a row.
func validateRecordEnd(r Record) error {
	return nil
}

// Field returns the nth field of a Record, or the zero value of a Field string.
func (r Record) Field(n int) Field {
	if n < 0 || n >= len(r) {
		return Field{}
	}

	return r[n]
}

// Indicator returns the indicator for the record.
func (r Record) Indicator() RecordIndicator {
	if len(r) == 0 {
		return RecordUndefined
	}

	ri, err := NewRecordIndicator(r[0].Value)
	if err != nil {
		return RecordUndefined
	}

	return ri
}

// RecordIndicator provides the indicator of the type record of record within
// the NEM12 data structure.
type RecordIndicator int

// RecordIndicators returns all record indicators.
func RecordIndicators() []RecordIndicator {
	return recordIndicators
}

// NewRecordIndicator returns a new record indicator, along with errors if not valid.
func NewRecordIndicator(s string) (RecordIndicator, error) {
	if s == "" {
		return RecordUndefined, ErrRecordIndicatorNil
	}

	r, ok := RecordIndicatorValue[strings.ToUpper(s)]
	if !ok {
		return RecordUndefined, ErrRecordIndicatorInvalid
	}

	return r, nil
}

// Validate ensures a record indicator is valid.
func (r RecordIndicator) Validate() error {
	if _, ok := RecordIndicatorName[r]; !ok {
		return ErrRecordIndicatorInvalid
	}

	return nil
}

// Identifier returns the identifier for the record indicator.
func (r RecordIndicator) Identifier() string {
	id, ok := RecordIndicatorName[r]
	if !ok {
		return fmt.Sprintf("RecordIndicator(%d)", r)
	}

	return id
}

// Description returns the description of a record indicator, along with an error if
// it is an unknown value.
func (r RecordIndicator) Description() (string, error) {
	s, ok := recordIndicatorDescriptions[r]
	if !ok {
		return fmt.Sprintf("%%!RecordIndicator(%d)", r), fmt.Errorf("record indicator description '%d': %w", r, ErrRecordIndicatorInvalid)
	}

	return s, nil
}

// String returns a text representation of the reason.
func (r RecordIndicator) String() string {
	desc, err := r.Description()
	if err != nil {
		return desc
	}

	return fmt.Sprintf("\"%s: %s\"", r.Identifier(), desc)
}

// GoString returns a text representation of the reason to satisfy the GoStringer
// interface.
func (r RecordIndicator) GoString() string {
	return fmt.Sprintf("%%!RecordIndicator(%d)", r)
}

// Fields returns the expected set of fields, in order, for a record.
func (r RecordIndicator) Fields() []FieldType {
	f, ok := recordFields[r]
	if !ok {
		return nil
	}

	return f
}

// IntervalDataFields returns the expected set of fields, in order, for an interval
// data record, given the expected number of intervals.
func (r RecordIndicator) IntervalDataFields(n int) ([]FieldType, error) {
	fields, ok := recordFields[r]
	if !ok {
		return nil, fmt.Errorf("record indicator interval fields '%d': %w", r, ErrRecordIndicatorInvalid)
	}

	if r != RecordIntervalData {
		return fields, nil
	}

	resp := []FieldType{}

	for _, f := range fields {
		if f != FieldIntervalValue {
			resp = append(resp, f)

			continue
		}

		for i := 0; i < n; i++ {
			resp = append(resp, f)
		}
	}

	return resp, nil
}
