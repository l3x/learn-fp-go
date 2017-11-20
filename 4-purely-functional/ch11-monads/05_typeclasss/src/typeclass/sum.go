package typeclass

type Sum interface {
	Sum(Sum) int64
}

type Int32 int32
type Int64 int64
type Float32 float32
type IntSlice []int

func (i Int32) Sum(s Sum) int64 {
	it := int64(i)
	switch x := s.(type) {
	case Int64:
		return it + int64(x)
	case Int32:
		return it + int64(x)
	case Float32:
		return it + int64(x)
	case IntSlice:
		sum := int64(0)
		for _, num := range x {
			sum += int64(num)
		}
		return it + sum
	default:
		return 0
	}
}

func (i Int64) Sum(s Sum) int64 {
	it := int64(i)
	switch x := s.(type) {
	case Int64:
		return it + int64(x)
	case Int32:
		return it + int64(x)
	case Float32:
		return it + int64(x)
	case IntSlice:
		sum := int64(0)
		for _, num := range x {
			sum += int64(num)
		}
		return it + sum
	default:
		return 0
	}
}

func (i Float32) Sum(s Sum) int64 {
	it := int64(i)
	switch x := s.(type) {
	case Int64:
		return it + int64(x)
	case Int32:
		return it + int64(x)
	case Float32:
		return it + int64(x)
	case IntSlice:
		sum := int64(0)
		for _, num := range x {
			sum += int64(num)
		}
		return it + sum
	default:
		return 0
	}
}

func (i IntSlice) Sum(s Sum) int64 {
	it := i
	switch x := s.(type) {
	case Int64:
		sum := int64(0)
		for _, num := range it {
			sum += int64(num)
		}
		return int64(x) + sum
	case Int32:
		sum := int64(0)
		for _, num := range it {
			sum += int64(num)
		}
		return int64(x) + sum
	case Float32:
		sum := int64(0)
		for _, num := range it {
			sum += int64(num)
		}
		return int64(x) + sum
	case IntSlice:
		sum := int64(0)
		for _, num := range it {
			sum += int64(num)
		}
		for _, num := range x {
			sum += int64(num)
		}
		return sum
	default:
		return 0
	}
}
