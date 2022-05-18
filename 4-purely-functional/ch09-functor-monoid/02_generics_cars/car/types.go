package car

// +gen slice:"Where,Sum[Dollars],GroupBy[string],Select[Dollars]"
type Car struct {
	Make string
	Model string
	Price Dollars
}

type Dollars int
