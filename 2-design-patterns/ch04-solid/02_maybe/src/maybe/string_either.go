package maybe

import (
	"fmt"
)

type StringEither interface {
	SuccessOrFailure
	Succeeded() StringOption
	Failed() ErrorOption
}

type eitherString struct {
	val string
	err error
}

func EitherSuccess(i string) StringEither {
	return eitherString{val: i}
}

func EitherFailure(err error) StringEither {
	return eitherString{err: err}
}

func (e eitherString) Success() bool {
	return e.err == nil
}

func (e eitherString) Failure() bool {
	return e.err != nil
}

func (e eitherString) Succeeded() StringOption {
	if e.err == nil {
		return SomeString(e.val)
	}
	return EmptyString()
}

func (e eitherString) Failed() ErrorOption {
	if e.err != nil {
		return SomeError(e.err)
	}
	return EmptyError()
}

func (e eitherString) String() string {
	if e.Success() {
		return fmt.Sprintf("Succeeded(%s)", e.val)
	}
	return fmt.Sprintf("Failed(%s)", e.err)
}
