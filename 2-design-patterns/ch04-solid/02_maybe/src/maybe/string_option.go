package maybe

import "fmt"

type StringOption interface {
	Option
	Val() string
}

type optionalString struct {
	exists bool
	val    string
}

func EmptyString() StringOption {
	return optionalString{exists: false}
}

func SomeString(val string) StringOption {
	return optionalString{exists: true, val: val}
}

func (o optionalString) Empty() bool {
	return !o.exists
}

func (o optionalString) Val() string {
	return o.val
}

func (o optionalString) String() string {
	if o.Empty() {
		return "<EMPTY>"
	}
	return fmt.Sprintf("%s", o.val)
}
