package linuxid

import "testing"

func BenchmarkConstructor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewID()
	}
}
