package set

import (
	"sync"
)

type (
	filterFunc       filterCallback
	equalityTestFunc (func(*StoreDatum, *StoreDatum) bool)
)

var (
	defaultEqualityTest equalityTestFunc = func(d1, d2 *StoreDatum) bool {
		return d1.Hash == d2.Hash
	}
)

type StoreDatum struct {
	Value interface{}
	Hash  uint64
	index *int
}

type storeData []*StoreDatum

func (s storeData) Len() int           { return len(s) }
func (s storeData) Less(i, j int) bool { return s[i].Hash < s[j].Hash }
func (s storeData) Swap(i, j int) {
	sI, sJ := s[i], s[j]
	sI.index, sJ.index = sJ.index, sI.index
	s[i], s[j] = sJ, sI
}

type datum struct {
	store             storeData
	equalityTest      equalityTestFunc
	cachedHash        *uint64
	maximumValueCount *int
	filter            filterFunc
	sorted            bool
	mutex             sync.RWMutex
}
