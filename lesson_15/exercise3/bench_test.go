package main

import "testing"

var data = map[string]any{"id": 1, "name": "test", "email": "test@test.com"}

func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = formatJSONWithPool(data)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = formatJSONWithoutPool(data)
	}
}
