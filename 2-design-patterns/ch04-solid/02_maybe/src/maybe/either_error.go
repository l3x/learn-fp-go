package maybe

import "fmt"

type Either interface {
	SuccessOrFailure
	Succeeded() StringOption
	Failed() ErrorOption
}

type either struct {
	val   string
	err error
}

func EitherSuccessError(val string) Either {
	return either{val: val}
}

func EitherFailureError(err error) Either {
	return either{err: err}
}

func (e either) Success() bool {
	return e.err == nil
}

func (e either) Failure() bool {
	return e.err != nil
}

func (e either) Succeeded() StringOption {
	if e.err == nil {
		return SomeString(e.val)
	}
	return EmptyString()
}

func (e either) Failed() ErrorOption {
	if e.err != nil {
		return SomeError(e.err)
	}
	return EmptyError()
}

func (e either) String() string {
	if e.Success() {
		return fmt.Sprintf("Succeeded(%s)", e.val)
	}
	return fmt.Sprintf("Failed(%s)", e.err)
}
