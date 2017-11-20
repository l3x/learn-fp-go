package sum

import "testing"

var sumTests = []struct {
	a     []int
	expected int
}{
	{[]int{1}, 1},
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3}, 6},
	{[]int{1, 2, 3, 4}, 10},
}

func TestSumLoop(t *testing.T) {
	for _, st := range sumTests {
		if v := SumLoop(st.a); v != st.expected {
			t.Errorf("SumLoop(%d) returned %d, expected %d", st.a, v, st.expected)
		}
	}
}

func BenchmarkSumLoop(b *testing.B) {
	fn := SumLoop
	for i := 0; i < b.N; i++ {
		_ = fn([]int{1, 2, 3})
	}
}

func benchmarkSumLoop(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumLoop(s)
	}
}

func BenchmarkSumLoop1(b *testing.B)  { benchmarkSumLoop([]int{1}, b) }
func BenchmarkSumLoop2(b *testing.B)  { benchmarkSumLoop([]int{1, 2}, b) }
func BenchmarkSumLoop3(b *testing.B)  { benchmarkSumLoop([]int{1, 2, 3}, b) }
func BenchmarkSumLoop10(b *testing.B) { benchmarkSumLoop([]int{1, 2, 3, 4}, b) }
func BenchmarkSumLoop20(b *testing.B) { benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, b) }
func BenchmarkSumLoop40(b *testing.B) { benchmarkSumLoop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}, b) }