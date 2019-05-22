package set

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
