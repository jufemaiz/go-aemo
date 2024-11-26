package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethods(t *testing.T) {
	Convey("Method", t, func() {
		got := Methods()

		So(got, ShouldHaveLength, len(MethodName))
	})
}

func TestNewMethod(t *testing.T) {
	type test struct {
		arg      string
		expected Method
		err      error
	}

	Convey("NewMethod", t, func() {
		tests := map[string]test{
			"empty method": {
				err: ErrMethodNil,
			},
			"invalid method": {
				arg: "Type-1",
				err: ErrMethodInvalid,
			},
			"valid method": {
				arg:      "14",
				expected: Method14LikeDay,
			},
		}

		for r, n := range MethodName {
			tests[fmt.Sprintf("method '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewMethodFlag(tc.arg)

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

func TestMethodValidate(t *testing.T) {
	type test struct {
		arg Method
		err error
	}

	Convey("method.Validate()", t, func() {
		tests := map[string]test{
			"invalid method": {
				arg: Method(-1),
				err: ErrMethodInvalid,
			},
			"undefined method": {
				arg: MethodUndefined,
				err: ErrMethodInvalid,
			},
			"valid method": {
				arg: Method14LikeDay,
			},
		}

		for r, n := range MethodName {
			tests[fmt.Sprintf("method '%s'", n)] = test{arg: r}
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

func TestMethodGoString(t *testing.T) {
	Convey("method.GoString()", t, func() {
		tests := map[string]struct {
			arg      Method
			expected string
		}{
			"UNDEFINED": {
				arg:      MethodUndefined,
				expected: "Method(0)",
			},
			"14": {
				arg:      Method14LikeDay,
				expected: "Method(14)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestMethodString(t *testing.T) {
	Convey("method.String()", t, func() {
		tests := map[string]struct {
			arg      Method
			expected string
		}{
			"UNDEFINED": {
				arg:      MethodUndefined,
				expected: "\"Method(0)\"",
			},
			"14": {
				arg:      Method14LikeDay,
				expected: "\"14: retrospective like day\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestMethodMethodTypes(t *testing.T) {
	Convey("method.Types()", t, func() {
		tests := map[string]struct {
			arg      Method
			expected []MethodType
		}{
			"UNDEFINED": {
				arg: MethodUndefined,
			},
			"14": {
				arg:      Method14LikeDay,
				expected: []MethodType{MethodTypeSubstituted},
			},
			"51": {
				arg:      Method51PreviousYear,
				expected: []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
			},
			"61": {
				arg:      Method61PreviousYear,
				expected: []MethodType{MethodTypeEstimated, MethodTypeSubstituted},
			},
			"73": {
				arg:      Method73RevisedAlgorithm,
				expected: []MethodType{MethodTypeSubstituted},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method of '%s'", name), func() {
				got := tc.arg.Types()

				So(got, ShouldResemble, tc.expected)
			})
		}
	})
}

func TestMethodInstallationTypes(t *testing.T) {
	Convey("method.InstallationTypes()", t, func() {
		tests := map[string]struct {
			arg      Method
			expected []Install
		}{
			"UNDEFINED": {
				arg: MethodUndefined,
			},
			"14": {
				arg:      Method14LikeDay,
				expected: []Install{InstallComms1, InstallComms2, InstallComms3, InstallComms4},
			},
			"51": {
				arg:      Method51PreviousYear,
				expected: []Install{InstallMRIM},
			},
			"61": {
				arg:      Method61PreviousYear,
				expected: []Install{InstallBasic},
			},
			"73": {
				arg:      Method73RevisedAlgorithm,
				expected: []Install{InstallUnmetered},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method of '%s'", name), func() {
				got := tc.arg.InstallationTypes()

				So(got, ShouldResemble, tc.expected)
			})
		}
	})
}
