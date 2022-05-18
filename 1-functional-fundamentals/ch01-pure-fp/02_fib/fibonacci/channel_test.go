package fibonacci

import "testing"

func TestChanneled(t *testing.T) {
	for _, ft := range FibTests {
		if v := FibChanneled(ft.a); v != ft.expected {
			t.Errorf("FibChanneled(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibChanneled(b *testing.B) {
	fn := FibChanneled
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibChanneled(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibChanneled(i)
	}
}

func BenchmarkFibChanneled0(b *testing.B)  { benchmarkFibChanneled(0, b) }
func BenchmarkFibChanneled1(b *testing.B)  { benchmarkFibChanneled(1, b) }
func BenchmarkFibChanneled2(b *testing.B)  { benchmarkFibChanneled(2, b) }
func BenchmarkFibChanneled3(b *testing.B)  { benchmarkFibChanneled(3, b) }
func BenchmarkFibChanneled4(b *testing.B)  { benchmarkFibChanneled(4, b) }
func BenchmarkFibChanneled5(b *testing.B)  { benchmarkFibChanneled(5, b) }
func BenchmarkFibChanneled6(b *testing.B)  { benchmarkFibChanneled(6, b) }
func BenchmarkFibChanneled10(b *testing.B) { benchmarkFibChanneled(10, b) }
func BenchmarkFibChanneled21(b *testing.B) { benchmarkFibChanneled(21, b) }
func BenchmarkFibChanneled43(b *testing.B) { benchmarkFibChanneled(43, b) }