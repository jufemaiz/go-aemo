package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQualitys(t *testing.T) {
	Convey("Qualities", t, func() {
		got := Qualities()

		So(got, ShouldHaveLength, len(QualityName))
	})
}

func TestNewQuality(t *testing.T) {
	type test struct {
		arg      string
		expected Quality
		err      error
	}

	Convey("NewQuality", t, func() {
		tests := map[string]test{
			"empty quality": {
				err: ErrQualityNil,
			},
			"invalid quality": {
				arg: "Type-1",
				err: ErrQualityInvalid,
			},
			"valid quality": {
				arg:      "A",
				expected: QualityActual,
			},
		}

		for r, n := range QualityName {
			tests[fmt.Sprintf("quality '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewQualityFlag(tc.arg)

				if tc.err != nil {
					So(got, ShouldBeZeroValue)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(got, ShouldEqual, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestQualityValidate(t *testing.T) {
	type test struct {
		arg Quality
		err error
	}

	Convey("quality.Validate()", t, func() {
		tests := map[string]test{
			"invalid quality": {
				arg: Quality(-1),
				err: ErrQualityInvalid,
			},
			"undefined quality": {
				arg: QualityUndefined,
				err: ErrQualityInvalid,
			},
			"valid quality": {
				arg: QualityActual,
			},
		}

		for r, n := range QualityName {
			tests[fmt.Sprintf("quality '%s'", n)] = test{arg: r}
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

func TestQualityGoString(t *testing.T) {
	Convey("quality.GoString()", t, func() {
		tests := map[string]struct {
			arg      Quality
			expected string
		}{
			"UNDEFINED": {
				arg:      QualityUndefined,
				expected: "Quality(0)",
			},
			"QualityActual": {
				arg:      QualityActual,
				expected: "Quality(1)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a quality of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestQualityString(t *testing.T) {
	Convey("quality.String()", t, func() {
		tests := map[string]struct {
			arg      Quality
			expected string
		}{
			"UNDEFINED": {
				arg:      QualityUndefined,
				expected: "\"Method(0)\"",
			},
			"QualityActual": {
				arg:      QualityActual,
				expected: "\"A: actual data\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a quality of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestQualityMustNotHaveReason(t *testing.T) {
	Convey("quality.MustNotHaveReason()", t, func() {
		tests := map[string]struct {
			arg      Quality
			expected bool
		}{
			"UNDEFINED": {
				arg: QualityUndefined,
			},
			"QualityActual": {
				arg: QualityActual,
			},
			"QualityVariable": {
				arg:      QualityVariable,
				expected: true,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a quality of '%s'", name), func() {
				got := tc.arg.MustNotHaveReason()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestQualityRequiresMethod(t *testing.T) {
	Convey("quality.RequiresMethod()", t, func() {
		tests := map[string]struct {
			arg      Quality
			expected bool
		}{
			"UNDEFINED": {
				arg: QualityUndefined,
			},
			"QualityActual": {
				arg: QualityActual,
			},
			"QualityEstimated": {
				arg:      QualityEstimated,
				expected: true,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a quality of '%s'", name), func() {
				got := tc.arg.RequiresMethod()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestQualityRequiresReason(t *testing.T) {
	Convey("quality.RequiresReason()", t, func() {
		tests := map[string]struct {
			arg      Quality
			expected bool
		}{
			"UNDEFINED": {
				arg: QualityUndefined,
			},
			"QualityActual": {
				arg: QualityActual,
			},
			"QualityFinal": {
				arg:      QualityFinal,
				expected: true,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a quality of '%s'", name), func() {
				got := tc.arg.RequiresReason()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
