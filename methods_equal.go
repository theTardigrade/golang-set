package set

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
