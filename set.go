package set

import (
	"sync"

	"github.com/theTardigrade/set/internal/hash"
)

type equalityTestFunc (func(*storeDatum, *storeDatum) bool)

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

func newStoreDatum(value interface{}) *storeDatum {
	return &storeDatum{
		value: value,
		hash:  hash.Get(value),
	}
}

func newStoreDatumWithIndex(value interface{}, index int) *storeDatum {
	s := newStoreDatum(value)
	s.index = &index
	return s
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
	storeMutex        sync.RWMutex
	cachedHash        *uint64
	cachedHashMutex   sync.RWMutex
	equalityTest      equalityTestFunc
	equalityTestMutex sync.RWMutex
}

func New() *datum {
	return &datum{
		equalityTest: defaultEqualityTest,
	}
}

func NewWithCapacity(c int) (d *datum) {
	d = New()

	d.makeStore(c)

	return
}

func NewFromSlice(s []interface{}) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromIntSlice(s []int) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromStringSlice(s []string) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func Clone(d *datum) (d2 *datum) {
	d2 = New()

	defer d.equalityTestMutex.RUnlock()
	defer d.cachedHashMutex.RUnlock()
	defer d.storeMutex.RUnlock()
	d.equalityTestMutex.RLock()
	d.cachedHashMutex.RLock()
	d.storeMutex.RLock()

	d2.store = d.store[:]
	d2.cachedHash = d.cachedHash
	d2.equalityTest = d.equalityTest

	return
}

func Union(d1, d2 *datum) (d3 *datum) {
	defer d1.storeMutex.RUnlock()
	defer d2.storeMutex.RUnlock()
	d1.storeMutex.RLock()
	d2.storeMutex.RLock()

	d3 = New()

	if len(d2.store) > len(d1.store) {
		d2, d1 = d1, d2
	}

	d3.store = d1.store[:]

	for _, s := range d2.store {
		d3.addOneFromDatum(s)
	}

	return
}

func Intersection(d1, d2 *datum) (d3 *datum) {
	defer d1.storeMutex.RUnlock()
	defer d2.storeMutex.RUnlock()
	d1.storeMutex.RLock()
	d2.storeMutex.RLock()

	d3 = New()

	if len(d2.store) > len(d1.store) {
		d2, d1 = d1, d2
	}

	for _, v := range d1.store {
		if d2.contains(v) {
			d3.addOne(v)
		}
	}

	return
}

func Difference(d1, d2 *datum) (d3 *datum) {
	defer d1.storeMutex.RUnlock()
	defer d2.storeMutex.RUnlock()
	d1.storeMutex.RLock()
	d2.storeMutex.RLock()

	d3 = New()

	for _, v := range d1.store {
		if !d2.contains(v) {
			d3.addOne(v)
		}
	}

	return
}

func Subset(d1, d2 *datum) bool {
	defer d1.storeMutex.RUnlock()
	defer d2.storeMutex.RUnlock()
	d1.storeMutex.RLock()
	d2.storeMutex.RLock()

	for _, v := range d1.store {
		if !d2.contains(v) {
			return false
		}
	}

	return true
}
