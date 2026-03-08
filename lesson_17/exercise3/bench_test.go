package exercise3

import "testing"

var items = []string{
	"user123", "hello", "order456", "world", "item789",
	"test", "foo42", "bar", "baz99", "qux",
}

func BenchmarkSlow(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = processItemsSlow(items)
	}
}

func BenchmarkFast(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = processItemsFast(items)
	}
}
