package deepcopy

// Slice is a shortcut for []interface{}
// which implements the Copyable interface.
type Slice []interface{}

// DeepCopy will create a deep copy of this slice.
// Maps and slices will be taken into account when copying.
func (s Slice) DeepCopy() interface{} {
	clone := make([]interface{}, len(s))
	copy(clone, s)

	for k, v := range clone {
		clone[k] = DeepCopy(v)
	}

	return clone
}

// Clone is a wrapper for Slice.DeepCopy() that returns
// the correct type for convenience.
func (s Slice) Clone() []interface{} {
	return s.DeepCopy().([]interface{})
}
