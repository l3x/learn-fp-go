package packagea

import b "02_circulardep/packageb"

func Atask() {
	println("A")
	b.Btask()
}

