package set

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards;
// s.value should not equal nil
func (d *datum) addOneFromDatum(s *storeDatum) (success bool) {
	if m := d.maximumValueCount; m != nil {
		if l, n := len(d.store), *m; l >= n {
			if l > n {
				d.store = d.store[:n]
			}

			return
		}
	}

	if f := d.filter; f != nil {
		if !f(s.value) {
			return
		}
	}

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
