package set

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
