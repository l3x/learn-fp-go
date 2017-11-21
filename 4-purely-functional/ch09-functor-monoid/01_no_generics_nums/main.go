package main

import (
	"fmt"
)

func int8Sum(list []int8) (int8) {
	var result int8 = 0
	for x := 0; x < len(list); x++ {
		result += list[x]
	}
	return result
}

func int32Sum(list []int32) (int32) {
	var result int32 = 0
	for x := 0; x < len(list); x++ {
		result += list[x]
	}
	return result
}

func float64Sum(list []float64) (float64) {
	var result float64 = 0
	for x := 0; x < len(list); x++ {
		result += list[x]
	}
	return result
}

func complex128Sum(list []complex128) (complex128) {
	var result complex128 = 0
	for x := 0; x < len(list); x++ {
		result += list[x]
	}
	return result
}

func main() {
	fmt.Println("int8Sum:", int8Sum([]int8 {1, 2, 3}))
	fmt.Println("int32Sum:", int32Sum([]int32{1, 2, 3}))
	fmt.Println("float64Sum:", float64Sum([]float64{1, 2, 3}))
	fmt.Println("complex128Sum:", complex128Sum([]complex128{1, 2, 3}))
}
