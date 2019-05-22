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

	if s := d.store; s != nil {
		if l := len(s); l > n {
			d.store = s[:n]
		}
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

	if s := d.store; s != nil {
		if l := len(d.store); l > 0 {
			var modified bool

			for l--; l >= 0; l-- {
				if v := s[l].Value; !f(v) {
					d.removeOneFromIndex(l)
					modified = true
				}
			}

			if modified {
				d.clearCachedHash()
			}
		}
	}

	return
}
