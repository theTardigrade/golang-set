package set

func (d *datum) Contains(value interface{}) bool {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return d.contains(value)
}

func (d *datum) ContainsEvery(values ...interface{}) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range values {
		if !d.contains(v) {
			return
		}
	}

	success = true
	return
}

func (d *datum) ContainsSome(values ...interface{}) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, v := range values {
		if d.contains(v) {
			success = true
			break
		}
	}

	return
}
