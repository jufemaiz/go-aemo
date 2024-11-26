package nem12

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInstalls(t *testing.T) {
	Convey("Install", t, func() {
		got := Installs()

		So(got, ShouldHaveLength, len(InstallName))
	})
}

func TestNewInstall(t *testing.T) {
	type test struct {
		arg      string
		expected Install
		err      error
	}

	Convey("NewInstall", t, func() {
		tests := map[string]test{
			"empty string": {
				err: ErrInstallNil,
			},
			"invalid string": {
				arg: "Not an install",
				err: ErrInstallInvalid,
			},
			"lowercase value": {
				arg:      "comms1",
				expected: InstallComms1,
			},
			"mixed case value": {
				arg:      "Comms1",
				expected: InstallComms1,
			},
			"upper case value": {
				arg:      "COMMS1",
				expected: InstallComms1,
			},
		}

		for i, n := range InstallName {
			tests[fmt.Sprintf("install '%s'", n)] = test{arg: n, expected: i}
		}

		for name, tc := range tests {
			tc := tc

			Convey("Given "+name, func() {
				got, err := NewInstall(tc.arg)

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

func TestInstallGoString(t *testing.T) {
	Convey("install.GoString()", t, func() {
		tests := map[string]struct {
			arg      Install
			expected string
		}{
			"UNDEFINED": {
				arg:      InstallUndefined,
				expected: "Install(0)",
			},
			"COMMS1": {
				arg:      InstallComms1,
				expected: "Install(1)",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a install of '%s'", name), func() {
				got := tc.arg.GoString()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestInstallString(t *testing.T) {
	Convey("install.String()", t, func() {
		tests := map[string]struct {
			arg      Install
			expected string
		}{
			"UNDEFINED": {
				arg:      InstallUndefined,
				expected: "%!Install(0)",
			},
			"COMMS1": {
				arg:      InstallComms1,
				expected: "COMMS1: Interval meter with communications – Type 1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a install of '%s'", name), func() {
				got := tc.arg.String()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestInstallIdentifier(t *testing.T) {
	Convey("install.Identifier()", t, func() {
		tests := map[string]struct {
			arg      Install
			expected string
		}{
			"UNDEFINED": {
				arg:      InstallUndefined,
				expected: "%!Install(0)",
			},
			"COMMS1": {
				arg:      InstallComms1,
				expected: "COMMS1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a install of '%s'", name), func() {
				got := tc.arg.Identifier()

				So(got, ShouldEqual, tc.expected)
			})
		}
	})
}

func TestInstallDescription(t *testing.T) {
	Convey("install.Description()", t, func() {
		tests := map[string]struct {
			arg      Install
			expected string
			err      error
		}{
			"UNDEFINED": {
				arg:      InstallUndefined,
				expected: "%!Install(0)",
				err:      ErrInstallInvalid,
			},
			"COMMS1": {
				arg:      InstallComms1,
				expected: "Interval meter with communications – Type 1",
			},
		}

		for name, tc := range tests {
			tc := tc

			Convey(fmt.Sprintf("Given a install of '%s'", name), func() {
				got, err := tc.arg.Description()

				So(got, ShouldEqual, tc.expected)
				So(errors.Is(err, tc.err), ShouldBeTrue)
			})
		}
	})
}
