package functor

import (
	"fmt"
	"strings"
)

type (
	Car struct {
		Make  string `json:"make"`
		Model string `json:"model"`
	}
)

type CarFunctor interface {
	Map(f func(Car) Car) CarFunctor
}

type carContainer struct {
	cars []Car
}

func (box carContainer) Map(f func(Car) Car) CarFunctor {
	for i, el := range box.cars {
		box.cars[i] = f(el)
	}
	return box
}

func Wrap(cars []Car) CarFunctor {
	return carContainer{cars: cars}
}

var (
	Zero = func(i Car) Car {
		return i
	}

	Upgrade = func(car Car) Car {
		if !strings.Contains(car.Model, " LX") {
			car.Model += " LX"
		} else if !strings.Contains(car.Model, " Limited") {
			car.Model += " Limited"
		}
		return car
	}

	Downgrade = func(car Car) Car {
		if strings.Contains(car.Model, " Limited") {
			car.Model = strings.Replace(car.Model, " Limited", "", -1)
		} else if strings.Contains(car.Model, " LX") {
			car.Model = strings.Replace(car.Model, " LX", "", -1)
		}
		return car
	}
)

func (box carContainer) String() string {
	return fmt.Sprintf("%+v", box.cars)
}
