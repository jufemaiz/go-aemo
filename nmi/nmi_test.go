package nmi

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/jufemaiz/go-aemo/region"
)

func TestFuncChecksum(t *testing.T) {
	Convey("Checksum()", t, func() {
		tests := map[string]struct {
			arg      string
			expected int
		}{
			"2001985732": {arg: "2001985732", expected: 8},
			"QAAAVZZZZZ": {arg: "QAAAVZZZZZ", expected: 3},
			"2001985733": {arg: "2001985733", expected: 6},
			"QCDWW00010": {arg: "QCDWW00010", expected: 2},
			"3075621875": {arg: "3075621875", expected: 8},
			"SMVEW00085": {arg: "SMVEW00085", expected: 8},
			"3075621876": {arg: "3075621876", expected: 6},
			"VAAA000065": {arg: "VAAA000065", expected: 7},
			"4316854005": {arg: "4316854005", expected: 9},
			"VAAA000066": {arg: "VAAA000066", expected: 5},
			"4316854006": {arg: "4316854006", expected: 7},
			"VAAA000067": {arg: "VAAA000067", expected: 2},
			"6305888444": {arg: "6305888444", expected: 6},
			"VAAASTY576": {arg: "VAAASTY576", expected: 8},
			"6350888444": {arg: "6350888444", expected: 2},
			"VCCCX00009": {arg: "VCCCX00009", expected: 1},
			"7001888333": {arg: "7001888333", expected: 8},
			"VEEEX00009": {arg: "VEEEX00009", expected: 1},
			"7102000001": {arg: "7102000001", expected: 7},
			"VKTS786150": {arg: "VKTS786150", expected: 2},
			"NAAAMYS582": {arg: "NAAAMYS582", expected: 6},
			"VKTS867150": {arg: "VKTS867150", expected: 5},
			"NBBBX11110": {arg: "NBBBX11110", expected: 0},
			"VKTS871650": {arg: "VKTS871650", expected: 7},
			"NBBBX11111": {arg: "NBBBX11111", expected: 8},
			"VKTS876105": {arg: "VKTS876105", expected: 7},
			"NCCC519495": {arg: "NCCC519495", expected: 5},
			"VKTS876150": {arg: "VKTS876150", expected: 3},
			"NGGG000055": {arg: "NGGG000055", expected: 4},
			"VKTS876510": {arg: "VKTS876510", expected: 8},
			"IIII001100": {arg: "IIII001100", expected: NMICHECKSUMINVALID},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a nmi of '%s'", name), func() {
				got := Checksum(tc.arg)
				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestNewNmi(t *testing.T) {
	Convey("NewNmi()", t, func() {
		tests := map[string]struct {
			arg      string
			expected *Nmi
			err      error
		}{
			"valid nmi": {
				arg:      "4123456789",
				expected: &Nmi{Identifier: "4123456789"},
			},
			"invalid character nmi": {
				arg: "4OIOIIOIO4",
				err: ErrNmiInvalidChar,
			},
			"invalid length nmi": {
				arg: "4123456789123456789123456789",
				err: ErrNmiInvalidLength,
			},
			"empty nmi": {
				err: ErrNmiInvalidLength,
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given %s", name), func() {
				got, err := NewNmi(tc.arg)

				if tc.err != nil {
					So(got, ShouldBeNil)
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
					So(got, ShouldNotBeNil)
					So(got, ShouldResemble, tc.expected)
				}
			})
		}
	})
}

func TestNmiGoString(t *testing.T) {
	Convey("GoString()", t, func() {
		tests := map[string]struct {
			arg      *Nmi
			expected string
		}{
			"valid nmi": {
				arg:      &Nmi{Identifier: "4123456789"},
				expected: "Nmi{Identifier: \"4123456789\"}",
			},
			"invalid nmi": {
				arg:      &Nmi{Identifier: "-1"},
				expected: "Nmi{Identifier: \"-1\"}",
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

func TestNmiString(t *testing.T) {
	Convey("String()", t, func() {
		tests := map[string]struct {
			arg      *Nmi
			expected string
		}{
			"valid nmi": {
				arg:      &Nmi{Identifier: "4123456789"},
				expected: "4123456789",
			},
			"invalid nmi": {
				arg:      &Nmi{Identifier: "-1"},
				expected: "-1",
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

func TestNmiChecksum(t *testing.T) {
	Convey("nmi.Checksum()", t, func() {
		tests := map[string]struct {
			arg      string
			expected int
		}{
			"2001985732": {arg: "2001985732", expected: 8},
			"QAAAVZZZZZ": {arg: "QAAAVZZZZZ", expected: 3},
			"2001985733": {arg: "2001985733", expected: 6},
			"QCDWW00010": {arg: "QCDWW00010", expected: 2},
			"3075621875": {arg: "3075621875", expected: 8},
			"SMVEW00085": {arg: "SMVEW00085", expected: 8},
			"3075621876": {arg: "3075621876", expected: 6},
			"VAAA000065": {arg: "VAAA000065", expected: 7},
			"4316854005": {arg: "4316854005", expected: 9},
			"VAAA000066": {arg: "VAAA000066", expected: 5},
			"4316854006": {arg: "4316854006", expected: 7},
			"VAAA000067": {arg: "VAAA000067", expected: 2},
			"6305888444": {arg: "6305888444", expected: 6},
			"VAAASTY576": {arg: "VAAASTY576", expected: 8},
			"6350888444": {arg: "6350888444", expected: 2},
			"VCCCX00009": {arg: "VCCCX00009", expected: 1},
			"7001888333": {arg: "7001888333", expected: 8},
			"VEEEX00009": {arg: "VEEEX00009", expected: 1},
			"7102000001": {arg: "7102000001", expected: 7},
			"VKTS786150": {arg: "VKTS786150", expected: 2},
			"NAAAMYS582": {arg: "NAAAMYS582", expected: 6},
			"VKTS867150": {arg: "VKTS867150", expected: 5},
			"NBBBX11110": {arg: "NBBBX11110", expected: 0},
			"VKTS871650": {arg: "VKTS871650", expected: 7},
			"NBBBX11111": {arg: "NBBBX11111", expected: 8},
			"VKTS876105": {arg: "VKTS876105", expected: 7},
			"NCCC519495": {arg: "NCCC519495", expected: 5},
			"VKTS876150": {arg: "VKTS876150", expected: 3},
			"NGGG000055": {arg: "NGGG000055", expected: 4},
			"VKTS876510": {arg: "VKTS876510", expected: 8},
			"IIII001100": {arg: "IIII001100", expected: NMICHECKSUMINVALID},
		}

		for name, tc := range tests {
			tc := tc

			nmi := &Nmi{Identifier: tc.arg}

			Convey(fmt.Sprintf("Given a nmi of '%s'", name), func() {
				got := nmi.Checksum()
				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestNmiChecksumValid(t *testing.T) {
	Convey("nmi.ChecksumValid()", t, func() {
		type arg struct {
			nmi      string
			checksum int
		}

		tests := map[string]struct {
			arg      arg
			expected bool
		}{
			"2001985732 w/ 0": {arg: arg{nmi: "2001985732", checksum: 0}, expected: false},
			"2001985732 w/ 1": {arg: arg{nmi: "2001985732", checksum: 1}, expected: false},
			"2001985732 w/ 2": {arg: arg{nmi: "2001985732", checksum: 2}, expected: false},
			"2001985732 w/ 3": {arg: arg{nmi: "2001985732", checksum: 3}, expected: false},
			"2001985732 w/ 4": {arg: arg{nmi: "2001985732", checksum: 4}, expected: false},
			"2001985732 w/ 5": {arg: arg{nmi: "2001985732", checksum: 5}, expected: false},
			"2001985732 w/ 6": {arg: arg{nmi: "2001985732", checksum: 6}, expected: false},
			"2001985732 w/ 7": {arg: arg{nmi: "2001985732", checksum: 7}, expected: false},
			"2001985732 w/ 8": {arg: arg{nmi: "2001985732", checksum: 8}, expected: true},
			"2001985732 w/ 9": {arg: arg{nmi: "2001985732", checksum: 9}, expected: false},
			"IIII001100 w/ 0": {arg: arg{nmi: "IIII001100", checksum: 0}, expected: false},
		}

		for name, tc := range tests {
			tc := tc

			nmi := &Nmi{Identifier: tc.arg.nmi}

			Convey(fmt.Sprintf("Given a nmi of '%s'", name), func() {
				got := nmi.ChecksumValid(tc.arg.checksum)
				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestNmiAllMeters(t *testing.T) {
	Convey("nmi.AllMeters()", t, func() {
		tests := map[string]struct {
			arg      *Nmi
			expected []*Meter
			err      error
		}{
			"nmi with no meters": {
				arg: &Nmi{
					Identifier: "2001985732",
					Meters:     Meters{},
				},
				expected: []*Meter{},
			},
			"nmi with meters": {
				arg: &Nmi{
					Identifier: "2001985732",
					Meters: Meters{
						"1": {Identifier: "1"},
						"2": {Identifier: "2"},
					},
				},
				expected: []*Meter{
					{Identifier: "1"},
					{Identifier: "2"},
				},
			},
			"nmi with meters and a nil meter": {
				arg: &Nmi{
					Identifier: "2001985732",
					Meters: Meters{
						"1":   {Identifier: "1"},
						"2":   {Identifier: "2"},
						"nil": nil,
					},
				},
				expected: []*Meter{
					{Identifier: "1"},
					{Identifier: "2"},
				},
			},
			"nmi is nil ": {arg: nil, err: ErrNmiNil},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given '%s'", name), func() {
				got, err := tc.arg.AllMeters()

				if tc.err != nil {
					So(got, ShouldBeZeroValue)
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					t.Log("")
					t.Logf("expected: %#v", tc.expected)
					t.Logf("got: %#v", got)

					So(err, ShouldBeNil)
					So(got, ShouldResemble, tc.expected)
				}
			})
		}
	})
}

func TestNmiAddMeter(t *testing.T) {
	Convey("nmi.AddMeter(*Meter)", t, func() {
		type arg struct {
			nmi   *Nmi
			meter *Meter
		}
		tests := map[string]struct {
			arg      arg
			expected Meters
			err      error
		}{
			"nmi adding nil meter": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters:     Meters{},
					},
					meter: nil,
				},
				expected: Meters{},
				err:      ErrMeterNil,
			},
			"nmi adding meter with empty identifier": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters:     Meters{},
					},
					meter: &Meter{Identifier: ""},
				},
				expected: Meters{},
				err:      ErrNmiMeterIdentifierEmpty,
			},
			"nmi adding meter that already exists": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters: Meters{
							"1": {Identifier: "1"},
							"2": {Identifier: "2"},
						},
					},
					meter: &Meter{Identifier: "1"},
				},
				expected: Meters{
					"1": {Identifier: "1"},
					"2": {Identifier: "2"},
				},
				err: ErrNmiMeterFound,
			},
			"nmi adding new meter": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters: Meters{
							"1": {Identifier: "1"},
							"2": {Identifier: "2"},
						},
					},
					meter: &Meter{Identifier: "3"},
				},
				expected: Meters{
					"1": {Identifier: "1"},
					"2": {Identifier: "2"},
					"3": {Identifier: "3"},
				},
			},
			"nmi is nil ": {arg: arg{}, err: ErrNmiNil},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given '%s'", name), func() {
				nmi := tc.arg.nmi

				t.Logf("arg: %#v", tc.arg)

				err := nmi.AddMeter(tc.arg.meter)
				if tc.err != nil {
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
					So(nmi.Meters, ShouldResemble, tc.expected)
				}
			})
		}
	})
}

func TestNmiRemoveMeter(t *testing.T) {
	Convey("nmi.RemoveMeter(*Meter)", t, func() {
		type arg struct {
			nmi   *Nmi
			meter *Meter
		}
		tests := map[string]struct {
			arg      arg
			expected Meters
			err      error
		}{
			"nmi is nil ": {arg: arg{}, err: ErrNmiNil},
			"nmi removing nil meter": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters:     Meters{},
					},
					meter: nil,
				},
				expected: Meters{},
				err:      ErrMeterNil,
			},
			"nmi removing meter with empty identifier": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters:     Meters{},
					},
					meter: &Meter{Identifier: ""},
				},
				expected: Meters{},
				err:      ErrNmiMeterIdentifierEmpty,
			},
			"nmi removing meter that does not exist": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters: Meters{
							"1": {Identifier: "1"},
							"2": {Identifier: "2"},
						},
					},
					meter: &Meter{Identifier: "3"},
				},
				expected: Meters{
					"1": {Identifier: "1"},
					"2": {Identifier: "2"},
				},
				err: ErrNmiMeterNotFound,
			},
			"nmi removing meter that exists": {
				arg: arg{
					nmi: &Nmi{
						Identifier: "2001985732",
						Meters: Meters{
							"1": {Identifier: "1"},
							"2": {Identifier: "2"},
						},
					},
					meter: &Meter{Identifier: "1"},
				},
				expected: Meters{
					"2": {Identifier: "2"},
				},
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given '%s'", name), func() {
				nmi := tc.arg.nmi

				t.Logf("arg: %#v", tc.arg)

				err := nmi.RemoveMeter(tc.arg.meter)
				if tc.err != nil {
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
					So(nmi.Meters, ShouldResemble, tc.expected)
				}
			})
		}
	})
}

func TestNmiParticipant(t *testing.T) {
	Convey("nmi.Participant()", t, func() {
		tests := map[string]struct {
			arg      *Nmi
			expected Participant
			err      error
		}{
			"2001985732": {arg: &Nmi{Identifier: "2001985732"}, expected: ParticipantUMPLP},
			"IIII001100": {arg: &Nmi{Identifier: "IIII001100"}, err: ErrParticipantInvalid},
			"nil":        {err: ErrNmiNil},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a nmi of '%s'", name), func() {
				got, err := tc.arg.Participant()

				if tc.err != nil {
					So(got, ShouldBeZeroValue)
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
					So(got, ShouldEqual, tc.expected)
				}
			})
		}
	})
}

func TestNmiRegion(t *testing.T) {
	Convey("nmi.Region()", t, func() {
		tests := map[string]struct {
			arg      *Nmi
			expected region.Region
			err      error
		}{
			"2001985732": {arg: &Nmi{Identifier: "2001985732"}, expected: region.RegionSA},
			"IIII001100": {arg: &Nmi{Identifier: "IIII001100"}, err: ErrParticipantInvalid},
			"nil":        {arg: nil, err: ErrNmiNil},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a nmi of '%s'", name), func() {
				got, err := tc.arg.Region()

				if tc.err != nil {
					So(got, ShouldBeZeroValue)
					So(err, ShouldBeError)
					So(errors.As(err, &tc.err), ShouldBeTrue)
				} else {
					So(err, ShouldBeNil)
					So(got, ShouldEqual, tc.expected)
				}
			})
		}
	})
}
