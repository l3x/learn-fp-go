package packagea

import (
	b "packageb"
	"fmt"
	u "github.com/go-goodies/go_utils"
)

func Atask() {
	fmt.Println(u.PadLeft("A", 3))
	b.Btask()
}

