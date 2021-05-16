package nmi

import (
	"fmt"
	"strings"
)

const (
	EnergyUndefined Energy = iota
	EnergyElectricity
	EnergyGas
)

var (
	energies = []Energy{
		EnergyUndefined,
		EnergyElectricity,
		EnergyGas,
	}

	Energy_name = map[Energy]string{
		EnergyUndefined:   "UNDEFINED",
		EnergyElectricity: "ELECTRICITY",
		EnergyGas:         "GAS",
	}

	Energy_value = map[string]Energy{
		"UNDEFINED":   EnergyUndefined,
		"ELECTRICITY": EnergyElectricity,
		"GAS":         EnergyGas,
	}
)

// Energy represents the type of energy for a nmi.
type Energy int32

// NewEnergy returns an energy based on the string provided.
func NewEnergy(s string) (Energy, error) {
	e, ok := Energy_value[strings.ToUpper(s)]
	if !ok {
		return EnergyUndefined, ErrEnergyInvalid
	}

	return e, nil
}

// String satisfies the stringer interface.
func (e Energy) GoString() string {
	return fmt.Sprintf("\"%s\"", e.String())
}

// String satisfies the stringer interface.
func (e Energy) String() string {
	s, ok := Energy_name[e]
	if !ok {
		return Energy_name[EnergyUndefined]
	}

	return s
}

// Energies provides a slice of valid energies.
func Energies() []Energy {
	return energies
}
