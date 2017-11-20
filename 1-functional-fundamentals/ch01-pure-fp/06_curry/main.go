package main

func curryAddTwo(n int) (ret int) {
	defer func(){ret = n + 2}()
	return n
}

func main()  {
	println(curryAddTwo(1))
}
