package sum

import "testing"

func TestSumTailCall(t *testing.T) {
	for _, st := range sumTests {
		if v := SumTailCall(st.a); v != st.expected {
			t.Errorf("SumTailCall(%d) returned %d, expected %d", st.a, v, st.expected)
		}
	}
}

func BenchmarkSumTailCall(b *testing.B) {
	fn := SumTailCall
	for i := 0; i < b.N; i++ {
		_ = fn([]int{1, 2, 3})
	}
}

func benchmarkSumTailCall(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumTailCall(s)
	}
}

func BenchmarkSumTailCall1(b *testing.B)  { benchmarkSumTailCall([]int{1}, b) }
func BenchmarkSumTailCall2(b *testing.B)  { benchmarkSumTailCall([]int{1, 2}, b) }
func BenchmarkSumTailCall3(b *testing.B)  { benchmarkSumTailCall([]int{1, 2, 3}, b) }
func BenchmarkSumTailCall10(b *testing.B) { benchmarkSumTailCall([]int{1, 2, 3, 4}, b) }
func BenchmarkSumTailCall20(b *testing.B) { benchmarkSumTailCall([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, b) }
func BenchmarkSumTailCall40(b *testing.B) { benchmarkSumTailCall([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b) }