package packageb

import a "packagea"

func Btask() {
	println("B")
	a.Atask()
}
