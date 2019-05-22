package set

func Union(d1, d2 *datum) (d3 *datum) {
	d3 = New()

	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3.copyConfig(d1)

	for _, s := range d1.store {
		d3.addOneFromDatum(s)
	}

	for _, s := range d2.store {
		d3.addOneFromDatum(s)
	}

	return
}

func Intersection(d1, d2 *datum) (d3 *datum) {
	d3 = New()

	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3.copyConfig(d1)

	if len(d2.store) < len(d1.store) {
		d2, d1 = d1, d2
	}

	for _, s := range d1.store {
		if d2.containsFromDatum(s) {
			d3.addOneFromDatum(s)
		}
	}

	return
}

func Difference(d1, d2 *datum) (d3 *datum) {
	d3 = New()

	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	d3.copyConfig(d1)

	for _, s := range d2.store {
		if !d1.containsFromDatum(s) {
			d3.addOneFromDatum(s)
		}
	}

	return
}

func Subset(d1, d2 *datum) (success bool) {
	defer d1.mutex.RUnlock()
	defer d2.mutex.RUnlock()
	d1.mutex.RLock()
	d2.mutex.RLock()

	for _, s := range d1.store {
		if !d2.containsFromDatum(s) {
			return
		}
	}

	success = true
	return
}
