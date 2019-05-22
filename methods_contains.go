package set

// mutex should be read-locked before calling
func (d *Datum) containsFromDatum(s *StoreDatum) (success bool) {
	for _, s2 := range d.store {
		if d.equalityTest(s, s2) {
			success = true
			break
		}
	}

	return
}

// mutex should be read-locked before calling
func (d *Datum) contains(value interface{}) bool {
	s1 := newStoreDatum(value)
	return d.containsFromDatum(s1)
}

func (d *Datum) Contains(value interface{}) bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return d.contains(value)
}

func (d *Datum) ContainsEvery(values ...interface{}) (success bool) {
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

func (d *Datum) ContainsSome(values ...interface{}) (success bool) {
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
