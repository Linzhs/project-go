package main

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Fib(20)
	}
}

func BenchmarkConcatString1(b *testing.B) {
	strId := uuid.New().String()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {

	}
}
