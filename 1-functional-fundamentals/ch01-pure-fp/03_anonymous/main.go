package main

import "fmt"

func namedGreeting(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func anonymousGreeting() func(string) {
	return func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}
}

func main() {
	namedGreeting("Alice")

	greet := anonymousGreeting()
	greet("Bob")

	func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}("Cindy")
}
