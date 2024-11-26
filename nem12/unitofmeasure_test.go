package nem12

import (
	"errors"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitOfMeasures(t *testing.T) {
	Convey("Units", t, func() {
		got := Units()

		So(got, ShouldHaveLength, len(UnitName))
	})
}

func TestNewUnit(t *testing.T) {
	type test struct {
		arg      string
		expected UnitOfMeasure
		err      error
	}

	Convey("NewUnit", t, func() {
		tests := map[string]test{
			"empty unit of measure": {
				err: ErrUnitOfMeasureNil,
			},
			"invalid unit of measure": {
				arg: "Type-1",
				err: ErrUnitOfMeasureInvalid,
			},
			"valid unit of measure": {
				arg:      "wh",
				expected: UnitWattHour,
			},
		}

		for r, n := range UnitName {
			tests[fmt.Sprintf("unit of measure '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewUnit(tc.arg)

				if tc.err != nil {
					So(got, ShouldEqual, UnitUndefined)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(got, ShouldEqual, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestUnitOfMeasureValidate(t *testing.T) {
	type test struct {
		arg UnitOfMeasure
		err error
	}

	Convey("unit.Validate()", t, func() {
		tests := map[string]test{
			"invalid unit of measure": {
				arg: UnitOfMeasure(-1),
				err: ErrUnitOfMeasureInvalid,
			},
			"undefined unit of measure": {
				arg: UnitUndefined,
				err: ErrUnitOfMeasureInvalid,
			},
			"valid unit of measure": {
				arg: UnitWattHour,
			},
		}

		for r, n := range UnitName {
			tests[fmt.Sprintf("unit of measure '%s'", n)] = test{arg: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := tc.arg.Validate()

				if tc.err != nil {
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestUnitOfMeasureGoString(t *testing.T) {
	Convey("unit.GoString()", t, func() {
		tests := map[string]struct {
			arg      UnitOfMeasure
			expected string
		}{
			"UNDEFINED": {
				arg:      UnitUndefined,
				expected: "%!UnitOfMeasure(0)",
			},
			"UnitWattHour": {
				arg:      UnitWattHour,
				expected: "UnitOfMeasure(3)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a unit of measure of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestUnitOfMeasureString(t *testing.T) {
	Convey("unit.String()", t, func() {
		tests := map[string]struct {
			arg      UnitOfMeasure
			expected string
		}{
			"UNDEFINED": {
				arg:      UnitUndefined,
				expected: "%!UnitOfMeasure(0)",
			},
			"UnitWattHour": {
				arg:      UnitWattHour,
				expected: "\"WH: watt hour\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a unit of measure of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestUnitOfMeasureMultiplier(t *testing.T) {
	Convey("unit.Multiplier()", t, func() {
		tests := map[string]struct {
			arg      UnitOfMeasure
			expected float64
		}{
			"UNDEFINED": {
				arg:      UnitUndefined,
				expected: 0.0,
			},
			"UnitWattHour": {
				arg:      UnitWattHour,
				expected: 1.0,
			},
			"UnitKilowattHour": {
				arg:      UnitKilowattHour,
				expected: 1.0e3,
			},
			"UnitMegawattHour": {
				arg:      UnitMegawattHour,
				expected: 1.0e6,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a unit of measure of '%s'", name), func() {
				got := tc.arg.Multiplier()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestUnitOfMeasureDecimalMultiplier(t *testing.T) {
	Convey("unit.DecimalMultiplier()", t, func() {
		tests := map[string]struct {
			arg      UnitOfMeasure
			expected decimal.Decimal
		}{
			"UNDEFINED": {
				arg:      UnitUndefined,
				expected: decimal.Zero,
			},
			"UnitWattHour": {
				arg:      UnitWattHour,
				expected: decimal.NewFromFloat(1.0),
			},
			"UnitKilowattHour": {
				arg:      UnitKilowattHour,
				expected: decimal.NewFromFloat(1.0e3),
			},
			"UnitMegawattHour": {
				arg:      UnitMegawattHour,
				expected: decimal.NewFromFloat(1.0e6),
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a unit of measure of '%s'", name), func() {
				got := tc.arg.DecimalMultiplier()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
