package set

func Union(d1, d2 *datum) (d3 *datum) {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3 = New()

	if len(d2.store) > len(d1.store) {
		d2, d1 = d1, d2
	}

	d3.store = d1.store[:]

	for _, s := range d2.store {
		d3.addOneFromDatum(s)
	}

	return
}

func Intersection(d1, d2 *datum) (d3 *datum) {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3 = New()

	if len(d2.store) > len(d1.store) {
		d2, d1 = d1, d2
	}

	for _, v := range d1.store {
		if d2.contains(v) {
			d3.addOne(v)
		}
	}

	return
}

func Difference(d1, d2 *datum) (d3 *datum) {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3 = New()

	for _, v := range d1.store {
		if !d2.contains(v) {
			d3.addOne(v)
		}
	}

	return
}

func Subset(d1, d2 *datum) bool {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	for _, v := range d1.store {
		if !d2.contains(v) {
			return false
		}
	}

	return true
}
