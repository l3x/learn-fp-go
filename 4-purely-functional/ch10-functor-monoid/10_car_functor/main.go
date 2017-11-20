package main

import (
	"encoding/json"
	"fmt"
	"functor"
	"strings"
)

func main() {

	cars := []functor.Car{
		{"Honda", "Accord"},
		{"Lexus", "IS250"}}

	str := `{"make": "Toyota", "model": "Highlander"}`
	highlander := functor.Car{}
	json.Unmarshal([]byte(str), &highlander)
	cars = append(cars, highlander)

	fmt.Println("initial state   :", functor.Wrap(cars))
	fmt.Println("unit application:", functor.Wrap(cars).Map(functor.Zero))
	fmt.Println("one  upgrade    :", functor.Wrap(cars).Map(functor.Upgrade))
	fmt.Println("chain upgrades  :", functor.Wrap(cars).Map(functor.Upgrade).Map(functor.Upgrade))
	fmt.Println("one downgrade   :", functor.Wrap([]functor.Car{{"Honda", "Accord"}, {"Lexus", "IS250 LX"}, {"Toyota", "Highlander LX Limited"}}).Map(functor.Downgrade))

	fmt.Println("up and downgrade:", functor.Wrap(cars).Map(functor.Upgrade).Map(functor.Downgrade))

	cars2 := []functor.Car{}
	for _, car := range cars {
		if !strings.Contains(car.Model, " LX") {
			car.Model += " LX"
		} else if !strings.Contains(car.Model, " Limited") {
			car.Model += " Limited"
		}
		cars2 = append(cars2, car)
	}
	cars3 := []functor.Car{}
	for _, car := range cars2 {
		if strings.Contains(car.Model, " Limited") {
			car.Model = strings.Replace(car.Model, " Limited", "", -1)
		} else if strings.Contains(car.Model, " LX") {
			car.Model = strings.Replace(car.Model, " LX", "", -1)
		}
		cars3 = append(cars3, car)
	}
	fmt.Println("up and downgrade:", cars3)
}
