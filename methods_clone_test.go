package set_test

import (
	"testing"
)

func BenchmarkClone0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set0.Clone()
	}
}

func BenchmarkClone10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10.Clone()
	}
}

func BenchmarkClone100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set100.Clone()
	}
}

func BenchmarkClone1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set1000.Clone()
	}
}

func BenchmarkClone10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10000.Clone()
	}
}
