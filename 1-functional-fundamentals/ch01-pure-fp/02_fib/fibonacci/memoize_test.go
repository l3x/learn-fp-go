package fibonacci

import "testing"

func TestMemoized(t *testing.T) {
	for _, ft := range FibTests {
		if v := FibMemoized(ft.a); v != ft.expected {
			t.Errorf("FibMemoized(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibMemoized(b *testing.B) {
	fn := FibMemoized
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibMemoized(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibMemoized(i)
	}
}

func BenchmarkFibMemoized0(b *testing.B)  { benchmarkFibMemoized(0, b) }
func BenchmarkFibMemoized1(b *testing.B)  { benchmarkFibMemoized(1, b) }
func BenchmarkFibMemoized2(b *testing.B)  { benchmarkFibMemoized(2, b) }
func BenchmarkFibMemoized3(b *testing.B)  { benchmarkFibMemoized(3, b) }
func BenchmarkFibMemoized4(b *testing.B)  { benchmarkFibMemoized(4, b) }
func BenchmarkFibMemoized5(b *testing.B)  { benchmarkFibMemoized(5, b) }
func BenchmarkFibMemoized6(b *testing.B)  { benchmarkFibMemoized(6, b) }
func BenchmarkFibMemoized10(b *testing.B) { benchmarkFibMemoized(10, b) }
func BenchmarkFibMemoized21(b *testing.B) { benchmarkFibMemoized(21, b) }
func BenchmarkFibMemoized43(b *testing.B) { benchmarkFibMemoized(43, b) }