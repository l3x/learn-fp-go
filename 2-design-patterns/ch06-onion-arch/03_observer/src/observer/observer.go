package observer

type Observable struct {
	Name string
}

type Observer struct {
}

func (ob *Observer) Notify(o *Observable) {
	println(o.Name)
}

type Callback interface {
	Notify(o *Observable)
}

