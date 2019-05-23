package set_test

import (
	"testing"
)

func BenchmarkString0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set0.String()
	}
}

func BenchmarkString10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10.String()
	}
}

func BenchmarkString100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set100.String()
	}
}

func BenchmarkString1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set1000.String()
	}
}

func BenchmarkString10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10000.String()
	}
}
