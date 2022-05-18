package packagea

import b "01_dependency-rule-good/packageb"

func Atask() {
	println("A")
	b.Btask()
}

