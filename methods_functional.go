package set

import "reflect"

type ForEachCallback (func(interface{}))

func (d *datum) ForEach(callback ForEachCallback) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		callback(s.value)
	}
}

type MapCallback (func(interface{}) interface{})

func (d *datum) Map(callback MapCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	for i, s := range d.store {
		d.store[i].value = callback(s.value)
	}
}

type FilterCallback (func(interface{}) bool)

func (d *datum) Filter(callback FilterCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var modified bool

	for i, s := range d.store {
		if !callback(s.value) {
			d.removeOneFromIndex(i)
			modified = true
		}
	}

	if modified {
		d.clearCachedHash()
	}
}

type ReduceCallback (func(interface{}, interface{}) interface{})

func (d *datum) Reduce(initialValue interface{}, callback ReduceCallback) (accumulator interface{}) {
	accumulator = initialValue

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		accumulator = callback(accumulator, s.value)
	}

	return
}

var (
	int64Type = reflect.TypeOf(int64(0))
)

func (d *datum) Int64Sum() (accumulator int64) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue int64

		switch value := s.value.(type) {
		case int64:
			newValue = value
		case int, int8, int16, int32,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(int64Type) {
				newValue = v.Convert(int64Type).Int()
			}
		default:
			continue
		}

		accumulator += newValue
	}

	return
}

func (d *datum) Int64Product() (accumulator int64) {
	accumulator = 1

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue int64

		switch value := s.value.(type) {
		case int64:
			newValue = value
		case int, int8, int16, int32,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(int64Type) {
				newValue = v.Convert(int64Type).Int()
			}
		default:
			continue
		}

		accumulator *= newValue
	}

	return
}

var (
	float64Type = reflect.TypeOf(float64(0))
)

func (d *datum) Float64Sum() (accumulator float64) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue float64

		switch value := s.value.(type) {
		case float64:
			newValue = value
		case float32,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(float64Type) {
				newValue = v.Convert(float64Type).Float()
			}
		default:
			continue
		}

		accumulator += newValue
	}

	return
}

func (d *datum) Float64Product() (accumulator float64) {
	accumulator = 1

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		var newValue float64

		switch value := s.value.(type) {
		case float64:
			newValue = value
		case float32,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			if v := reflect.ValueOf(value); v.Type().ConvertibleTo(float64Type) {
				newValue = v.Convert(float64Type).Float()
			}
		default:
			continue
		}

		accumulator *= newValue
	}

	return
}

type EveryCallback FilterCallback

func (d *datum) Every(callback EveryCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		if !callback(s.value) {
			return
		}
	}

	success = true
	return
}

type SomeCallback EveryCallback

func (d *datum) Some(callback SomeCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		if callback(s.value) {
			success = true
			break
		}
	}

	return
}
