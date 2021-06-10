package nem12

import "fmt"

const (
	// MethodUndefined for undefined methods.
	MethodUndefined = Method(0)
	// Method11Check for check.
	Method11Check = Method(11)
	// Method12Calculated for calculated.
	Method12Calculated = Method(12)
	// Method13SCADA for scada.
	Method13SCADA = Method(13)
	// Method14LikeDay for like day.
	Method14LikeDay = Method(14)
	// Method15AverageLikeDay for average like day.
	Method15AverageLikeDay = Method(15)
	// Method16Agreed for agreed.
	Method16Agreed = Method(16)
	// Method17Linear for linear.
	Method17Linear = Method(17)
	// Method18Alternate for alternate.
	Method18Alternate = Method(18)
	// Method19Zero for zero.
	Method19Zero = Method(19)
	// Method51PreviousYear for previous year.
	Method51PreviousYear = Method(51)
	// Method52PreviousRead for previous read.
	Method52PreviousRead = Method(52)
	// Method53Revision for revision.
	Method53Revision = Method(53)
	// Method54Linear for linear.
	Method54Linear = Method(54)
	// Method55Agreed for agreed.
	Method55Agreed = Method(55)
	// Method56PriortoFirstReadAgreed for prior to first read - agreed.
	Method56PriortoFirstReadAgreed = Method(56)
	// Method57CustomerClass for customer class.
	Method57CustomerClass = Method(57)
	// Method58Zero for zero.
	Method58Zero = Method(58)
	// Method61PreviousYear for previous year.
	Method61PreviousYear = Method(61)
	// Method62PreviousRead for previous read.
	Method62PreviousRead = Method(62)
	// Method63CustomerClass for customer class.
	Method63CustomerClass = Method(63)
	// Method64Agreed for agreed.
	Method64Agreed = Method(64)
	// Method65ADL for adl.
	Method65ADL = Method(65)
	// Method66Revision for revision.
	Method66Revision = Method(66)
	// Method67CustomerRead for customer read.
	Method67CustomerRead = Method(67)
	// Method68Zero for zero.
	Method68Zero = Method(68)
	// Method71Recalculation for recalculation.
	Method71Recalculation = Method(71)
	// Method72RevisedTable for revised table.
	Method72RevisedTable = Method(72)
	// Method73RevisedAlgorithm for revised algorithm.
	Method73RevisedAlgorithm = Method(73)
	// Method74Agreed for agreed.
	Method74Agreed = Method(74)
	// Method75ExistingTable for existing table.
	Method75ExistingTable = Method(75)

	// MethodTypeUndefined
	MethodTypeUndefined MethodType = iota
	// MethodTypeEstimated for methods using estimation.
	MethodTypeEstimated
	// MethodTypeEstimated for methods using substitution.
	MethodTypeSubstituted
)

var (
	// methods is a slice of all valid methods.
	methods = []Method{ //nolint:gochecknoglobals
		MethodUndefined,
		Method11Check,
		Method12Calculated,
		Method13SCADA,
		Method14LikeDay,
		Method15AverageLikeDay,
		Method16Agreed,
		Method17Linear,
		Method18Alternate,
		Method19Zero,
		Method51PreviousYear,
		Method52PreviousRead,
		Method53Revision,
		Method54Linear,
		Method55Agreed,
		Method56PriortoFirstReadAgreed,
		Method57CustomerClass,
		Method58Zero,
		Method61PreviousYear,
		Method62PreviousRead,
		Method63CustomerClass,
		Method64Agreed,
		Method65ADL,
		Method66Revision,
		Method67CustomerRead,
		Method68Zero,
		Method71Recalculation,
		Method72RevisedTable,
		Method73RevisedAlgorithm,
		Method74Agreed,
		Method75ExistingTable,
	}

	// MethodName maps a method to its string equivalent.
	MethodName = map[Method]string{ //nolint:gochecknoglobals
		Method11Check:                  "Type11",
		Method12Calculated:             "Type12",
		Method13SCADA:                  "Type13",
		Method14LikeDay:                "Type14",
		Method15AverageLikeDay:         "Type15",
		Method16Agreed:                 "Type16",
		Method17Linear:                 "Type17",
		Method18Alternate:              "Type18",
		Method19Zero:                   "Type19",
		Method51PreviousYear:           "Type51",
		Method52PreviousRead:           "Type52",
		Method53Revision:               "Type53",
		Method54Linear:                 "Type54",
		Method55Agreed:                 "Type55",
		Method56PriortoFirstReadAgreed: "Type56",
		Method57CustomerClass:          "Type57",
		Method58Zero:                   "Type58",
		Method61PreviousYear:           "Type61",
		Method62PreviousRead:           "Type62",
		Method63CustomerClass:          "Type63",
		Method64Agreed:                 "Type64",
		Method65ADL:                    "Type65",
		Method66Revision:               "Type66",
		Method67CustomerRead:           "Type67",
		Method68Zero:                   "Type68",
		Method71Recalculation:          "Type71",
		Method72RevisedTable:           "Type72",
		Method73RevisedAlgorithm:       "Type73",
		Method74Agreed:                 "Type74",
		Method75ExistingTable:          "Type75",
	}

	// MethodValue maps a method from its string equivalent.
	MethodValue = map[string]Method{ //nolint:gochecknoglobals
		"Type11": Method11Check,
		"Type12": Method12Calculated,
		"Type13": Method13SCADA,
		"Type14": Method14LikeDay,
		"Type15": Method15AverageLikeDay,
		"Type16": Method16Agreed,
		"Type17": Method17Linear,
		"Type18": Method18Alternate,
		"Type19": Method19Zero,
		"Type51": Method51PreviousYear,
		"Type52": Method52PreviousRead,
		"Type53": Method53Revision,
		"Type54": Method54Linear,
		"Type55": Method55Agreed,
		"Type56": Method56PriortoFirstReadAgreed,
		"Type57": Method57CustomerClass,
		"Type58": Method58Zero,
		"Type61": Method61PreviousYear,
		"Type62": Method62PreviousRead,
		"Type63": Method63CustomerClass,
		"Type64": Method64Agreed,
		"Type65": Method65ADL,
		"Type66": Method66Revision,
		"Type67": Method67CustomerRead,
		"Type68": Method68Zero,
		"Type71": Method71Recalculation,
		"Type72": Method72RevisedTable,
		"Type73": Method73RevisedAlgorithm,
		"Type74": Method74Agreed,
		"Type75": Method75ExistingTable,
	}

	// methodDescriptions maps each method to its description.
	methodDescriptions = map[Method]string{ //nolint:gochecknoglobals
		Method11Check:                  "check",
		Method12Calculated:             "calculated",
		Method13SCADA:                  "scada",
		Method14LikeDay:                "like day",
		Method15AverageLikeDay:         "average like day",
		Method16Agreed:                 "agreed",
		Method17Linear:                 "linear",
		Method18Alternate:              "alternate",
		Method19Zero:                   "zero",
		Method51PreviousYear:           "previous year",
		Method52PreviousRead:           "previous read",
		Method53Revision:               "revision",
		Method54Linear:                 "linear",
		Method55Agreed:                 "agreed",
		Method56PriortoFirstReadAgreed: "prior to first read - agreed",
		Method57CustomerClass:          "customer class",
		Method58Zero:                   "zero",
		Method61PreviousYear:           "previous year",
		Method62PreviousRead:           "previous read",
		Method63CustomerClass:          "customer class",
		Method64Agreed:                 "agreed",
		Method65ADL:                    "adl",
		Method66Revision:               "revision",
		Method67CustomerRead:           "customer read",
		Method68Zero:                   "zero",
		Method71Recalculation:          "recalculation",
		Method72RevisedTable:           "revised table",
		Method73RevisedAlgorithm:       "revised algorithm",
		Method74Agreed:                 "agreed",
		Method75ExistingTable:          "existing table",
	}

	// methodInstallationTypes maps each method to the installation types.
	methodInstallationTypes = map[Method][]Install{ //nolint:gochecknoglobals
		Method11Check:                  []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method12Calculated:             []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method13SCADA:                  []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method14LikeDay:                []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method15AverageLikeDay:         []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method16Agreed:                 []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method17Linear:                 []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method18Alternate:              []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method19Zero:                   []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method51PreviousYear:           []Install{InstallMRIM},
		Method52PreviousRead:           []Install{InstallMRIM},
		Method53Revision:               []Install{InstallMRIM},
		Method54Linear:                 []Install{InstallMRIM},
		Method55Agreed:                 []Install{InstallMRIM},
		Method56PriortoFirstReadAgreed: []Install{InstallMRIM},
		Method57CustomerClass:          []Install{InstallMRIM},
		Method58Zero:                   []Install{InstallMRIM},
		Method61PreviousYear:           []Install{InstallBasic},
		Method62PreviousRead:           []Install{InstallBasic},
		Method63CustomerClass:          []Install{InstallBasic},
		Method64Agreed:                 []Install{InstallBasic},
		Method65ADL:                    []Install{InstallBasic},
		Method66Revision:               []Install{InstallBasic},
		Method67CustomerRead:           []Install{InstallBasic},
		Method68Zero:                   []Install{InstallBasic},
		Method71Recalculation:          []Install{InstallUnmetered},
		Method72RevisedTable:           []Install{InstallUnmetered},
		Method73RevisedAlgorithm:       []Install{InstallUnmetered},
		Method74Agreed:                 []Install{InstallUnmetered},
		Method75ExistingTable:          []Install{InstallUnmetered},
	}

	methodMethodTypes = map[Method][]MethodType{ //nolint:gochecknoglobals
		Method11Check:                  []MethodType{MethodTypeSubstituted},
		Method12Calculated:             []MethodType{MethodTypeSubstituted},
		Method13SCADA:                  []MethodType{MethodTypeSubstituted},
		Method14LikeDay:                []MethodType{MethodTypeSubstituted},
		Method15AverageLikeDay:         []MethodType{MethodTypeSubstituted},
		Method16Agreed:                 []MethodType{MethodTypeSubstituted},
		Method17Linear:                 []MethodType{MethodTypeSubstituted},
		Method18Alternate:              []MethodType{MethodTypeSubstituted},
		Method19Zero:                   []MethodType{MethodTypeSubstituted},
		Method51PreviousYear:           []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method52PreviousRead:           []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method53Revision:               []MethodType{MethodTypeSubstituted},
		Method54Linear:                 []MethodType{MethodTypeSubstituted},
		Method55Agreed:                 []MethodType{MethodTypeSubstituted},
		Method56PriortoFirstReadAgreed: []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method57CustomerClass:          []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method58Zero:                   []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method61PreviousYear:           []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method62PreviousRead:           []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method63CustomerClass:          []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method64Agreed:                 []MethodType{MethodTypeSubstituted},
		Method65ADL:                    []MethodType{MethodTypeEstimated},
		Method66Revision:               []MethodType{MethodTypeSubstituted},
		Method67CustomerRead:           []MethodType{MethodTypeSubstituted},
		Method68Zero:                   []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
		Method71Recalculation:          []MethodType{MethodTypeSubstituted},
		Method72RevisedTable:           []MethodType{MethodTypeSubstituted},
		Method73RevisedAlgorithm:       []MethodType{MethodTypeSubstituted},
		Method74Agreed:                 []MethodType{MethodTypeSubstituted},
		Method75ExistingTable:          []MethodType{MethodTypeEstimated},
	}

	// methodTypes lists all valid method types.
	methodTypes = []MethodType{ //nolint:gochecknoglobals
		MethodTypeUndefined,
		MethodTypeEstimated,
		MethodTypeSubstituted,
	}

	// MethodTypeName maps method types to their string version.
	MethodTypeName = map[MethodType]string{ //nolint:gochecknoglobals
		MethodTypeEstimated:   "EST",
		MethodTypeSubstituted: "SUB",
	}

	// MethodTypeName maps method types from their string version.
	MethodTypeValue = map[string]MethodType{ //nolint:gochecknoglobals
		"EST": MethodTypeEstimated,
		"SUB": MethodTypeSubstituted,
	}

	// methodTypeDescriptions maps method types from their descriptions.
	methodTypeDescriptions = map[MethodType]string{ //nolint:gochecknoglobals
		MethodTypeEstimated:   "estimated",
		MethodTypeSubstituted: "substituted",
	}
)

// A Method represents the value of the method flag section of a QualityMethod field
// of a NEM12 interval value.
type Method int

// NewMethodFlag returns a new method flag if valid, and an error if not.
func NewMethodFlag(s string) (Method, error) {
	m, ok := MethodValue[s]

	if !ok {
		return m, ErrMethodInvalid
	}

	return m, nil
}

// Validate returns an error if the method flag is invalid.
func (m Method) Validate() error {
	if _, ok := MethodName[m]; !ok {
		return ErrMethodInvalid
	}

	return nil
}

// Identifier to meet the interface specification for a Flag.
func (m Method) Identifier() string {
	return string(m)
}

// String returns a text representation of the Method.
func (m Method) String() string {
	s, _ := m.Description()

	return fmt.Sprintf("\"%s: %s\"", m.Identifier(), s)
}

// GoString returns a text representation of the Method to satisfy the GoStringer
// interface.
func (m Method) GoString() string {
	return fmt.Sprintf("Method(%d)", m)
}

// Description returns the description of a method flag, along with an error if it is an unknown value.
func (m Method) Description() (string, error) {
	desc, ok := methodDescriptions[m]
	if !ok {
		return fmt.Sprintf("%!Method(%d)", m), fmt.Errorf("method description '%d': ", m, ErrReasonCodeInvalid)
	}

	return desc, nil
}

// A MethodType represents the value of the method flag section of a QualityMethodType field
// of a NEM12 interval value.
type MethodType int

// NewMethodType returns a new method flag if valid, and an error if not.
func NewMethodType(s string) (MethodType, error) {
	m, ok := MethodTypeValue[s]

	if !ok {
		return m, ErrMethodTypeInvalid
	}

	return m, nil
}

// Validate returns an error if the method flag is invalid.
func (m MethodType) Validate() error {
	if _, ok := MethodTypeName[m]; !ok {
		return ErrMethodTypeInvalid
	}

	return nil
}

// Identifier to meet the interface specification for a Flag.
func (m MethodType) Identifier() string {
	return string(m)
}

// String returns a text representation of the MethodType.
func (m MethodType) String() string {
	s, _ := m.Description()

	return fmt.Sprintf("\"%s: %s\"", m.Identifier(), s)
}

// GoString returns a text representation of the MethodType to satisfy the GoStringer
// interface.
func (m MethodType) GoString() string {
	return fmt.Sprintf("MethodType(%d)", m)
}

// Description returns the description of a method flag, along with an error if it is an unknown value.
func (m MethodType) Description() (string, error) {
	desc, ok := methodTypeDescriptions[m]
	if !ok {
		return fmt.Sprintf("%!MethodType(%d)", m), fmt.Errorf("method type description '%d': ", m, ErrReasonCodeInvalid)
	}

	return desc, nil
}
