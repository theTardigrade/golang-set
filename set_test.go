package set_test

import (
	"github.com/theTardigrade/set"
)

var (
	set0     = set.NewWithCapacity(0)
	set10    = set.NewWithCapacity(10)
	set100   = set.NewWithCapacity(1e2)
	set1000  = set.NewWithCapacity(1e3)
	set10000 = set.NewWithCapacity(1e4)
)

func init() {
	for _, s := range []*set.Datum{
		set0,
		set10,
		set100,
		set1000,
		set10000,
	} {
		c := s.Cap()
		i := make([]int, c)

		for c--; c >= 0; c-- {
			i = append(i, c)
		}

		s.AddFromIntSlice(i)
	}
}
