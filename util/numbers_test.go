package util

import "testing"

func BenchmarkExplodeBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitExplode(uint(1337))
	}
}

func BenchmarkDecimalExplode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalExplode(1337)
	}
}
