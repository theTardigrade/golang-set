package hash

import (
	"sync"

	"github.com/mitchellh/hashstructure"
)

const (
	maxHashCacheLen = 1 << 12
)

type hashCacheType map[interface{}]uint64

var (
	hashCacheBackup hashCacheType
	hashCache       = make(hashCacheType)
	hashCacheMutex  sync.Mutex
)

func Get(value interface{}) (hashedValue uint64) {
	var useCache bool

	defer hashCacheMutex.Unlock()
	hashCacheMutex.Lock()

	switch value.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		string, bool:
		{
			var found bool

			if hashedValue, found = hashCache[value]; found {
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
		if l := len(hashCache); l >= maxHashCacheLen {
			hashCache = make(hashCacheType)

			for k, v := range hashCacheBackup {
				hashCache[k] = v
			}

			hashCacheBackup = nil
		} else if hashCacheBackup == nil && l >= maxHashCacheLen/2 {
			hashCacheBackup = make(hashCacheType)

			for k, v := range hashCache {
				hashCacheBackup[k] = v
			}
		}

		hashCache[value] = hashedValue
	}

	return
}
