package set

import (
	"github.com/theTardigrade/set/internal/hash"
)

func newStoreDatum(value interface{}) *StoreDatum {
	return &StoreDatum{
		Value: value,
		Hash:  hash.Get(value),
	}
}

func newStoreDatumWithIndex(value interface{}, index int) *StoreDatum {
	s := newStoreDatum(value)
	s.index = &index
	return s
}

func New() *datum {
	return &datum{
		equalityTest: DefaultEqualityTest,
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
