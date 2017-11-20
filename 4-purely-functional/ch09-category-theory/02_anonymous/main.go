package main

func addTwo(x int) int {
	return x + 2
}

func main() {
	println(addTwo(5))

	println(func(x int) int {return x + 2}(5))

	val := func(x int) int {return x + 2}(5)
	println(val)
}

