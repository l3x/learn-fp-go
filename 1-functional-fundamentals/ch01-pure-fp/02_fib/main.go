package main

import (
    "02_fib/fibonacci"
)

func main() {
    println(fibonacci.FibSimple(4))
    println(fibonacci.FibMemoized(5))
    println(fibonacci.FibChanneled(6))
}
