package hash

import "github.com/mitchellh/hashstructure"

func Get(value interface{}) (hashedValue uint64) {
	hashedValue, err := hashstructure.Hash(value, nil)
	if err != nil {
		panic(err)
	}

	return
}
