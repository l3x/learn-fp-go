package packageb

import a "02_circulardep/packagea"

func Btask() {
	println("B")
	a.Atask()
}
