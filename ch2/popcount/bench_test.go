package main_test

import (
	"testing"

	"github.com/benjohns1/go-programming-language-exercises/ch2/popcount/popcount"
	"github.com/benjohns1/go-programming-language-exercises/ch2/popcount/popcountclear"
	"github.com/benjohns1/go-programming-language-exercises/ch2/popcount/popcountloop"
	"github.com/benjohns1/go-programming-language-exercises/ch2/popcount/popcountshift"
)

func BenchmarkPopCounts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoops(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountloop.PopCount(uint64(i))
	}
}

func BenchmarkPopCountNoTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountshift.PopCount(uint64(i))
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountclear.PopCount(uint64(i))
	}
}
