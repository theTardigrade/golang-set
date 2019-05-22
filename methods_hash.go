package set

import (
	"github.com/theTardigrade/set/internal/hash"
)

// mutex should be locked before calling
func (d *datum) hash() (value uint64) {
	if d.cachedHash != nil {
		value = *d.cachedHash
	} else {
		value = hash.Get(d.store)
		d.cachedHash = &value
	}

	return
}

func (d *datum) Hash() uint64 {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.hash()
}
