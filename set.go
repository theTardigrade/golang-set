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
	store        storeData
	equalityTest equalityTestFunc
	cachedHash   *uint64
	sorted       bool
	mutex        sync.RWMutex
}

func New() *datum {
	return &datum{
		equalityTest: defaultEqualityTest,
		sorted:       true,
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

func NewFromBoolSlice(s []bool) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt8Slice(s []int8) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt16Slice(s []int16) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt32Slice(s []int32) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt64Slice(s []int64) (d *datum) {
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

func NewFromUint8Slice(s []uint8) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint16Slice(s []uint16) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint32Slice(s []uint32) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint64Slice(s []uint64) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUintSlice(s []uint) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUintptrSlice(s []uintptr) (d *datum) {
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

func NewFromByteSlice(s []byte) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromRuneSlice(s []rune) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromFloat32Slice(s []float32) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromFloat64Slice(s []float64) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromComplex64Slice(s []complex64) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromComplex128Slice(s []complex128) (d *datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func Clone(d *datum) (d2 *datum) {
	d2 = New()

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	d2.store = d.store[:]
	d2.cachedHash = d.cachedHash
	d2.equalityTest = d.equalityTest

	return
}

func Union(d1, d2 *datum) (d3 *datum) {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

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
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

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
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3 = New()

	for _, v := range d1.store {
		if !d2.contains(v) {
			d3.addOne(v)
		}
	}

	return
}

func Subset(d1, d2 *datum) bool {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	for _, v := range d1.store {
		if !d2.contains(v) {
			return false
		}
	}

	return true
}
