package main

import (
    "fmt"
)

var sum = 5

func addTwo() func() int {
    sum := 0
    return func() int {  // anonymous function
        sum += 2
        return sum
    }
}

func addTwoDynamic() func() int {
    return func() int {
        sum += 2
        return sum
    }
}

func main() {
    twoMore := addTwo()
    fmt.Println(twoMore())
    fmt.Println(twoMore())

    twoMoreDynamic := addTwoDynamic()
    fmt.Println(twoMoreDynamic())
    fmt.Println(twoMoreDynamic())
}
