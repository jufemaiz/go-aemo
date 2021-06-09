package nmi

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMeterGoString(t *testing.T) {
	additionalSiteInformation := "Additional site information."
	now := time.Now()
	serialNumber := "ABC123"

	Convey("meter.GoString()", t, func() {
		tests := map[string]struct {
			arg      *Meter
			expected string
		}{
			"nil": {
				arg:      nil,
				expected: "nil",
			},
			"minimal": {
				arg:      &Meter{Nmi: "4123456789", Identifier: "1"},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil)}",
			},
			"with serial": {
				arg:      &Meter{Nmi: "4123456789", Identifier: "1", SerialNumber: &serialNumber},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil), SerialNumber: \"ABC123\"}",
			},
			"with registers": {
				arg: &Meter{
					Nmi:        "4123456789",
					Identifier: "1",
					Registers: []*MeterRegister{
						{RegisterID: "123", MeasurementStream: "E1", UnitOfMeasure: "KWH"},
					},
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister{MeterRegister{RegisterID: \"123\", MeasurementStream: \"E1\", NetworkTariffCode: \"\", UnitOfMeasure: \"KWH\", TimeOfDay: \"\", Multiplier: 0, DialFormat: \"\", ControlledLoad: false, ConsumptionType: \"\", Status: \"\"}}}",
			},
			"with from date": {
				arg: &Meter{
					Nmi:          "4123456789",
					Identifier:   "1",
					FromDateTime: &now,
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil)}",
			},
			"with to date": {
				arg: &Meter{
					Nmi:        "4123456789",
					Identifier: "1",
					ToDateTime: &now,
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil)}",
			},
			"with last test date": {
				arg: &Meter{
					Nmi:          "4123456789",
					Identifier:   "1",
					LastTestDate: &now,
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil)}",
			},
			"with additional site information": {
				arg: &Meter{
					Nmi:                "4123456789",
					Identifier:         "1",
					AdditionalSiteInfo: &additionalSiteInformation,
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister(nil)}",
			},
			"with all": {
				arg: &Meter{
					Nmi:        "4123456789",
					Identifier: "1",
					Registers: []*MeterRegister{
						{RegisterID: "123", MeasurementStream: "E1", UnitOfMeasure: "KWH"},
					},
				},
				expected: "Meter{Nmi: \"4123456789\", Identifier: \"1\", Registers: []*nmi.MeterRegister{MeterRegister{RegisterID: \"123\", MeasurementStream: \"E1\", NetworkTariffCode: \"\", UnitOfMeasure: \"KWH\", TimeOfDay: \"\", Multiplier: 0, DialFormat: \"\", ControlledLoad: false, ConsumptionType: \"\", Status: \"\"}}}",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a meter of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestMeterString(t *testing.T) {
	Convey("meter.String()", t, func() {
		tests := map[string]struct {
			arg      *Meter
			expected string
		}{
			"minimal": {
				arg:      &Meter{Nmi: "4123456789", Identifier: "1"},
				expected: "1",
			},
			"with nil": {
				arg:      nil,
				expected: "",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a meter of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
