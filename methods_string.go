package set

import (
	"fmt"
	"sort"
	"strings"
)

// mutex should be read-locked before calling
func (d *datum) storeValueStringFromIndex(i int) (s string) {
	if i >= 0 && i < len(d.store) {
		s = fmt.Sprintf("%v", d.store[i].value)
	}

	return
}

func (d *datum) String() string {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	l := len(d.store)

	if l > 1 && !d.sorted {
		sort.Sort(d.store)
		d.sorted = true
	}

	var builder strings.Builder

	builder.WriteByte('[')

	if l > 0 {
		builder.WriteString(d.storeValueStringFromIndex(0))

		for l--; l > 0; l-- {
			builder.WriteByte(' ')
			builder.WriteString(d.storeValueStringFromIndex(l))
		}
	}

	builder.WriteByte(']')

	return builder.String()
}
