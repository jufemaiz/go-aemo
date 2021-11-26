package nmi

import "fmt"

// MeterRegister for a Meter for a NMI.
type MeterRegister struct {
	RegisterID        string `json:"registerID,omitempty"`
	MeasurementStream string `json:"measurementStream,omitempty"`
	NetworkTariffCode string `json:"networkTariffCode,omitempty"`
	UnitOfMeasure     string `json:"unitOfMeasure,omitempty"`
	TimeOfDay         string `json:"timeOfDay,omitempty"`
	Multiplier        int    `json:"multiplier,omitempty"`
	DialFormat        string `json:"dialFormat,omitempty"`
	ControlledLoad    bool   `json:"controlledLoad,omitempty"`
	ConsumptionType   string `json:"consumptionType,omitempty"`
	Status            string `json:"status,omitempty"`
}

// GoString meets the gostring interface.
func (mr *MeterRegister) GoString() string {
	if mr == nil {
		return nilstr
	}

	return fmt.Sprintf(
		"MeterRegister{"+
			"RegisterID: \"%s\", MeasurementStream: \"%s\", NetworkTariffCode: \"%s\", "+
			"UnitOfMeasure: \"%s\", TimeOfDay: \"%s\", Multiplier: %d, "+
			"DialFormat: \"%s\", ControlledLoad: %t, ConsumptionType: \"%s\", Status: \"%s\""+
			"}",
		mr.RegisterID, mr.MeasurementStream, mr.NetworkTariffCode, mr.UnitOfMeasure,
		mr.TimeOfDay, mr.Multiplier, mr.DialFormat, mr.ControlledLoad, mr.ConsumptionType,
		mr.Status,
	)
}

// String meets the stringer interface.
func (mr *MeterRegister) String() string {
	if mr == nil {
		return ""
	}

	return mr.RegisterID
}
