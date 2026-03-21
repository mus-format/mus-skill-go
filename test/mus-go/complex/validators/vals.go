package validators

import "errors"

var (
	ErrNegativeInt       = errors.New("negative int")
	ErrEmptyString       = errors.New("empty string")
	ErrNegativeSliceElem = errors.New("negative slice element")
	ErrEmptyMapKey       = errors.New("empty map key")
)

func ValidateInt(v int) error {
	if v < 0 {
		return ErrNegativeInt
	}
	return nil
}

func ValidateString(v string) error {
	if len(v) == 0 {
		return ErrEmptyString
	}
	return nil
}

func ValidateSliceElem(v int) error {
	if v < 0 {
		return ErrNegativeSliceElem
	}
	return nil
}

func ValidateMapKey(v string) error {
	if len(v) == 0 {
		return ErrEmptyMapKey
	}
	return nil
}
