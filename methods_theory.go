package set

func (d *Datum) Union(d2 *Datum) {
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
		d.clearCachedFields()
	}
}

func (d *Datum) Intersection(d2 *Datum) {
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
		d.clearCachedFields()
	}
}

func (d *Datum) Difference(d2 *Datum) {
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
	d.clearCachedFields()
}

func (d *Datum) Subset(d2 *Datum) bool {
	return Subset(d, d2)
}
