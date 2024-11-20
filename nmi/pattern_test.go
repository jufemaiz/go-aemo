package nmi

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPattern(t *testing.T) {
	Convey("NewPattern()", t, func() {
		tests := map[string]struct {
			p   string
			err error
		}{
			"valid pattern": {
				p: "^bob$",
			},
			"invalid pattern": {
				p:   "1u40[1lg hocq89 1 98r 34x",
				err: ErrPatternInvalid,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("When %s", name), func() {
				p, err := NewPattern(tc.p)

				if tc.err != nil {
					So(err, ShouldBeError)
					So(err, ShouldWrap, tc.err)
				} else {
					So(err, ShouldBeNil)
					So(string(p), ShouldEqual, tc.p)
				}
			})
		}
	})
}

func TestPatternMatch(t *testing.T) {
	Convey("pattern.Match(string)", t, func() {
		tests := map[string]struct {
			p        Pattern
			s        string
			expected bool
		}{
			"valid with match": {
				p:        Pattern("^GOT$"),
				s:        "GOT",
				expected: true,
			},
			"valid without match": {
				p:        Pattern("^GOT$"),
				s:        "NOTGOT",
				expected: false,
			},
			"invalid pattern": {
				p:        Pattern("1u40[1lg hocq89 1 98r 34x"),
				s:        "NOTGOT",
				expected: false,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got := tc.p.Match(tc.s)

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}
