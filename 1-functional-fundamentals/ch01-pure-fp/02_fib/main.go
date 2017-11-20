package main

import (
    "fibonacci"
)

func main() {
    println(fibonacci.FibSimple(4))
    println(fibonacci.FibMemoized(5))
    println(fibonacci.FibChanneled(6))
}
