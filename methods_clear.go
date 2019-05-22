package set

// mutex should be locked before calling
func (d *Datum) makeStore(capacity int) {
	d.store = make(storeData, 0, capacity)
}

// mutex should be locked before calling
func (d *Datum) clearCachedHash() {
	d.cachedHash = nil
	d.sorted = false
}

// mutex should be locked before calling
func (d *Datum) clear(capacity int) {
	d.makeStore(capacity)
	d.clearCachedHash()
}

func (d *Datum) Clear(keepCapacity bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var capacity int

	if keepCapacity {
		capacity = cap(d.store)
	}

	d.clear(capacity)
}
