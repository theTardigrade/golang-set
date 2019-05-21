package set

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/theTardigrade/set/internal/hash"
)

func (d *datum) SetEqualityTest(equalityTest equalityTestFunc) {
	if equalityTest == nil {
		return
	}

	d.equalityTestMutex.Lock()
	d.equalityTest = equalityTest
	d.equalityTestMutex.Unlock()
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling;
// clearCachedHash method should be called afterwards;
// s.value should not equal nil
func (d *datum) addOneFromDatum(s *storeDatum) {
	for _, s2 := range d.store {
		if d.equalityTest(s, s2) {
			return
		}
	}

	d.store = append(d.store, s)
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) addOne(value interface{}) {
	if value == nil {
		return
	}

	s := newStoreDatumWithIndex(value, len(d.store))
	d.addOneFromDatum(s)
}

func (d *datum) clearCachedHash() {
	d.cachedHashMutex.Lock()
	d.cachedHash = nil
	d.cachedHashMutex.Unlock()
}

func (d *datum) Add(values ...interface{}) {
	defer d.equalityTestMutex.RUnlock()
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()
	d.equalityTestMutex.RLock()

	for _, v := range values {
		d.addOne(v)
	}

	d.clearCachedHash()
}

func (d *datum) AddFromSlice(values []interface{}) {
	d.Add(values...)
}

func (d *datum) AddFromIntSlice(values []int) {
	defer d.equalityTestMutex.RUnlock()
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()
	d.equalityTestMutex.RLock()

	for i := len(values) - 1; i >= 0; i-- {
		d.addOne(values[i])
	}

	d.clearCachedHash()
}

func (d *datum) AddFromStringSlice(values []string) {
	defer d.equalityTestMutex.RUnlock()
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()
	d.equalityTestMutex.RLock()

	for i := len(values) - 1; i >= 0; i-- {
		d.addOne(values[i])
	}

	d.clearCachedHash()
}

// storeMutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOneFromIndex(i int) {
	if j := len(d.store) - 1; i <= j {
		d.store[j], d.store[i] = d.store[i], d.store[j]
		d.store = d.store[:j]
	}
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOneFromDatum(s *storeDatum) (success bool) {
	for i, s2 := range d.store {
		if d.equalityTest(s, s2) {
			d.removeOneFromIndex(i)
			success = true
			break
		}
	}

	return
}

// storeMutex should be locked before calling;
// equalityTestMutex should be read-locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOne(value interface{}) bool {
	s := newStoreDatum(value)
	return d.removeOneFromDatum(s)
}

func (d *datum) Remove(values ...interface{}) {
	defer d.equalityTestMutex.RUnlock()
	defer d.storeMutex.Unlock()
	d.equalityTestMutex.RLock()
	d.storeMutex.Lock()

	var modified bool

	for _, v := range values {
		if d.removeOne(v) {
			modified = true
		}
	}

	if modified {
		d.clearCachedHash()
	}
}

// storeMutex should be read-locked before calling
func (d *datum) contains(value interface{}) bool {
	s1 := newStoreDatum(value)

	defer d.equalityTestMutex.RUnlock()
	d.equalityTestMutex.RLock()

	for _, s2 := range d.store {
		if d.equalityTest(s1, s2) {
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
		d.removeOneFromIndex(index)
		d.clearCachedHash()
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

// storeMutex should be locked before calling
func (d *datum) makeStore(capacity int) {
	d.store = make(storeData, 0, capacity)
}

// storeMutex should be locked before calling
func (d *datum) clear(capacity int) {
	d.makeStore(capacity)
	d.clearCachedHash()
}

func (d *datum) Clear(keepCapacity bool) {
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()

	var capacity int

	if keepCapacity {
		capacity = cap(d.store)
	}

	d.clear(capacity)
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

	for i, s := range d.store {
		d.store[i].value = callback(s.value)
	}
}

type FilterCallback (func(interface{}) bool)

func (d *datum) FilterCallback(callback FilterCallback) {
	defer d.storeMutex.Unlock()
	d.storeMutex.Lock()

	var modified bool

	for i, v := range d.store {
		if !callback(v) {
			d.removeOneFromIndex(i)
			modified = true
		}
	}

	if modified {
		d.clearCachedHash()
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

	for _, s := range d.store {
		if value, ok := s.value.(int); ok {
			accumulator += value
		}
	}

	return
}

func (d *datum) IntProduct() (accumulator int) {
	accumulator = 1

	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	for _, s := range d.store {
		if value, ok := s.value.(int); ok {
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
		value = hash.Get(d.store)
		d.cachedHash = &value
	}

	return
}

func (d *datum) Hash() uint64 {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	return d.hash()
}

func (d *datum) Slice() (values []interface{}) {
	defer d.storeMutex.RUnlock()
	d.storeMutex.RLock()

	l := len(d.store)
	values = make([]interface{}, 0, l)

	for l--; l >= 0; l-- {
		values = append(values, d.store[l].value)
	}

	return
}

// storeMutex should be read-locked before calling
func (d *datum) valueStringFromIndex(i int) (s string) {
	if i >= 0 && i < len(d.store) {
		s = fmt.Sprintf("%v", d.store[i].value)
	}

	return
}

func (d *datum) String() string {
	var builder strings.Builder

	builder.WriteByte('[')

	func() {
		defer d.storeMutex.RUnlock()
		d.storeMutex.RLock()

		if l := len(d.store); l > 0 {
			builder.WriteString(d.valueStringFromIndex(0))

			for l--; l > 0; l-- {
				builder.WriteByte(' ')
				builder.WriteString(d.valueStringFromIndex(l))
			}
		}
	}()

	builder.WriteByte(']')

	return builder.String()
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

// seed the random-number generator for use in the pick method
func init() {
	rand.Seed(time.Now().UnixNano())
}
