package main

import (
	"fmt"
	"time"
)

func greeting(name string) {
	// variable in outer function
	msg := name + fmt.Sprintf(" (at %v)", time.Now().String())

	// foo is a inner function and has access to text variable, is a closure
	// closures have access to variables even after exiting this block
	foo := func() {
		fmt.Printf("Hey %s!\n", msg)
	}

	// calling the closure
	foo()
}

func main() {
	greeting("alice")
}
