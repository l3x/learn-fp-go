package main

import (
	"log"
	"functor"
)


func main() {

	amPmMapper := func(i int) int {
		return (i + 12) % 24
	}

	log.Printf("initial state     : %s", fun.Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}))

	unit := func(i int) int {
		return i
	}

	log.Printf("unit application  : %s", fun.Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}).Map(unit))

	log.Printf("1st application   : %s", fun.Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}).Map(amPmMapper))

	log.Printf("chain applications: %s", fun.Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}).Map(amPmMapper).Map(amPmMapper))

	log.Printf("chain applications: %s", fun.Wrap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}).Map(amPmMapper))
}

