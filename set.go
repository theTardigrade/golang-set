package set

import (
	"sync"
)

type (
	filterFunc       filterCallback
	equalityTestFunc (func(*storeDatum, *storeDatum) bool)
)

var (
	defaultEqualityTest equalityTestFunc = func(d1, d2 *storeDatum) bool {
		return d1.hash == d2.hash
	}
)

type storeDatum struct {
	value interface{}
	hash  uint64
	index *int
}

type storeData []*storeDatum

func (s storeData) Len() int           { return len(s) }
func (s storeData) Less(i, j int) bool { return s[i].hash < s[j].hash }
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
