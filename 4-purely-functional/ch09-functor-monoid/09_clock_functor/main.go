package main

import (
	. "functor"
	"fmt"
)

func main() {
	fmt.Println("Initial state     :", Wrap(AmHoursFn()))
	fmt.Println("Zero application  :", Wrap(AmHoursFn()).Map(Zero))
	fmt.Println("1st application   :", Wrap(AmHoursFn()).Map(AmPmMapper))
	fmt.Println("Chain applications:", Wrap(AmHoursFn()).Map(AmPmMapper).Map(AmPmMapper))
	fmt.Println("Chain applications:", Wrap(AmHoursFn()).Map(AmPmMapper))
}
