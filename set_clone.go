package set

func Clone(d *datum) (d2 *datum) {
	d2 = New()

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	d2.store = d.store[:]
	d2.cachedHash = d.cachedHash
	d2.equalityTest = d.equalityTest

	return
}
