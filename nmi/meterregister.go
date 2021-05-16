package nmi

import "fmt"

// MeterRegister for a Meter for a NMI
type MeterRegister struct {
	RegisterID        string `json:"registerID"`
	MeasurementStream string `json:"measurementStream"`
	NetworkTariffCode string `json:"networkTariffCode"`
	UnitOfMeasure     string `json:"unitOfMeasure"`
	TimeOfDay         string `json:"timeOfDay"`
	Multiplier        int    `json:"multiplier"`
	DialFormat        string `json:"dialFormat"`
	ControlledLoad    bool   `json:"controlledLoad"`
	ConsumptionType   string `json:"consumptionType"`
	Status            string `json:"status"`
}

// GoString meets the gostring interface.
func (mr *MeterRegister) GoString() string {
	if mr == nil {
		return "nil"
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

	return string(mr.RegisterID)
}
