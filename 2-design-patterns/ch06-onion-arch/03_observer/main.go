package main

import (
	. "observer"
)

func main() {

	subject := Subject{}
	oa := Observable{Name: "A"}
	ob := Observable{Name: "B"}
	subject.AddObserver(&Observer{})
	subject.NotifyObservers(oa, ob)

	oc := Observable{Name: "C"}
	subject.NotifyObservers(oa, ob, oc)

	subject.DeleteObserver(&Observer{})
	subject.NotifyObservers(oa, ob, oc)

	od := Observable{Name: "D"}
	subject.NotifyObservers(oa, ob, oc, od)
}
