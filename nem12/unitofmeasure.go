package nem12

import "fmt"

const (
	// UnitUndefined for undefined units.
	UnitUndefined UnitOfMeasure = iota
	// UnitMegawattHour for the unit of megawatt hours.
	UnitMegawattHour
	// UnitKilowattHour for the unit of kilowatt hours.
	UnitKilowattHour
	// UnitWattHour for the unit of watt hours.
	UnitWattHour
	// UnitMegawatt for the unit of megawatts.
	UnitMegawatt
	// UnitKilowatt for the unit of kilowatts.
	UnitKilowatt
	// UnitWatt for the unit of watts.
	UnitWatt
	// UnitMegavoltAmpereReactiveHour for the unit of megavolt ampere reactive hours.
	UnitMegavoltAmpereReactiveHour
	// UnitKilovoltAmpereReactiveHour for the unit of kilovolt ampere reactive hours.
	UnitKilovoltAmpereReactiveHour
	// UnitVoltAmpereReactiveHour for the unit of volt ampere reactive hours.
	UnitVoltAmpereReactiveHour
	// UnitMegavoltAmpereReactive for the unit of megavolt ampere reactives.
	UnitMegavoltAmpereReactive
	// UnitKilovoltAmpereReactive for the unit of kilovolt ampere reactives.
	UnitKilovoltAmpereReactive
	// UnitVoltAmpereReactive for the unit of volt ampere reactives.
	UnitVoltAmpereReactive
	// UnitMegavoltAmpereHour for the unit of megavolt ampere hours.
	UnitMegavoltAmpereHour
	// UnitKilovoltAmpereHour for the unit of kilovolt ampere hours.
	UnitKilovoltAmpereHour
	// UnitVoltAmpereHour for the unit of volt ampere hours.
	UnitVoltAmpereHour
	// UnitMegavoltAmpere for the unit of megavolt amperes.
	UnitMegavoltAmpere
	// UnitKilovoltAmpere for the unit of kilovolt amperes.
	UnitKilovoltAmpere
	// UnitVoltAmpere for the unit of volt amperes.
	UnitVoltAmpere
	// UnitKilovolt for the unit of kilovolts.
	UnitKilovolt
	// UnitVolt for the unit of volts.
	UnitVolt
	// UnitKiloampere for the unit of kiloamperes.
	UnitKiloampere
	// UnitAmpere for the unit of amperes.
	UnitAmpere
	// UnitPowerFactor for the unit of power factors.
	UnitPowerFactor
)

var (
	units = []UnitOfMeasure{
		UnitUndefined,
		UnitMegawattHour,
		UnitKilowattHour,
		UnitWattHour,
		UnitMegawatt,
		UnitKilowatt,
		UnitWatt,
		UnitMegavoltAmpereReactiveHour,
		UnitKilovoltAmpereReactiveHour,
		UnitVoltAmpereReactiveHour,
		UnitMegavoltAmpereReactive,
		UnitKilovoltAmpereReactive,
		UnitVoltAmpereReactive,
		UnitMegavoltAmpereHour,
		UnitKilovoltAmpereHour,
		UnitVoltAmpereHour,
		UnitMegavoltAmpere,
		UnitKilovoltAmpere,
		UnitVoltAmpere,
		UnitKilovolt,
		UnitVolt,
		UnitKiloampere,
		UnitAmpere,
		UnitPowerFactor,
	}

	// UnitName maps a unit of measure to a name.
	UnitName = map[UnitOfMeasure]string{ //nolint:gochecknoglobals
		UnitMegawattHour:               "MWH",
		UnitKilowattHour:               "KWH",
		UnitWattHour:                   "WH",
		UnitMegawatt:                   "MW",
		UnitKilowatt:                   "KW",
		UnitWatt:                       "W",
		UnitMegavoltAmpereReactiveHour: "MVARH",
		UnitKilovoltAmpereReactiveHour: "KVARH",
		UnitVoltAmpereReactiveHour:     "VARH",
		UnitMegavoltAmpereReactive:     "MVAR",
		UnitKilovoltAmpereReactive:     "KVAR",
		UnitVoltAmpereReactive:         "VAR",
		UnitMegavoltAmpereHour:         "MVAH",
		UnitKilovoltAmpereHour:         "KVAH",
		UnitVoltAmpereHour:             "VAH",
		UnitMegavoltAmpere:             "MVA",
		UnitKilovoltAmpere:             "KVA",
		UnitVoltAmpere:                 "VA",
		UnitKilovolt:                   "KV",
		UnitVolt:                       "V",
		UnitKiloampere:                 "KA",
		UnitAmpere:                     "A",
		UnitPowerFactor:                "PF",
	}

	// UnitValue maps a unit of measure from a name.
	UnitValue = map[string]UnitOfMeasure{ //nolint:gochecknoglobals
		"MWH":   UnitMegawattHour,
		"KWH":   UnitKilowattHour,
		"WH":    UnitWattHour,
		"MW":    UnitMegawatt,
		"KW":    UnitKilowatt,
		"W":     UnitWatt,
		"MVARH": UnitMegavoltAmpereReactiveHour,
		"KVARH": UnitKilovoltAmpereReactiveHour,
		"VARH":  UnitVoltAmpereReactiveHour,
		"MVAR":  UnitMegavoltAmpereReactive,
		"KVAR":  UnitKilovoltAmpereReactive,
		"VAR":   UnitVoltAmpereReactive,
		"MVAH":  UnitMegavoltAmpereHour,
		"KVAH":  UnitKilovoltAmpereHour,
		"VAH":   UnitVoltAmpereHour,
		"MVA":   UnitMegavoltAmpere,
		"KVA":   UnitKilovoltAmpere,
		"VA":    UnitVoltAmpere,
		"KV":    UnitKilovolt,
		"V":     UnitVolt,
		"KA":    UnitKiloampere,
		"A":     UnitAmpere,
		"PF":    UnitPowerFactor,
	}

	unitDescriptions = map[UnitOfMeasure]string{ //nolint:gochecknoglobals
		UnitMegawattHour:               "megawatt hour",
		UnitKilowattHour:               "kilowatt hour",
		UnitWattHour:                   "watt hour",
		UnitMegawatt:                   "megawatt",
		UnitKilowatt:                   "kilowatt",
		UnitWatt:                       "watt",
		UnitMegavoltAmpereReactiveHour: "megavolt ampere reactive hour",
		UnitKilovoltAmpereReactiveHour: "kilovolt ampere reactive hour",
		UnitVoltAmpereReactiveHour:     "volt ampere reactive hour",
		UnitMegavoltAmpereReactive:     "megavolt ampere reactive",
		UnitKilovoltAmpereReactive:     "kilovolt ampere reactive",
		UnitVoltAmpereReactive:         "volt ampere reactive",
		UnitMegavoltAmpereHour:         "megavolt ampere hour",
		UnitKilovoltAmpereHour:         "kilovolt ampere hour",
		UnitVoltAmpereHour:             "volt ampere hour",
		UnitMegavoltAmpere:             "megavolt ampere",
		UnitKilovoltAmpere:             "kilovolt ampere",
		UnitVoltAmpere:                 "volt ampere",
		UnitKilovolt:                   "kilovolt",
		UnitVolt:                       "volt",
		UnitKiloampere:                 "kiloampere",
		UnitAmpere:                     "ampere",
		UnitPowerFactor:                "power factor",
	}

	unitMultipliers = map[UnitOfMeasure]float64{
		UnitMegawattHour:               1.0e6,
		UnitKilowattHour:               1.0e3,
		UnitWattHour:                   1.0,
		UnitMegawatt:                   1.0e6,
		UnitKilowatt:                   1.0e3,
		UnitWatt:                       1.0,
		UnitMegavoltAmpereReactiveHour: 1.0e6,
		UnitKilovoltAmpereReactiveHour: 1.0e3,
		UnitVoltAmpereReactiveHour:     1.0,
		UnitMegavoltAmpereReactive:     1.0e6,
		UnitKilovoltAmpereReactive:     1.0e3,
		UnitVoltAmpereReactive:         1.0,
		UnitMegavoltAmpereHour:         1.0e6,
		UnitKilovoltAmpereHour:         1.0e3,
		UnitVoltAmpereHour:             1.0,
		UnitMegavoltAmpere:             1.0e6,
		UnitKilovoltAmpere:             1.0e3,
		UnitVoltAmpere:                 1.0,
		UnitKilovolt:                   1.0e3,
		UnitVolt:                       1.0,
		UnitKiloampere:                 1.0e3,
		UnitAmpere:                     1.0,
		UnitPowerFactor:                1.0,
	}

	unitNames = map[UnitOfMeasure]string{
		UnitMegawattHour:               "MWh",
		UnitKilowattHour:               "kWh",
		UnitWattHour:                   "Wh",
		UnitMegawatt:                   "MW",
		UnitKilowatt:                   "kW",
		UnitWatt:                       "W",
		UnitMegavoltAmpereReactiveHour: "MVArh",
		UnitKilovoltAmpereReactiveHour: "kVArh",
		UnitVoltAmpereReactiveHour:     "VArh",
		UnitMegavoltAmpereReactive:     "MVAr",
		UnitKilovoltAmpereReactive:     "kVAr",
		UnitVoltAmpereReactive:         "VAr",
		UnitMegavoltAmpereHour:         "MVAh",
		UnitKilovoltAmpereHour:         "kVAh",
		UnitVoltAmpereHour:             "VAh",
		UnitMegavoltAmpere:             "MVA",
		UnitKilovoltAmpere:             "kVA",
		UnitVoltAmpere:                 "VA",
		UnitKilovolt:                   "kV",
		UnitVolt:                       "V",
		UnitKiloampere:                 "kA",
		UnitAmpere:                     "A",
		UnitPowerFactor:                "pf",
	}
)

// A UnitOfMeasure represents a unit of measure as specified by the UOM field of a NMIDataDetails record.
type UnitOfMeasure int

// NewUnit returns a new reason, along with errors if not valid.
func NewUnit(s string) (UnitOfMeasure, error) {
	if s == "" {
		return UnitUndefined, ErrUnitOfMeasureEmpty
	}

	u, ok := UnitValue[s]
	if !ok {
		return UnitUndefined, ErrUnitOfMeasureInvalid
	}

	return u, nil
}

// Validate ensures a reason is valid.
func (u UnitOfMeasure) Validate() error {
	_, ok := UnitName[u]
	if !ok {
		return ErrUnitOfMeasureInvalid
	}

	return nil
}

// Identifier returns the identifier.
func (u UnitOfMeasure) Identifier() string {
	id, ok := UnitName[u]
	if !ok {
		return fmt.Sprintf("UnitOfMeasure(%d)", u)
	}

	return id
}

// String returns a text representation of the reason.
func (u UnitOfMeasure) String() string {
	name, err := u.Name()
	if err != nil {
		return name
	}

	desc, err := u.Description()
	if err != nil {
		return fmt.Sprintf("%d: %s %s", u.Identifier(), name, desc)
	}

	return fmt.Sprintf("%d: %s", u.Identifier(), desc)
}

// GoString returns a text representation of the reason to satisfy the GoStringer
// interface.
func (u UnitOfMeasure) GoString() string {
	return fmt.Sprintf("UnitOfMeasure(%d)", u)
}

// Name returns the name of the unit.
func (u UnitOfMeasure) Name() (string, error) {
	s, ok := unitNames[u]
	if !ok {
		return fmt.Sprintf("%!UnitOfMeasure(%d)", u), fmt.Errorf("unit of measure name '%d': ", u, ErrUnitOfMeasureInvalid)
	}

	return s, nil
}

// Description returns the description of a reason code, along with an error if
// it is an unknown value.
func (u UnitOfMeasure) Description() (string, error) {
	s, ok := unitDescriptions[u]
	if !ok {
		return fmt.Sprintf("%!UnitOfMeasure(%d)", u), fmt.Errorf("unit of measure description '%d': ", u, ErrUnitOfMeasureInvalid)
	}

	return s, nil
}

// Multiplier for the unit of measure to SI unit.
func (u UnitOfMeasure) Multiplier() float64 {
	m, ok := unitMultipliers[u]
	if !ok {
		return 1.0
	}

	return m
}
