package fibonacci

func FibSimple(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return FibSimple(x-2) + FibSimple(x-1)
	}
}