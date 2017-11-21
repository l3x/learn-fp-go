package main

import (
	"fmt"
	. "car"
)

func main() {
	var cars = CarSlice{
		Car{"Honda", "Accord", 3000},
		Car{"Lexus", "IS250", 40000},
		Car{"Toyota", "Highlander", 3500},
		Car{"Honda", "Accord ES", 3500},
	}
	fmt.Println("cars:", cars)

	honda := func (c Car) bool {
		return c.Make == "Honda"
	}
	fmt.Println("filter cars by 'Honda':", cars.Where(honda))

	price := func (c Car) Dollars {
		return c.Price
	}
	fmt.Println("Hondas prices:", cars.Where(honda).SelectDollars(price))

	fmt.Println("Hondas sum(prices):", cars.Where(honda).SumDollars(price))
}
