package sum

func SumTailCall(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return vs[0] + SumTailCall(vs[1:])
}
