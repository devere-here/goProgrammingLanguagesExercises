package main

import (
	"testing"

	"github.com/sideProjects/goProgrammingLanguagesExercises/ch2/popcount"
)

func BenchmarkForPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkForPopCountEx1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountEx1(uint64(i))
	}
}

func BenchmarkForPopCountEx2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountEx2(uint64(i))
	}
}
