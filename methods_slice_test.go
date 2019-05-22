package set_test

import (
	"testing"
)

func BenchmarkSlice0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set0.Slice()
	}
}

func BenchmarkSlice10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10.Slice()
	}
}

func BenchmarkSlice100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set100.Slice()
	}
}

func BenchmarkSlice1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set1000.Slice()
	}
}

func BenchmarkSlice10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10000.Slice()
	}
}
