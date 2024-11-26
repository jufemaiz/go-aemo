package nmi

import (
	"fmt"
	"time"
)

// Meters a collection of meters, using a map, with the identifier.
type Meters map[string]*Meter

// Meter for a NMI.
type Meter struct {
	Nmi                string           `json:"nmi,omitempty"`
	Identifier         string           `json:"identifier,omitempty"`
	Registers          []*MeterRegister `json:"registers,omitempty"`
	SerialNumber       *string          `json:"serialNumber,omitempty"`
	FromDateTime       *time.Time       `json:"fromDateTime,omitempty"`
	ToDateTime         *time.Time       `json:"toDateTime,omitempty"`
	LastTestDate       *time.Time       `json:"lastTestDate,omitempty"`
	AdditionalSiteInfo *string          `json:"additionalSiteInformation,omitempty"` //nolint:tagliatelle
}

// GoString meets the gostring interface.
func (m *Meter) GoString() string {
	if m == nil {
		return nilstr
	}

	str := fmt.Sprintf(
		"Meter{Nmi: %q, Identifier: %q, Registers: %#v",
		m.Nmi, m.Identifier, m.Registers,
	)

	if m.SerialNumber != nil {
		str += fmt.Sprintf(", SerialNumber: %q", *m.SerialNumber)
	}

	str += "}"

	return str
}

// String meets the stringer interface.
func (m *Meter) String() string {
	if m == nil {
		return ""
	}

	return m.Identifier
}
