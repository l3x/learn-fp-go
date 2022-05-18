package observer

type Subject struct {
	callbacks []Callback
}

func (o *Subject) AddObserver(c Callback) {
	o.callbacks = append(o.callbacks, c)
}
func (o *Subject) DeleteObserver(c Callback) {
	o.callbacks = append(o.callbacks, c)

	newCallbacks := []Callback{}
	for _, cb := range o.callbacks {
		if cb != c {
			newCallbacks = append(newCallbacks, cb)
		}
	}
	o.callbacks = newCallbacks
}

func (o *Subject) NotifyObservers(oes ...Observable) {
	for _, oe := range oes {
		for _, c := range o.callbacks {
			c.Notify(&oe)
		}
	}
}
