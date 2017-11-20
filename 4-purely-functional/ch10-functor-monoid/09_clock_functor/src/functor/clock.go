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

var Zero = func(i int) int {
	return i
}

var AmPmMapper = func(i int) int {
	return (i + 12) % 24
}

func AmHoursFn()  []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
}

func (box hourContainer) String() string {
	return fmt.Sprintf("%+v", box.hours)
}
