package functor

import (
	"fmt"
)

type ClockFunctor interface {
	Map(f func(int) int) ClockFunctor
}

type hourContainer struct {
	hours []int
}

func (box hourContainer) Map(f func(int) int) ClockFunctor {
	for i, el := range box.hours {
		box.hours[i] = f(el)
	}
	return box
}

func Wrap(hours []int) ClockFunctor {
	return hourContainer{hours: hours}
}

func (hc hourContainer) String() string {
	return fmt.Sprintf("%+v", hc.hours)
}
