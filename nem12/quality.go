package nem12

import "fmt"

const (
	// QualityUndefined is for undefined quality flags.
	QualityUndefined Quality = iota
	// Actual is the quality flag value for actual data.
	QualityActual
	// Estimated is the quality flag value for forward estimated data.
	QualityEstimated
	// Final is the quality flag value for final substituted data.
	QualityFinal
	// Null is the quality flag value for null data.
	QualityNull
	// Substituted is the quality flag value for substituted data.
	QualitySubstituted
	// Variable is the quality flag value for variable data.
	QualityVariable
)

var (
	qualities = []Quality{ //nolint:gochecknoglobals
		QualityUndefined,
		QualityActual,
		QualityEstimated,
		QualityFinal,
		QualityNull,
		QualitySubstituted,
		QualityVariable,
	}

	// QualityKey maps a Quality to its name.
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

// NewQualityFlag returns a new quality flag if valid, and an error if not.
func NewQualityFlag(s string) (Quality, error) {
	q, ok := QualityValue[s]

	if !ok {
		return q, ErrQualityInvalid
	}

	return q, nil
}

// Validate returns an error if the quality flag is invalid.
func (q Quality) Validate() error {
	switch q {
	case QualityActual, QualityEstimated, QualityFinal, QualityNull, QualitySubstituted, QualityVariable:
		return nil
	}

	if q == QualityUndefined {
		return ErrQualityEmpty
	}

	return ErrQualityInvalid
}

// Identifier to meet the interface specification for a Flag.
func (q Quality) Identifier() string {
	return QualityName[q]
}

// GoString returns a text representation of the Quality to satisfy the GoStringer
// interface.
func (q Quality) GoString() string {
	s, _ := q.Description()

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

// MustNoteHaveReason indicates if a quality flag must not have a reason.
func (q Quality) MustNoteHaveReason() bool {
	return q == QualityVariable
}

// RequiresMethod indicates if a quality flag requires an accompanying method.
func (q Quality) RequiresMethod() bool {
	switch q {
	case QualityActual, QualityNull, QualityVariable:
		return false
	}

	return true
}

// RequiresReason indicates if a quality flag requires a reason.
func (q Quality) RequiresReason() bool {
	switch q {
	case QualityActual, QualityNull, QualityEstimated, QualityVariable:
		return false
	}

	return true
}
