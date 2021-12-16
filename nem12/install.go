package nem12

import (
	"fmt"
	"strings"
)

// Reference: https://aemo.com.au/-/media/archive/files/other/msats/msats-procedures-cats-v4-0-final-determination-v1.pdf
// page 43.

const (
	// InstallUndefined if install is undefined.
	InstallUndefined Install = iota
	// InstallComms1 for a type 1 interval meter with communications .
	InstallComms1
	// InstallComms2 for a type 2 interval meter with communications .
	InstallComms2
	// InstallComms3 for a type 3 interval meter with communications .
	InstallComms3
	// InstallComms4 for a type 4 interval meter with communications .
	InstallComms4
	// InstallMRIM for a type 5 manually read interval meter.
	InstallMRIM
	// InstallBasic for a type 6 basic consumption meter.
	InstallBasic
	// InstallUnmetered for type 7 unmetered supply.
	InstallUnmetered
)

var (
	installs = []Install{ //nolint:gochecknoglobals
		InstallComms1,
		InstallComms2,
		InstallComms3,
		InstallComms4,
		InstallMRIM,
		InstallBasic,
		InstallUnmetered,
	}

	// InstallName maps an install to its name.
	InstallName = map[Install]string{ //nolint:gochecknoglobals
		InstallComms1:    "COMMS1",
		InstallComms2:    "COMMS2",
		InstallComms3:    "COMMS3",
		InstallComms4:    "COMMS4",
		InstallMRIM:      "MRIM",
		InstallBasic:     "BASIC",
		InstallUnmetered: "UMCP",
	}

	// InstallValue maps an install from its name.
	InstallValue = map[string]Install{ //nolint:gochecknoglobals
		"COMMS1": InstallComms1,
		"COMMS2": InstallComms2,
		"COMMS3": InstallComms3,
		"COMMS4": InstallComms4,
		"MRIM":   InstallMRIM,
		"BASIC":  InstallBasic,
		"UMCP":   InstallUnmetered,
	}

	installDescriptions = map[Install]string{ //nolint:gochecknoglobals
		InstallComms1:    "Interval meter with communications – Type 1",
		InstallComms2:    "Interval meter with communications – Type 2",
		InstallComms3:    "Interval meter with communications – Type 3",
		InstallComms4:    "Interval meter with communications – Type 4",
		InstallMRIM:      "Manually Read Interval Meter – Type 5",
		InstallBasic:     "Basic Consumption Meter – Type 6",
		InstallUnmetered: "Unmetered Supply – Type 7 ",
	}
)

// Install represents an installation type for the metering at a nmi.
type Install int

// Installs returns all the installs.
func Installs() []Install {
	return installs
}

// NewInstall returns a new install, and error if not found.
func NewInstall(s string) (Install, error) {
	if s == "" {
		return InstallUndefined, ErrInstallNil
	}

	i, ok := InstallValue[strings.ToUpper(s)]
	if !ok {
		return InstallUndefined, ErrInstallInvalid
	}

	return i, nil
}

// Identifier returns the code used by AEMO.
func (i Install) Identifier() string {
	str, ok := InstallName[i]
	if !ok {
		return fmt.Sprintf("%%!Install(%d)", i)
	}

	return str
}

// Description returns the description of the install from AEMO.
func (i Install) Description() (string, error) {
	desc, ok := installDescriptions[i]
	if !ok {
		return fmt.Sprintf("%%!Install(%d)", i), ErrInstallInvalid
	}

	return desc, nil
}

// String returns a text representation of the install.
func (i Install) String() string {
	desc, err := i.Description()
	if err != nil {
		return desc
	}

	return fmt.Sprintf("%s: %s", i.Identifier(), desc)
}

// GoString returns a text representation of the install to satisfy the GoStringer
// interface.
func (i Install) GoString() string {
	return fmt.Sprintf("Install(%d)", i)
}
