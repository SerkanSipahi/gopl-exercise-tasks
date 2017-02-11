package main

import "testing"

/**
 * @Howto:
 * 1.) go test -bench=.
 */

var (
	args = []string{"a", "b", "c"}
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(args)
	}
}