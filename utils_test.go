package iban

import "testing"

func BenchmarkMod97Small(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mod97("123")
	}
}

func BenchmarkMod97Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mod97("00123014764700968325")
	}
}

func BenchmarkMod97SmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mod97("123")
		}
	})
}

func BenchmarkMod97LargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mod97("00123014764700968325")
		}
	})
}
