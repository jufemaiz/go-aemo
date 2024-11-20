package nem12

import (
	"fmt"
)

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
	Method16Agreed = Method(16) // [OBSOLETE]
	// Method17Linear for linear.
	Method17Linear = Method(17)
	// Method18Alternate for alternate.
	Method18Alternate = Method(18)
	// Method19Zero for zero.
	Method19Zero = Method(19)
	// Method20ChurnCorrection for churn correction (like day).
	Method20ChurnCorrection = Method(20)
	// Method21FiveMinuteNoHistoricalData for five-minute no historical data.
	Method21FiveMinuteNoHistoricalData = Method(21)
	// Method22ProspectiveAverageDay for prospective average day.
	Method22ProspectiveAverageDay = Method(22)
	// Method23PreviousYear for use of previous year.
	Method23PreviousYear = Method(23)
	// Method24DataScaling for data  scaling.
	Method24DataScaling = Method(24)
	// Method25ADL for average daily load.
	Method25ADL = Method(25)
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
	// Method59FiveMinuteNoHistoricalData for five-minute no historical data.
	Method59FiveMinuteNoHistoricalData = Method(59)
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
	// Method69LinearExtrapolation for linear extrapolation.
	MetMethod69LinearExtrapolation = Method(69)
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
)

var (
	// methods is a slice of all valid methods.
	methods = []Method{ //nolint:gochecknoglobals
		Method11Check,
		Method12Calculated,
		Method13SCADA,
		Method14LikeDay,
		Method15AverageLikeDay,
		Method16Agreed,
		Method17Linear,
		Method18Alternate,
		Method19Zero,
		Method20ChurnCorrection,
		Method21FiveMinuteNoHistoricalData,
		Method22ProspectiveAverageDay,
		Method23PreviousYear,
		Method24DataScaling,
		Method25ADL,
		Method51PreviousYear,
		Method52PreviousRead,
		Method53Revision,
		Method54Linear,
		Method55Agreed,
		Method56PriortoFirstReadAgreed,
		Method57CustomerClass,
		Method58Zero,
		Method59FiveMinuteNoHistoricalData,
		Method61PreviousYear,
		Method62PreviousRead,
		Method63CustomerClass,
		Method64Agreed,
		Method65ADL,
		Method66Revision,
		Method67CustomerRead,
		Method68Zero,
		MetMethod69LinearExtrapolation,
		Method71Recalculation,
		Method72RevisedTable,
		Method73RevisedAlgorithm,
		Method74Agreed,
		Method75ExistingTable,
	}

	// MethodName maps a method to its string equivalent.
	MethodName = map[Method]string{ //nolint:gochecknoglobals
		Method11Check:                      "11",
		Method12Calculated:                 "12",
		Method13SCADA:                      "13",
		Method14LikeDay:                    "14",
		Method15AverageLikeDay:             "15",
		Method16Agreed:                     "16",
		Method17Linear:                     "17",
		Method18Alternate:                  "18",
		Method19Zero:                       "19",
		Method20ChurnCorrection:            "20",
		Method21FiveMinuteNoHistoricalData: "21",
		Method22ProspectiveAverageDay:      "22",
		Method23PreviousYear:               "23",
		Method24DataScaling:                "24",
		Method25ADL:                        "25",
		Method51PreviousYear:               "51",
		Method52PreviousRead:               "52",
		Method53Revision:                   "53",
		Method54Linear:                     "54",
		Method55Agreed:                     "55",
		Method56PriortoFirstReadAgreed:     "56",
		Method57CustomerClass:              "57",
		Method58Zero:                       "58",
		Method59FiveMinuteNoHistoricalData: "59",
		Method61PreviousYear:               "61",
		Method62PreviousRead:               "62",
		Method63CustomerClass:              "63",
		Method64Agreed:                     "64",
		Method65ADL:                        "65",
		Method66Revision:                   "66",
		Method67CustomerRead:               "67",
		Method68Zero:                       "68",
		MetMethod69LinearExtrapolation:     "69",
		Method71Recalculation:              "71",
		Method72RevisedTable:               "72",
		Method73RevisedAlgorithm:           "73",
		Method74Agreed:                     "74",
		Method75ExistingTable:              "75",
	}

	// MethodValue maps a method from its string equivalent.
	MethodValue = map[string]Method{ //nolint:gochecknoglobals
		"11": Method11Check,
		"12": Method12Calculated,
		"13": Method13SCADA,
		"14": Method14LikeDay,
		"15": Method15AverageLikeDay,
		"16": Method16Agreed,
		"17": Method17Linear,
		"18": Method18Alternate,
		"19": Method19Zero,
		"20": Method20ChurnCorrection,
		"21": Method21FiveMinuteNoHistoricalData,
		"22": Method22ProspectiveAverageDay,
		"23": Method23PreviousYear,
		"24": Method24DataScaling,
		"25": Method25ADL,
		"51": Method51PreviousYear,
		"52": Method52PreviousRead,
		"53": Method53Revision,
		"54": Method54Linear,
		"55": Method55Agreed,
		"56": Method56PriortoFirstReadAgreed,
		"57": Method57CustomerClass,
		"58": Method58Zero,
		"59": Method59FiveMinuteNoHistoricalData,
		"61": Method61PreviousYear,
		"62": Method62PreviousRead,
		"63": Method63CustomerClass,
		"64": Method64Agreed,
		"65": Method65ADL,
		"66": Method66Revision,
		"67": Method67CustomerRead,
		"68": Method68Zero,
		"69": MetMethod69LinearExtrapolation,
		"71": Method71Recalculation,
		"72": Method72RevisedTable,
		"73": Method73RevisedAlgorithm,
		"74": Method74Agreed,
		"75": Method75ExistingTable,
	}

	// methodDescriptions maps each method to its description.
	methodDescriptions = map[Method]string{ //nolint:gochecknoglobals
		Method11Check:                      "check",
		Method12Calculated:                 "calculated",
		Method13SCADA:                      "scada",
		Method14LikeDay:                    "retrospective like day",
		Method15AverageLikeDay:             "retrospective average like day",
		Method16Agreed:                     "[OBSOLETE] agreed",
		Method17Linear:                     "linear",
		Method18Alternate:                  "alternate",
		Method19Zero:                       "zero",
		Method20ChurnCorrection:            "prospective like day",
		Method21FiveMinuteNoHistoricalData: "five-minute no historical data",
		Method22ProspectiveAverageDay:      "prospective average day",
		Method23PreviousYear:               "previous year",
		Method24DataScaling:                "data scaling",
		Method25ADL:                        "adl",
		Method51PreviousYear:               "previous year",
		Method52PreviousRead:               "previous read",
		Method53Revision:                   "revision",
		Method54Linear:                     "linear",
		Method55Agreed:                     "agreed",
		Method56PriortoFirstReadAgreed:     "prior to first read - agreed",
		Method57CustomerClass:              "customer class",
		Method58Zero:                       "zero",
		Method59FiveMinuteNoHistoricalData: "five-minute no historical data",
		Method61PreviousYear:               "previous year",
		Method62PreviousRead:               "previous read",
		Method63CustomerClass:              "customer class",
		Method64Agreed:                     "agreed",
		Method65ADL:                        "adl",
		Method66Revision:                   "revision",
		Method67CustomerRead:               "customer read",
		Method68Zero:                       "zero",
		MetMethod69LinearExtrapolation:     "linear extrapolation",
		Method71Recalculation:              "recalculation",
		Method72RevisedTable:               "revised table",
		Method73RevisedAlgorithm:           "revised algorithm",
		Method74Agreed:                     "agreed",
		Method75ExistingTable:              "existing table",
	}

	// methodInstallationTypes maps each method to the installation types.
	methodInstallationTypes = map[Method][]Install{ //nolint:gochecknoglobals
		Method11Check:                      {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method12Calculated:                 {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method13SCADA:                      {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method14LikeDay:                    {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method15AverageLikeDay:             {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method16Agreed:                     {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method17Linear:                     {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method18Alternate:                  {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method19Zero:                       {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method20ChurnCorrection:            {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method21FiveMinuteNoHistoricalData: {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method22ProspectiveAverageDay:      {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method23PreviousYear:               {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method24DataScaling:                {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method25ADL:                        {InstallComms1, InstallComms2, InstallComms3, InstallComms4},
		Method51PreviousYear:               {InstallMRIM},
		Method52PreviousRead:               {InstallMRIM},
		Method53Revision:                   {InstallMRIM},
		Method54Linear:                     {InstallMRIM},
		Method55Agreed:                     {InstallMRIM},
		Method56PriortoFirstReadAgreed:     {InstallMRIM},
		Method57CustomerClass:              {InstallMRIM},
		Method58Zero:                       {InstallMRIM},
		Method59FiveMinuteNoHistoricalData: {InstallMRIM},
		Method61PreviousYear:               {InstallBasic},
		Method62PreviousRead:               {InstallBasic},
		Method63CustomerClass:              {InstallBasic},
		Method64Agreed:                     {InstallBasic},
		Method65ADL:                        {InstallBasic},
		Method66Revision:                   {InstallBasic},
		Method67CustomerRead:               {InstallBasic},
		Method68Zero:                       {InstallBasic},
		MetMethod69LinearExtrapolation:     {InstallBasic},
		Method71Recalculation:              {InstallUnmetered},
		Method72RevisedTable:               {InstallUnmetered},
		Method73RevisedAlgorithm:           {InstallUnmetered},
		Method74Agreed:                     {InstallUnmetered},
		Method75ExistingTable:              {InstallUnmetered},
	}

	methodMethodTypes = map[Method][]MethodType{ //nolint:gochecknoglobals
		Method11Check:                      {MethodTypeSubstituted},
		Method12Calculated:                 {MethodTypeSubstituted},
		Method13SCADA:                      {MethodTypeSubstituted},
		Method14LikeDay:                    {MethodTypeSubstituted},
		Method15AverageLikeDay:             {MethodTypeSubstituted},
		Method16Agreed:                     {MethodTypeSubstituted},
		Method17Linear:                     {MethodTypeSubstituted},
		Method18Alternate:                  {MethodTypeSubstituted},
		Method19Zero:                       {MethodTypeSubstituted},
		Method20ChurnCorrection:            {MethodTypeSubstituted},
		Method21FiveMinuteNoHistoricalData: {MethodTypeSubstituted},
		Method22ProspectiveAverageDay:      {MethodTypeSubstituted},
		Method23PreviousYear:               {MethodTypeSubstituted},
		Method24DataScaling:                {MethodTypeSubstituted},
		Method25ADL:                        {MethodTypeSubstituted},
		Method51PreviousYear:               {MethodTypeEstimated, MethodTypeSubstituted},
		Method52PreviousRead:               {MethodTypeEstimated, MethodTypeSubstituted},
		Method53Revision:                   {MethodTypeSubstituted},
		Method54Linear:                     {MethodTypeSubstituted},
		Method55Agreed:                     {MethodTypeSubstituted},
		Method56PriortoFirstReadAgreed:     {MethodTypeEstimated, MethodTypeSubstituted},
		Method57CustomerClass:              {MethodTypeEstimated, MethodTypeSubstituted},
		Method58Zero:                       {MethodTypeEstimated, MethodTypeSubstituted},
		Method59FiveMinuteNoHistoricalData: {MethodTypeEstimated, MethodTypeSubstituted},
		Method61PreviousYear:               {MethodTypeEstimated, MethodTypeSubstituted},
		Method62PreviousRead:               {MethodTypeEstimated, MethodTypeSubstituted},
		Method63CustomerClass:              {MethodTypeEstimated, MethodTypeSubstituted},
		Method64Agreed:                     {MethodTypeSubstituted},
		Method65ADL:                        {MethodTypeEstimated},
		Method66Revision:                   {MethodTypeSubstituted},
		Method67CustomerRead:               {MethodTypeSubstituted},
		Method68Zero:                       {MethodTypeEstimated, MethodTypeSubstituted},
		MetMethod69LinearExtrapolation:     {MethodTypeSubstituted},
		Method71Recalculation:              {MethodTypeSubstituted},
		Method72RevisedTable:               {MethodTypeSubstituted},
		Method73RevisedAlgorithm:           {MethodTypeSubstituted},
		Method74Agreed:                     {MethodTypeSubstituted},
		Method75ExistingTable:              {MethodTypeEstimated},
	}
)

// A Method represents the value of the method flag section of a QualityMethod field
// of a NEM12 interval value.
type Method int

// Methods returns all methods.
func Methods() []Method {
	return methods
}

// NewMethodFlag returns a new method flag if valid, and an error if not.
func NewMethodFlag(s string) (Method, error) {
	if s == "" {
		return MethodUndefined, ErrMethodNil
	}

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
	id, ok := MethodName[m]
	if !ok {
		return fmt.Sprintf("Method(%d)", m)
	}

	return id
}

// Description returns the description of a method flag, along with an error if it is an unknown value.
func (m Method) Description() (string, error) {
	desc, ok := methodDescriptions[m]
	if !ok {
		return fmt.Sprintf("%%!Method(%d)", m), fmt.Errorf("method description '%d': %w", m, ErrMethodInvalid)
	}

	return desc, nil
}

// Types returns the installation types for a method flag.
func (m Method) Types() []MethodType {
	mt, ok := methodMethodTypes[m]
	if !ok {
		return nil
	}

	return mt
}

// InstallationTypes returns the installation types for a method flag.
func (m Method) InstallationTypes() []Install {
	it, ok := methodInstallationTypes[m]
	if !ok {
		return nil
	}

	return it
}

// String returns a text representation of the Method.
func (m Method) String() string {
	s, err := m.Description()
	if err != nil {
		return fmt.Sprintf("%q", m.Identifier())
	}

	return fmt.Sprintf("\"%s: %s\"", m.Identifier(), s)
}

// GoString returns a text representation of the Method to satisfy the GoStringer
// interface.
func (m Method) GoString() string {
	return fmt.Sprintf("Method(%d)", m)
}
