package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReasons(t *testing.T) {
	Convey("Reasons", t, func() {
		got := Reasons()

		So(got, ShouldHaveLength, len(ReasonName))
	})
}

func TestActiveReasons(t *testing.T) {
	Convey("ActiveReasons()", t, func() {
		got := ActiveReasons()

		So(got, ShouldHaveLength, len(reasons)-len(reasonsDeprecated))
	})
}

func TestDeprecatedReasons(t *testing.T) {
	Convey("DeprecatedReasons()", t, func() {
		got := DeprecatedReasons()

		So(got, ShouldHaveLength, len(reasonsDeprecated))
	})
}

func TestNewReason(t *testing.T) {
	type test struct {
		arg      string
		expected Reason
		err      error
	}

	Convey("NewReason", t, func() {
		tests := map[string]test{
			"empty reason": {
				err: ErrReasonCodeNil,
			},
			"invalid reason": {
				arg: "Type-1",
				err: ErrReasonCodeInvalid,
			},
			"valid reason": {
				arg:      "4",
				expected: ReasonDangerousDog,
			},
		}

		for r, n := range ReasonName {
			tests[fmt.Sprintf("reason '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewReason(tc.arg)

				if tc.err != nil {
					So(got, ShouldEqual, ReasonUndefined)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(got, ShouldEqual, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestReasonValidate(t *testing.T) {
	type test struct {
		arg Reason
		err error
	}

	Convey("reason.Validate()", t, func() {
		tests := map[string]test{
			"invalid reason": {
				arg: Reason(-1),
				err: ErrReasonCodeInvalid,
			},
			"undefined reason": {
				arg: ReasonUndefined,
				err: ErrReasonCodeInvalid,
			},
			"valid reason": {
				arg: ReasonDangerousDog,
			},
		}

		for r, n := range ReasonName {
			tests[fmt.Sprintf("reason '%s'", n)] = test{arg: r}
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

func TestReasonGoString(t *testing.T) {
	Convey("reason.GoString()", t, func() {
		tests := map[string]struct {
			arg      Reason
			expected string
		}{
			"UNDEFINED": {
				arg:      ReasonUndefined,
				expected: "%!Reason(-1)",
			},
			"ReasonDangerousDog": {
				arg:      ReasonDangerousDog,
				expected: "%!Reason(4)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a reason of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestReasonString(t *testing.T) {
	Convey("reason.String()", t, func() {
		tests := map[string]struct {
			arg      Reason
			expected string
		}{
			"UNDEFINED": {
				arg:      ReasonUndefined,
				expected: "%!Reason(-1)",
			},
			"ReasonDangerousDog": {
				arg:      ReasonDangerousDog,
				expected: "\"4: Dangerous dog\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a reason of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestReasonDeprecated(t *testing.T) {
	Convey("reason.Deprecated()", t, func() {
		tests := map[string]struct {
			arg      Reason
			expected bool
		}{
			"UNDEFINED": {
				arg: ReasonUndefined,
			},
			"ReasonDangerousDog": {
				arg: ReasonDangerousDog,
			},
			"ReasonROMChecksumError": {
				arg:      ReasonROMChecksumError,
				expected: true,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a reason of '%s'", name), func() {
				got := tc.arg.Deprecated()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestRequiresDescription(t *testing.T) {
	Convey("reason.RequiresDescription()", t, func() {
		tests := map[string]struct {
			arg      Reason
			expected bool
		}{
			"UNDEFINED": {
				arg: ReasonUndefined,
			},
			"ReasonDangerousDog": {
				arg: ReasonDangerousDog,
			},
			"ReasonFreeTextDescription": {
				arg:      ReasonFreeTextDescription,
				expected: true,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a reason of '%s'", name), func() {
				got := tc.arg.RequiresDescription()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
