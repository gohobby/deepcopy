package deepcopy

// Map is a shortcut for map[string]interface{}
// which implements the Copyable interface.
type Map map[string]interface{}

// DeepCopy will create a deep copy of this map.
// Maps and slices will be taken into account when copying.
func (m Map) DeepCopy() interface{} {
	clone := make(map[string]interface{}, len(m))

	for k, v := range m {
		clone[k] = DeepCopy(v)
	}

	return clone
}

// Clone is a wrapper for Map.DeepCopy() that returns
// the correct type for convenience.
func (m Map) Clone() map[string]interface{} {
	return m.DeepCopy().(map[string]interface{})
}
