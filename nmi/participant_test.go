package nmi

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jufemaiz/go-aemo/region"
)

func TestParticipants(t *testing.T) {
	Convey("Participants()", t, func() {
		got := Participants()

		So(got, ShouldHaveLength, len(participants))
	})
}

func TestParticipantsForRegion(t *testing.T) {
	Convey("ParticipantsForRegion(region.Region)", t, func() {
		tests := map[string]struct {
			len int
			err error
		}{
			"ACT":       {len: 1},
			"NSW":       {len: 7},
			"NT":        {len: 1},
			"QLD":       {len: 4},
			"SA":        {len: 3},
			"TAS":       {len: 3},
			"VIC":       {len: 7},
			"WA":        {len: 3},
			"UNDEFINED": {len: 2},
			"BOB":       {err: region.ErrRegionInvalid},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given region '%s'", name), func() {
				r, err := region.NewRegion(name)
				if tc.err != nil {
					So(err, ShouldBeError, tc.err)
				} else {
					So(err, ShouldBeNil)

					resp := ParticipantsForRegion(r)
					So(resp, ShouldHaveLength, tc.len)
				}
			})
		}
	})
}

func TestNewParticipant(t *testing.T) {
	Convey("NewParticipant()", t, func() {
		Convey("When I have a known participant", func() {
			for _, pid := range ParticipantIDs {
				Convey(fmt.Sprintf("Given the participant id '%s'", pid), func() {
					p, err := NewParticipant(pid)

					So(p.ParticipantID(), ShouldEqual, pid)
					So(err, ShouldBeNil)
				})
			}
		})

		Convey("When I have an unknown participant ID", func() {
			p, err := NewParticipant("THIS_IS_FAKE")

			So(p, ShouldEqual, ParticipantUndefined)
			So(err, ShouldBeError)
			So(errors.As(err, &ErrParticipantInvalid), ShouldBeTrue)
		})
	})
}

func TestParticipantGoString(t *testing.T) {
	Convey("GoString()", t, func() {
		tests := map[string]struct {
			arg      Participant
			expected string
		}{
			"valid participant": {
				arg:      ParticipantENERGYAP,
				expected: "Participant{Participant: 8, ParticipantID: \"ENERGYAP\", Region: {Region: 3, MarketNode: \"NSW\", Name: \"NSW\", LongName: \"New South Wales\", ISOCode: \"AU-NSW\"}, LongName: \"Ausgrid\", ShortName: \"Ausgrid\", Energy: \"ELECTRICITY\", Allocations: nmi.Allocations{\"^(NCCC[A-HJ-NP-VX-Z\\\\d][A-HJ-NP-Z\\\\d]{5})$\", \"^(410[234]\\\\d{6})$\"}}",
			},
			"undefined participant": {
				arg:      ParticipantUndefined,
				expected: "Participant{Participant: 0, ParticipantID: \"UNDEFINED\", Region: {Region: 0, MarketNode: \"UNDEFINED\", Name: \"UNDEFINED\", LongName: \"Undefined\", ISOCode: \"UNDEFINED\"}, LongName: \"UNDEFINED\", ShortName: \"UNDEFINED\", Energy: \"UNDEFINED\", Allocations: nmi.Allocations{}}",
			},
			"invalid participant": {
				arg:      Participant(-1),
				expected: "Participant{Participant: 0, ParticipantID: \"UNDEFINED\", Region: {Region: 0, MarketNode: \"UNDEFINED\", Name: \"UNDEFINED\", LongName: \"Undefined\", ISOCode: \"UNDEFINED\"}, LongName: \"UNDEFINED\", ShortName: \"UNDEFINED\", Energy: \"UNDEFINED\", Allocations: nmi.Allocations{}}",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestParticipantString(t *testing.T) {
	Convey("String()", t, func() {
		tests := map[string]struct {
			arg      Participant
			expected string
		}{
			"valid participant": {
				arg:      ParticipantENERGYAP,
				expected: "ENERGYAP",
			},
			"undefined participant": {
				arg:      ParticipantUndefined,
				expected: "UNDEFINED",
			},
			"invalid participant": {
				arg:      Participant(-1),
				expected: "UNDEFINED",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got := tc.arg.String()
				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestParticipantInfo(t *testing.T) {
	Convey("Info()", t, func() {
		tests := map[string]struct {
			arg      Participant
			expected ParticipantInfo
			err      error
		}{
			"valid participant": {
				arg: ParticipantENERGYAP,
				expected: ParticipantInfo{
					Participant:   8,
					ParticipantID: "ENERGYAP",
					Region:        region.RegionNSW,
					LongName:      "Ausgrid",
					ShortName:     "Ausgrid",
					Energy:        EnergyElectricity,
					Allocations:   allocationsENERGYAP,
				},
			},
			"undefined participant": {
				arg:      ParticipantUndefined,
				expected: ParticipantInfo{},
				err:      ErrParticipantInvalid,
			},
			"invalid participant": {
				arg:      Participant(-1),
				expected: ParticipantInfo{},
				err:      ErrParticipantInvalid,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got, err := tc.arg.Info()
				if tc.err != nil {
					So(got, ShouldBeNil)
					So(err, ShouldBeError)
					So(errors.As(err, &ErrParticipantInvalid), ShouldBeTrue)
				} else {
					So(*got, ShouldResemble, tc.expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
