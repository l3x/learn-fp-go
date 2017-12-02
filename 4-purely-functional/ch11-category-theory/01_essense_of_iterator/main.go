package main

import "fmt"

type Car struct {
	Make  string
	Model  string
}
type MapFunc func(string) string
func (car Car) Upgrade() *Car {
	car.Model += " LX"
	return &car
}

func main() {
	is250 := &Car{"Lexus", "IS250"}
	accord := &Car{"Honda", "Accord"}
	cars := []*Car{is250, accord}
	upgradedCars := []*Car{}
	count := 0
	for _, car := range cars {
		upgradedCars = append(upgradedCars, car.Upgrade())
		count ++
	}
	for _, upgradedCar := range upgradedCars {
		fmt.Printf("upgradedCar: %+v\n", upgradedCar)
	}
}
