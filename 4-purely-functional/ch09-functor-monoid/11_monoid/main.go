package main

import (
	"monoid"
	"fmt"
)

func main() {

	const name = "Alice"
	stringMonoid := monoid.WrapName(name)
	fmt.Println("NameMonoid")
	fmt.Println("Initial state:", stringMonoid)
	fmt.Println("Zero:", stringMonoid.Zero())
	fmt.Println("1st application:", stringMonoid.Append(name))
	fmt.Println("Chain applications:", stringMonoid.Append(name).Append(name))

	ints := []int{1, 2, 3}
	intMonoid := monoid.WrapInt(ints)
	fmt.Println("\nIntMonoid")
	fmt.Println("Initial state:", intMonoid)
	fmt.Println("Zero:", intMonoid.Zero())
	fmt.Println("1st application:", intMonoid.Append(ints...))
	fmt.Println("Chain applications:", intMonoid.Append(ints...).Append(ints...))
	fmt.Println("Reduce chain:", intMonoid.Append(ints...).Append(ints...).Reduce())

	lineitems := []monoid.Lineitem{
		{1, 12978, 22330},
		{2, 530, 786},
		{5, 270, 507},
	}
	lineitemMonoid := monoid.WrapLineitem(lineitems)
	fmt.Println("\nLineItemMonoid")
	fmt.Println("Initial state:", lineitemMonoid)
	fmt.Println("Zero:", lineitemMonoid.Zero())
	fmt.Println("1st application:", lineitemMonoid.Append(lineitems...))
	fmt.Println("Chain applications:", lineitemMonoid.Append(lineitems...).Append(lineitems...))
	fmt.Println("Reduce chain:", lineitemMonoid.Append(lineitems...).Append(lineitems...).Reduce())
}
