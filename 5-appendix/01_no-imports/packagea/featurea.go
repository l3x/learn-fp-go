package packagea

import b "01_no-imports/packageb"

func Atask() {
	println("A")
	b.Btask()
}
