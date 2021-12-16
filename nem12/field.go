package nem12

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/jufemaiz/go-aemo/nmi"
)

const (
	// nem12Str is used in the version header of the header record.
	nem12Str = "NEM12"
	// fieldParticipantStrLen for the string length of the participant.
	fieldParticipantStrLen = 10
	// fieldRegisterIDStrLen for the string length of the register.
	fieldRegisterIDStrLen = 10
	// fieldMeterSerialNumberStrLen for the string length of the meter serial number.
	fieldMeterSerialNumberStrLen = 12
	// fieldReasonDescriptionStrLen for the string length of the reason description.
	fieldReasonDescriptionStrLen = 240
	// fieldIntervalMax maximum value of the field interval.
	fieldIntervalMax = 288
	// fieldRetServiceOrderStrLen for the string length of the ret service order.
	fieldRetServiceOrderStrLen = 15
	// fieldIndexReadStrLen for the string  length of the index read.
	fieldIndexReadStrLen = 15
)

const (
	// FieldUndefined is for any undefined fields.
	FieldUndefined FieldType = iota
	// FieldRecordIndicator is the field at the start of each row.
	FieldRecordIndicator
	// FieldVersionHeader is the field ... .
	FieldVersionHeader
	// FieldDateTime is the field ... .
	FieldDateTime
	// FieldFromParticipant is the field ... .
	FieldFromParticipant
	// FieldToParticipant is the field ... .
	FieldToParticipant
	// FieldNMI is the field ... .
	FieldNMI
	// FieldNMIConfiguration is the field ... .
	FieldNMIConfiguration
	// FieldRegisterID is the field ... .
	FieldRegisterID
	// FieldNMISuffix is the field ... .
	FieldNMISuffix
	// FieldMDMDataStreamIdentifier is the field ... .
	FieldMDMDataStreamIdentifier
	// FieldMeterSerialNumber is the field ... .
	FieldMeterSerialNumber
	// FieldUnitOfMeasurement is the field ... .
	FieldUnitOfMeasurement
	// FieldIntervalLength is the field ... .
	FieldIntervalLength
	// FieldNextScheduledReadDate is the field ... .
	FieldNextScheduledReadDate
	// FieldIntervalDate is the field ... .
	FieldIntervalDate
	// FieldIntervalValue is the field ... .
	FieldIntervalValue
	// FieldQualityMethod is the field ... .
	FieldQualityMethod
	// FieldReasonCode is the field ... .
	FieldReasonCode
	// FieldReasonDescription is the field ... .
	FieldReasonDescription
	// FieldUpdateDateTime is the field ... .
	FieldUpdateDateTime
	// FieldMSATSLoadDateTime is the field ... .
	FieldMSATSLoadDateTime
	// FieldStartInterval is the field ... .
	FieldStartInterval
	// FieldFinishInterval is the field ... .
	FieldFinishInterval
	// FieldTransactionCode is the field ... .
	FieldTransactionCode
	// FieldRetServiceOrder is the field ... .
	FieldRetServiceOrder
	// FieldReadDateTime is the field ... .
	FieldReadDateTime
	// FieldIndexRead is the field ... .
	FieldIndexRead
)

var (
	fields = []FieldType{ //nolint:gochecknoglobals
		FieldRecordIndicator,
		FieldVersionHeader,
		FieldDateTime,
		FieldFromParticipant,
		FieldToParticipant,
		FieldNMI,
		FieldNMIConfiguration,
		FieldRegisterID,
		FieldNMISuffix,
		FieldMDMDataStreamIdentifier,
		FieldMeterSerialNumber,
		FieldUnitOfMeasurement,
		FieldIntervalLength,
		FieldNextScheduledReadDate,
		FieldIntervalDate,
		FieldIntervalValue,
		FieldQualityMethod,
		FieldReasonCode,
		FieldReasonDescription,
		FieldUpdateDateTime,
		FieldMSATSLoadDateTime,
		FieldStartInterval,
		FieldFinishInterval,
		FieldTransactionCode,
		FieldRetServiceOrder,
		FieldReadDateTime,
		FieldIndexRead,
	}

	fieldName = map[FieldType]string{ //nolint:gochecknoglobals
		FieldRecordIndicator:         "record indicator",
		FieldVersionHeader:           "version header",
		FieldDateTime:                "datetime",
		FieldFromParticipant:         "from participant",
		FieldToParticipant:           "to participant",
		FieldNMI:                     "NMI",
		FieldNMIConfiguration:        "NMI configuration",
		FieldRegisterID:              "register ID",
		FieldNMISuffix:               "NMI suffix",
		FieldMDMDataStreamIdentifier: "MDM data stream identifier",
		FieldMeterSerialNumber:       "meter serial number",
		FieldUnitOfMeasurement:       "unit of measurement",
		FieldIntervalLength:          "interval length",
		FieldNextScheduledReadDate:   "next scheduled read date",
		FieldIntervalDate:            "interval date",
		FieldIntervalValue:           "interval value",
		FieldQualityMethod:           "quality method",
		FieldReasonCode:              "reason code",
		FieldReasonDescription:       "reason description",
		FieldUpdateDateTime:          "update datetime",
		FieldMSATSLoadDateTime:       "MSATS load datetime",
		FieldStartInterval:           "start interval",
		FieldFinishInterval:          "finish interval",
		FieldTransactionCode:         "trans code",
		FieldRetServiceOrder:         "ret service order",
		FieldReadDateTime:            "read datetime",
		FieldIndexRead:               "index read",
	}
)

// Field is the type of Field in a Record of the NEM12 file.
type Field struct {
	Type  FieldType
	Value string
}

// NewField returns a new Field, or errors.
func NewField(ft FieldType, v string) (f Field, err error) {
	f = Field{Type: ft, Value: v}

	if !validFieldType(ft) {
		err = ErrFieldTypeInvalid

		return f, err
	}

	err = f.Validate()

	return f, err
}

// GoString provides go string.
func (f Field) GoString() string {
	return fmt.Sprintf("Field{Type: %s, Value: %q}", f.Type.GoString(), f.Value)
}

// Validate returns any errors for the value of the field.
//nolint:funlen
func (f Field) Validate() error { //nolint:cyclop,gocyclo
	switch f.Type {
	case FieldRecordIndicator:
		return validateFieldRecordIndicator(f.Value)
	case FieldVersionHeader:
		return validateFieldVersionHeader(f.Value)
	case FieldDateTime:
		return validateFieldDateTime(f.Value)
	case FieldFromParticipant:
		return validateFieldFromParticipant(f.Value)
	case FieldToParticipant:
		return validateFieldToParticipant(f.Value)
	case FieldNMI:
		return validateFieldNMI(f.Value)
	case FieldNMIConfiguration:
		return validateFieldNMIConfiguration(f.Value)
	case FieldRegisterID:
		return validateFieldRegisterID(f.Value)
	case FieldNMISuffix:
		return validateFieldNMISuffix(f.Value)
	case FieldMDMDataStreamIdentifier:
		return validateFieldMDMDataStreamIdentifier(f.Value)
	case FieldMeterSerialNumber:
		return validateFieldMeterSerialNumber(f.Value)
	case FieldUnitOfMeasurement:
		return validateFieldUnitOfMeasurement(f.Value)
	case FieldIntervalLength:
		return validateFieldIntervalLength(f.Value)
	case FieldNextScheduledReadDate:
		return validateFieldNextScheduledReadDate(f.Value)
	case FieldIntervalDate:
		return validateFieldIntervalDate(f.Value)
	case FieldIntervalValue:
		return validateFieldIntervalValue(f.Value)
	case FieldQualityMethod:
		return validateFieldQualityMethod(f.Value)
	case FieldReasonCode:
		return validateFieldReasonCode(f.Value)
	case FieldReasonDescription:
		return validateFieldReasonDescription(f.Value)
	case FieldUpdateDateTime:
		return validateFieldUpdateDateTime(f.Value)
	case FieldMSATSLoadDateTime:
		return validateFieldMSATSLoadDateTime(f.Value)
	case FieldStartInterval:
		return validateFieldStartInterval(f.Value)
	case FieldFinishInterval:
		return validateFieldFinishInterval(f.Value)
	case FieldTransactionCode:
		return validateFieldTransactionCode(f.Value)
	case FieldRetServiceOrder:
		return validateFieldRetServiceOrder(f.Value)
	case FieldReadDateTime:
		return validateFieldReadDateTime(f.Value)
	case FieldIndexRead:
		return validateFieldIndexRead(f.Value)
	case FieldUndefined:
		fallthrough //nolint:gocritic
	default:
		return ErrFieldTypeInvalid
	}
}

func validateFieldRecordIndicator(v string) error {
	if v == "" {
		return fmt.Errorf("field record indicator: %w", ErrFieldNil)
	}

	if _, err := NewRecordIndicator(v); err != nil {
		return fmt.Errorf("field record indicator '%s': %w", v, err)
	}

	return nil
}

func validateFieldVersionHeader(v string) error {
	if v == "" {
		return fmt.Errorf("field version: %w", ErrFieldNil)
	}

	if v != nem12Str {
		return fmt.Errorf("field version '%s': %w", v, ErrFieldVersionHeaderInvalid)
	}

	return nil
}

func validateFieldDateTime(v string) error {
	if v == "" {
		return fmt.Errorf("field date time: %w", ErrFieldNil)
	}

	if err := validateDateTime12(v); err != nil {
		return fmt.Errorf("field date time '%s': %w", v, err)
	}

	return nil
}

func validateFieldFromParticipant(v string) error {
	if v == "" {
		return fmt.Errorf("field from participant: %w", ErrFieldNil)
	}

	if len(v) > fieldParticipantStrLen {
		return fmt.Errorf("field from participant '%s': %w", v, ErrFieldParticipantLengthInvalid)
	}

	return nil
}

func validateFieldToParticipant(v string) error {
	if v == "" {
		return fmt.Errorf("field to participant: %w", ErrFieldNil)
	}

	if len(v) > fieldParticipantStrLen {
		return fmt.Errorf("field from participant '%s': %w", v, ErrFieldParticipantLengthInvalid)
	}

	return nil
}

func validateFieldNMI(v string) error {
	if v == "" {
		return fmt.Errorf("field nmi: %w", ErrFieldNil)
	}

	if _, err := nmi.NewNmi(v); err != nil {
		return fmt.Errorf("field nmi '%s': %w", v, err)
	}

	return nil
}

func validateFieldNMIConfiguration(v string) error {
	if v == "" {
		return fmt.Errorf("field nmi configuration: %w", ErrFieldNil)
	}

	pairSize := 2
	pairs := chunkString(v, pairSize)
	exists := map[string]bool{}

	for i, pair := range pairs {
		_, err := NewSuffix(pair)
		if err != nil {
			return fmt.Errorf("field nmi configuration '%s' pair number %d '%s': %w", v, i, pair, err)
		}

		if _, ok := exists[pair]; ok {
			return fmt.Errorf("field nmi configuration '%s' pair number %d '%s': %w", v, i, pair, ErrFieldNMIConfigurationNMISuffixDuplicate)
		}

		exists[pair] = true
	}

	return nil
}

func validateFieldRegisterID(v string) error {
	if len(v) > fieldRegisterIDStrLen {
		return fmt.Errorf("field register ID '%s': %w", v, ErrFieldRegisterIDInvalid)
	}

	return nil
}

func validateFieldNMISuffix(v string) error {
	if v == "" {
		return fmt.Errorf("field nmi suffix: %w", ErrFieldNil)
	}

	if _, err := NewSuffix(v); err != nil {
		return fmt.Errorf("field nmi suffix '%s': %w", v, err)
	}

	return nil
}

// Standing Data details:
//
// Metering Datastream identifier (for MDM). Identifies the Datastream as
// delivered to AEMO for settlements purposes. The value must be a valid suffix
// for this NMI and is active for this date range. The value must comply with
// requirements of the NMI Procedure.
//
// If the MeterInstallCode is COMMSn, MRIM, MRAM, VICAMI or UMCP, the Suffix
// value must be in the form `Nx` where DataStreamType is I or P for an interval
// Datastream. If the MeterInstallCode is BASIC, the Suffix value must be
// numeric.
//
// Ref:
//nolint:lll
// <https://www.aemo.com.au/-/media/Files/Electricity/NEM/Retail_and_Metering/Market_Settlement_And_Transfer_Solutions/2017/Standing-Data-for-MSATS.pdf>
func validateFieldMDMDataStreamIdentifier(v string) error {
	if v == "" {
		// No longer return an error. // return fmt.Errorf("field MDM data stream identifier: %w", ErrFieldNil)
		return nil
	}

	runes := []rune(v)

	for i, r := range runes {
		switch i {
		case 0:
			if string(r) != "N" {
				return fmt.Errorf("field MDM data stream identifier '%s': %w", v, ErrFieldMDMDataStreamIdentifierInvalid)
			}
		case 1:
			if err := ValidateSuffixMeter(string(r)); err != nil {
				return fmt.Errorf("field MDM data stream identifier '%s': %w", v, err)
			}
		}
	}

	return nil
}

func validateFieldMeterSerialNumber(v string) error {
	if len(v) > fieldMeterSerialNumberStrLen {
		return fmt.Errorf("field meter serial number '%s': %w", v, ErrFieldMeterSerialNumberInvalid)
	}

	return nil
}

func validateFieldUnitOfMeasurement(v string) error {
	if v == "" {
		return fmt.Errorf("field unit of measurement: %w", ErrFieldNil)
	}

	if _, err := NewUnit(v); err != nil {
		return fmt.Errorf("field unit of measurement '%s': %w", v, err)
	}

	return nil
}

func validateFieldIntervalLength(v string) error {
	if v == "" {
		return fmt.Errorf("field interval length: %w", ErrFieldNil)
	}

	il, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("field interval length '%s': %w", v, err)
	}

	if il != 5 && il != 15 && il != 30 {
		return fmt.Errorf("field interval length '%s': %w", v, ErrFieldIntervalLengthInvalid)
	}

	return nil
}

func validateFieldNextScheduledReadDate(v string) error {
	if v == "" {
		return nil
	}

	if err := validateDate8(v); err != nil {
		return fmt.Errorf("field next scheduled read date '%s': %w", v, err)
	}

	return nil
}

func validateFieldIntervalDate(v string) error {
	if v == "" {
		return fmt.Errorf("field interval date: %w", ErrFieldNil)
	}

	if err := validateDate8(v); err != nil {
		return fmt.Errorf("field interval date '%s': %w", v, err)
	}

	return nil
}

func validateFieldIntervalValue(v string) error {
	if v == "" {
		return fmt.Errorf("field interval value: %w", ErrFieldNil)
	}

	val, err := decimal.NewFromString(v)
	if err != nil {
		return fmt.Errorf("field interval value '%s': %w", v, err)
	}

	if val.LessThan(decimal.Zero) {
		return fmt.Errorf("field interval value '%s': %w", v, ErrFieldIntervalValueNegative)
	}

	return nil
}

func validateFieldQualityMethod(v string) error {
	if v == "" {
		return fmt.Errorf("field quality method: %w", ErrFieldNil)
	}

	if _, err := NewQualityMethod(v); err != nil {
		return fmt.Errorf("field quality method '%s': %w", v, err)
	}

	return nil
}

func validateFieldReasonCode(v string) error {
	if v == "" {
		return nil
	}

	if _, err := NewReason(v); err != nil {
		return fmt.Errorf("field reason code '%s': %w", v, err)
	}

	return nil
}

func validateFieldReasonDescription(v string) error {
	if len(v) > fieldReasonDescriptionStrLen {
		return fmt.Errorf("field reason description '%s': %w", v, ErrFieldReasonDescriptionLengthInvalid)
	}

	return nil
}

func validateFieldUpdateDateTime(v string) error {
	if v == "" {
		return fmt.Errorf("field update date time: %w", ErrFieldNil)
	}

	if err := validateDateTime14(v); err != nil {
		return fmt.Errorf("field update datetime '%s': %w", v, err)
	}

	return nil
}

func validateFieldMSATSLoadDateTime(v string) error {
	if v == "" {
		return nil
	}

	if err := validateDateTime14(v); err != nil {
		return fmt.Errorf("field update datetime '%s': %w", v, err)
	}

	return nil
}

func validateFieldStartInterval(v string) error {
	if v == "" {
		return fmt.Errorf("field start interval: %w", ErrFieldNil)
	}

	il, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("field start interval '%s': %w", v, err)
	}

	if il > fieldIntervalMax {
		return fmt.Errorf("field start interval '%s': %w", v, ErrFieldIntervalExceedsMaximum)
	}

	if il < 1 {
		return fmt.Errorf("field start interval '%s': %w", v, ErrFieldIntervalNegativeInvalid)
	}

	return nil
}

func validateFieldFinishInterval(v string) error {
	if v == "" {
		return fmt.Errorf("field finish interval: %w", ErrFieldNil)
	}

	il, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("field finish interval '%s': %w", v, err)
	}

	if il > fieldIntervalMax {
		return fmt.Errorf("field finish interval '%s': %w", v, ErrFieldIntervalExceedsMaximum)
	}

	if il < 1 {
		return fmt.Errorf("field start interval '%s': %w", v, ErrFieldIntervalNegativeInvalid)
	}

	return nil
}

func validateFieldTransactionCode(v string) error {
	if v == "" {
		return fmt.Errorf("field transaction code: %w", ErrFieldNil)
	}

	if _, err := NewTransactionCode(v); err != nil {
		return fmt.Errorf("field transaction code '%s': %w", v, err)
	}

	return nil
}

func validateFieldRetServiceOrder(v string) error {
	if len(v) > fieldRetServiceOrderStrLen {
		return fmt.Errorf("field ret service order '%s': %w", v, ErrFieldRetServiceOrderLengthInvalid)
	}

	return nil
}

func validateFieldReadDateTime(v string) error {
	if v == "" {
		return nil
	}

	if err := validateDateTime14(v); err != nil {
		return fmt.Errorf("field read datetime '%s': %w", v, err)
	}

	return nil
}

func validateFieldIndexRead(v string) error {
	if len(v) > fieldIndexReadStrLen {
		return fmt.Errorf("field index read '%s': %w", v, ErrFieldIndexReadLengthInvalid)
	}

	return nil
}

func validateDate8(v string) error {
	if _, err := time.Parse(Date8Format, v); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func validateDateTime12(v string) error {
	if _, err := time.Parse(DateTime12Format, v); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func validateDateTime14(v string) error {
	if _, err := time.Parse(DateTime14Format, v); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// chunkString returns a slice of strings from s.
func chunkString(s string, n int) []string {
	var chunks []string

	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += n {
		j := i + n

		if j > len(runes) {
			j = len(runes)
		}

		chunks = append(chunks, string(runes[i:j]))
	}

	return chunks
}

// FieldType is the type of Field in a Record of the NEM12 file.
type FieldType int

// Fields returns all the fields.
func Fields() []FieldType {
	return fields
}

// validFieldType returns true if the field type is valid.
func validFieldType(f FieldType) bool {
	b, ok := fieldName[f]

	return b != "" && ok
}

// Identifier returns the code used by AEMO.
func (f FieldType) Identifier() string {
	str, ok := fieldName[f]
	if !ok {
		return fmt.Sprintf("%%!FieldType(%d)", f)
	}

	return str
}

// GoString provides the go string.
func (f FieldType) GoString() string {
	return fmt.Sprintf("FieldType(%d)", f)
}
