package maybe

type Option interface {
	Empty() bool
}

func None() Option {
	return emptyOption{}
}

type emptyOption struct{}

func (e emptyOption) Empty() bool { return true }

func (e emptyOption) String() string {
	return "<EMPTY>"
}

