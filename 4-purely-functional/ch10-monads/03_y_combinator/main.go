package main

import "fmt"

type Func func(int) int
type FuncFunc func(Func) Func
type RecursiveFunc func (RecursiveFunc) Func

// @tco yCombinator implements the lambda expression that
// enables composition of unpure components in our workflow
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

func Recurse(f FuncFunc) Func {
    return func(x int) int {
        return f(Recurse(f))(x)
    }
}

func fib(x int) int {
    if x == 0 {
        return 0
    } else if x <= 2 {
        return 1
    } else {
        return fib(x-2) + fib(x-1)
    }
}

func main() {
    yCombo := yCombinator(fibFuncFunc)
    recurse := Recurse(fibFuncFunc)
    for x := 0; x < 50; x++ {
        fmt.Printf("yCombo(%d) = %d\n", x, yCombo(x))
        fmt.Printf("recurse(%d) = %d\n", x, recurse(x))
        fmt.Printf("plainRe(%d) = %d\n", x, fib(x))
        fmt.Println("---")
    }
}
