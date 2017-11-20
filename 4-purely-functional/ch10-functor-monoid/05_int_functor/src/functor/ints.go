package functor

import (
	"fmt"
)

type IntFunctor interface {
	Map(f func(int) int) IntFunctor
}

type intBox struct {
	ints []int
}

func (box intBox) Map(f func(int) int) IntFunctor {
	for i, el := range box.ints {
		box.ints[i] = f(el)
	}
	return box
}

func Functor(ints []int) IntFunctor {
	return intBox{ints: ints}
}

func (box intBox) String() string {
	return fmt.Sprintf("%+v", box.ints)
}
