package set

func Clone(d *Datum) (d2 *Datum) {
	d2 = NewWithCapacity(d.Cap())

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		d2.store = append(d2.store, s)
	}

	d2.copyConfig(d)

	d.cachedHash = d2.cachedHash
	d.cachedInstancesLen = d2.cachedInstancesLen
	d.sorted = d2.sorted

	return
}
