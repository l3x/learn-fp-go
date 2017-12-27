package main

import (
    "oop"
    "fmt"
)

func main() {
    oop.MyCars.Add(oop.Car{"IS250"})
    oop.MyCars.Add(oop.Car{"Accord"})
    oop.MyCars.Add(oop.Car{"Highlander"})

    car, err := oop.MyCars.Find("Highlander")
    if err != nil {
        fmt.Printf("ERROR: %+v", err)
    } else {
        fmt.Printf("Found %+v", car)
    }
}
