package nmi

import (
	"errors"
	"fmt"
)

var (
	// ErrIsDuplicated if is duplicated.
	ErrIsDuplicated = errors.New("is duplicated")
	// ErrIsInvalid if is invalid.
	ErrIsInvalid = errors.New("is invalid")
	// ErrIsMissing if is missing.
	ErrIsMissing = errors.New("is missing")
	// ErrIsNil if is nil.
	ErrIsNil = errors.New("is nil")
	// ErrIsNotFound if is not found.
	ErrIsNotFound = errors.New("is not found")
	// ErrParseFailed if parse has failed.
	ErrParseFailed = errors.New("parse has failed")

	// ErrEnergyInvalid if energy is invalid.
	ErrEnergyInvalid = fmt.Errorf("energy %w", ErrIsInvalid)
	// ErrNmiInvalid if nmi is invalid.
	ErrNmiInvalid = fmt.Errorf("nmi %w", ErrIsInvalid)
	// ErrNmiInvalidChar if nmi has invalid character.
	ErrNmiInvalidChar = fmt.Errorf("nmi has character that %w", ErrIsInvalid)
	// ErrNmiInvalidLength if nmi has invalid length.
	ErrNmiInvalidLength = fmt.Errorf("nmi length %w", ErrIsInvalid)
	// ErrMeterNil if meter is nil.
	ErrMeterNil = fmt.Errorf("meter %w", ErrIsNil)
	// ErrNmiMeterFound if nmi meter is found.
	ErrNmiMeterFound = errors.New("nmi meter is found")
	// ErrNmiMeterNotFound if nmi meter is not found.
	ErrNmiMeterNotFound = fmt.Errorf("nmi meter %w", ErrIsNotFound)
	// ErrNmiMeterIdentifierEmpty if nmi meter identifier is empty.
	ErrNmiMeterIdentifierEmpty = fmt.Errorf("nmi meter identifier %w", ErrIsMissing)
	// ErrNmiNil if nmi is nil.
	ErrNmiNil = fmt.Errorf("nmi %w", ErrIsNil)
	// ErrNmiParticipantNotFound if nmi participant is not found.
	ErrNmiParticipantNotFound = fmt.Errorf("nmi participant %w", ErrIsNotFound)
	// ErrParticipantInvalid if participant is invalid.
	ErrParticipantInvalid = fmt.Errorf("participant %w", ErrIsInvalid)
	// ErrPatternInvalid if pattern is invalid.
	ErrPatternInvalid = fmt.Errorf("pattern %w", ErrIsInvalid)
	// ErrRegionInvalid if region is invalid.
	ErrRegionInvalid = fmt.Errorf("region %w", ErrIsInvalid)
)
