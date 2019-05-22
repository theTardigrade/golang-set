package set

import "reflect"

type forEachCallback (func(interface{}))

func (d *datum) ForEach(callback forEachCallback) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		callback(s.Value)
	}
}

type mapCallback (func(interface{}) interface{})

func (d *datum) Map(callback mapCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	for i, s := range d.store {
		d.store[i].Value = callback(s.Value)
	}
}

type filterCallback (func(interface{}) bool)

func (d *datum) Filter(callback filterCallback) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var modified bool

	for i, s := range d.store {
		if !callback(s.Value) {
			d.removeOneFromIndex(i)
			modified = true
		}
	}

	if modified {
		d.clearCachedHash()
	}
}

type reduceCallback (func(interface{}, interface{}) interface{})

func (d *datum) Reduce(initialValue interface{}, callback reduceCallback) (accumulator interface{}) {
	accumulator = initialValue

	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		accumulator = callback(accumulator, s.Value)
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

		switch value := s.Value.(type) {
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

		switch value := s.Value.(type) {
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

		switch value := s.Value.(type) {
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

		switch value := s.Value.(type) {
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

type everyCallback filterCallback

func (d *datum) Every(callback everyCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		if !callback(s.Value) {
			return
		}
	}

	success = true
	return
}

type someCallback filterCallback

func (d *datum) Some(callback someCallback) (success bool) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	for _, s := range d.store {
		if callback(s.Value) {
			success = true
			break
		}
	}

	return
}
