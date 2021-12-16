package nem12

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	// ErrIsDuplicated if is duplicated.
	ErrIsDuplicated = errors.New("is duplicated")
	// ErrIsInvalid if is invalid.
	ErrIsInvalid = errors.New("is invalid")
	// ErrIsMissing if is missing.
	ErrIsMissing = errors.New("is missing")
	// ErrIsNil if is nil.
	ErrIsNil = errors.New("is nil")
	// ErrParseFailed if parse has failed.
	ErrParseFailed = errors.New("parse has failed")

	// ErrFieldNil if field is empty.
	ErrFieldNil = fmt.Errorf("field %w", ErrIsNil)
	// ErrFieldInvalid if field is invalid.
	ErrFieldInvalid = fmt.Errorf("field %w", ErrIsInvalid)
	// ErrFieldParticipantLengthInvalid if field participant length is invalid.
	ErrFieldParticipantLengthInvalid = fmt.Errorf("field participan %w", ErrLengthInvalid)
	// ErrFieldIndexReadLengthInvalid if field index read length is invalid.
	ErrFieldIndexReadLengthInvalid = fmt.Errorf("field index read %w", ErrLengthInvalid)
	// ErrFieldIntervalExceedsMaximum if field interval exceeds maximum value.
	ErrFieldIntervalExceedsMaximum = errors.New("field interval exceeds maximum value")
	// ErrFieldIntervalLengthInvalid if field interval length is invalid.
	ErrFieldIntervalLengthInvalid = fmt.Errorf("field interval %w", ErrLengthInvalid)
	// ErrFieldIntervalNegativeInvalid if field interval negative is invalid.
	ErrFieldIntervalNegativeInvalid = fmt.Errorf("field interval negative %w", ErrIsInvalid)
	// ErrFieldIntervalValueNegative if field interval value negative.
	ErrFieldIntervalValueNegative = fmt.Errorf("field interval value negative %w", ErrIsInvalid)
	// ErrFieldMDMDataStreamIdentifierInvalid if field MDM data stream identifier invalid.
	ErrFieldMDMDataStreamIdentifierInvalid = fmt.Errorf("field MDM data stream identifier %w", ErrIsInvalid)
	// ErrFieldMeterSerialNumberInvalid if field meter serial number invalid.
	ErrFieldMeterSerialNumberInvalid = fmt.Errorf("field meter serial number %w", ErrIsInvalid)
	// ErrFieldNMIConfigurationNMISuffixDuplicate if nmi configuration has a nmi suffix that is duplicated.
	ErrFieldNMIConfigurationNMISuffixDuplicate = fmt.Errorf("nmi configuration has a nmi suffix that %w", ErrIsDuplicated)
	// ErrFieldNMIConfigurationMissingNMISuffix if nmi configuration nmi suffix is missing.
	ErrFieldNMIConfigurationMissingNMISuffix = fmt.Errorf("nmi configuration nmi suffix %w", ErrIsMissing)
	// ErrFieldReasonDescriptionLengthInvalid if field reason description length is invalid.
	ErrFieldReasonDescriptionLengthInvalid = fmt.Errorf("field reason description %w", ErrLengthInvalid)
	// ErrFieldRegisterIDInvalid if field register id invalid.
	ErrFieldRegisterIDInvalid = fmt.Errorf("field register id %w", ErrIsInvalid)
	// ErrFieldRetServiceOrderLengthInvalid if field ret service order length is invalid.
	ErrFieldRetServiceOrderLengthInvalid = fmt.Errorf("field ret service order %w", ErrLengthInvalid)
	// ErrFieldVersionHeaderInvalid if field version header invalid.
	ErrFieldVersionHeaderInvalid = fmt.Errorf("field version header %w", ErrIsInvalid)
	// ErrFieldTypeInvalid if field type is invalid.
	ErrFieldTypeInvalid = fmt.Errorf("field type %w", ErrIsInvalid)
	// ErrInstallNil if install is empty.
	ErrInstallNil = fmt.Errorf("install %w", ErrIsNil)
	// ErrInstallInvalid if install is invalid.
	ErrInstallInvalid = fmt.Errorf("install %w", ErrIsInvalid)
	// ErrIntervalMetadataNil if interval metadata is nil.
	ErrIntervalMetadataNil = fmt.Errorf("interval metadata %w", ErrIsNil)
	// ErrIntervalNil if interval is nil.
	ErrIntervalNil = fmt.Errorf("interval %w", ErrIsNil)
	// ErrLengthInvalid if length is invalid.
	ErrLengthInvalid = fmt.Errorf("length %w", ErrIsInvalid)
	// ErrMethodNil if method flag is empty.
	ErrMethodNil = fmt.Errorf("method flag %w", ErrIsNil)
	// ErrMethodInvalid if method flag is invalid.
	ErrMethodInvalid = fmt.Errorf("method flag %w", ErrIsInvalid)
	// ErrMethodTypeNil if method type is empty.
	ErrMethodTypeNil = fmt.Errorf("method type %w", ErrIsNil)
	// ErrMethodTypeInvalid if method type is invalid.
	ErrMethodTypeInvalid = fmt.Errorf("method type %w", ErrIsInvalid)
	// ErrParseFieldCountInvalid if wrong number of fields.
	ErrParseFieldCountInvalid = fmt.Errorf("parse fields count %w", ErrIsInvalid)
	// ErrParseIntervalDataLengthInvalid if parse error: interval data length invalid.
	ErrParseIntervalDataLengthInvalid = fmt.Errorf("interval data length invalid: %w", ErrParseFailed)
	// ErrParseIntervalLengthInvalid if parse error: interval length invalid.
	ErrParseIntervalLengthInvalid = fmt.Errorf("interval length invalid: %w", ErrParseFailed)
	// ErrParseInvalidState if parse error: invalid state.
	ErrParseInvalidState = fmt.Errorf("invalid state: %w", ErrParseFailed)
	// ErrParseRecordHeaderMissing if parse error: record header missing.
	ErrParseRecordHeaderMissing = fmt.Errorf("record header missing: %w", ErrParseFailed)
	// ErrParseRecordNMIDataDetailsMissing if parse error: record NMI data details missing.
	ErrParseRecordNMIDataDetailsMissing = fmt.Errorf("record NMI data details missing: %w", ErrParseFailed)
	// ErrParseRecordIntervalDataMissing if parse error: record interval data missing.
	ErrParseRecordIntervalDataMissing = fmt.Errorf("record interval data missing: %w", ErrParseFailed)
	// ErrParseRecordIntervalEventMissing if parse error: record interval event missing.
	ErrParseRecordIntervalEventMissing = fmt.Errorf("record interval event missing: %w", ErrParseFailed)
	// ErrParseRecordEndMissing if parse error: record end missing.
	ErrParseRecordEndMissing = fmt.Errorf("record end missing: %w", ErrParseFailed)
	// ErrParseUnexpectedEOF if parse error: unexpected EOF.
	ErrParseUnexpectedEOF = fmt.Errorf("unexpected EOF: %w", ErrParseFailed)
	// ErrQualityNil if quality flag is empty.
	ErrQualityNil = fmt.Errorf("quality flag %w", ErrIsNil)
	// ErrQualityInvalid if quality flag is invalid.
	ErrQualityInvalid = fmt.Errorf("quality flag %w", ErrIsInvalid)
	// ErrQualityMethodNil if quality method is empty.
	ErrQualityMethodNil = fmt.Errorf("quality method %w", ErrIsNil)
	// ErrQualityMethodInvalid if quality method is invalid.
	ErrQualityMethodInvalid = fmt.Errorf("quality method %w", ErrIsInvalid)
	// ErrQualityMethodLengthInvalid if quality method length is invalid.
	ErrQualityMethodLengthInvalid = fmt.Errorf("quality method %w", ErrLengthInvalid)
	// ErrQualityMissingMethod if quality method missing required method.
	ErrQualityMissingMethod = fmt.Errorf("quality method required method %w", ErrIsMissing)
	// ErrReaderNil if reader is nil.
	ErrReaderNil = fmt.Errorf("reader %w", ErrIsNil)
	// ErrRecordFieldLengthInvalid if record field length is invalid.
	ErrRecordFieldLengthInvalid = fmt.Errorf("record field %w", ErrLengthInvalid)
	// ErrRecordNil if record is empty.
	ErrRecordNil = fmt.Errorf("record %w", ErrIsNil)
	// ErrReasonCodeNil if reason code is empty.
	ErrReasonCodeNil = fmt.Errorf("reason code %w", ErrIsNil)
	// ErrReasonCodeInvalid if reason code is invalid.
	ErrReasonCodeInvalid = fmt.Errorf("reason code %w", ErrIsInvalid)
	// ErrRecordIndicatorNil if record indicator is empty.
	ErrRecordIndicatorNil = fmt.Errorf("record indicator %w", ErrIsNil)
	// ErrRecordIndicatorInvalid if record indicator is invalid.
	ErrRecordIndicatorInvalid = fmt.Errorf("record indicator %w", ErrIsInvalid)
	// ErrRecordIntervalDataWithoutIntervalCountInvalid if record interval data without interval count is invalid.
	ErrRecordIntervalDataWithoutIntervalCountInvalid = fmt.Errorf("record interval data without interval count %w", ErrIsInvalid)
	// ErrSuffixNil if suffix is empty.
	ErrSuffixNil = fmt.Errorf("suffix %w", ErrIsNil)
	// ErrSuffixLengthInvalid if suffix is invalid.
	ErrSuffixLengthInvalid = fmt.Errorf("suffix %w", ErrIsInvalid)
	// ErrSuffixMeterInvalid if suffix meter is invalid.
	ErrSuffixMeterInvalid = fmt.Errorf("suffix meter %w", ErrIsInvalid)
	// ErrSuffixTypeNil if suffix type is empty.
	ErrSuffixTypeNil = fmt.Errorf("suffix type %w", ErrIsNil)
	// ErrSuffixTypeInvalid if suffix type is invalid.
	ErrSuffixTypeInvalid = fmt.Errorf("suffix type %w", ErrIsInvalid)
	// ErrTransactionCodeNil if transaction code is empty.
	ErrTransactionCodeNil = fmt.Errorf("transaction code %w", ErrIsNil)
	// ErrTransactionCodeInvalid if transaction code is invalid.
	ErrTransactionCodeInvalid = fmt.Errorf("transaction code %w", ErrIsInvalid)
	// ErrUnitOfMeasureNil if unit of measure is empty.
	ErrUnitOfMeasureNil = fmt.Errorf("unit of measure %w", ErrIsNil)
	// ErrUnitOfMeasureInvalid if unit of measure is invalid.
	ErrUnitOfMeasureInvalid = fmt.Errorf("unit of measure %w", ErrIsInvalid)
)

// A ParseError is returned for parsing errors. Line numbers are 1-indexed and columns
// are 0-indexed.
type ParseError struct {
	Line  int   // Line where the error occurred.
	Field int   // Field where the error occurred.
	Err   error // The actual error.
}

// Error to meet the error interface.
func (e *ParseError) Error() string {
	if e == nil {
		return ""
	}

	if errors.Is(e.Err, ErrParseFieldCountInvalid) {
		return fmt.Sprintf("record on line %d: %v", e.Line, e.Err)
	}

	return fmt.Sprintf("parse error on line %d, field %d: %v", e.Line, e.Field, e.Err)
}

// Unwrap for the interface.
func (e *ParseError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}
