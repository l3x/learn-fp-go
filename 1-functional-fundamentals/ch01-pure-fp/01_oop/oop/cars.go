package oop

import (
	"errors"
)

type Car struct {
	Model string
}
type Cars []Car
var MyCars Cars

func (cars *Cars) Add(car Car) {
	MyCars = append(MyCars, car)
}

func (cars *Cars) Find(model string) (*Car, error) {
	for _, car := range *cars {
		if car.Model == model {
			return &car, nil
		}
	}
	return nil, errors.New("car not found")
}
