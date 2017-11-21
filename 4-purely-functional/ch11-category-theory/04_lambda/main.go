package main

import "fmt"

func main() {
	fmt.Println(compose(mul3, add1)(10))

	//	fmt.Println(compose(mul3, func(n int) int { return n + 1 })(10))
	//	fmt.Println(compose(func(n int) int { return n * 3 }, func(n int) int { return n + 1 })(10))

	fmt.Println(compose(mul3, func(x int) fii { return func(v int) int { return func(n int) int { return n + x }(v) } }(1))(10))

	fmt.Println(compose(func(n int) int { return func(z int) int { return n * 3 }(1234567) }, func(x int) fii { return func(v int) int { return func(n int) int { return n + x }(v) } }(1))(10))

	fmt.Println(func(f, g fii) fii { return func(n int) int { return f(g(n)) } }(func(n int) int { return func(z int) int { return n * 3 }(1234567) }, func(x int) fii { return func(v int) int { return func(n int) int { return n + x }(v) } }(1))(10))
}

type fii func(int) int

func compose(f, g fii) fii {
	return func(n int) int { return f(g(n)) }
}

var add1 = makeAdder(1)

func mul3(n int) int {
	return func(z int) int { return n * 3 }(1234567)
}

func makeAdder(x int) fii {
	return func(v int) int { return func(n int) int { return n + x }(v) }
}