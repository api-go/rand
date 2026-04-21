package rand_test

import (
	"testing"

	"apigo.cc/go/rand"
)

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rand.Int(1, 100)
	}
}

func BenchmarkFastIntParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = rand.FastInt(1, 100)
		}
	})
}
