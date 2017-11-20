package packagea

import (
	b "packageb"
	"fmt"
)

func Atask() {
	fmt.Println("A")
	b.Btask()
}

