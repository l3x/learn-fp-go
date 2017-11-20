package typeclass

import (
	"strconv"
)

type Equals interface {
	Equals(Equals) bool
}

type Int int

func (i Int) Equals(e Equals) bool {
	intVal := int(i)
	switch x := e.(type) {
	case Int:
		return intVal == int(x)
	case String:
		convertedInt, err := strconv.Atoi(string(x))
		if err != nil {
			return false
		}
		return intVal == convertedInt
	default:
		return false
	}
}

type String string

func (s String) Equals(e Equals) bool {
	stringVal := string(s)
	switch x := e.(type) {
	case String:
		return stringVal == string(x)
	case Int:
		return stringVal == strconv.Itoa(int(x))
	default:
		return false
	}
}
