package set

import (
	"reflect"
	"sync"
)

type (
	filterFunc       filterCallback
	equalityTestFunc (func(*StoreDatum, *StoreDatum) bool)
)

var (
	DefaultEqualityTest equalityTestFunc = func(s1, s2 *StoreDatum) bool {
		if s1.Hash != s2.Hash {
			return false
		}

		return reflect.DeepEqual(s1.Value, s2.Value)
	}
)

type StoreDatum struct {
	Value       interface{}
	valueString *string
	Hash        uint64
	Instances   int
	index       *int
}

type storeData []*StoreDatum

func (s storeData) Len() int           { return len(s) }
func (s storeData) Less(i, j int) bool { return s[i].Hash < s[j].Hash }
func (s storeData) Swap(i, j int) {
	sI, sJ := s[i], s[j]
	sI.index, sJ.index = sJ.index, sI.index
	s[i], s[j] = sJ, sI
}

type Datum struct {
	store              storeData
	equalityTest       equalityTestFunc
	cachedHash         *uint64
	cachedInstancesLen *int
	maximumValueCount  *int
	filter             filterFunc
	sorted             bool
	multiMode          bool
	mutex              sync.RWMutex
}
