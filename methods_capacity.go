package set

import (
	"math"
	"sort"
)

func (d *Datum) Cap() int {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	return cap(d.store)
}

// mutex should be locked before calling
func (d *Datum) setCapacity(newCapacity int) (success bool) {
	if newCapacity < 0 {
		return
	}

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
	} else if newCapacity < oldCapacity {
		d.makeStore(newCapacity)

		if newCapacity > 0 {
			if !d.sorted {
				sort.Sort(oldStore)
				d.sorted = true
			}

			l := int(math.Min(
				float64(newCapacity),
				float64(len(oldStore)),
			))

			for i := 0; i < l; i++ {
				d.store = append(d.store, oldStore[i])
			}
		}

		success = true
	}

	return
}

func (d *Datum) SetCapacity(newCapacity int) bool {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.setCapacity(newCapacity)
}
