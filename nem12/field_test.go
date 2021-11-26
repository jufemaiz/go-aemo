package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFields(t *testing.T) {
	Convey("Fields", t, func() {
		got := Fields()

		So(got, ShouldHaveLength, len(fieldName))
	})
}

func TestNewField(t *testing.T) {
	type test struct {
		ft       FieldType
		val      string
		expected Field
		err      error
	}

	Convey("NewField", t, func() {
		tests := map[string]test{
			"field type is invalid": {
				ft:       FieldType(-1),
				val:      "BOB",
				expected: Field{Type: FieldType(-1), Value: "BOB"},
				err:      ErrFieldTypeInvalid,
			},
			"valid field": {
				ft:       FieldVersionHeader,
				val:      "",
				expected: Field{Type: FieldVersionHeader, Value: ""},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got, err := NewField(tc.ft, tc.val)

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
