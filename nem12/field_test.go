package nem12

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jufemaiz/go-aemo/nmi"
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

	Convey("nem12/NewField", t, func() {
		tests := map[string]test{
			"field type is invalid": {
				ft:       FieldType(-1),
				val:      "BOB",
				expected: Field{Type: FieldType(-1), Value: "BOB"},
				err:      ErrFieldTypeInvalid,
			},
			"valid field with nil value": {
				ft:       FieldVersionHeader,
				val:      "",
				expected: Field{Type: FieldVersionHeader, Value: ""},
				err:      ErrFieldNil,
			},
			"valid field": {
				ft:       FieldVersionHeader,
				val:      "NEM12",
				expected: Field{Type: FieldVersionHeader, Value: "NEM12"},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
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

func TestField_GoString(t *testing.T) {
	Convey("nem12/Field.GoString", t, func() {
		tests := map[string]struct {
			field func() Field
			resp  string
		}{
			"empty": {
				field: func() Field {
					return Field{}
				},
				resp: "Field{Type: FieldType(0), Value: \"\"}",
			},
			"invalid value": {
				field: func() Field {
					return Field{
						Type:  FieldVersionHeader,
						Value: "NEM13",
					}
				},
				resp: "Field{Type: FieldType(2), Value: \"NEM13\"}",
			},
			"valid value": {
				field: func() Field {
					return Field{
						Type:  FieldVersionHeader,
						Value: "NEM12",
					}
				},
				resp: "Field{Type: FieldType(2), Value: \"NEM12\"}",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				f := tc.field()
				resp := f.GoString()

				So(resp, ShouldEqual, tc.resp)
			})
		}
	})
}

func TestFieldType_Identifier(t *testing.T) {
	Convey("nem12/FieldType.Identifier", t, func() {
		tests := map[string]struct {
			ft   FieldType
			resp string
		}{
			"empty": {
				ft:   FieldType(0),
				resp: "%!FieldType(0)",
			},
			"valid value": {
				ft:   FieldVersionHeader,
				resp: "version header",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				ft := tc.ft
				resp := ft.Identifier()

				So(resp, ShouldEqual, tc.resp)
			})
		}
	})
}
func TestField_Validate(t *testing.T) {
	Convey("nem12/Field.Validate", t, func() {
		tests := map[string]struct {
			field func() Field
			err   error
		}{
			"empty": {
				field: func() Field {
					return Field{}
				},
				err: ErrFieldTypeInvalid,
			},
			"invalid value": {
				field: func() Field {
					return Field{
						Type:  FieldVersionHeader,
						Value: "NEM13",
					}
				},
				err: ErrIsInvalid,
			},
			"valid value": {
				field: func() Field {
					return Field{
						Type:  FieldVersionHeader,
						Value: "NEM12",
					}
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				f := tc.field()
				err := f.Validate()

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldRecordIndicator(t *testing.T) {
	Convey("nem12/validateFieldRecordIndicator", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid": {
				val: "INVALID",
				err: ErrIsInvalid,
			},
			"100": {
				val: "100",
			},
			"200": {
				val: "200",
			},
			"300": {
				val: "300",
			},
			"400": {
				val: "400",
			},
			"500": {
				val: "500",
			},
			"900": {
				val: "900",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldRecordIndicator(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldVersionHeader(t *testing.T) {
	Convey("nem12/validateFieldVersionHeader", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid": {
				val: "NEM13",
				err: ErrIsInvalid,
			},
			"NEM12": {
				val: "NEM12",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldVersionHeader(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldDateTime(t *testing.T) {
	Convey("nem12/validateFieldDateTime", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool // needed because golang time does not export standard errors to check.
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field date time 'NOT A DATE': parsing time \"NOT A DATE\" as \"200601021504\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "20211010",
				err: errors.New("field date time '20211010': parsing time \"20211010\" as \"200601021504\": cannot parse \"\" as \"15\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "999999999999",
				err: errors.New("field date time '999999999999': parsing time \"999999999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "202112011234",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldDateTime(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldFromParticipant(t *testing.T) {
	Convey("nem12/validateFieldFromParticipant", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid length": {
				val: "THIS STRING IS TOO LONG",
				err: ErrIsInvalid,
			},
			"valid value of 'ENRGYAUST'": {
				val: "ENRGYAUST",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldFromParticipant(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldToParticipant(t *testing.T) {
	Convey("nem12/validateFieldToParticipant", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid length": {
				val: "THIS STRING IS TOO LONG",
				err: ErrIsInvalid,
			},
			"valid value of 'ENRGYAUST'": {
				val: "ENRGYAUST",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldToParticipant(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldNMI(t *testing.T) {
	Convey("nem12/validateFieldNMI", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid NMI": {
				val: "NOTANMI",
				err: nmi.ErrIsInvalid,
			},
			"invalid length": {
				val: "THIS STRING IS TOO LONG",
				err: nmi.ErrIsInvalid,
			},
			"valid value of '4123456789'": {
				val: "4123456789",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldNMI(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldNMIConfiguration(t *testing.T) {
	Convey("nem12/validateFieldNMIConfiguration", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid NMI configration": {
				val: "NOTANMICONFIG",
				err: ErrIsInvalid,
			},
			"invalid NMI configuration due to duplicate": {
				val: "E1B1E1",
				err: ErrIsDuplicated,
			},
			"invalid length": {
				val: "E1B1K1Q1E2B2K2Q2E3B3K3Q3E4B4K4Q4E5B5K5Q5E6B6K6Q6E7B7K7Q7E8B8K8Q8E9B9K9Q9EABAKAQAEBBBKBQBECBCKCQCEDBDKDQDEEBEKEQEEFBFKFQFEGBGKGQGEHBHKHQHEJBJKJQJEKBKKKQKELBLKLQLEMBMKMQMENBNKNQNEPBPKPQPEQBQKQQQERBRKRQRESBSKSQSETBTKTQTEUBUKUQUEVBVKVQVEWBWKWQWEXBXKXQXEYBYKYQYEZBZKZQZ",
				// err: ErrIsInvalid, // is currently being accepted as it is an artefact limit.
			},
			"valid value of 'E1B1E2B2K2Q2'": {
				val: "E1B1E2B2K2Q2",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldNMIConfiguration(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldRegisterID(t *testing.T) {
	Convey("nem12/validateFieldRegisterID", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "123456789123456789",
				err: ErrIsInvalid,
			},
			"valid value of '123'": {
				val: "123",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldRegisterID(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldNMISuffix(t *testing.T) {
	Convey("nem12/validateFieldNMISuffix", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid suffix": {
				val: "NOT A SUFFIX",
				err: ErrIsInvalid,
			},
			"valid value of 'E1'": {
				val: "E1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldNMISuffix(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldMDMDataStreamIdentifier(t *testing.T) {
	Convey("nem12/validateFieldMDMDataStreamIdentifier", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid mdm data stream identifier": {
				val: "NOT A DATA STREAM",
				err: ErrIsInvalid,
			},
			"valid value of 'E1'": {
				val: "E1",
			},
			"valid value of 'NA'": {
				val: "NA",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldMDMDataStreamIdentifier(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldMeterSerialNumber(t *testing.T) {
	Convey("nem12/validateFieldMeterSerialNumber", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "A REALLY LONG METER SERIAL NUMBER",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "ABC123",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldMeterSerialNumber(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldUnitOfMeasurement(t *testing.T) {
	Convey("nem12/validateFieldUnitOfMeasurement", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid value": {
				val: "NANOMETERS",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "KWH",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldUnitOfMeasurement(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldIntervalLength(t *testing.T) {
	Convey("nem12/validateFieldIntervalLength", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid type": {
				val: "NOTANUMBER",
				err: errors.New("field interval length 'NOTANUMBER': strconv.Atoi: parsing \"NOTANUMBER\": invalid syntax"), //nolint:goerr113
			},
			"invalid value": {
				val:        "42",
				err:        ErrIsInvalid,
				shouldWrap: true,
			},
			"valid": {
				val: "30",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldIntervalLength(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldNextScheduledReadDate(t *testing.T) {
	Convey("nem12/validateFieldNextScheduledReadDate", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field next scheduled read date 'NOT A DATE': parsing time \"NOT A DATE\" as \"20060102\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "202110101234",
				err: errors.New("field next scheduled read date '202110101234': parsing time \"202110101234\": extra text: \"1234\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "99999999",
				err: errors.New("field next scheduled read date '99999999': parsing time \"99999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "20211201",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldNextScheduledReadDate(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldIntervalDate(t *testing.T) {
	Convey("nem12/validateFieldIntervalDate", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field interval date 'NOT A DATE': parsing time \"NOT A DATE\" as \"20060102\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "202110101234",
				err: errors.New("field interval date '202110101234': parsing time \"202110101234\": extra text: \"1234\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "99999999",
				err: errors.New("field interval date '99999999': parsing time \"99999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "20211201",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldIntervalDate(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldIntervalValue(t *testing.T) {
	Convey("nem12/validateFieldIntervalValue", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid": {
				val: "NOT A VALUE",
				err: errors.New("field interval value 'NOT A VALUE': can't convert NOT A VALUE to decimal: exponent is not numeric"), //nolint:goerr113
			},
			"invalid when negative": {
				val:        "-1.23456789",
				err:        ErrFieldIntervalValueNegative,
				shouldWrap: true,
			},
			"valid": {
				val: "1234.56789",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldIntervalValue(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldQualityMethod(t *testing.T) {
	Convey("nem12/validateFieldQualityMethod", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid length": {
				val: "QMINVALIDLENGTH",
				err: ErrIsInvalid,
			},
			"invalid forward estimate": {
				val: "E",
				err: ErrIsMissing,
			},
			"invalid final substituted": {
				val: "F",
				err: ErrIsMissing,
			},
			"invalid substituted": {
				val: "S",
				err: ErrIsMissing,
			},
			"valid actual": {
				val: "A",
			},
			"valid variable": {
				val: "V",
			},
			"valid forward estimate": {
				val: "E12",
			},
			"valid final substituted": {
				val: "F13",
			},
			"valid substituted": {
				val: "S14",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldQualityMethod(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldReasonCode(t *testing.T) {
	Convey("nem12/validateFieldReasonCode", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "RCINVALIDLENGTH",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldReasonCode(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldReasonDescription(t *testing.T) {
	Convey("nem12/validateFieldReasonDescription", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "A REALLY LONG STRING OVER TWO HUNDRED AND FOURTY CHARACTERS IS NOT PERMISSIBLE UNDER AEMO'S METER DATA FILE FORMAT SPECIFICATION NEM12 & NEM13 AS AT THE FIRST OF OCTOBER 2021 WHICH IS VERSION 2.4 OF THE SPECIFICATION. SO NOW WE JUST NEED TO FILL IN A WHOLE HEAP MORE CHARACTERS", //nolint:misspell
				err: ErrIsInvalid,
			},
			"valid": {
				val: "A DESCRIPTION TO GO WITH THE REASON CODE",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldReasonDescription(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldUpdateDateTime(t *testing.T) {
	Convey("nem12/validateFieldUpdateDateTime", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field update datetime 'NOT A DATE': parsing time \"NOT A DATE\" as \"20060102150405\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "202110101234",
				err: errors.New("field update datetime '202110101234': parsing time \"202110101234\" as \"20060102150405\": cannot parse \"\" as \"05\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "99999999",
				err: errors.New("field update datetime '99999999': parsing time \"99999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "20211201123456",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldUpdateDateTime(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldMSATSLoadDateTime(t *testing.T) {
	Convey("nem12/validateFieldMSATSLoadDateTime", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field update datetime 'NOT A DATE': parsing time \"NOT A DATE\" as \"20060102150405\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "202110101234",
				err: errors.New("field update datetime '202110101234': parsing time \"202110101234\" as \"20060102150405\": cannot parse \"\" as \"05\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "99999999",
				err: errors.New("field update datetime '99999999': parsing time \"99999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "20211201123456",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldMSATSLoadDateTime(tc.val)

				if tc.err != nil {
					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldStartInterval(t *testing.T) {
	Convey("nem12/validateFieldStartInterval", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid not a number": {
				val: "NOT A NUMBER",
				err: errors.New("field start interval 'NOT A NUMBER': strconv.Atoi: parsing \"NOT A NUMBER\": invalid syntax"), //nolint:goerr113
			},
			"invalid negative number": {
				val:        "-48",
				err:        ErrFieldIntervalNegativeInvalid,
				shouldWrap: true,
			},
			"invalid too big a number": {
				val:        "300",
				err:        ErrFieldIntervalExceedsMaximum,
				shouldWrap: true,
			},
			"valid": {
				val: "1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldStartInterval(tc.val)

				if tc.err != nil {
					So(err, ShouldNotBeNil)

					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldFinishInterval(t *testing.T) {
	Convey("nem12/validateFieldFinishInterval", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {
				err:        ErrIsNil,
				shouldWrap: true,
			},
			"invalid not a number": {
				val: "NOT A NUMBER",
				err: errors.New("field finish interval 'NOT A NUMBER': strconv.Atoi: parsing \"NOT A NUMBER\": invalid syntax"), //nolint:goerr113
			},
			"invalid negative number": {
				val:        "-48",
				err:        ErrFieldIntervalNegativeInvalid,
				shouldWrap: true,
			},
			"invalid too big a number": {
				val:        "300",
				err:        ErrFieldIntervalExceedsMaximum,
				shouldWrap: true,
			},
			"valid": {
				val: "1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldFinishInterval(tc.val)

				if tc.err != nil {
					So(err, ShouldNotBeNil)

					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldTransactionCode(t *testing.T) {
	Convey("nem12/validateFieldTransactionCode", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {
				err: ErrIsNil,
			},
			"invalid": {
				val: "NOT A TRANSACTION CODE",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "A",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldTransactionCode(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldRetServiceOrder(t *testing.T) {
	Convey("nem12/validateFieldRetServiceOrder", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "RET SERVICE ORDER CANNOT BE A LONG STRING OR IT FAILS TO VALIDATE",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "A",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldRetServiceOrder(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldReadDateTime(t *testing.T) {
	Convey("nem12/validateFieldReadDateTime", t, func() {
		tests := map[string]struct {
			val        string
			err        error
			shouldWrap bool
		}{
			"empty": {},
			"invalid": {
				val: "NOT A DATE",
				err: errors.New("field read datetime 'NOT A DATE': parsing time \"NOT A DATE\" as \"20060102150405\": cannot parse \"NOT A DATE\" as \"2006\""), //nolint:goerr113
			},
			"invalid length": {
				val: "202110101234",
				err: errors.New("field read datetime '202110101234': parsing time \"202110101234\" as \"20060102150405\": cannot parse \"\" as \"05\""), //nolint:goerr113
			},
			"invalid numbers": {
				val: "99999999",
				err: errors.New("field read datetime '99999999': parsing time \"99999999\": month out of range"), //nolint:goerr113
			},
			"valid": {
				val: "20211201123456",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldReadDateTime(tc.val)

				if tc.err != nil {
					So(err, ShouldNotBeNil)

					if tc.shouldWrap {
						So(err, ShouldWrap, tc.err)
					} else {
						So(err.Error(), ShouldEqual, tc.err.Error())
					}
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestField_validateFieldIndexRead(t *testing.T) {
	Convey("nem12/validateFieldIndexRead", t, func() {
		tests := map[string]struct {
			val string
			err error
		}{
			"empty": {},
			"invalid length": {
				val: "INDEX READ FIELD VALUE CANNOT BE A LONG STRING OR IT FAILS TO VALIDATE",
				err: ErrIsInvalid,
			},
			"valid": {
				val: "A",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				err := validateFieldIndexRead(tc.val)

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func Test_chunkString(t *testing.T) {
	Convey("nem12/chunkString", t, func() {
		tests := map[string]struct {
			arg  string
			resp []string
		}{
			"empty": {
				arg:  "",
				resp: []string{""},
			},
			"1 char": {
				arg:  "1",
				resp: []string{"1"},
			},
			"4 char": {
				arg:  "1234",
				resp: []string{"1", "2", "3", "4"},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				resp := chunkString(tc.arg, 1)

				So(resp, ShouldResemble, tc.resp)
			})
		}
	})
}
