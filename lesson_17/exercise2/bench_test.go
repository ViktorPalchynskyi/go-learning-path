package exercise2

import "testing"

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := newPointValue(1.0, 2.0)
		_ = p
	}
}

func BenchmarkPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := newPointPtr(1.0, 2.0)
		_ = p
	}
}
