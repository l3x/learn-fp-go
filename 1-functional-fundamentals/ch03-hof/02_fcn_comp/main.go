package main

import (
	"fmt"
	"strings"
)

type StrFunc func(string) string

func Compose(f StrFunc, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func main() {
	var recognize = func(name string) string {
			return fmt.Sprintf("Hey %s", name)
		}

	var emphasize = func(statement string) string {
		return fmt.Sprintf(strings.ToUpper(statement) + "!")
		}

	var greetFoG = Compose(recognize, emphasize)
	fmt.Println(greetFoG("Gopher"))

	var greetGoF = Compose(emphasize, recognize)
	fmt.Println(greetGoF("Gopher"))
}
