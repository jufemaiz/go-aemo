package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewQualityMethod(t *testing.T) {
	type test struct {
		arg      string
		expected QualityMethod
		err      error
	}

	Convey("NewQualityMethod", t, func() {
		tests := map[string]test{
			"empty quality method": {
				expected: QualityMethod(""),
				err:      ErrQualityMethodNil,
			},
			"invalid quality method - length short": {
				arg:      "T2",
				expected: QualityMethod("T2"),
				err:      ErrQualityMethodLengthInvalid,
			},
			"invalid quality method - length long": {
				arg:      "T23456",
				expected: QualityMethod("T23456"),
				err:      ErrQualityMethodLengthInvalid,
			},
			"invalid quality method - estimate quality without method": {
				arg:      "E",
				expected: QualityMethod("E"),
				err:      ErrQualityMissingMethod,
			},
			"invalid quality method - quality invalid": {
				arg:      "Z",
				expected: QualityMethod("Z"),
				err:      ErrQualityInvalid,
			},
			"invalid quality method - method invalid": {
				arg:      "E99",
				expected: QualityMethod("E99"),
				err:      ErrMethodInvalid,
			},
			"valid quality method": {
				arg:      "A",
				expected: QualityMethod("A"),
			},
			"valid quality method - with method": {
				arg:      "E14",
				expected: QualityMethod("E14"),
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got, err := NewQualityMethod(tc.arg)

				if tc.err != nil {
					So(got, ShouldEqual, tc.expected)
					t.Logf("err: %v", err)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(got, ShouldEqual, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestQualityMethod_Validate(t *testing.T) {
	type test struct {
		arg QualityMethod
		err error
	}

	Convey("nem12/QualityMethod.Validate()", t, func() {
		tests := map[string]test{
			"empty quality method": {
				arg: QualityMethod(""),
				err: ErrQualityMethodNil,
			},
			"invalid quality method - length short": {
				arg: QualityMethod("T2"),
				err: ErrQualityMethodLengthInvalid,
			},
			"invalid quality method - length long": {
				arg: QualityMethod("T23456"),
				err: ErrQualityMethodLengthInvalid,
			},
			"invalid quality method - estimate quality without method": {
				arg: QualityMethod("E"),
				err: ErrQualityMissingMethod,
			},
			"invalid quality method - quality invalid": {
				arg: QualityMethod("Z"),
				err: ErrQualityInvalid,
			},
			"invalid quality method - method invalid": {
				arg: QualityMethod("E99"),
				err: ErrMethodInvalid,
			},
			"valid quality method": {
				arg: QualityMethod("A"),
			},
			"valid quality method - with method": {
				arg: QualityMethod("E14"),
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
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

func TestQualityMethod_Method(t *testing.T) {
	Convey("nem12/QualityMethod.Method()", t, func() {
		tests := map[string]struct {
			arg        func() QualityMethod
			err        error
			assertions func(arg QualityMethod, res Method, err error)
		}{
			"invalid method": {
				arg: func() QualityMethod {
					return QualityMethod("A1B2C3D4E5F6")
				},
				err:        ErrMethodInvalid,
				assertions: func(arg QualityMethod, res Method, err error) {},
			},
			"valid without method": {
				arg: func() QualityMethod {
					return QualityMethod("A")
				},
				assertions: func(arg QualityMethod, res Method, err error) {
					So(res, ShouldEqual, MethodUndefined)
				},
			},
			"valid with method": {
				arg: func() QualityMethod {
					return QualityMethod("S12")
				},
				assertions: func(arg QualityMethod, res Method, err error) {
					So(res, ShouldEqual, Method12Calculated)
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				arg := tc.arg()

				res, err := arg.Method()
				if tc.err != nil {
					So(res, ShouldEqual, MethodUndefined)
					So(err, ShouldBeError)
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}

				tc.assertions(arg, res, err)
			})
		}
	})
}
