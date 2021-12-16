package nem12

import (
	"fmt"
	"strings"
)

const (
	// MethodTypeUndefined for undefined method types.
	MethodTypeUndefined MethodType = iota
	// MethodTypeEstimated for methods using estimation.
	MethodTypeEstimated
	// MethodTypeSubstituted for methods using substitution.
	MethodTypeSubstituted
)

var (
	// methodTypes lists all valid method types.
	methodTypes = []MethodType{ //nolint:gochecknoglobals
		MethodTypeEstimated,
		MethodTypeSubstituted,
	}

	// MethodTypeName maps method types to their string version.
	MethodTypeName = map[MethodType]string{ //nolint:gochecknoglobals
		MethodTypeEstimated:   "EST",
		MethodTypeSubstituted: "SUB",
	}

	// MethodTypeValue maps method types from their string version.
	MethodTypeValue = map[string]MethodType{ //nolint:gochecknoglobals
		"EST": MethodTypeEstimated,
		"SUB": MethodTypeSubstituted,
	}

	// methodTypeDescriptions maps method types from their descriptions.
	methodTypeDescriptions = map[MethodType]string{ //nolint:gochecknoglobals
		MethodTypeEstimated:   "estimated",
		MethodTypeSubstituted: "substituted",
	}
)

// A MethodType represents the value of the method flag section of a QualityMethodType field
// of a NEM12 interval value.
type MethodType int

// MethodTypes returns all method types.
func MethodTypes() []MethodType {
	return methodTypes
}

// NewMethodType returns a new method flag if valid, and an error if not.
func NewMethodType(s string) (MethodType, error) {
	if s == "" {
		return MethodTypeUndefined, ErrMethodTypeNil
	}

	m, ok := MethodTypeValue[strings.ToUpper(s)]

	if !ok {
		return m, ErrMethodTypeInvalid
	}

	return m, nil
}

// Validate returns an error if the method flag is invalid.
func (m MethodType) Validate() error {
	if _, ok := MethodTypeName[m]; !ok {
		return ErrMethodTypeInvalid
	}

	return nil
}

// Identifier to meet the interface specification for a Flag.
func (m MethodType) Identifier() string {
	id, ok := MethodTypeName[m]
	if !ok {
		return fmt.Sprintf("MethodType(%d)", m)
	}

	return id
}

// Description returns the description of a method flag, along with an error if it is an unknown value.
func (m MethodType) Description() (string, error) {
	desc, ok := methodTypeDescriptions[m]
	if !ok {
		return fmt.Sprintf("%%!MethodType(%d)", m), fmt.Errorf("method type description '%d': %w", m, ErrMethodTypeInvalid)
	}

	return desc, nil
}

// String returns a text representation of the MethodType.
func (m MethodType) String() string {
	s, _ := m.Description()

	return fmt.Sprintf("\"%s: %s\"", m.Identifier(), s)
}

// GoString returns a text representation of the MethodType to satisfy the GoStringer
// interface.
func (m MethodType) GoString() string {
	return fmt.Sprintf("MethodType(%d)", m)
}
