package set

func (d *datum) SetEqualityTest(equalityTest equalityTestFunc) (success bool) {
	if equalityTest != nil {
		return
	}

	success = true

	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.equalityTest = equalityTest

	return
}

func (d *datum) SetMaximumValueCount(n int) (success bool) {
	if n < 0 {
		return
	}

	success = true

	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.maximumValueCount = &n

	if l := len(d.store); l > n {
		d.store = d.store[:n]
	}

	return
}

func (d *datum) SetFilter(f filterFunc) (success bool) {
	if f == nil {
		return
	}

	success = true

	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.filter = f

	if l := len(d.store); l > 0 {
		var modified bool

		for i, s := range d.store {
			if !f(s.value) {
				d.removeOneFromIndex(i)
				modified = true
			}
		}

		if modified {
			d.clearCachedHash()
		}
	}

	return
}
