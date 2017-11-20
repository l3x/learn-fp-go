package main

import "fmt"

type Collection []string
type MapFunc func(string) string
type MapFunc2 func(string, int) string

func (cars Collection) Map(fn MapFunc) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car))
	}
	return mappedCars
}
func (cars Collection) Map2(fn MapFunc2) Collection {
	mappedCars := make(Collection, 0, len(cars))
	for _, car := range cars {
		mappedCars = append(mappedCars, fn(car, 2))
	}
	return mappedCars
}

func Upgrade() MapFunc {
	return func(car string) string {
		return fmt.Sprintf("%s %s", car, "LX")
	}
}

// For an explanation of what's going on below, see chapter 11.
type Func func(int) int
type FuncFunc func(Func) Func
type RecursiveFunc func (RecursiveFunc) Func
func yCombinator(f FuncFunc) Func {
	g := func(r RecursiveFunc) Func {
		return f(func(x int) int {
			return r(r)(x)
		})
	}
	return g(g)
}
func fibFuncFunc(f Func) Func {
	return func(x int) int {
		if x == 0 {
			return 0
		} else if x <= 2 {
			return 1
		} else {
			return f(x-2) + f(x-1)
		}
	}
}

func main() {
	cars := &Collection{"Honda Accord", "Lexus IS 250"}

	fmt.Println("Upgrade() is not a Lambda Expression:")
	fmt.Printf("> cars.Map(Upgrade()): %+v\n\n", cars.Map(Upgrade()))

	fmt.Println("Anonymous function is not a Lambda Expression:")
	fmt.Printf("> cars.Map(func(...{...}): %+v\n\n", cars.Map2(func(car string, num int) string {
		return fmt.Sprintf("%s %s%d", car, "LX", num)
	}))

	yCombo := yCombinator(fibFuncFunc)
	fmt.Println("r(r)(x) in anonymous function in yCombinator is a Lambda Expression:")
	fmt.Println("> yCombo(5):", yCombo(5))
}
