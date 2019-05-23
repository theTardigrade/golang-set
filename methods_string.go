package set

import (
	"fmt"
	"sort"
	"strings"
)

func (d *Datum) valueStringFromIndex(i int) string {
	s := d.store[i]

	if s.valueString == nil {
		v := fmt.Sprintf("%v", s.Value)
		s.valueString = &v
	}

	return *s.valueString
}

func (d *Datum) String() string {
	var builder strings.Builder

	defer d.mutex.Unlock()
	d.mutex.Lock()

	l := len(d.store)

	if l > 1 && !d.sorted {
		sort.Sort(d.store)
		d.sorted = true
	}

	builder.Grow(l * 8)
	builder.WriteByte('[')

	if l > 0 {
		builder.WriteString(d.valueStringFromIndex(0))

		for l--; l > 0; l-- {
			builder.WriteByte(' ')
			builder.WriteString(d.valueStringFromIndex(l))
		}
	}

	builder.WriteByte(']')

	return builder.String()
}
