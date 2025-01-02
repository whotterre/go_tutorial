package popcount

import (
	"testing"

	"2.3-2.5/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(43)
	}
}

func BenchmarkLoopCount(b *testing.B){
	for i := 0; i < b.N; i++ {
		popcount.LoopCount(43)
	}
}
func BenchmarkPopCountShift(b *testing.B){
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(43)
	}
}