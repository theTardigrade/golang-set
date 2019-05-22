package set

func (d *datum) Len() (value int) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store)
}

func (d *datum) Empty() bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return len(d.store) == 0
}
