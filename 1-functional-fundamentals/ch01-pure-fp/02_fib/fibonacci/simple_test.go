package fibonacci

import "testing"

var FibTests = []struct {
	a     int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{21, 10946},
	{43, 433494437},
}

func TestSimple(t *testing.T) {
	for _, ft := range FibTests {
		if v := FibSimple(ft.a); v != ft.expected {
			t.Errorf("FibSimple(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibSimple(b *testing.B) {
	fn := FibSimple
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibSimple(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibSimple(i)
	}
}

func BenchmarkFibSimple0(b *testing.B)  { benchmarkFibSimple(0, b) }
func BenchmarkFibSimple1(b *testing.B)  { benchmarkFibSimple(1, b) }
func BenchmarkFibSimple2(b *testing.B)  { benchmarkFibSimple(2, b) }
func BenchmarkFibSimple3(b *testing.B)  { benchmarkFibSimple(3, b) }
func BenchmarkFibSimple4(b *testing.B) { benchmarkFibSimple(4, b) }
func BenchmarkFibSimple5(b *testing.B) { benchmarkFibSimple(5, b) }
func BenchmarkFibSimple6(b *testing.B) { benchmarkFibSimple(6, b) }
func BenchmarkFibSimple21(b *testing.B) { benchmarkFibSimple(21, b) }
func BenchmarkFibSimple43(b *testing.B) { benchmarkFibSimple(43, b) }