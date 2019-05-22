package set

// mutex should be read-locked before calling
func (d *Datum) instancesFromDatum(s *StoreDatum) (n int) {
	if d.multiMode {
		for _, s2 := range d.store {
			if d.equalityTest(s, s2) {
				n = s2.Instances
				break
			}
		}
	} else if d.containsFromDatum(s) {
		n = 1
	}

	return
}

// mutex should be read-locked before calling
func (d *Datum) instances(value interface{}) int {
	s1 := newStoreDatum(value)
	return d.instancesFromDatum(s1)
}

func (d *Datum) Instances(value interface{}) int {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return d.instances(value)
}
