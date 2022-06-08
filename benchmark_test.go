package linuxid

import "testing"

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkConstructor-12    	 1896780	       794.5 ns/op	      71 B/op	       3 allocs/op
func BenchmarkNewIDRandom(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDRandom()
	}
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkConstructor-12    	 4363184	       349.8 ns/op	      71 B/op	       4 allocs/op
func BenchmarkConstructor(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncremental10K()
	}
}
