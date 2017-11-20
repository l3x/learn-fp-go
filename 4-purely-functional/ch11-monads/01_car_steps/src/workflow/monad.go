package workflow

type Data interface{}
type Monad func(error) (Data, error)

func Get(d Data) Monad {
	return func(e error) (Data, error) {
		return d, e
	}
}

func Next(m Monad, f func(Data) Monad) Monad {
	return func(e error) (Data, error) {
		newData, newError := m(e)
		if newError != nil {
			return nil, newError
		}
		return f(newData)(newError)
	}
}
