package nem12

import "github.com/pkg/errors"

var (
	// ErrInstallInvalid if install is invalid.
	ErrInstallInvalid = errors.New("install is invalid")
	// ErrMethodEmpty if method flag is empty.
	ErrMethodEmpty = errors.New("method flag is empty")
	// ErrMethodInvalid if method flag is invalid.
	ErrMethodInvalid = errors.New("method flag is invalid")
	// ErrMethodTypeEmpty if method type is empty.
	ErrMethodTypeEmpty = errors.New("method type is empty")
	// ErrMethodTypeInvalid if method type is invalid.
	ErrMethodTypeInvalid = errors.New("method type is invalid")
	// ErrQualityEmpty if quality flag is empty.
	ErrQualityEmpty = errors.New("quality flag is empty")
	// ErrQualityInvalid if quality flag is invalid.
	ErrQualityInvalid = errors.New("quality flag is invalid")
	// ErrQualityMethodEmpty if quality method is empty.
	ErrQualityMethodEmpty = errors.New("quality method is empty")
	// ErrQualityMethodInvalid if quality method is invalid.
	ErrQualityMethodInvalid = errors.New("quality method is invalid")
	// ErrQualityMethodInvalidLength if quality flag has invalid length.
	ErrQualityMethodInvalidLength = errors.New("quality method has invalid length")
	// ErrQualityMissingMethod if quality method missing required method.
	ErrQualityMissingMethod = errors.New("quality method missing required method")
	// ErrReasonCodeEmpty if reason code is empty.
	ErrReasonCodeEmpty = errors.New("reason code is empty")
	// ErrReasonCodeInvalid if reason code is invalid.
	ErrReasonCodeInvalid = errors.New("reason code is invalid")
	// ErrUnitOfMeasureEmpty if unit of measure is empty.
	ErrUnitOfMeasureEmpty = errors.New("unit of measure is empty")
	// ErrUnitOfMeasureInvalid if unit of measure is invalid.
	ErrUnitOfMeasureInvalid = errors.New("unit of measure is invalid")
)
