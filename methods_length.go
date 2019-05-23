package set

func (d *Datum) Len(includeInstances bool) (value int) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	if includeInstances && d.multiMode {
		if d.cachedInstancesLen != nil {
			value = *d.cachedInstancesLen
		} else {
			for _, s := range d.store {
				value += s.Instances
			}

			d.cachedInstancesLen = &value
		}
	} else {
		value = len(d.store)
	}

	return
}

func (d *Datum) Cap() int {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return cap(d.store)
}

func (d *Datum) Grow(newCapacity int) (success bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var oldCapacity int
	oldStore := d.store

	if oldStore != nil {
		oldCapacity = cap(oldStore)
	}

	if newCapacity > oldCapacity {
		d.makeStore(newCapacity)

		for _, s := range oldStore {
			d.store = append(d.store, s)
		}

		success = true
	}

	return
}

func (d *Datum) Empty() bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store) == 0
}
