package set

func Clone(d *Datum) (d2 *Datum) {
	d2 = New()

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	d2.store = d.store[:]
	d2.copyConfig(d)
	d.cachedHash = d2.cachedHash
	d.sorted = d2.sorted

	return
}
