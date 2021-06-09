package nmi

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAllocationsCompile(t *testing.T) {
	Convey("Should all compile", t, func() {
		for i, a := range participantAllocations {
			Convey(fmt.Sprintf("returns a slice of regular expressions for %d", i), func() {
				regexps, err := a.Compile()
				So(err, ShouldBeNil)

				Convey(fmt.Sprintf("with %d entities", len(a)), func() {
					So(len(regexps), ShouldEqual, len(a))
				})
			})
		}
	})

	Convey("With invalid", t, func() {
		a := Allocations{Pattern("1/;l1jt90b[eosinb09t2	90nql;kvn0923fh")}

		resp, err := a.Compile()
		So(resp, ShouldBeEmpty)
		So(err, ShouldBeError)
		So(errors.As(err, &ErrPatternInvalid), ShouldBeTrue)
	})
}
