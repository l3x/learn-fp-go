package maybe

import (
	"fmt"
	"errors"
)

var NilError = errors.New("nil error occurred")

type ErrorOption interface {
	Option
	Err() error
}

type optionalError struct {
	err error
}

func EmptyError() ErrorOption {
	return optionalError{err: NilError}
}

func SomeError(err error) ErrorOption {
	return optionalError{err: err}
}

func (o optionalError) Empty() bool {
	return o.err == NilError
}

func (o optionalError) Err() error {
	return o.err
}

func (o optionalError) String() string {
	if o.err == NilError {
		return "<EMPTY>"
	}
	return fmt.Sprintf("%s", o.err.Error())
}
