package set

func (d *datum) Union(d2 *datum) {
	var modified bool

	defer d2.mutex.RUnlock()
	defer d.mutex.Unlock()
	d2.mutex.RLock()
	d.mutex.Lock()

	for _, s := range d2.store {
		if d.addOneFromDatum(s) {
			modified = true
		}
	}

	if modified {
		d.clearCachedHash()
	}
}

func (d *datum) Intersection(d2 *datum) {
	var modified bool

	defer d2.mutex.RUnlock()
	defer d.mutex.Unlock()
	d2.mutex.RLock()
	d.mutex.Lock()

	for _, s := range d.store {
		if !d2.containsFromDatum(s) {
			if d.removeOneFromDatum(s) {
				modified = true
			}
		}
	}

	if modified {
		d.clearCachedHash()
	}
}

func (d *datum) Difference(d2 *datum) {
	var data storeData

	defer d2.mutex.RUnlock()
	defer d.mutex.Unlock()
	d2.mutex.RLock()
	d.mutex.Lock()

	for _, s := range d2.store {
		if !d.containsFromDatum(s) {
			data = append(data, s)
		}
	}

	d.store = data
	d.clearCachedHash()
}

func (d *datum) Subset(d2 *datum) bool {
	return Subset(d, d2)
}
