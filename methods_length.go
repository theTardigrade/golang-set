package set

func (d *datum) Len(includeInstances bool) (value int) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	if includeInstances {
		for _, s := range d.store {
			value += s.Instances
		}
	} else {
		value = len(d.store)
	}

	return
}

func (d *datum) Empty() bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store) == 0
}
