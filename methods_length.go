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

func (d *Datum) Empty() bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store) == 0
}
