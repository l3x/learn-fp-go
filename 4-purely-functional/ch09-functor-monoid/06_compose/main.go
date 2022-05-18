package main

import (
	"fmt"
	. "06_compose/compose"
)

func main() {
	fmt.Println("A to B - Humanize(true):", Humanize(true))
	fmt.Println("B to C - Emphasize(\"yes\"):", Emphasize("yes"))
	fmt.Println("A to C - EmphasizeHumanize(true)", EmphasizeHumanize(true))
}
