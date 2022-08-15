// This file is generated - do not edit.

package deepcopy

type Copyable interface {
	DeepCopy() interface{}
}

// DeepCopy will create a deep copy of the source object.
// Maps and slices will be taken into account when copying.
func DeepCopy(object interface{}) interface{} {
	switch t := object.(type) {
	case Copyable, *Copyable:
		var value Copyable

		if val, ok := t.(*Copyable); ok {
			value = *val
		} else {
			value = t.(Copyable)
		}

		return value.DeepCopy()

	// Deep copy for types map[string]T
	case map[string]bool, *map[string]bool:
		var value map[string]bool

		if val, ok := t.(*map[string]bool); ok {
			value = *val
		} else {
			value = t.(map[string]bool)
		}

		clone := make(map[string]bool, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]int, *map[string]int:
		var value map[string]int

		if val, ok := t.(*map[string]int); ok {
			value = *val
		} else {
			value = t.(map[string]int)
		}

		clone := make(map[string]int, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]int8, *map[string]int8:
		var value map[string]int8

		if val, ok := t.(*map[string]int8); ok {
			value = *val
		} else {
			value = t.(map[string]int8)
		}

		clone := make(map[string]int8, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]int16, *map[string]int16:
		var value map[string]int16

		if val, ok := t.(*map[string]int16); ok {
			value = *val
		} else {
			value = t.(map[string]int16)
		}

		clone := make(map[string]int16, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]int64, *map[string]int64:
		var value map[string]int64

		if val, ok := t.(*map[string]int64); ok {
			value = *val
		} else {
			value = t.(map[string]int64)
		}

		clone := make(map[string]int64, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]uint, *map[string]uint:
		var value map[string]uint

		if val, ok := t.(*map[string]uint); ok {
			value = *val
		} else {
			value = t.(map[string]uint)
		}

		clone := make(map[string]uint, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]uint8, *map[string]uint8:
		var value map[string]uint8

		if val, ok := t.(*map[string]uint8); ok {
			value = *val
		} else {
			value = t.(map[string]uint8)
		}

		clone := make(map[string]uint8, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]uint16, *map[string]uint16:
		var value map[string]uint16

		if val, ok := t.(*map[string]uint16); ok {
			value = *val
		} else {
			value = t.(map[string]uint16)
		}

		clone := make(map[string]uint16, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]uint64, *map[string]uint64:
		var value map[string]uint64

		if val, ok := t.(*map[string]uint64); ok {
			value = *val
		} else {
			value = t.(map[string]uint64)
		}

		clone := make(map[string]uint64, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]uintptr, *map[string]uintptr:
		var value map[string]uintptr

		if val, ok := t.(*map[string]uintptr); ok {
			value = *val
		} else {
			value = t.(map[string]uintptr)
		}

		clone := make(map[string]uintptr, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]float32, *map[string]float32:
		var value map[string]float32

		if val, ok := t.(*map[string]float32); ok {
			value = *val
		} else {
			value = t.(map[string]float32)
		}

		clone := make(map[string]float32, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]float64, *map[string]float64:
		var value map[string]float64

		if val, ok := t.(*map[string]float64); ok {
			value = *val
		} else {
			value = t.(map[string]float64)
		}

		clone := make(map[string]float64, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]complex64, *map[string]complex64:
		var value map[string]complex64

		if val, ok := t.(*map[string]complex64); ok {
			value = *val
		} else {
			value = t.(map[string]complex64)
		}

		clone := make(map[string]complex64, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]complex128, *map[string]complex128:
		var value map[string]complex128

		if val, ok := t.(*map[string]complex128); ok {
			value = *val
		} else {
			value = t.(map[string]complex128)
		}

		clone := make(map[string]complex128, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]string, *map[string]string:
		var value map[string]string

		if val, ok := t.(*map[string]string); ok {
			value = *val
		} else {
			value = t.(map[string]string)
		}

		clone := make(map[string]string, len(value))

		for k, v := range value {
			clone[k] = v
		}

		return clone
	case map[string]interface{}, *map[string]interface{}:
		var value map[string]interface{}

		if val, ok := t.(*map[string]interface{}); ok {
			value = *val
		} else {
			value = t.(map[string]interface{})
		}

		clone := make(map[string]interface{}, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v)
		}

		return clone
	// Deep copy for types map[string][]T
	case map[string][]bool, *map[string][]bool:
		var value map[string][]bool

		if val, ok := t.(*map[string][]bool); ok {
			value = *val
		} else {
			value = t.(map[string][]bool)
		}

		clone := make(map[string][]bool, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]bool)
		}

		return clone
	case map[string][]int, *map[string][]int:
		var value map[string][]int

		if val, ok := t.(*map[string][]int); ok {
			value = *val
		} else {
			value = t.(map[string][]int)
		}

		clone := make(map[string][]int, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]int)
		}

		return clone
	case map[string][]int8, *map[string][]int8:
		var value map[string][]int8

		if val, ok := t.(*map[string][]int8); ok {
			value = *val
		} else {
			value = t.(map[string][]int8)
		}

		clone := make(map[string][]int8, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]int8)
		}

		return clone
	case map[string][]int16, *map[string][]int16:
		var value map[string][]int16

		if val, ok := t.(*map[string][]int16); ok {
			value = *val
		} else {
			value = t.(map[string][]int16)
		}

		clone := make(map[string][]int16, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]int16)
		}

		return clone
	case map[string][]int64, *map[string][]int64:
		var value map[string][]int64

		if val, ok := t.(*map[string][]int64); ok {
			value = *val
		} else {
			value = t.(map[string][]int64)
		}

		clone := make(map[string][]int64, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]int64)
		}

		return clone
	case map[string][]uint, *map[string][]uint:
		var value map[string][]uint

		if val, ok := t.(*map[string][]uint); ok {
			value = *val
		} else {
			value = t.(map[string][]uint)
		}

		clone := make(map[string][]uint, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]uint)
		}

		return clone
	case map[string][]uint8, *map[string][]uint8:
		var value map[string][]uint8

		if val, ok := t.(*map[string][]uint8); ok {
			value = *val
		} else {
			value = t.(map[string][]uint8)
		}

		clone := make(map[string][]uint8, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]uint8)
		}

		return clone
	case map[string][]uint16, *map[string][]uint16:
		var value map[string][]uint16

		if val, ok := t.(*map[string][]uint16); ok {
			value = *val
		} else {
			value = t.(map[string][]uint16)
		}

		clone := make(map[string][]uint16, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]uint16)
		}

		return clone
	case map[string][]uint64, *map[string][]uint64:
		var value map[string][]uint64

		if val, ok := t.(*map[string][]uint64); ok {
			value = *val
		} else {
			value = t.(map[string][]uint64)
		}

		clone := make(map[string][]uint64, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]uint64)
		}

		return clone
	case map[string][]uintptr, *map[string][]uintptr:
		var value map[string][]uintptr

		if val, ok := t.(*map[string][]uintptr); ok {
			value = *val
		} else {
			value = t.(map[string][]uintptr)
		}

		clone := make(map[string][]uintptr, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]uintptr)
		}

		return clone
	case map[string][]float32, *map[string][]float32:
		var value map[string][]float32

		if val, ok := t.(*map[string][]float32); ok {
			value = *val
		} else {
			value = t.(map[string][]float32)
		}

		clone := make(map[string][]float32, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]float32)
		}

		return clone
	case map[string][]float64, *map[string][]float64:
		var value map[string][]float64

		if val, ok := t.(*map[string][]float64); ok {
			value = *val
		} else {
			value = t.(map[string][]float64)
		}

		clone := make(map[string][]float64, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]float64)
		}

		return clone
	case map[string][]complex64, *map[string][]complex64:
		var value map[string][]complex64

		if val, ok := t.(*map[string][]complex64); ok {
			value = *val
		} else {
			value = t.(map[string][]complex64)
		}

		clone := make(map[string][]complex64, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]complex64)
		}

		return clone
	case map[string][]complex128, *map[string][]complex128:
		var value map[string][]complex128

		if val, ok := t.(*map[string][]complex128); ok {
			value = *val
		} else {
			value = t.(map[string][]complex128)
		}

		clone := make(map[string][]complex128, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]complex128)
		}

		return clone
	case map[string][]string, *map[string][]string:
		var value map[string][]string

		if val, ok := t.(*map[string][]string); ok {
			value = *val
		} else {
			value = t.(map[string][]string)
		}

		clone := make(map[string][]string, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]string)
		}

		return clone
	case map[string][]interface{}, *map[string][]interface{}:
		var value map[string][]interface{}

		if val, ok := t.(*map[string][]interface{}); ok {
			value = *val
		} else {
			value = t.(map[string][]interface{})
		}

		clone := make(map[string][]interface{}, len(value))

		for k, v := range value {
			clone[k] = DeepCopy(v).([]interface{})
		}

		return clone

	// Deep copy for types []T
	case []bool, *[]bool:
		var value []bool

		if val, ok := t.(*[]bool); ok {
			value = *val
		} else {
			value = t.([]bool)
		}

		clone := make([]bool, len(value))
		copy(clone, value)

		return clone
	case []int, *[]int:
		var value []int

		if val, ok := t.(*[]int); ok {
			value = *val
		} else {
			value = t.([]int)
		}

		clone := make([]int, len(value))
		copy(clone, value)

		return clone
	case []int8, *[]int8:
		var value []int8

		if val, ok := t.(*[]int8); ok {
			value = *val
		} else {
			value = t.([]int8)
		}

		clone := make([]int8, len(value))
		copy(clone, value)

		return clone
	case []int16, *[]int16:
		var value []int16

		if val, ok := t.(*[]int16); ok {
			value = *val
		} else {
			value = t.([]int16)
		}

		clone := make([]int16, len(value))
		copy(clone, value)

		return clone
	case []int64, *[]int64:
		var value []int64

		if val, ok := t.(*[]int64); ok {
			value = *val
		} else {
			value = t.([]int64)
		}

		clone := make([]int64, len(value))
		copy(clone, value)

		return clone
	case []uint, *[]uint:
		var value []uint

		if val, ok := t.(*[]uint); ok {
			value = *val
		} else {
			value = t.([]uint)
		}

		clone := make([]uint, len(value))
		copy(clone, value)

		return clone
	case []uint8, *[]uint8:
		var value []uint8

		if val, ok := t.(*[]uint8); ok {
			value = *val
		} else {
			value = t.([]uint8)
		}

		clone := make([]uint8, len(value))
		copy(clone, value)

		return clone
	case []uint16, *[]uint16:
		var value []uint16

		if val, ok := t.(*[]uint16); ok {
			value = *val
		} else {
			value = t.([]uint16)
		}

		clone := make([]uint16, len(value))
		copy(clone, value)

		return clone
	case []uint64, *[]uint64:
		var value []uint64

		if val, ok := t.(*[]uint64); ok {
			value = *val
		} else {
			value = t.([]uint64)
		}

		clone := make([]uint64, len(value))
		copy(clone, value)

		return clone
	case []uintptr, *[]uintptr:
		var value []uintptr

		if val, ok := t.(*[]uintptr); ok {
			value = *val
		} else {
			value = t.([]uintptr)
		}

		clone := make([]uintptr, len(value))
		copy(clone, value)

		return clone
	case []float32, *[]float32:
		var value []float32

		if val, ok := t.(*[]float32); ok {
			value = *val
		} else {
			value = t.([]float32)
		}

		clone := make([]float32, len(value))
		copy(clone, value)

		return clone
	case []float64, *[]float64:
		var value []float64

		if val, ok := t.(*[]float64); ok {
			value = *val
		} else {
			value = t.([]float64)
		}

		clone := make([]float64, len(value))
		copy(clone, value)

		return clone
	case []complex64, *[]complex64:
		var value []complex64

		if val, ok := t.(*[]complex64); ok {
			value = *val
		} else {
			value = t.([]complex64)
		}

		clone := make([]complex64, len(value))
		copy(clone, value)

		return clone
	case []complex128, *[]complex128:
		var value []complex128

		if val, ok := t.(*[]complex128); ok {
			value = *val
		} else {
			value = t.([]complex128)
		}

		clone := make([]complex128, len(value))
		copy(clone, value)

		return clone
	case []string, *[]string:
		var value []string

		if val, ok := t.(*[]string); ok {
			value = *val
		} else {
			value = t.([]string)
		}

		clone := make([]string, len(value))
		copy(clone, value)

		return clone
	case []interface{}, *[]interface{}:
		var value []interface{}

		if val, ok := t.(*[]interface{}); ok {
			value = *val
		} else {
			value = t.([]interface{})
		}

		clone := make([]interface{}, len(value))
		copy(clone, value)

		for k, v := range clone {
			clone[k] = DeepCopy(v)
		}

		return clone
	}

	return object
}
