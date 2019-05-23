package set_test

import "testing"

func BenchmarkClear0_loseCapacity(b *testing.B) {
	s := set0.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(false)
	}
}

func BenchmarkClear10_loseCapacity(b *testing.B) {
	s := set10.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(false)
	}
}

func BenchmarkClear100_loseCapacity(b *testing.B) {
	s := set100.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(false)
	}
}

func BenchmarkClear1000_loseCapacity(b *testing.B) {
	s := set1000.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(false)
	}
}

func BenchmarkClear10000_loseCapacity(b *testing.B) {
	s := set10000.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(false)
	}
}

func BenchmarkClear0_keepCapacity(b *testing.B) {
	s := set0.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(true)
	}
}

func BenchmarkClear10_keepCapacity(b *testing.B) {
	s := set100.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(true)
	}
}

func BenchmarkClear100_keepCapacity(b *testing.B) {
	s := set100.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(true)
	}
}

func BenchmarkClear1000_keepCapacity(b *testing.B) {
	s := set1000.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(true)
	}
}

func BenchmarkClear10000_keepCapacity(b *testing.B) {
	s := set10000.Clone()
	for n := 0; n < b.N; n++ {
		s.Clear(true)
	}
}
