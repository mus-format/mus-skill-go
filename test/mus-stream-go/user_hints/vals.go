package userhints

import "errors"

var (
	ErrNegativeNum     = errors.New("negative num")
	ErrEmptyStr        = errors.New("empty str")
	ErrNegativeLength  = errors.New("negative length")
	ErrEmptyKey        = errors.New("empty key")
	ErrNegativeValue   = errors.New("negative value")
	ErrNegativeElement = errors.New("negative element")
	ErrZeroNum         = errors.New("zero num")
)

func ValidateStruct(v Valid) error {
	if v.num == 0 {
		return ErrZeroNum
	}
	return nil
}

func ValidateNum(n int) error {
	if n < 0 {
		return ErrNegativeNum
	}
	return nil
}

func ValidateStr(s string) error {
	if len(s) == 0 {
		return ErrEmptyStr
	}
	return nil
}

func ValidateLength(l int) error {
	if l < 0 {
		return ErrNegativeLength
	}
	return nil
}

func ValidateKey(k string) error {
	if len(k) == 0 {
		return ErrEmptyKey
	}
	return nil
}

func ValidateValue(v int) error {
	if v < 0 {
		return ErrNegativeValue
	}
	return nil
}

func ValidateElement(e int) error {
	if e < 0 {
		return ErrNegativeElement
	}
	return nil
}
