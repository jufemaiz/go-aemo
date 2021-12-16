package nmi

import (
	"fmt"
	"strings"
)

const (
	// EnergyUndefined is undefined.
	EnergyUndefined Energy = iota
	// EnergyElectricity is electricity.
	EnergyElectricity
	// EnergyGas is gas.
	EnergyGas
)

var (
	energies = []Energy{ //nolint:gochecknoglobals
		EnergyUndefined,
		EnergyElectricity,
		EnergyGas,
	}

	// EnergyName maps Energy to strings.
	EnergyName = map[Energy]string{ //nolint:gochecknoglobals
		EnergyUndefined:   "UNDEFINED",
		EnergyElectricity: "ELECTRICITY",
		EnergyGas:         "GAS",
	}

	// EnergyValue maps strings to Energy.
	EnergyValue = map[string]Energy{ //nolint:gochecknoglobals
		"UNDEFINED":   EnergyUndefined,
		"ELECTRICITY": EnergyElectricity,
		"GAS":         EnergyGas,
	}
)

// Energy represents the type of energy for a nmi.
type Energy int32

// NewEnergy returns an energy based on the string provided.
func NewEnergy(s string) (Energy, error) {
	e, ok := EnergyValue[strings.ToUpper(s)]
	if !ok {
		return EnergyUndefined, ErrEnergyInvalid
	}

	return e, nil
}

// GoString satisfies the GoString interface.
func (e Energy) GoString() string {
	return fmt.Sprintf("%q", e.String())
}

// String satisfies the stringer interface.
func (e Energy) String() string {
	s, ok := EnergyName[e]
	if !ok {
		return EnergyName[EnergyUndefined]
	}

	return s
}

// Energies provides a slice of valid energies.
func Energies() []Energy {
	return energies
}
