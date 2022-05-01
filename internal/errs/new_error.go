package errs

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedType = errors.New("unsupported type.")
	ErrEmptyString     = errors.New("the given string is empty.")
	ErrNilReference    = errors.New("the given parameter is nil.")
	ErrOutOfRange      = errors.New("converter: out of range.")
)

func NewUnsupportedTypeError(kind string) error {
	return fmt.Errorf("the given parameter(%s) is unsupported type ", kind)
}

func NilReferenceTypeError() error {
	return fmt.Errorf("the given parameter is nil ")
}

func EmptyStringError() error {
	return fmt.Errorf("the given string is empty")
}

func Int64Error(kind string) error {
	return fmt.Errorf("the entered value cannot convert from %s to int64", kind)
}

func Int32Error(kind string) error {
	return fmt.Errorf("the entered value cannot convert from %s to int32", kind)
}

func Int16Error(kind string) error {
	return fmt.Errorf("the entered value cannot convert from %s to int16", kind)
}

func CustomError(message string) error {
	return fmt.Errorf(message)
}

func NewUnsupportedParameterTypeError(kind string, desc string) error {
	return fmt.Errorf("the given parameter/s(%s) is unsupported type. %s", kind, desc)
}

func NewMissingParameterError(kind string, desc string) error {
	return fmt.Errorf("the given parameter/s(%s) is missing. %s", kind, desc)
}
