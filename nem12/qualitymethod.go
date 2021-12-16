package nem12

import "fmt"

const (
	methodLength  = 2
	qualityLength = 1
)

// QualityMethod of the NEM12 interval data record (300) or interval event record
// (400) row.
type QualityMethod string

// NewQualityMethod returns a new quality method, with error if not valid.
func NewQualityMethod(s string) (QualityMethod, error) {
	qm := QualityMethod(s)

	if err := qm.Validate(); err != nil {
		return qm, err
	}

	return qm, nil
}

// Validate returns an error if a quality method is invalid.
func (qm QualityMethod) Validate() error {
	strLen := len(qm)

	if strLen == 0 {
		return ErrQualityMethodNil
	}

	if strLen != qualityLength && strLen != (qualityLength+methodLength) {
		return ErrQualityMethodLengthInvalid
	}

	q, err := qm.Quality()
	if err != nil {
		return fmt.Errorf("validate '%s': %w", qm, err)
	}

	if strLen == qualityLength && (q == QualityActual || q == QualityNull || q == QualityVariable) {
		return nil
	}

	if strLen != (qualityLength + methodLength) {
		return fmt.Errorf("validate '%s': %w", qm, ErrQualityMissingMethod)
	}

	if _, err := qm.Method(); err != nil {
		return fmt.Errorf("validate '%s': %w", qm, err)
	}

	return nil
}

// Quality returns the quality for a given QualityMethod.
func (qm QualityMethod) Quality() (Quality, error) {
	q, err := NewQualityFlag(string(qm)[0:1])
	if err != nil {
		err = fmt.Errorf("quality from '%s': %w", qm, err)
	}

	return q, err
}

// Method returns the quality for a given Method.
func (qm QualityMethod) Method() (Method, error) {
	if len(qm) < (qualityLength + methodLength) {
		q, err := qm.Quality()
		if err != nil {
			return MethodUndefined, err
		}

		switch q {
		case QualityActual, QualityNull, QualityVariable:
			return MethodUndefined, nil

		case QualityEstimated, QualitySubstituted, QualityFinal:
			return MethodUndefined, ErrQualityMissingMethod

		case QualityUndefined:
			fallthrough //nolint:gocritic

		default:
			return MethodUndefined, ErrQualityMissingMethod
		}
	}

	m, err := NewMethodFlag(string(qm)[1:3])
	if err != nil {
		err = fmt.Errorf("method for '%s': %w", qm, err)
	}

	return m, err
}
