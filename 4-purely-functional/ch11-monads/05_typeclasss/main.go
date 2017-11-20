package main

import (
    "typeclass"
    "fmt"
)

func main() {
    int42 := typeclass.Int(42)
    str42 := typeclass.String("42")
    fmt.Println("str42.Equals(int42):", str42.Equals(int42))

    int64One := typeclass.Int64(1)
    int64Two := typeclass.Int64(2)
    fmt.Println("int64Two.Sum(int64One):", int64Two.Sum(int64One))

    int32One := typeclass.Int32(1)
    fmt.Println("int32One.Sum(int64One):", int32One.Sum(int64One))

    float32Five := typeclass.Float32(5)
    fmt.Println("int32One.Sum(int64One):", float32Five.Sum(int64One))

    int64Slice123 := typeclass.IntSlice([]int{1, 2, 3})
    int64Slice234 := typeclass.IntSlice([]int{2, 3, 4})
    fmt.Println("int64Slice123.Sum(int64Slice234):", int64Slice123.Sum(int64Slice234))
}
