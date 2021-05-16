package nmi

import "errors"

var (
	// ErrEnergyInvalid if energy is invalid.
	ErrEnergyInvalid = errors.New("energy is invalid")
	// ErrNmiInvalid if nmi is invalid.
	ErrNmiInvalid = errors.New("nmi is invalid")
	// ErrNmiInvalidChar if nmi has invalid character.
	ErrNmiInvalidChar = errors.New("nmi has invalid character")
	// ErrNmiInvalidLength if nmi has invalid length.
	ErrNmiInvalidLength = errors.New("nmi has invalid length")
	// ErrMeterNil if meter is nil
	ErrMeterNil = errors.New("meter is nil")
	// ErrNmiMeterFound if nmi meter is found.
	ErrNmiMeterFound = errors.New("nmi meter is found")
	// ErrNmiMeterNotFound if nmi meter is not found.
	ErrNmiMeterNotFound = errors.New("nmi meter is not found")
	// ErrNmiMeterIdentifierEmpty if nmi meter identifier is empty.
	ErrNmiMeterIdentifierEmpty = errors.New("nmi meter identifier is empty")
	// ErrNmiNil if nmi is nil.
	ErrNmiNil = errors.New("nmi is nil")
	// ErrNmiParticipantNotFound if nmi participant is not found.
	ErrNmiParticipantNotFound = errors.New("nmi participant is not found")
	// ErrParticipantInvalid if participant is invalid.
	ErrParticipantInvalid = errors.New("participant is invalid")
	// ErrPatternInvalid if pattern is invalid.
	ErrPatternInvalid = errors.New("pattern is invalid")
	// ErrRegionInvalid if region is invalid.
	ErrRegionInvalid = errors.New("region is invalid")
)
