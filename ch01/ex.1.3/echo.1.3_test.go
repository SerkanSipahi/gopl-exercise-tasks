package main

import "testing"

var (
	args = []string{"a", "b", "c"}
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(args)
	}
}