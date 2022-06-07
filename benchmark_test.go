package linuxid

import "testing"

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkConstructor-12    	  983395	      2140 ns/op	     320 B/op	      11 allocs/op

func BenchmarkConstructor(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewID()
	}
}
