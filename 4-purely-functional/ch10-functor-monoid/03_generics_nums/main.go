package main

import (
	"fmt"
	. "num"
)

func main() {
	fmt.Println("int8Sum:", Int8Slice{1, 2, 3}.SumInt8(Int8fn))
	fmt.Println("int32Sum:", Int32Slice{1, 2, 3}.SumInt32(Int32fn))
	fmt.Println("float64Sum:", Float64Slice{1, 2, 3}.SumFloat64(Float64fn))
	fmt.Println("complex128Sum:", Complex128Slice{1, 2, 3}.SumComplex128(Complex128fn))
}
