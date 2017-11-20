package hof

import (
	"sync"
	"log"
)


func carGenerator(iterator func(int) int, lower int, upper int) func() (int, bool) {
	return func() (int, bool) {
		lower = iterator(lower)
		return lower, lower > upper
	}
}

func iterator(i int) int {
	i += 1
	return i
}

func (cars Collection) GenerateCars(start, limit int) Collection {

	carChannel := make(chan *IndexedCar)

	var waitGroup sync.WaitGroup

	numCarsToGenerate := start + limit - 1
	generatedCars := Collection{}

	waitGroup.Add(numCarsToGenerate)

	next := carGenerator(iterator, start -1, numCarsToGenerate)

	carIndex, done := next()
	for !done {
		go func(carIndex int) {
			thisCar, err := GetThisCar(carIndex)
			//log.Printf("GetThisCar(%v): %v\n", carIndex, thisCar)
			if err != nil {
				panic(err)
			}
			carChannel <- thisCar
			generatedCars = append(generatedCars, thisCar.Car)
			waitGroup.Done()
		}(carIndex)

		carIndex, done = next()
	}

	go func() {
		waitGroup.Wait()
		close(carChannel)
	}()

	printCars(carChannel, start, limit)
	return generatedCars
}

func printCars(indexedCars chan *IndexedCar, start, limit int) {
	log.Printf("\nGenerated Cars (%d to %d)\n%s\n", start, start + limit, DASHES)
	var cars Collection
	for car := range indexedCars {
		log.Printf("car: %s\n", car.Car)
		cars = append(cars, car.Car)
	}
}
