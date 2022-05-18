package monoid

type LineitemMonoid interface {
	Zero() []int
	Append(i ...int) LineitemMonoid
	Reduce() int
}

func WrapLineitem(lineitems []Lineitem) lineitemContainer {
	return lineitemContainer{lineitems: lineitems}
}

type Lineitem struct {
	Quantity 	int
	Price		int
	ListPrice 	int
}


type lineitemContainer struct {
	lineitems []Lineitem
}

func (lineitemContainer) Zero() []Lineitem {
	return nil
}

func (i lineitemContainer) Append(lineitems ...Lineitem) lineitemContainer {
	i.lineitems = append(i.lineitems, lineitems...)
	return i
}

func (i lineitemContainer) Reduce() Lineitem {
	totalQuantity := 0
	totalPrice := 0
	totalListPrice := 0
	for _, item := range i.lineitems {
		totalQuantity += item.Quantity
		totalPrice += item.Price
		totalListPrice += item.ListPrice
	}
	return Lineitem{totalQuantity, totalPrice, totalListPrice}
}

