package main

import (
	. "03_car/car"
	"fmt"
)

func main() {
	accord := Car{"Honda", "Accord"}
	accord.PrintInfo()
	highlander := CarWithSpare{Car{"Toyota", "Highlander"}}
	highlander.PrintInfo()
	fmt.Printf("%v has %d tires\n", highlander.Car, highlander.Tires())
	accord.PrintInfo()
}
