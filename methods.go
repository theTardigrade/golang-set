package set

import (
	"math/rand"
	"time"

	"github.com/mitchellh/hashstructure"
)

func (d *datum) SetEqualityTest(equalityTest equalityTestFunc) {
	d.equalityTestMutex.Lock()
	d.equalityTest = equalityTest
	d.equalityTestMutex.Unlock()
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling
func (d *datum) addOne(value interface{}) {
	if value == nil {
		return
	}

	for _, v := range d.store {
		if d.equalityTest(v, value) {
			return
		}
	}

	d.store = append(d.store, value)

	d.cachedHashMutex.Lock()
	d.cachedHash = nil
	d.cachedHashMutex.Unlock()
}

func (d *datum) Add(values ...interface{}) {
	d.storeMutex.Lock()
	d.equalityTestMutex.RLock()

	for _, v := range values {
		d.addOne(v)
	}

	d.equalityTestMutex.RUnlock()
	d.storeMutex.Unlock()
}

// storeMutex should be locked before calling
func (d *datum) removeOneByIndex(i int) {
	if j := len(d.store) - 1; i <= j {
		d.store[j], d.store[i] = d.store[i], d.store[j]
		d.store = d.store[:j]

		d.cachedHashMutex.Lock()
		d.cachedHash = nil
		d.cachedHashMutex.Unlock()
	}
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling
func (d *datum) removeOne(value interface{}) {
	for i, v := range d.store {
		if d.equalityTest(v, value) {
			d.removeOneByIndex(i)
			return
		}
	}
}

func (d *datum) Remove(values ...interface{}) {
	d.storeMutex.Lock()
	d.equalityTestMutex.RLock()

	for _, v := range values {
		d.removeOne(v)
	}

	d.equalityTestMutex.RUnlock()
	d.storeMutex.Unlock()
}

// storeMutex should be read-locked before calling
func (d *datum) contains(value interface{}) bool {
	defer d.equalityTestMutex.RUnlock()
	d.equalityTestMutex.RLock()

	for _, v := range d.store {
		if d.equalityTest(v, value) {
			return true
		}
	}

	return false
}

func (d *datum) Contains(value interface{}) bool {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return d.contains(value)
}

func (d *datum) Pop() (value interface{}) {
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()

	value, index := d.pick()
	if value != nil {
		d.removeOneByIndex(index)
	}

	return
}

// storeMutex should be read-locked before calling
func (d *datum) pick() (value interface{}, index int) {
	if l := len(d.store); l > 0 {
		index = rand.Intn(l)
		value = d.store[index]
	}

	return
}

func (d *datum) Pick() (value interface{}) {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	value, _ = d.pick()

	return
}

func (d *datum) Clear() {
	d.storeMutex.Lock()
	d.cachedHashMutex.Lock()

	d.store = make(InterfaceSlice, 0, cap(d.store))
	d.cachedHash = nil

	d.cachedHashMutex.Unlock()
	d.storeMutex.Unlock()
}

type ForEachCallback (func(interface{}))

func (d *datum) ForEach(callback ForEachCallback) {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		callback(v)
	}
}

type MapCallback (func(interface{}) interface{})

func (d *datum) Map(callback MapCallback) {
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()

	for i, v := range d.store {
		d.store[i] = callback(v)
	}
}

type FilterCallback (func(interface{}) bool)

func (d *datum) FilterCallback(callback FilterCallback) {
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()

	for i, v := range d.store {
		if !callback(v) {
			d.removeOneByIndex(i)
		}
	}
}

type ReduceCallback (func(interface{}, interface{}) interface{})

func (d *datum) Reduce(initialValue interface{}, callback ReduceCallback) (accumulator interface{}) {
	accumulator = initialValue

	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		accumulator = callback(accumulator, v)
	}

	return
}

func (d *datum) IntSum() (accumulator int) {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		if value, ok := v.(int); ok {
			accumulator += value
		}
	}

	return
}

func (d *datum) IntProduct() (accumulator int) {
	accumulator = 1

	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		if value, ok := v.(int); ok {
			accumulator *= value
		}
	}

	return
}

type EveryCallback FilterCallback

func (d *datum) Every(callback EveryCallback) bool {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		if !callback(v) {
			return false
		}
	}

	return true
}

type SomeCallback EveryCallback

func (d *datum) Some(callback SomeCallback) bool {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, v := range d.store {
		if callback(v) {
			return true
		}
	}

	return false
}

func (d *datum) Equal(d2 *datum) bool {
	if d2 == d {
		return true
	}

	defer d.storeMutex.RUnlock()
	defer d2.storeMutex.RUnlock()
	d.storeMutex.RLock()
	d2.storeMutex.RLock()

	l, l2 := len(d.store), len(d2.store)

	if l != l2 {
		return false
	}

	if l == 0 {
		return true
	}

	return d.hash() == d2.hash()
}

// storeMutex should be read-locked before calling
func (d *datum) hash() (value uint64) {
	defer d.cachedHashMutex.Unlock()
	d.cachedHashMutex.Lock()

	if d.cachedHash != nil {
		value = *d.cachedHash
	} else {
		var err error

		value, err = hashstructure.Hash(d.store, nil)
		if err != nil {
			panic(err)
		}

		d.cachedHash = &value
	}

	return
}

func (d *datum) Hash() uint64 {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return d.hash()
}

func (d *datum) Slice() InterfaceSlice {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return InterfaceSlice(d.store[:])
}

func (d *datum) Empty() bool {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return len(d.store) == 0
}

func (d *datum) Len() (value int) {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return len(d.store)
}

func (d *datum) Clone() *datum {
	return Clone(d)
}

func (d *datum) Union(d2 *datum) *datum {
	return Union(d, d2)
}

func (d *datum) Intersection(d2 *datum) *datum {
	return Intersection(d, d2)
}

func (d *datum) Difference(d2 *datum) *datum {
	return Difference(d, d2)
}

func (d *datum) Subset(d2 *datum) bool {
	return Subset(d, d2)
}

// seed the random-number generator for use in Pick and Pop functions
func init() {
	rand.Seed(time.Now().UnixNano())
}
