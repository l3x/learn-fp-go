package main

import (
	"fmt"
	. "functor"
)

func main() {
	ints := []int{1,2,3}
	impInts := []int{}
	for _, v := range ints {
		impInts = append(impInts, v + 2)
	}
	fmt.Println("imperative loop:", impInts)

	add2 := func(i int) int { return i + 2 }
	fpInts := Functor(ints).Map(add2)
	fmt.Println("fp map:", fpInts)
}

