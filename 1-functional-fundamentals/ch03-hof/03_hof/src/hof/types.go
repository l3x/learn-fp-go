package hof

type (
	FilterFunc func(string) bool
	MapFunc func(string) string
	Collection []string
	ReducerFunc func(string, Collection) Collection
	ReducerFunc2 func(string, CarCollection) CarCollection
	Car string
	CarType struct {
		Make  string `json:"make"`
		Model string `json:"model"`
	}
	CarCollection []CarType
	IndexedCar struct {
		Index int `json:"index"`
		Car   string` json:"car"`
	}
	Payload struct {
		IndexedCars []IndexedCar
	}
)
