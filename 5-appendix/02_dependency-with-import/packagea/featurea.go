package packagea

import (
	b "02_dependency-with-import/packageb"
	"fmt"
)

func Atask() {
	fmt.Println("A")
	b.Btask()
}
