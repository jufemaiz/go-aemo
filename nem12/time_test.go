package nem12

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNEMTime(t *testing.T) {
	Convey("NEMTime", t, func() {
		got := NEMTime()

		So(got, ShouldNotBeNil)
	})
}
