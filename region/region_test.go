package region

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRegion(t *testing.T) {
	type test struct {
		arg      string
		expected Region
		err      error
	}

	Convey("NewRegion", t, func() {
		tests := map[string]test{
			"empty string": {
				err: ErrRegionInvalid,
			},
			"invalid string": {
				arg: "Not a region",
				err: ErrRegionInvalid,
			},
			"lowercase value": {
				arg:      "nsw",
				expected: RegionNSW,
			},
			"mixed case value": {
				arg:      "nSw",
				expected: RegionNSW,
			},
			"upper case value": {
				arg:      "NSW",
				expected: RegionNSW,
			},
		}

		for r, n := range RegionName {
			tests[fmt.Sprintf("state '%s'", n)] = test{arg: n, expected: r}
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got, err := NewRegion(tc.arg)

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

func TestRegions(t *testing.T) {
	Convey("Region", t, func() {
		got := Regions()

		So(got, ShouldHaveLength, len(RegionName)-1)
	})
}

func TestRegionGoString(t *testing.T) {
	Convey("region.GoString()", t, func() {
		tests := map[string]struct {
			arg      Region
			expected string
		}{
			"UNDEFINED": {
				arg:      RegionUndefined,
				expected: "{Region: 0, MarketNode: \"UNDEFINED\", Name: \"UNDEFINED\", LongName: \"Undefined\", ISOCode: \"UNDEFINED\"}",
			},
			"NSW": {
				arg:      RegionNSW,
				expected: "{Region: 3, MarketNode: \"NSW\", Name: \"NSW\", LongName: \"New South Wales\", ISOCode: \"AU-NSW\"}",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a region of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestInfo(t *testing.T) {
	Convey("region.Info()", t, func() {
		tests := map[string]struct {
			arg      Region
			expected *Info
			err      error
		}{
			"UNDEFINED": {
				arg: RegionUndefined,
				err: ErrRegionInvalid,
			},
			"NSW": {
				arg: RegionNSW,
				expected: &Info{
					Region:     RegionNSW,
					MarketNode: RegionNSW,
					Name:       "NSW",
					LongName:   "New South Wales",
					ISOCode:    "AU-NSW",
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a region of '%s'", name), func() {
				got, err := tc.arg.Info()

				if tc.err != nil {
					So(got, ShouldResemble, tc.expected)
					So(err, ShouldBeError)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				} else {
					So(*got, ShouldResemble, *tc.expected)
					So(err, ShouldBeNil)
					So(errors.Is(err, tc.err), ShouldBeTrue)
				}
			})
		}
	})
}
