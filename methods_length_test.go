package set_test

import "testing"

func BenchmarkLen0_native(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set0.Len(false)
	}
}

func BenchmarkLen10_native(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10.Len(false)
	}
}

func BenchmarkLen100_native(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set100.Len(false)
	}
}

func BenchmarkLen1000_native(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set1000.Len(false)
	}
}

func BenchmarkLen10000_native(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10000.Len(false)
	}
}

func BenchmarkLen0_instances(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set0.Len(true)
	}
}

func BenchmarkLen10_instances(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10.Len(true)
	}
}

func BenchmarkLen100_instances(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set100.Len(true)
	}
}

func BenchmarkLen1000_instances(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set1000.Len(true)
	}
}

func BenchmarkLen10000_instances(b *testing.B) {
	for n := 0; n < b.N; n++ {
		set10000.Len(true)
	}
}
