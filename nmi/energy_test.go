package nmi

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnergies(t *testing.T) {
	Convey("TestEnergies()", t, func() {
		resp := Energies()

		So(resp, ShouldHaveLength, 3)
		So(resp[0], ShouldEqual, EnergyUndefined)
	})
}

func TestNewEnergy(t *testing.T) {
	Convey("NewEnergy()", t, func() {
		tests := map[string]struct {
			s   string
			err error
		}{
			"valid with 'electricity'": {
				s: "electricity",
			},
			"valid with 'gas'": {
				s: "gas",
			},
			"valid with 'ELECTRICITY'": {
				s: "ELECTRICITY",
			},
			"valid with 'GAS'": {
				s: "GAS",
			},
			"invalid with 'WATER'": {
				s:   "WATER",
				err: ErrEnergyInvalid,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("When %s", name), func() {
				resp, err := NewEnergy(tc.s)

				if tc.err != nil {
					So(err, ShouldBeError)
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
					So(resp.String(), ShouldEqual, strings.ToUpper(tc.s))
				}
			})
		}
	})
}

func TestEnergyGoString(t *testing.T) {
	Convey("energy.GoString()", t, func() {
		tests := map[string]struct {
			e Energy
			s string
		}{
			"valid with EnergyElectricity": {
				e: EnergyElectricity,
				s: "\"ELECTRICITY\"",
			},
			"valid with EnergyGas": {
				e: EnergyGas,
				s: "\"GAS\"",
			},
			"invalid value": {
				e: Energy(-1),
				s: "\"UNDEFINED\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("When %s", name), func() {
				resp := tc.e.GoString()

				So(resp, ShouldEqual, tc.s)
			})
		}
	})
}

func TestEnergyString(t *testing.T) {
	Convey("energy.String()", t, func() {
		tests := map[string]struct {
			e Energy
			s string
		}{
			"valid with EnergyElectricity": {
				e: EnergyElectricity,
				s: "ELECTRICITY",
			},
			"valid with EnergyGas": {
				e: EnergyGas,
				s: "GAS",
			},
			"invalid value": {
				e: Energy(-1),
				s: "UNDEFINED",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("When %s", name), func() {
				resp := tc.e.String()

				So(resp, ShouldEqual, tc.s)
			})
		}
	})
}
