package set

import (
	"github.com/theTardigrade/set/internal/hash"
)

func newStoreDatum(value interface{}) *StoreDatum {
	return &StoreDatum{
		Value:     value,
		Hash:      hash.Get(value),
		Instances: 1,
	}
}

func newStoreDatumWithIndex(value interface{}, index int) *StoreDatum {
	s := newStoreDatum(value)
	s.index = &index
	return s
}

func New() *Datum {
	return &Datum{
		equalityTest: DefaultEqualityTest,
	}
}

func NewWithCapacity(c int) (d *Datum) {
	d = New()

	d.makeStore(c)

	return
}

func NewFromSlice(s []interface{}) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromBoolSlice(s []bool) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt8Slice(s []int8) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt16Slice(s []int16) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt32Slice(s []int32) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromInt64Slice(s []int64) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromIntSlice(s []int) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint8Slice(s []uint8) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint16Slice(s []uint16) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint32Slice(s []uint32) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUint64Slice(s []uint64) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUintSlice(s []uint) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromUintptrSlice(s []uintptr) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromStringSlice(s []string) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromByteSlice(s []byte) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromRuneSlice(s []rune) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromFloat32Slice(s []float32) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromFloat64Slice(s []float64) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromComplex64Slice(s []complex64) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}

func NewFromComplex128Slice(s []complex128) (d *Datum) {
	l := len(s)
	d = NewWithCapacity(l)

	for l--; l >= 0; l-- {
		d.addOne(s[l])
	}

	return
}
