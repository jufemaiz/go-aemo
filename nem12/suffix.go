package nem12

import "fmt"

const (
	// SuffixUndefined if suffix is undefined.
	SuffixUndefined SuffixType = iota
	// SuffixImportWattHourAverage is for SuffixImportWattHourAverage.
	SuffixImportWattHourAverage
	// SuffixImportWattHourMaster is for SuffixImportWattHourMaster.
	SuffixImportWattHourMaster
	// SuffixImportWattHourCheck is for SuffixImportWattHourCheck.
	SuffixImportWattHourCheck
	// SuffixExportWattHourAverage is for SuffixExportWattHourAverage.
	SuffixExportWattHourAverage
	// SuffixExportWattHourMaster is for SuffixExportWattHourMaster.
	SuffixExportWattHourMaster
	// SuffixExportWattHourCheck is for SuffixExportWattHourCheck.
	SuffixExportWattHourCheck
	// SuffixWattHourNet is for SuffixWattHourNet.
	SuffixWattHourNet
	// SuffixImportVoltAmpReactiveHourAverage is for SuffixImportVoltAmpReactiveHourAverage.
	SuffixImportVoltAmpReactiveHourAverage
	// SuffixImportVoltAmpReactiveHourMaster is for SuffixImportVoltAmpReactiveHourMaster.
	SuffixImportVoltAmpReactiveHourMaster
	// SuffixImportVoltAmpReactiveHourCheck is for SuffixImportVoltAmpReactiveHourCheck.
	SuffixImportVoltAmpReactiveHourCheck
	// SuffixExportVoltAmpReactiveHourAverage is for SuffixExportVoltAmpReactiveHourAverage.
	SuffixExportVoltAmpReactiveHourAverage
	// SuffixExportVoltAmpReactiveHourMaster is for SuffixExportVoltAmpReactiveHourMaster.
	SuffixExportVoltAmpReactiveHourMaster
	// SuffixExportVoltAmpReactiveHourCheck is for SuffixExportVoltAmpReactiveHourCheck.
	SuffixExportVoltAmpReactiveHourCheck
	// SuffixVoltAmpReactiveHourNet is for SuffixVoltAmpReactiveHourNet.
	SuffixVoltAmpReactiveHourNet
	// SuffixVoltAmpHourAverage is for SuffixVoltAmpHourAverage.
	SuffixVoltAmpHourAverage
	// SuffixVoltAmpHourMaster is for SuffixVoltAmpHourMaster.
	SuffixVoltAmpHourMaster
	// SuffixVoltAmpHourCheck is for SuffixVoltAmpHourCheck.
	SuffixVoltAmpHourCheck
	// SuffixPowerFactorMaster is for SuffixPowerFactorMaster.
	SuffixPowerFactorMaster
	// SuffixQMeteringMaster is for SuffixQMeteringMaster.
	SuffixQMeteringMaster
	// SuffixQMeteringCheck is for SuffixQMeteringCheck.
	SuffixQMeteringCheck
	// SuffixParMeteringMaster is for SuffixParMeteringMaster.
	SuffixParMeteringMaster
	// SuffixParMeteringCheck is for SuffixParMeteringCheck.
	SuffixParMeteringCheck
	// SuffixVoltsOrAmpsMaster is for SuffixVoltsOrAmpsMaster.
	SuffixVoltsOrAmpsMaster
	// SuffixVoltsOrAmpsCheck is for SuffixVoltsOrAmpsCheck.
	SuffixVoltsOrAmpsCheck
)

var (
	suffixes = []SuffixType{ //nolint:gochecknoglobals
		SuffixImportWattHourAverage,
		SuffixImportWattHourMaster,
		SuffixImportWattHourCheck,
		SuffixExportWattHourAverage,
		SuffixExportWattHourMaster,
		SuffixExportWattHourCheck,
		SuffixWattHourNet,
		SuffixImportVoltAmpReactiveHourAverage,
		SuffixImportVoltAmpReactiveHourMaster,
		SuffixImportVoltAmpReactiveHourCheck,
		SuffixExportVoltAmpReactiveHourAverage,
		SuffixExportVoltAmpReactiveHourMaster,
		SuffixExportVoltAmpReactiveHourCheck,
		SuffixVoltAmpReactiveHourNet,
		SuffixVoltAmpHourAverage,
		SuffixVoltAmpHourMaster,
		SuffixVoltAmpHourCheck,
		SuffixPowerFactorMaster,
		SuffixQMeteringMaster,
		SuffixQMeteringCheck,
		SuffixParMeteringMaster,
		SuffixParMeteringCheck,
		SuffixVoltsOrAmpsMaster,
		SuffixVoltsOrAmpsCheck,
	}

	// SuffixValue maps the string value to the suffix.
	SuffixValue = map[string]SuffixType{ //nolint:gochecknoglobals
		"A": SuffixImportWattHourAverage,
		"B": SuffixImportWattHourMaster,
		"C": SuffixImportWattHourCheck,
		"D": SuffixExportWattHourAverage,
		"E": SuffixExportWattHourMaster,
		"F": SuffixExportWattHourCheck,
		"N": SuffixWattHourNet,
		"J": SuffixImportVoltAmpReactiveHourAverage,
		"K": SuffixImportVoltAmpReactiveHourMaster,
		"L": SuffixImportVoltAmpReactiveHourCheck,
		"P": SuffixExportVoltAmpReactiveHourAverage,
		"Q": SuffixExportVoltAmpReactiveHourMaster,
		"R": SuffixExportVoltAmpReactiveHourCheck,
		"X": SuffixVoltAmpReactiveHourNet,
		"S": SuffixVoltAmpHourAverage,
		"T": SuffixVoltAmpHourMaster,
		"U": SuffixVoltAmpHourCheck,
		"G": SuffixPowerFactorMaster,
		"H": SuffixQMeteringMaster,
		"Y": SuffixQMeteringCheck,
		"M": SuffixParMeteringMaster,
		"W": SuffixParMeteringCheck,
		"V": SuffixVoltsOrAmpsMaster,
		"Z": SuffixVoltsOrAmpsCheck,
	}

	// SuffixName maps the suffix to its string value.
	SuffixName = map[SuffixType]string{ //nolint:gochecknoglobals
		SuffixImportWattHourAverage:            "A",
		SuffixImportWattHourMaster:             "B",
		SuffixImportWattHourCheck:              "C",
		SuffixExportWattHourAverage:            "D",
		SuffixExportWattHourMaster:             "E",
		SuffixExportWattHourCheck:              "F",
		SuffixWattHourNet:                      "N",
		SuffixImportVoltAmpReactiveHourAverage: "J",
		SuffixImportVoltAmpReactiveHourMaster:  "K",
		SuffixImportVoltAmpReactiveHourCheck:   "L",
		SuffixExportVoltAmpReactiveHourAverage: "P",
		SuffixExportVoltAmpReactiveHourMaster:  "Q",
		SuffixExportVoltAmpReactiveHourCheck:   "R",
		SuffixVoltAmpReactiveHourNet:           "X",
		SuffixVoltAmpHourAverage:               "S",
		SuffixVoltAmpHourMaster:                "T",
		SuffixVoltAmpHourCheck:                 "U",
		SuffixPowerFactorMaster:                "G",
		SuffixQMeteringMaster:                  "H",
		SuffixQMeteringCheck:                   "Y",
		SuffixParMeteringMaster:                "M",
		SuffixParMeteringCheck:                 "W",
		SuffixVoltsOrAmpsMaster:                "V",
		SuffixVoltsOrAmpsCheck:                 "Z",
	}

	suffixDescriptions = map[SuffixType]string{ //nolint:gochecknoglobals
		SuffixImportWattHourAverage:            "import",
		SuffixImportWattHourMaster:             "import",
		SuffixImportWattHourCheck:              "import",
		SuffixExportWattHourAverage:            "export",
		SuffixExportWattHourMaster:             "export",
		SuffixExportWattHourCheck:              "export",
		SuffixWattHourNet:                      "net",
		SuffixImportVoltAmpReactiveHourAverage: "import",
		SuffixImportVoltAmpReactiveHourMaster:  "import",
		SuffixImportVoltAmpReactiveHourCheck:   "import",
		SuffixExportVoltAmpReactiveHourAverage: "export",
		SuffixExportVoltAmpReactiveHourMaster:  "export",
		SuffixExportVoltAmpReactiveHourCheck:   "export",
		SuffixVoltAmpReactiveHourNet:           "net",
		// SuffixVoltAmpHourAverage:               "",
		// SuffixVoltAmpHourMaster:                "",
		// SuffixVoltAmpHourCheck:                 "",
		SuffixPowerFactorMaster: "power factor",
		SuffixQMeteringMaster:   "q metering",
		SuffixQMeteringCheck:    "q metering",
		SuffixParMeteringMaster: "par metering",
		SuffixParMeteringCheck:  "par metering",
		SuffixVoltsOrAmpsMaster: "volts or amps",
		SuffixVoltsOrAmpsCheck:  "volts or amps",
	}

	suffixStreams = map[SuffixType]string{ //nolint:gochecknoglobals
		SuffixImportWattHourAverage:            "average",
		SuffixImportWattHourMaster:             "master",
		SuffixImportWattHourCheck:              "check",
		SuffixExportWattHourAverage:            "average",
		SuffixExportWattHourMaster:             "master",
		SuffixExportWattHourCheck:              "check",
		SuffixWattHourNet:                      "net",
		SuffixImportVoltAmpReactiveHourAverage: "average",
		SuffixImportVoltAmpReactiveHourMaster:  "master",
		SuffixImportVoltAmpReactiveHourCheck:   "check",
		SuffixExportVoltAmpReactiveHourAverage: "average",
		SuffixExportVoltAmpReactiveHourMaster:  "master",
		SuffixExportVoltAmpReactiveHourCheck:   "check",
		SuffixVoltAmpReactiveHourNet:           "net",
		SuffixVoltAmpHourAverage:               "average",
		SuffixVoltAmpHourMaster:                "master",
		SuffixVoltAmpHourCheck:                 "check",
		SuffixPowerFactorMaster:                "master",
		SuffixQMeteringMaster:                  "master",
		SuffixQMeteringCheck:                   "check",
		SuffixParMeteringMaster:                "master",
		SuffixParMeteringCheck:                 "check",
		SuffixVoltsOrAmpsMaster:                "master",
	}

	suffixUnits = map[SuffixType]UnitOfMeasure{ //nolint:gochecknoglobals
		SuffixImportWattHourAverage:            UnitWattHour,
		SuffixImportWattHourMaster:             UnitWattHour,
		SuffixImportWattHourCheck:              UnitWattHour,
		SuffixExportWattHourAverage:            UnitWattHour,
		SuffixExportWattHourMaster:             UnitWattHour,
		SuffixExportWattHourCheck:              UnitWattHour,
		SuffixWattHourNet:                      UnitWattHour,
		SuffixImportVoltAmpReactiveHourAverage: UnitVoltAmpereReactiveHour,
		SuffixImportVoltAmpReactiveHourMaster:  UnitVoltAmpereReactiveHour,
		SuffixImportVoltAmpReactiveHourCheck:   UnitVoltAmpereReactiveHour,
		SuffixExportVoltAmpReactiveHourAverage: UnitVoltAmpereReactiveHour,
		SuffixExportVoltAmpReactiveHourMaster:  UnitVoltAmpereReactiveHour,
		SuffixExportVoltAmpReactiveHourCheck:   UnitVoltAmpereReactiveHour,
		SuffixVoltAmpReactiveHourNet:           UnitVoltAmpereReactiveHour,
		SuffixVoltAmpHourAverage:               UnitVoltAmpereHour,
		SuffixVoltAmpHourMaster:                UnitVoltAmpereHour,
		SuffixVoltAmpHourCheck:                 UnitVoltAmpereHour,
		SuffixPowerFactorMaster:                UnitPowerFactor,
		// SuffixQMeteringMaster:                  UnitUndefined,
		// SuffixQMeteringCheck:                   UnitUndefined,
		// SuffixParMeteringMaster:                UnitUndefined,
		// SuffixParMeteringCheck:                 UnitUndefined,
		SuffixVoltsOrAmpsMaster: UnitVolt,
		SuffixVoltsOrAmpsCheck:  UnitVolt,
	}
)

// SuffixType for suffix types.
type SuffixType int

// SuffixTypes returns all suffixes.
func SuffixTypes() []SuffixType {
	return suffixes
}

// NewSuffixType to create a new suffix from the string value.
func NewSuffixType(s string) (SuffixType, error) {
	if s == "" {
		return SuffixUndefined, ErrSuffixTypeNil
	}

	st, ok := SuffixValue[s]
	if !ok {
		return SuffixUndefined, ErrSuffixTypeInvalid
	}

	return st, nil
}

// Identifier returns the identifier.
func (s SuffixType) Identifier() string {
	id, ok := SuffixName[s]
	if !ok {
		return fmt.Sprintf("SuffixType(%d)", s)
	}

	return id
}

// String returns a text representation of the reason.
func (s SuffixType) String() string {
	return fmt.Sprintf("%q", s.Identifier())
}

// GoString returns a text representation of the reason to satisfy the GoStringer
// interface.
func (s SuffixType) GoString() string {
	if _, ok := SuffixName[s]; !ok {
		return fmt.Sprintf("%%!SuffixType(%d)", s)
	}

	return fmt.Sprintf("SuffixType(%d)", s)
}

// MarshalJSON marshals for JSON.
func (s *SuffixType) MarshalJSON() ([]byte, error) {
	id, ok := SuffixName[*s]
	if !ok {
		return []byte(fmt.Sprintf("\"%d\"", *s)), nil
	}

	return []byte(fmt.Sprintf("%q", id)), nil
}

// UnmarshalJSON unmarshals the JSON data to a suffix type.
func (s *SuffixType) UnmarshalJSON(data []byte) error {
	v, ok := SuffixValue[string(data)]
	if !ok {
		return ErrSuffixTypeInvalid
	}

	*s = v

	return nil
}

// Description returns the description for a suffix type.
func (s SuffixType) Description() string {
	str, ok := suffixDescriptions[s]
	if !ok {
		return fmt.Sprintf("SuffixType(%d)", s)
	}

	return str
}

// Stream returns the stream for a suffix type.
func (s SuffixType) Stream() string {
	str, ok := suffixStreams[s]
	if !ok {
		return fmt.Sprintf("SuffixType(%d)", s)
	}

	return str
}

// Unit returns the base unit for a suffix type.
func (s SuffixType) Unit() UnitOfMeasure {
	unit, ok := suffixUnits[s]
	if !ok {
		return UnitUndefined
	}

	return unit
}

// Suffix is a datastream suffix, made up first of a single character flag for
// the type of datastream, and a second character flag for a meter.
type Suffix struct {
	Type  SuffixType `json:"suffixType,omitempty"` //nolint:tagliatelle
	Meter rune       `json:"meter,omitempty"`
}

// NewSuffix returns a new suffix, with errors raised if invalid.
func NewSuffix(s string) (Suffix, error) {
	sfx := Suffix{}

	if s == "" {
		return sfx, ErrSuffixNil
	}

	runes := []rune(s)

	if len(runes) == 0 {
		return sfx, ErrSuffixNil
	}

	if pairCount := 2; len(runes) != pairCount {
		return sfx, fmt.Errorf("suffix '%s': %w", s, ErrSuffixLengthInvalid)
	}

	t, err := NewSuffixType(string(runes[0]))
	if err != nil {
		return sfx, fmt.Errorf("suffix '%s': %w", s, err)
	}

	sfx.Type = t

	if err := ValidateSuffixMeter(string(runes[1])); err != nil {
		return sfx, fmt.Errorf("suffix '%s': %w", s, err)
	}

	sfx.Meter = runes[1]

	return sfx, nil
}

// ValidateSuffixMeter checks the validity of a NMI suffix's meter character.
func ValidateSuffixMeter(s string) error {
	valid := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
		"A": true,
		"B": true,
		"C": true,
		"D": true,
		"E": true,
		"F": true,
		"G": true,
		"H": true,
		"J": true,
		"K": true,
		"L": true,
		"M": true,
		"N": true,
		"P": true,
		"Q": true,
		"R": true,
		"S": true,
		"T": true,
		"U": true,
		"V": true,
		"W": true,
		"X": true,
		"Y": true,
		"Z": true,
	}

	if _, ok := valid[s]; !ok {
		return fmt.Errorf("meter '%s': %w", s, ErrSuffixMeterInvalid)
	}

	return nil
}
