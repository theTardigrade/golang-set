package set

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/theTardigrade/set/internal/hash"
)

func (d *datum) SetEqualityTest(equalityTest equalityTestFunc) {
	if equalityTest == nil {
		return
	}

	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.equalityTest = equalityTest
}

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards;
// s.value should not equal nil
func (d *datum) addOneFromDatum(s *storeDatum) (success bool) {
	for _, s2 := range d.store {
		if d.equalityTest(s, s2) {
			return
		}
	}

	d.store = append(d.store, s)
	success = true

	return
}

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) addOne(value interface{}) (success bool) {
	if value != nil {
		s := newStoreDatumWithIndex(value, len(d.store))
		success = d.addOneFromDatum(s)
	}

	return
}

// mutex should be locked before calling
func (d *datum) clearCachedHash() {
	d.cachedHash = nil
	d.sorted = false
}

func (d *datum) Add(values ...interface{}) (success bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	if l := len(values); l > 0 {
		for l--; l >= 0; l-- {
			if d.addOne(values[l]) {
				success = true
			}
		}

		if success {
			d.clearCachedHash()
		}
	}

	return
}

func (d *datum) AddFromSlice(values []interface{}) {
	d.Add(values...)
}

func (d *datum) AddFromIntSlice(values []int) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromInt8Slice(values []int8) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromInt16Slice(values []int16) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromInt32Slice(values []int32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromInt64Slice(values []int64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUintSlice(values []uint) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUint8Slice(values []uint8) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUint16Slice(values []uint16) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUint32Slice(values []uint32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUint64Slice(values []uint64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromUinptrSlice(values []uintptr) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromFloat32Slice(values []float32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromFloat64Slice(values []float64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromComplex64Slice(values []complex64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromComplex128Slice(values []complex128) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromByteSlice(values []byte) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromRuneSlice(values []rune) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromBoolSlice(values []bool) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *datum) AddFromStringSlice(values []string) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOneFromIndex(i int) {
	if j := len(d.store) - 1; i <= j {
		d.store[j], d.store[i] = d.store[i], d.store[j]
		d.store = d.store[:j]
	}
}

// mutex should be locked before calling;
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

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOne(value interface{}) bool {
	s := newStoreDatum(value)
	return d.removeOneFromDatum(s)
}

func (d *datum) Remove(values ...interface{}) (success bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	for _, v := range values {
		if d.removeOne(v) {
			success = true
		}
	}

	if success {
		d.clearCachedHash()
	}

	return
}

// mutex should be read-locked before calling
func (d *datum) contains(value interface{}) (success bool) {
	s1 := newStoreDatum(value)

	for _, s2 := range d.store {
		if d.equalityTest(s1, s2) {
			success = true
			break
		}
	}

	return
}

func (d *datum) Contains(value interface{}) bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return d.contains(value)
}

func (d *datum) ContainsEvery(values ...interface{}) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range values {
		if !d.contains(v) {
			return
		}
	}

	success = true
	return
}

func (d *datum) ContainsSome(values ...interface{}) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range values {
		if d.contains(v) {
			success = true
			break
		}
	}

	return
}

func (d *datum) Pop() (value interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	value, index := d.pick()
	if value != nil {
		d.removeOneFromIndex(index)
		d.clearCachedHash()
	}

	return
}

// mutex should be read-locked before calling
func (d *datum) pick() (value interface{}, index int) {
	if l := len(d.store); l > 0 {
		index = rand.Intn(l)
		value = d.store[index]
	}

	return
}

func (d *datum) Pick() (value interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	value, _ = d.pick()

	return
}

// mutex should be locked before calling
func (d *datum) makeStore(capacity int) {
	d.store = make(storeData, 0, capacity)
}

// mutex should be locked before calling
func (d *datum) clear(capacity int) {
	d.makeStore(capacity)
	d.clearCachedHash()
}

func (d *datum) Clear(keepCapacity bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var capacity int

	if keepCapacity {
		capacity = cap(d.store)
	}

	d.clear(capacity)
}

type ForEachCallback (func(interface{}))

func (d *datum) ForEach(callback ForEachCallback) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range d.store {
		callback(v)
	}
}

type MapCallback (func(interface{}) interface{})

func (d *datum) Map(callback MapCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	for i, s := range d.store {
		d.store[i].value = callback(s.value)
	}
}

type FilterCallback (func(interface{}) bool)

func (d *datum) FilterCallback(callback FilterCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

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

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range d.store {
		accumulator = callback(accumulator, v)
	}

	return
}

var (
	int64Type = reflect.TypeOf(int64(0))
)

func (d *datum) Int64Sum() (accumulator int64) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue int64

		switch value := s.value.(type) {
		case int64:
			newValue = value
		case int, int8, int16, int32,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(int64Type) {
				newValue = v.Convert(int64Type).Int()
			}
		default:
			continue
		}

		accumulator += newValue
	}

	return
}

func (d *datum) Int64Product() (accumulator int64) {
	accumulator = 1

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue int64

		switch value := s.value.(type) {
		case int64:
			newValue = value
		case int, int8, int16, int32,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(int64Type) {
				newValue = v.Convert(int64Type).Int()
			}
		default:
			continue
		}

		accumulator *= newValue
	}

	return
}

var (
	float64Type = reflect.TypeOf(float64(0))
)

func (d *datum) Float64Sum() (accumulator float64) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue float64

		switch value := s.value.(type) {
		case float64:
			newValue = value
		case float32,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(float64Type) {
				newValue = v.Convert(float64Type).Float()
			}
		default:
			continue
		}

		accumulator += newValue
	}

	return
}

func (d *datum) Float64Product() (accumulator float64) {
	accumulator = 1

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue float64

		switch value := s.value.(type) {
		case float64:
			newValue = value
		case float32,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(float64Type) {
				newValue = v.Convert(float64Type).Float()
			}
		default:
			continue
		}

		accumulator *= newValue
	}

	return
}

type EveryCallback FilterCallback

func (d *datum) Every(callback EveryCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range d.store {
		if !callback(v) {
			return
		}
	}

	success = true
	return
}

type SomeCallback EveryCallback

func (d *datum) Some(callback SomeCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range d.store {
		if callback(v) {
			success = true
			break
		}
	}

	return
}

func (d *datum) Equal(d2 *datum) bool {
	if d2 == d {
		return true
	}

	defer d.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d.mutex.RLock()
	d2.mutex.RLock()

	l, l2 := len(d.store), len(d2.store)

	if l != l2 {
		return false
	}

	if l == 0 {
		return true
	}

	return d.hash() == d2.hash()
}

// mutex should be locked before calling
func (d *datum) hash() (value uint64) {
	if d.cachedHash != nil {
		value = *d.cachedHash
	} else {
		value = hash.Get(d.store)
		d.cachedHash = &value
	}

	return
}

func (d *datum) Hash() uint64 {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.hash()
}

func (d *datum) Slice() (values []interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	l := len(d.store)
	values = make([]interface{}, 0, l)

	for l--; l >= 0; l-- {
		values = append(values, d.store[l].value)
	}

	return
}

// mutex should be read-locked before calling
func (d *datum) storeValueStringFromIndex(i int) (s string) {
	if i >= 0 && i < len(d.store) {
		s = fmt.Sprintf("%v", d.store[i].value)
	}

	return
}

func (d *datum) String() string {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	l := len(d.store)

	if l > 1 {
		if !d.sorted {
			sort.Sort(d.store)
			d.sorted = true
		}
	}

	var builder strings.Builder

	builder.WriteByte('[')

	if l > 0 {
		builder.WriteString(d.storeValueStringFromIndex(0))

		for l--; l > 0; l-- {
			builder.WriteByte(' ')
			builder.WriteString(d.storeValueStringFromIndex(l))
		}
	}

	builder.WriteByte(']')

	return builder.String()
}

func (d *datum) Empty() bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store) == 0
}

func (d *datum) Len() (value int) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

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
