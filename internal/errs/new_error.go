package errs

import (
	"errors"
	"fmt"
)

func NewUnsupportedTypeError(kind string) error {
	return errors.New(fmt.Sprintf("the given parameter(%s) is unsupported type ", kind))
}

func NilReferenceTypeError() error {
	return errors.New(fmt.Sprintf("the given parameter is nil "))
}

func EmptyStringError() error {
	return errors.New(fmt.Sprintf("the given string is empty"))
}
