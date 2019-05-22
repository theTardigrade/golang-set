package set

// mutex should be locked before calling
func (d *datum) removeDuplicates() (success bool) {
	if s := d.store; s != nil {
		if l := len(s); l > 0 {
			for i := 0; i < l; i++ {
				for j := i + 1; j < l; j++ {
					if d.equalityTest(s[i], s[j]) {
						d.removeOneFromIndex(j)
						l--
						j--
						if d.multiMode {
							s[i].Instances++
						}
						success = true
					}
				}
			}
		}
	}

	if success {
		d.clearCachedHash()
	}

	return
}

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOneFromIndex(i int) {
	if j := len(d.store) - 1; i <= j {
		s := d.store[i]
		if d.multiMode {
			if s.Instances--; s.Instances > 0 {
				return
			}
		}

		d.store[j], d.store[i] = d.store[i], d.store[j]
		d.store = d.store[:j]
	}
}

// mutex should be locked before calling;
// clearCachedHash method should be called afterwards
func (d *datum) removeOneFromDatum(s *StoreDatum) (success bool) {
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
