package main

type IntIterator interface {
	Next() (value string, ok bool)
}
const INVALID_INT_VAL = -1
const INVALID_STRING_VAL = ""

type Collection struct {
	index int
	List  []string
}

func (collection *Collection) Next() (value string, ok bool) {
	collection.index++
	if collection.index >= len(collection.List) {
		return INVALID_STRING_VAL, false
	}
	return collection.List[collection.index], true
}

func newSlice(s []string) *Collection {
	return &Collection{INVALID_INT_VAL, s}
}

func main() {
	var intCollection IntIterator
	intCollection = newSlice([]string{"CRV", "IS250", "Blazer"})
	value, ok := intCollection.Next()
	for ok {
		println(value)
		value, ok = intCollection.Next()
	}
}
