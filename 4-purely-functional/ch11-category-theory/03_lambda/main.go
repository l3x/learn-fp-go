package main

type lambda func(int) int

func add(a int) lambda {
	return func(free int) int {
		return func(b int) int {
			return a + b
		}(free)
	}
}

func main() {
	add2 := add(2)
	three := add2(1)
	println("Pass 1 to to add2 expression to get:", three)
	four := add2(2)
	println("Pass 2 to to add2 expression to get:", four)
}


/*
	var n = 10
	fmt.Println(n == func(z int) int { return n }(123))
	fmt.Println(add(1)(10))

 */
