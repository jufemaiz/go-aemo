package nmi

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMeterRegisterGoString(t *testing.T) {
	// additionalSiteInformation := "Additional site information."
	// now := time.Now()
	// serialNumber := "ABC123"

	Convey("meterregister.GoString()", t, func() {
		tests := map[string]struct {
			arg      *MeterRegister
			expected string
		}{
			"nil": {
				arg:      nil,
				expected: "nil",
			},
			"minimal": {
				arg:      &MeterRegister{RegisterID: "123"},
				expected: "MeterRegister{RegisterID: \"123\", MeasurementStream: \"\", NetworkTariffCode: \"\", UnitOfMeasure: \"\", TimeOfDay: \"\", Multiplier: 0, DialFormat: \"\", ControlledLoad: false, ConsumptionType: \"\", Status: \"\"}",
			},
			// "with all": {
			// 	arg: &Meter{
			// 		Nmi:        "4123456789",
			// 		Identifier: "1",
			// 		Registers: []*MeterRegister{
			// 			{RegisterID: "123", MeasurementStream: "E1", UnitOfMeasure: "KWH"},
			// 		},
			// 	},
			// 	expected: "MeterRegister{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister{MeterRegister{RegisterID: \"123\", MeasurementStream: \"E1\", NetworkTariffCode: \"\", UnitOfMeasure: \"KWH\", TimeOfDay: \"\", Multiplier: 0, DialFormat: \"\", ControlledLoad: false, ConsumptionType: \"\", Status: \"\"}}}",
			// },
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a meterregister of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestMeterRegisterString(t *testing.T) {
	Convey("meterregister.String()", t, func() {
		tests := map[string]struct {
			arg      *MeterRegister
			expected string
		}{
			"minimal": {
				arg:      &MeterRegister{RegisterID: "123"},
				expected: "123",
			},
			"with nil": {
				arg:      nil,
				expected: "",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a meterregister of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
