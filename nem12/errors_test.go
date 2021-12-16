package nem12_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jufemaiz/go-aemo/nem12"
)

func TestParseErrorError(t *testing.T) {
	Convey("nem12/ParseError.Error", t, func() {
		tests := map[string]struct {
			parseErr *nem12.ParseError
			err      string
		}{
			"nil parse error": {},
			"parse error exists": {
				parseErr: &nem12.ParseError{
					Line:  1,
					Field: 1,
					Err:   nem12.ErrIsInvalid,
				},
				err: "parse error on line 1, field 1: is invalid",
			},
			"parse error is ErrParseFieldCountInvalid": {
				parseErr: &nem12.ParseError{
					Line: 22,
					Err:  nem12.ErrParseFieldCountInvalid,
				},
				err: "record on line 22: parse fields count is invalid",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				pe := tc.parseErr

				err := pe.Error()

				if tc.err != "" {
					So(err, ShouldEqual, tc.err)
				} else {
					So(err, ShouldBeBlank)
				}
			})
		}
	})
}

func TestParseErrorUnwrap(t *testing.T) {
	Convey("nem12/ParseError.Unwrap", t, func() {
		tests := map[string]struct {
			parseErr *nem12.ParseError
			err      error
		}{
			"nil parse error": {},
			"parse error exists": {
				parseErr: &nem12.ParseError{
					Line:  1,
					Field: 1,
					Err:   nem12.ErrIsInvalid,
				},
				err: nem12.ErrIsInvalid,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				pe := tc.parseErr

				err := pe.Unwrap()

				if tc.err != nil {
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
