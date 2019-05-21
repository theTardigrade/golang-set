package hash

import (
	"sync"

	"github.com/mitchellh/hashstructure"
)

const (
	cacheMaxLen = 1 << 12
)

type cacheType map[interface{}]uint64

var (
	cacheBackup cacheType
	cache       = make(cacheType)
	cacheMutex  sync.Mutex
)

func Get(value interface{}) (hashedValue uint64) {
	var useCache bool

	defer cacheMutex.Unlock()
	cacheMutex.Lock()

	switch value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		{
			var found bool

			if hashedValue, found = cache[value]; found {
				return
			}

			useCache = true
		}
	}

	hashedValue, err := hashstructure.Hash(value, nil)
	if err != nil {
		panic(err)
	}

	if useCache {
		if l, m := len(cache), cacheMaxLen; l >= m {
			cache = make(cacheType)

			for k, v := range cacheBackup {
				cache[k] = v
			}

			cacheBackup = nil
		} else if cacheBackup == nil {
			if m /= 2; l >= m {
				cacheBackup = make(cacheType)

				for k, v := range cache {
					cacheBackup[k] = v
				}
			}
		}

		cache[value] = hashedValue
	}

	return
}
