package monoid

type IntMonoid interface {
	Zero() []int
	Append(i ...int) IntMonoid
	Reduce() int
}

func WrapInt(ints []int) IntMonoid {
	return intContainer{ints: ints}
}

type intContainer struct {
	ints []int
}

func (intContainer) Zero() []int {
	return nil
}

func (i intContainer) Append(ints ...int) IntMonoid {
	i.ints = append(i.ints, ints...)
	return i
}

func (i intContainer) Reduce() int {
	total := 0
	for _, item := range i.ints {
		total += item
	}
	return total
}

