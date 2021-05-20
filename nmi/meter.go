package nmi

import (
	"fmt"
	"time"
)

// Meters a collection of meters, using a map, with the identifier.
type Meters map[string]*Meter

// Meter for a NMI.
type Meter struct {
	Nmi                string           `json:"nmi"`
	Identifier         string           `json:"idenifier"`
	Registers          []*MeterRegister `json:"registers"`
	SerialNumber       *string          `json:"serialNumber"`
	FromDateTime       *time.Time       `json:"fromDateTime"`
	ToDateTime         *time.Time       `json:"toDateTime"`
	LastTestDate       *time.Time       `json:"lastTestDate"`
	AdditionalSiteInfo *string          `json:"additionalSiteInformation"`
	// Location                  string
	// Point                     int
	// Status                    Status
}

// GoString meets the gostring interface.
func (m *Meter) GoString() string {
	if m == nil {
		return nilstr
	}

	str := fmt.Sprintf(
		"Meter{Nmi: \"%s\", Identifier: \"%s\", Registers: %#v",
		m.Nmi, m.Identifier, m.Registers,
	)

	if m.SerialNumber != nil {
		str += fmt.Sprintf(", SerialNumber: \"%s\"", *m.SerialNumber)
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
