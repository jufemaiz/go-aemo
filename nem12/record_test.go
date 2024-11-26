package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRecordIndicators(t *testing.T) {
	Convey("RecordIndicators", t, func() {
		got := RecordIndicators()

		So(got, ShouldHaveLength, len(RecordIndicatorName))
	})
}

func TestNewRecordIndicator(t *testing.T) {
	type test struct {
		arg      string
		expected RecordIndicator
		err      error
	}

	Convey("NewRecordIndicator", t, func() {
		tests := map[string]test{
			"empty record indicator": {
				err: ErrRecordIndicatorNil,
			},
			"invalid record indicator": {
				arg: "Type-1",
				err: ErrRecordIndicatorInvalid,
			},
			"valid record indicator": {
				arg:      "100",
				expected: RecordHeader,
			},
		}

		for r, n := range RecordIndicatorName {
			tests[fmt.Sprintf("record indicator '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewRecordIndicator(tc.arg)

				if tc.err != nil {
					So(got, ShouldEqual, RecordUndefined)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(got, ShouldEqual, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestRecordIndicatorValidate(t *testing.T) {
	type test struct {
		arg RecordIndicator
		err error
	}

	Convey("recordIndicator.Validate()", t, func() {
		tests := map[string]test{
			"invalid record indicator": {
				arg: RecordIndicator(-1),
				err: ErrRecordIndicatorInvalid,
			},
			"undefined record indicator": {
				arg: RecordUndefined,
				err: ErrRecordIndicatorInvalid,
			},
			"valid record indicator": {
				arg: RecordHeader,
			},
		}

		for r, n := range RecordIndicatorName {
			tests[fmt.Sprintf("record indicator '%s'", n)] = test{arg: r}
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

func TestRecordIndicatorGoString(t *testing.T) {
	Convey("recordIndicator.GoString()", t, func() {
		tests := map[string]struct {
			arg      RecordIndicator
			expected string
		}{
			"UNDEFINED": {
				arg:      RecordUndefined,
				expected: "%!RecordIndicator(0)",
			},
			"RecordHeader": {
				arg:      RecordHeader,
				expected: "%!RecordIndicator(1)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a record indicator of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestRecordIndicatorString(t *testing.T) {
	Convey("recordIndicator.String()", t, func() {
		tests := map[string]struct {
			arg      RecordIndicator
			expected string
		}{
			"UNDEFINED": {
				arg:      RecordUndefined,
				expected: "%!RecordIndicator(0)",
			},
			"RecordHeader": {
				arg:      RecordHeader,
				expected: "\"100: header\"",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a record indicator of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestRecordIndicatorFields(t *testing.T) {
	Convey("recordIndicator.Fields()", t, func() {
		tests := map[string]struct {
			arg      RecordIndicator
			expected []FieldType
			err      error
		}{
			"UNDEFINED": {
				arg: RecordUndefined,
			},
			"RecordHeader": {
				arg: RecordHeader,
				expected: []FieldType{
					FieldRecordIndicator, FieldVersionHeader, FieldDateTime, FieldFromParticipant,
					FieldToParticipant,
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a record indicator of '%s'", name), func() {
				got := tc.arg.Fields()

				So(got, ShouldResemble, tc.expected)
			})
		}
	})
}

func TestRecordIndicatorIntervalDataFields(t *testing.T) {
	Convey("recordIndicator.IntervalDataFields()", t, func() {
		tests := map[string]struct {
			arg      RecordIndicator
			i        int
			expected []FieldType
			err      error
		}{
			"UNDEFINED": {
				arg: RecordUndefined,
				err: ErrRecordIndicatorInvalid,
			},
			"RecordHeader": {
				arg: RecordHeader,
				expected: []FieldType{
					FieldRecordIndicator, FieldVersionHeader, FieldDateTime, FieldFromParticipant,
					FieldToParticipant,
				},
			},
			"RecordIntervalData w/ 0 data fields": {
				arg: RecordIntervalData,
				expected: []FieldType{
					FieldRecordIndicator, FieldIntervalDate, FieldQualityMethod,
					FieldReasonCode, FieldReasonDescription, FieldUpdateDateTime, FieldMSATSLoadDateTime,
				},
			},
			"RecordIntervalData w/ 48 data fields": {
				arg: RecordIntervalData,
				i:   48,
				expected: []FieldType{
					FieldRecordIndicator, FieldIntervalDate, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldQualityMethod, FieldReasonCode, FieldReasonDescription, FieldUpdateDateTime, FieldMSATSLoadDateTime,
				},
			},
			"RecordIntervalData w/ 96 data fields": {
				arg: RecordIntervalData,
				i:   96,
				expected: []FieldType{
					FieldRecordIndicator, FieldIntervalDate, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldQualityMethod, FieldReasonCode, FieldReasonDescription, FieldUpdateDateTime, FieldMSATSLoadDateTime,
				},
			},
			"RecordIntervalData w/ 288 data fields": {
				arg: RecordIntervalData,
				i:   288,
				expected: []FieldType{
					FieldRecordIndicator, FieldIntervalDate, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue, FieldIntervalValue,
					FieldQualityMethod, FieldReasonCode, FieldReasonDescription, FieldUpdateDateTime, FieldMSATSLoadDateTime,
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a record indicator of '%s'", name), func() {
				got, err := tc.arg.IntervalDataFields(tc.i)

				So(got, ShouldResemble, tc.expected)

				if tc.err != nil {
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
