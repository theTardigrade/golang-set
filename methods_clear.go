package set

// mutex should be locked before calling
func (d *datum) makeStore(capacity int) {
	d.store = make(storeData, 0, capacity)
}

// mutex should be locked before calling
func (d *datum) clearCachedHash() {
	d.cachedHash = nil
	d.sorted = false
}

// mutex should be locked before calling
func (d *datum) clear(capacity int) {
	d.makeStore(capacity)
	d.clearCachedHash()
}

func (d *datum) Clear(keepCapacity bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var capacity int

	if keepCapacity {
		capacity = cap(d.store)
	}

	d.clear(capacity)
}
