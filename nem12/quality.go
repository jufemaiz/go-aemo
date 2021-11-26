package nem12

import (
	"fmt"
	"strings"
)

const (
	// QualityUndefined is for undefined quality flags.
	QualityUndefined Quality = iota
	// QualityActual is the quality flag value for actual data.
	QualityActual
	// QualityEstimated is the quality flag value for forward estimated data.
	QualityEstimated
	// QualityFinal is the quality flag value for final substituted data.
	QualityFinal
	// QualityNull is the quality flag value for null data.
	QualityNull
	// QualitySubstituted is the quality flag value for substituted data.
	QualitySubstituted
	// QualityVariable is the quality flag value for variable data.
	QualityVariable
)

var (
	// qualities lists all qualities.
	qualities = []Quality{ //nolint:gochecknoglobals
		QualityActual,
		QualityEstimated,
		QualityFinal,
		QualityNull,
		QualitySubstituted,
		QualityVariable,
	}

	// QualityName maps a Quality to its name.
	QualityName = map[Quality]string{ //nolint:gochecknoglobals
		QualityActual:      "A",
		QualityEstimated:   "E",
		QualityFinal:       "F",
		QualityNull:        "N",
		QualitySubstituted: "S",
		QualityVariable:    "V",
	}

	// QualityValue maps a name to its value.
	QualityValue = map[string]Quality{ //nolint:gochecknoglobals
		"A": QualityActual,
		"E": QualityEstimated,
		"F": QualityFinal,
		"N": QualityNull,
		"S": QualitySubstituted,
		"V": QualityVariable,
	}

	// qualityDescriptions provides the descriptions for the quality flags.
	qualityDescriptions = map[Quality]string{ //nolint:gochecknoglobals
		QualityActual:      "actual data",
		QualityEstimated:   "forward estimated data",
		QualityFinal:       "final estimated data",
		QualityNull:        "null data",
		QualitySubstituted: "substituted data",
		QualityVariable:    "variable data",
	}
)

// Quality represents the value of the quality flag part of the QualityMethod field
// of an NEM12 interval.
type Quality int

// Qualities returns a slice of all the qualities.
func Qualities() []Quality {
	return qualities
}

// NewQualityFlag returns a new quality flag if valid, and an error if not.
func NewQualityFlag(s string) (Quality, error) {
	if s == "" {
		return QualityUndefined, ErrQualityNil
	}

	q, ok := QualityValue[strings.ToUpper(s)]

	if !ok {
		return q, ErrQualityInvalid
	}

	return q, nil
}

// Validate returns an error if the quality flag is invalid.
func (q Quality) Validate() (err error) {
	switch q {
	case QualityActual, QualityEstimated, QualityFinal, QualityNull, QualitySubstituted, QualityVariable:
		err = nil
	case QualityUndefined:
		err = ErrQualityInvalid
	default:
		err = ErrQualityInvalid
	}

	return err
}

// Identifier to meet the interface specification for a Flag.
func (q Quality) Identifier() string {
	id, ok := QualityName[q]
	if !ok {
		return fmt.Sprintf("Method(%d)", q)
	}

	return id
}

// GoString returns a text representation of the Quality to satisfy the GoStringer
// interface.
func (q Quality) GoString() string {
	return fmt.Sprintf("Quality(%d)", q)
}

// String returns a text representation of the Quality.
func (q Quality) String() string {
	s, err := q.Description()
	if err != nil {
		return fmt.Sprintf("\"%s\"", q.Identifier())
	}

	return fmt.Sprintf("\"%s: %s\"", q.Identifier(), s)
}

// Description returns the description of a quality flag. Error is returned if the
// flag is invalid.
func (q Quality) Description() (string, error) {
	d, ok := qualityDescriptions[q]

	if !ok {
		return "", ErrQualityInvalid
	}

	return d, nil
}

// MarshalJSON marshals for JSON.
func (q *Quality) MarshalJSON() ([]byte, error) {
	id, ok := QualityName[*q]
	if !ok {
		return []byte(fmt.Sprintf("\"%d\"", *q)), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", id)), nil
}

// UnmarshalJSON unmarshals json string.
func (q *Quality) UnmarshalJSON(data []byte) error {
	v, ok := QualityValue[string(data)]
	if !ok {
		return ErrSuffixTypeInvalid
	}

	*q = v

	return nil
}

// MustNotHaveReason indicates if a quality flag must not have a reason.
func (q Quality) MustNotHaveReason() bool {
	return q == QualityVariable
}

// RequiresMethod indicates if a quality flag requires an accompanying method.
func (q Quality) RequiresMethod() (b bool) {
	switch q {
	case QualityEstimated, QualityFinal, QualitySubstituted:
		b = true
	case QualityActual, QualityNull, QualityVariable, QualityUndefined:
		fallthrough
	default:
		b = false
	}

	return b
}

// RequiresReason indicates if a quality flag requires a reason.
func (q Quality) RequiresReason() (b bool) {
	switch q {
	case QualityFinal, QualitySubstituted:
		b = true
	case QualityActual, QualityNull, QualityEstimated, QualityVariable, QualityUndefined:
		fallthrough
	default:
		b = false
	}

	return b
}
