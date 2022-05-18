package hof

import (
	"fmt"
	s "strings"
	"regexp"
	"log"
	"encoding/json"
)

var CarsDB = initCarsDB()

func initCarsDB() []IndexedCar {
	var indexedCars []IndexedCar
	for i, car := range LoadCars() {
		indexedCars = append(indexedCars, IndexedCar{i, car})
	}
	lenCars := len(indexedCars)
	for i, car := range LoadMoreCars() {
		indexedCars = append(indexedCars, IndexedCar{i + lenCars, car})
	}
	return indexedCars
}

func LoadCars() Collection {
	return CsvToStruct("cars.csv")
}

func LoadMoreCars() Collection {
	return CsvToStruct("more_cars.csv")
}


func (cars Collection) AddCars(carsToAdd Collection) Collection {
	return append(cars, carsToAdd...)
}


// Reduce collection based on the function
func (cars Collection) Reduce(fn ReducerFunc, accumulator Collection) Collection {
	var result = accumulator
	for _, car := range cars {
		result = append(fn(car, result))
	}
	return result
}
func (cars Collection) Reduce2(fn ReducerFunc2, accumulator CarCollection) CarCollection {
	var result = accumulator
	for _, car := range cars {
		result = append(fn(car, result))
	}
	return result
}

func JsonReducer(cars Collection) ReducerFunc  {
	return func(car string, cars Collection) Collection {
		JSON := fmt.Sprintf("{\"car\": {\"make\": \"%s\", \"model\": \"%s\"}}", GetMake(car), GetModel(car))
		cars = append(cars, JSON)
		return cars
	}
}

func CarTypeReducer(cars Collection) ReducerFunc2 {

	return func(car string, cars CarCollection) CarCollection {
		JSON := fmt.Sprintf("{\"make\": \"%s\", \"model\": \"%s\"}", GetMake(car), GetModel(car))
		var c CarType
		err := json.Unmarshal([]byte(JSON), &c)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
		cars = append(cars, c)
		return cars
	}
}

func MakeReducer(make string, cars Collection) ReducerFunc  {
	return func(car string, cars Collection) Collection {
		if s.Contains(car, make) {
			cars = append(cars, car)
		}
		return cars
	}
}

func (cars Collection) Filter(fn FilterFunc) Collection {
	filteredCars := make(Collection, 0)
	for _, car := range cars {
		if fn(car) {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

func ByMake(make string) FilterFunc {
	return func(car string) bool {
		return s.Contains(car, make)
	}
}

func ByForeign() FilterFunc {
	return func(car string) bool {
		return !isDomestic(car)
	}
}

func isDomestic(car string) bool {
	return s.Contains(car, "Ford") || s.Contains(car, "GM") || s.Contains(car, "Chrysler")
}

func ByDomestic() FilterFunc {
	return func(car string) bool {
		return isDomestic(car)
	}
}

func ByHasNumber() FilterFunc {
	return func(car string) bool {
		match, _ := regexp.MatchString(".+[0-9].*", car)
		return match
	}
}

func ByModel(model string) FilterFunc {
	return func(car string) bool {
		return s.Contains(car, model)
	}
}

func (cars Collection) Map(fn MapFunc) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car))
	}
	return mappedCars
}

func Upgrade() MapFunc {
	return func(car string) string {
		return fmt.Sprintf("%s %s", car, UpgradeLabel(car))
	}
}

func Downgrade() MapFunc {
	return func(car string) string {
		return car[:len(car)-2]
	}
}

func UpgradeLabel(car string) string {
	return map[string]string{
		"Honda": "LX",
		"Lexus": "LS",
		"Toyota": "EV",
		"Ford": "XL",
		"GM": "X",
	}[GetMake(car)]
}

func (cars Collection) String() string {
	return s.Join(cars, ",")
}


func x() {
	val, err := buggyCode()
	// more code
	return

}