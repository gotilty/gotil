package errs

import (
	"fmt"
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

func CustomError(message string) error {
	return fmt.Errorf(message)
}
