package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodTypes(t *testing.T) {
	Convey("MethodType", t, func() {
		got := MethodTypes()

		So(got, ShouldHaveLength, len(MethodTypeName))
	})
}

func TestNewMethodType(t *testing.T) {
	type test struct {
		arg      string
		expected MethodType
		err      error
	}

	Convey("NewMethodType", t, func() {
		tests := map[string]test{
			"empty method type": {
				err: ErrMethodTypeNil,
			},
			"invalid method type": {
				arg: "made up",
				err: ErrMethodTypeInvalid,
			},
			"upper case method type": {
				arg:      "EST",
				expected: MethodTypeEstimated,
			},
			"lower case method type": {
				arg:      "est",
				expected: MethodTypeEstimated,
			},
			"mixed case method type": {
				arg:      "EsT",
				expected: MethodTypeEstimated,
			},
		}

		for r, n := range MethodTypeName {
			tests[fmt.Sprintf("method type '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewMethodType(tc.arg)

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

func TestMethodTypeValidate(t *testing.T) {
	type test struct {
		arg MethodType
		err error
	}

	Convey("method.Validate()", t, func() {
		tests := map[string]test{
			"invalid method": {
				arg: MethodType(-1),
				err: ErrMethodTypeInvalid,
			},
			"undefined method": {
				arg: MethodTypeUndefined,
				err: ErrMethodTypeInvalid,
			},
			"valid method": {
				arg: MethodTypeSubstituted,
			},
		}

		for r, n := range MethodTypeName {
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

func TestMethodTypeGoString(t *testing.T) {
	Convey("methodType.GoString()", t, func() {
		tests := map[string]struct {
			arg      MethodType
			expected string
		}{
			"UNDEFINED": {
				arg:      MethodTypeUndefined,
				expected: "MethodType(0)",
			},
			"NSW": {
				arg:      MethodTypeEstimated,
				expected: "MethodType(1)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method type of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestMethodTypeString(t *testing.T) {
	Convey("methodType.String()", t, func() {
		tests := map[string]struct {
			arg      MethodType
			expected string
		}{
			"UNDEFINED": {
				arg:      MethodTypeUndefined,
				expected: "\"MethodType(0): %!MethodType(0)\"",
			},
			"NSW": {
				arg:      MethodTypeEstimated,
				expected: "\"EST: estimated\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a method type of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

// func TestInfo(t *testing.T) {
// 	Convey("methodType.Info()", t, func() {
// 		tests := map[string]struct {
// 			arg      MethodType
// 			expected *Info
// 			err      error
// 		}{
// 			"UNDEFINED": {
// 				arg: MethodTypeUndefined,
// 				err: ErrMethodTypeInvalid,
// 			},
// 			"NSW": {
// 				arg: MethodTypeEstimated,
// 				expected: &Info{
// 					MethodType:     MethodTypeEstimated,
// 					MarketNode: MethodTypeEstimated,
// 					Name:       "NSW",
// 					LongName:   "New South Wales",
// 					ISOCode:    "AU-NSW",
// 				},
// 			},
// 		}

// 		for name, tc := range tests {
// 			tc := tc

// 			Convey(fmt.Sprintf("Given a method type of '%s'", name), func() {
// 				got, err := tc.arg.Info()

// 				if tc.err != nil {
// 					So(got, ShouldResemble, tc.expected)
// 					So(err, ShouldBeError)
// 					So(errors.Is(err, tc.err), ShouldBeTrue)
// 				} else {
// 					So(*got, ShouldResemble, *tc.expected)
// 					So(err, ShouldBeNil)
// 					So(errors.Is(err, tc.err), ShouldBeTrue)
// 				}
// 			})
// 		}
// 	})
// }
