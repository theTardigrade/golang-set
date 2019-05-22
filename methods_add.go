package set

// mutex should be locked before calling;
// clearCachedFields method should be called afterwards;
// s.Value should not equal nil
func (d *Datum) addOneFromDatum(s *StoreDatum) (success bool) {
	if m := d.maximumValueCount; m != nil {
		if l, n := len(d.store), *m; l >= n {
			if l > n {
				d.store = d.store[:n]
			}

			return
		}
	}

	if f := d.filter; f != nil {
		if !f(s.Value) {
			return
		}
	}

	for _, s2 := range d.store {
		if d.equalityTest(s, s2) {
			if d.multiMode {
				s2.Instances++
				success = true
			}

			return
		}
	}

	d.store = append(d.store, s)
	success = true

	return
}

// mutex should be locked before calling;
// clearCachedFields method should be called afterwards
func (d *Datum) addOne(value interface{}) (success bool) {
	if value != nil {
		s := newStoreDatumWithIndex(value, len(d.store))
		success = d.addOneFromDatum(s)
	}

	return
}

func (d *Datum) Add(values ...interface{}) (success bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	if l := len(values); l > 0 {
		for l--; l >= 0; l-- {
			if d.addOne(values[l]) {
				success = true
			}
		}

		if success {
			d.clearCachedFields()
		}
	}

	return
}

func (d *Datum) AddFromSlice(values []interface{}) {
	d.Add(values...)
}

func (d *Datum) AddFromIntSlice(values []int) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromInt8Slice(values []int8) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromInt16Slice(values []int16) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromInt32Slice(values []int32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromInt64Slice(values []int64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUintSlice(values []uint) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUint8Slice(values []uint8) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUint16Slice(values []uint16) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUint32Slice(values []uint32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUint64Slice(values []uint64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromUinptrSlice(values []uintptr) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromFloat32Slice(values []float32) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromFloat64Slice(values []float64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromComplex64Slice(values []complex64) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromComplex128Slice(values []complex128) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromByteSlice(values []byte) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromRuneSlice(values []rune) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromBoolSlice(values []bool) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}

func (d *Datum) AddFromStringSlice(values []string) bool {
	l := len(values)
	i := make([]interface{}, l)

	for l--; l >= 0; l-- {
		i[l] = values[l]
	}

	return d.Add(i...)
}
