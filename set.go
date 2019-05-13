package set

import (
	"reflect"
	"sync"
)

type InterfaceSlice []interface{}

type equalityTestFunc (func(interface{}, interface{}) bool)

type datum struct {
	store             InterfaceSlice
	storeMutex        sync.RWMutex
	cachedHash        *uint64
	cachedHashMutex   sync.RWMutex
	equalityTest      equalityTestFunc
	equalityTestMutex sync.RWMutex
}

func New() *datum {
	return &datum{
		equalityTest: reflect.DeepEqual,
	}
}

func NewWithCapacity(c int) (d *datum) {
	d = New()

	d.store = make(InterfaceSlice, 0, c)

	return
}

func NewFromSlice(s InterfaceSlice) (d *datum) {
	d = New()

	d.Add(s...)

	return
}

func NewFromIntSlice(s []int) (d *datum) {
	d = New()

	for _, v := range s {
		d.addOne(v)
	}

	return
}

func NewFromStringSlice(s []string) (d *datum) {
	d = New()

	for _, v := range s {
		d.addOne(v)
	}

	return
}

func Clone(d *datum) (d2 *datum) {
	d2 = New()

	d.storeMutex.RLock()
	d.cachedHashMutex.RLock()
	d.equalityTestMutex.RLock()

	d2.store = d.store[:]
	d2.cachedHash = d.cachedHash
	d2.equalityTest = d.equalityTest

	d.equalityTestMutex.RUnlock()
	d.cachedHashMutex.RUnlock()
	d.storeMutex.RUnlock()

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
	d3.Add(d2.store...)

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
