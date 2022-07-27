package deepcopy

import (
	"testing"

	"github.com/gohobby/assert"
)

var testCasesForMap = []struct {
	// original and expectedOriginal are the same value in each test case. We do
	// this to avoid unintentionally asserting against a mutated
	// expectedOriginal and having the test pass erroneously. We also do not
	// want to rely on the deep copy function we are testing to ensure this does
	// not happen.
	name             string
	original         map[string]interface{}
	transformer      func(m map[string]interface{}) map[string]interface{}
	expectedCopy     map[string]interface{}
	expectedOriginal map[string]interface{}
}{
	// reassignment of entire map, should be okay even without deepCopy.
	{
		name:     "1",
		original: nil,
		transformer: func(m map[string]interface{}) map[string]interface{} {
			return map[string]interface{}{}
		},
		expectedCopy:     map[string]interface{}{},
		expectedOriginal: nil,
	},
	{
		original: map[string]interface{}{},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			return nil
		},
		expectedCopy:     nil,
		expectedOriginal: map[string]interface{}{},
	},
	// mutation of map
	{
		name:     "Mutation of an empty map",
		original: map[string]interface{}{},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"] = "bar"
			return m
		},
		expectedCopy:     map[string]interface{}{"foo": "bar"},
		expectedOriginal: map[string]interface{}{},
	},
	{
		name:     "Mutation of a non-empty map",
		original: map[string]interface{}{"foo": "bar"},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"] = "baz"
			return m
		},
		expectedCopy:     map[string]interface{}{"foo": "baz"},
		expectedOriginal: map[string]interface{}{"foo": "bar"},
	},
	// mutation of nested maps
	{
		name:     "Mutation of empty nested maps",
		original: map[string]interface{}{},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"] = map[string]interface{}{"bar": "baz"}
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": map[string]interface{}{"bar": "baz"},
		},
		expectedOriginal: map[string]interface{}{},
	},
	{
		name: "Mutation of non-empty nested maps",
		original: map[string]interface{}{
			"foo": map[string]interface{}{"key1": "val1", "key2": "val2"},
			"bar": map[string]int{"key1": 1, "key2": 2},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"] = map[string]interface{}{"bar": "baz"}
			m["bar"].(map[string]int)["key1"] = 3
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": map[string]interface{}{"bar": "baz"},
			"bar": map[string]int{"key1": 3, "key2": 2},
		},
		expectedOriginal: map[string]interface{}{
			"foo": map[string]interface{}{"key1": "val1", "key2": "val2"},
			"bar": map[string]int{"key1": 1, "key2": 2},
		},
	},
	// mutation of nested map values
	{
		name: "Mutation of nested map values bool and []bool",
		original: map[string]interface{}{
			"foo": map[string]bool{"key1": true, "key2": false},
			"bar": map[string][]bool{"baz": {true, false}},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"].(map[string]bool)["key1"] = false
			m["bar"].(map[string][]bool)["baz"][0] = false
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": map[string]bool{"key1": false, "key2": false},
			"bar": map[string][]bool{"baz": {false, false}},
		},
		expectedOriginal: map[string]interface{}{
			"foo": map[string]bool{"key1": true, "key2": false},
			"bar": map[string][]bool{"baz": {true, false}},
		},
	},
	{
		name: "Mutation of nested int map values",
		original: map[string]interface{}{
			"int":     map[string]int{"key1": 1, "key2": 2},
			"int8":    map[string]int8{"key1": 1, "key2": 2},
			"int16":   map[string]int16{"key1": 1, "key2": 2},
			"int64":   map[string]int64{"key1": 1, "key2": 2},
			"uint":    map[string]uint{"key1": 1, "key2": 2},
			"uint8":   map[string]uint8{"key1": 1, "key2": 2},
			"uint16":  map[string]uint16{"key1": 1, "key2": 2},
			"uint64":  map[string]uint64{"key1": 1, "key2": 2},
			"uintptr": map[string]uintptr{"key1": 1, "key2": 2},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["int"].(map[string]int)["key1"] = 3
			m["int8"].(map[string]int8)["key1"] = 3
			m["int16"].(map[string]int16)["key1"] = 3
			m["int64"].(map[string]int64)["key1"] = 3
			m["uint"].(map[string]uint)["key1"] = 3
			m["uint8"].(map[string]uint8)["key1"] = 3
			m["uint16"].(map[string]uint16)["key1"] = 3
			m["uint64"].(map[string]uint64)["key1"] = 3
			m["uintptr"].(map[string]uintptr)["key1"] = 3
			return m
		},
		expectedCopy: map[string]interface{}{
			"int":     map[string]int{"key1": 3, "key2": 2},
			"int8":    map[string]int8{"key1": 3, "key2": 2},
			"int16":   map[string]int16{"key1": 3, "key2": 2},
			"int64":   map[string]int64{"key1": 3, "key2": 2},
			"uint":    map[string]uint{"key1": 3, "key2": 2},
			"uint8":   map[string]uint8{"key1": 3, "key2": 2},
			"uint16":  map[string]uint16{"key1": 3, "key2": 2},
			"uint64":  map[string]uint64{"key1": 3, "key2": 2},
			"uintptr": map[string]uintptr{"key1": 3, "key2": 2},
		},
		expectedOriginal: map[string]interface{}{
			"int":     map[string]int{"key1": 1, "key2": 2},
			"int8":    map[string]int8{"key1": 1, "key2": 2},
			"int16":   map[string]int16{"key1": 1, "key2": 2},
			"int64":   map[string]int64{"key1": 1, "key2": 2},
			"uint":    map[string]uint{"key1": 1, "key2": 2},
			"uint8":   map[string]uint8{"key1": 1, "key2": 2},
			"uint16":  map[string]uint16{"key1": 1, "key2": 2},
			"uint64":  map[string]uint64{"key1": 1, "key2": 2},
			"uintptr": map[string]uintptr{"key1": 1, "key2": 2},
		},
	},
	{
		name: "Mutation of nested []int map values",
		original: map[string]interface{}{
			"int":     map[string][]int{"sliceInt": {1, 2}},
			"int8":    map[string][]int8{"sliceInt8": {1, 2}},
			"int16":   map[string][]int16{"sliceInt16": {1, 2}},
			"int64":   map[string][]int64{"sliceInt64": {1, 2}},
			"uint":    map[string][]uint{"sliceUint": {1, 2}},
			"uint8":   map[string][]uint8{"sliceUint8": {1, 2}},
			"uint16":  map[string][]uint16{"sliceUint16": {1, 2}},
			"uint64":  map[string][]uint64{"sliceUint64": {1, 2}},
			"uintptr": map[string][]uintptr{"sliceUintptr": {1, 2}},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["int"].(map[string][]int)["sliceInt"][0] = 3
			m["int8"].(map[string][]int8)["sliceInt8"][0] = 3
			m["int16"].(map[string][]int16)["sliceInt16"][0] = 3
			m["int64"].(map[string][]int64)["sliceInt64"][0] = 3
			m["uint"].(map[string][]uint)["sliceUint"][0] = 3
			m["uint8"].(map[string][]uint8)["sliceUint8"][0] = 3
			m["uint16"].(map[string][]uint16)["sliceUint16"][0] = 3
			m["uint64"].(map[string][]uint64)["sliceUint64"][0] = 3
			m["uintptr"].(map[string][]uintptr)["sliceUintptr"][0] = 3
			return m
		},
		expectedCopy: map[string]interface{}{
			"int":     map[string][]int{"sliceInt": {3, 2}},
			"int8":    map[string][]int8{"sliceInt8": {3, 2}},
			"int16":   map[string][]int16{"sliceInt16": {3, 2}},
			"int64":   map[string][]int64{"sliceInt64": {3, 2}},
			"uint":    map[string][]uint{"sliceUint": {3, 2}},
			"uint8":   map[string][]uint8{"sliceUint8": {3, 2}},
			"uint16":  map[string][]uint16{"sliceUint16": {3, 2}},
			"uint64":  map[string][]uint64{"sliceUint64": {3, 2}},
			"uintptr": map[string][]uintptr{"sliceUintptr": {3, 2}},
		},
		expectedOriginal: map[string]interface{}{
			"int":     map[string][]int{"sliceInt": {1, 2}},
			"int8":    map[string][]int8{"sliceInt8": {1, 2}},
			"int16":   map[string][]int16{"sliceInt16": {1, 2}},
			"int64":   map[string][]int64{"sliceInt64": {1, 2}},
			"uint":    map[string][]uint{"sliceUint": {1, 2}},
			"uint8":   map[string][]uint8{"sliceUint8": {1, 2}},
			"uint16":  map[string][]uint16{"sliceUint16": {1, 2}},
			"uint64":  map[string][]uint64{"sliceUint64": {1, 2}},
			"uintptr": map[string][]uintptr{"sliceUintptr": {1, 2}},
		},
	},
	{
		name: "Mutation of nested map values float* and []float*",
		original: map[string]interface{}{
			"foo32": map[string]float32{"key1": 1.1, "key2": 2.2},
			"bar32": map[string][]float32{"baz32": {1.1, 2.2}},
			"foo64": map[string]float64{"key1": 1.1, "key2": 2.2},
			"bar64": map[string][]float64{"baz64": {1.1, 2.2}},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo32"].(map[string]float32)["key1"] = 3.3
			m["bar32"].(map[string][]float32)["baz32"][0] = 3.3
			m["foo64"].(map[string]float64)["key1"] = 3.3
			m["bar64"].(map[string][]float64)["baz64"][0] = 3.3
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo32": map[string]float32{"key1": 3.3, "key2": 2.2},
			"bar32": map[string][]float32{"baz32": {3.3, 2.2}},
			"foo64": map[string]float64{"key1": 3.3, "key2": 2.2},
			"bar64": map[string][]float64{"baz64": {3.3, 2.2}},
		},
		expectedOriginal: map[string]interface{}{
			"foo32": map[string]float32{"key1": 1.1, "key2": 2.2},
			"bar32": map[string][]float32{"baz32": {1.1, 2.2}},
			"foo64": map[string]float64{"key1": 1.1, "key2": 2.2},
			"bar64": map[string][]float64{"baz64": {1.1, 2.2}},
		},
	},
	{
		name: "Mutation of nested map values complex* and []complex*",
		original: map[string]interface{}{
			"foo64":  map[string]complex64{"key1": 1, "key2": 2},
			"bar64":  map[string][]complex64{"baz64": {1, 2}},
			"foo128": map[string]complex128{"key1": complex(1, 2), "key2": complex(1, 2)},
			"bar128": map[string][]complex128{"baz128": {complex(1, 2), complex(1, 2)}},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo64"].(map[string]complex64)["key1"] = 3
			m["bar64"].(map[string][]complex64)["baz64"][0] = 3
			m["foo128"].(map[string]complex128)["key1"] = complex(3, 4)
			m["bar128"].(map[string][]complex128)["baz128"][0] = complex(3, 4)
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo64":  map[string]complex64{"key1": 3, "key2": 2},
			"bar64":  map[string][]complex64{"baz64": {3, 2}},
			"foo128": map[string]complex128{"key1": complex(3, 4), "key2": complex(1, 2)},
			"bar128": map[string][]complex128{"baz128": {complex(3, 4), complex(1, 2)}},
		},
		expectedOriginal: map[string]interface{}{
			"foo64":  map[string]complex64{"key1": 1, "key2": 2},
			"bar64":  map[string][]complex64{"baz64": {1, 2}},
			"foo128": map[string]complex128{"key1": complex(1, 2), "key2": complex(1, 2)},
			"bar128": map[string][]complex128{"baz128": {complex(1, 2), complex(1, 2)}},
		},
	},
	{
		name: "Mutation of nested map values string and []string",
		original: map[string]interface{}{
			"foo": map[string]string{"key1": "val1", "key2": "val2"},
			"bar": map[string][]string{"key1": {"val1", "val2"}},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"].(map[string]string)["key1"] = "baz"
			m["bar"].(map[string][]string)["key1"][0] = "qux"
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": map[string]string{"key1": "baz", "key2": "val2"},
			"bar": map[string][]string{"key1": {"qux", "val2"}},
		},
		expectedOriginal: map[string]interface{}{
			"foo": map[string]string{"key1": "val1", "key2": "val2"},
			"bar": map[string][]string{"key1": {"val1", "val2"}},
		},
	},
	{
		name: "Mutation of nested []interface{} map values",
		original: map[string]interface{}{
			"foo": map[string][]interface{}{
				"slice": {
					[]interface{}{"bar"},
					[]interface{}{"baz"},
				},
			},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"].(map[string][]interface{})["slice"][0].([]interface{})[0] = "qux"
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": map[string][]interface{}{
				"slice": {
					[]interface{}{"qux"},
					[]interface{}{"baz"},
				},
			},
		},
		expectedOriginal: map[string]interface{}{
			"foo": map[string][]interface{}{
				"slice": {
					[]interface{}{"bar"},
					[]interface{}{"baz"},
				},
			},
		},
	},
	// mutation of nested slice values
	{
		name: "Mutation of nested bool slice values",
		original: map[string]interface{}{
			"foo": []bool{true, false},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"].([]bool)[0] = false
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": []bool{false, false},
		},
		expectedOriginal: map[string]interface{}{
			"foo": []bool{true, false},
		},
	},
	{
		name: "Mutation of nested int slice values",
		original: map[string]interface{}{
			"int":     []int{1, 2, 3},
			"int8":    []int8{1, 2, 3},
			"int16":   []int16{1, 2, 3},
			"int64":   []int64{1, 2, 3},
			"uint":    []uint{1, 2, 3},
			"uint8":   []uint8{1, 2, 3},
			"uint16":  []uint16{1, 2, 3},
			"uint64":  []uint64{1, 2, 3},
			"uintptr": []uintptr{1, 2, 3},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["int"].([]int)[0] = 4
			m["int8"].([]int8)[0] = 4
			m["int16"].([]int16)[0] = 4
			m["int64"].([]int64)[0] = 4
			m["uint"].([]uint)[0] = 4
			m["uint8"].([]uint8)[0] = 4
			m["uint16"].([]uint16)[0] = 4
			m["uint64"].([]uint64)[0] = 4
			m["uintptr"].([]uintptr)[0] = 4
			return m
		},
		expectedCopy: map[string]interface{}{
			"int":     []int{4, 2, 3},
			"int8":    []int8{4, 2, 3},
			"int16":   []int16{4, 2, 3},
			"int64":   []int64{4, 2, 3},
			"uint":    []uint{4, 2, 3},
			"uint8":   []uint8{4, 2, 3},
			"uint16":  []uint16{4, 2, 3},
			"uint64":  []uint64{4, 2, 3},
			"uintptr": []uintptr{4, 2, 3},
		},
		expectedOriginal: map[string]interface{}{
			"int":     []int{1, 2, 3},
			"int8":    []int8{1, 2, 3},
			"int16":   []int16{1, 2, 3},
			"int64":   []int64{1, 2, 3},
			"uint":    []uint{1, 2, 3},
			"uint8":   []uint8{1, 2, 3},
			"uint16":  []uint16{1, 2, 3},
			"uint64":  []uint64{1, 2, 3},
			"uintptr": []uintptr{1, 2, 3},
		},
	},
	{
		name: "Mutation of nested float slice values",
		original: map[string]interface{}{
			"float32": []float32{1.1, 2.2, 3.3},
			"float64": []float64{1.1, 2.2, 3.3},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["float32"].([]float32)[0] = 4.4
			m["float64"].([]float64)[0] = 4.4
			return m
		},
		expectedCopy: map[string]interface{}{
			"float32": []float32{4.4, 2.2, 3.3},
			"float64": []float64{4.4, 2.2, 3.3},
		},
		expectedOriginal: map[string]interface{}{
			"float32": []float32{1.1, 2.2, 3.3},
			"float64": []float64{1.1, 2.2, 3.3},
		},
	},
	{
		name: "Mutation of nested complex number slice values",
		original: map[string]interface{}{
			"complex64":  []complex64{1, 2, 3},
			"complex128": []complex128{complex(1, 2), complex(3, 4)},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["complex64"].([]complex64)[0] = 4
			m["complex128"].([]complex128)[0] = complex(5, 6)
			return m
		},
		expectedCopy: map[string]interface{}{
			"complex64":  []complex64{4, 2, 3},
			"complex128": []complex128{complex(5, 6), complex(3, 4)},
		},
		expectedOriginal: map[string]interface{}{
			"complex64":  []complex64{1, 2, 3},
			"complex128": []complex128{complex(1, 2), complex(3, 4)},
		},
	},
	{
		name: "Mutation of nested string slice values",
		original: map[string]interface{}{
			"foo": []string{"bar", "baz"},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			m["foo"].([]string)[0] = "qux"
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": []string{"qux", "baz"},
		},
		expectedOriginal: map[string]interface{}{
			"foo": []string{"bar", "baz"},
		},
	},
	{
		name: "Mutation of nested interface{} slice values",
		original: map[string]interface{}{
			"foo": []interface{}{
				true,
				123,
				3.14,
				complex(1, 2),
				"foobar",
				map[string]interface{}{"key1": 1, "key2": 2, "key3": nil},
				[]bool{true, false},
				[]int{1, 2, 3},
				[]float32{1.1, 2.2, 3.3},
				[]string{"foo", "bar"},
			},
		},
		transformer: func(m map[string]interface{}) map[string]interface{} {
			s := m["foo"].([]interface{})
			s[0] = false
			s[1] = 456
			s[2] = 6.28
			s[3] = complex(3, 4)
			s[4] = "bar"
			s[5].(map[string]interface{})["key1"] = 3
			s[6].([]bool)[0] = false
			s[7].([]int)[0] = 4
			s[8].([]float32)[0] = 4.4
			s[9].([]string)[0] = "baz"
			return m
		},
		expectedCopy: map[string]interface{}{
			"foo": []interface{}{
				false,
				456,
				6.28,
				complex(3, 4),
				"bar",
				map[string]interface{}{"key1": 3, "key2": 2, "key3": nil},
				[]bool{false, false},
				[]int{4, 2, 3},
				[]float32{4.4, 2.2, 3.3},
				[]string{"baz", "bar"},
			},
		},
		expectedOriginal: map[string]interface{}{
			"foo": []interface{}{
				true,
				123,
				3.14,
				complex(1, 2),
				"foobar",
				map[string]interface{}{"key1": 1, "key2": 2, "key3": nil},
				[]bool{true, false},
				[]int{1, 2, 3},
				[]float32{1.1, 2.2, 3.3},
				[]string{"foo", "bar"},
			},
		},
	},
}

func TestCopyableMap(t *testing.T) {
	for _, tc := range testCasesForMap {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			copyMap := Map(tc.original).DeepCopy().(map[string]interface{})
			assert.Equal(t, tc.expectedCopy, tc.transformer(copyMap), "DeepCopy was not mutated.")
			assert.Equal(t, tc.expectedOriginal, tc.original, "Original was mutated.")
		})
	}
}

func TestDeepCopy_Map(t *testing.T) {
	for _, tc := range testCasesForMap {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			copyMap := DeepCopy(tc.original).(map[string]interface{})
			assert.Equal(t, tc.expectedCopy, tc.transformer(copyMap), "DeepCopy was not mutated.")
			assert.Equal(t, tc.expectedOriginal, tc.original, "Original was mutated.")
		})
	}
}

func TestDeepCopy_Pointer(t *testing.T) {
	ptr := []string{"foo", "bar"}

	originalMap := map[string]interface{}{"ptr": &ptr}
	expectedOriginal := map[string]interface{}{"ptr": &ptr}
	expectedCopy := map[string]interface{}{"ptr": []string{"foo", "baz"}}
	transformer := func(m map[string]interface{}) map[string]interface{} {
		m["ptr"].([]string)[1] = "baz"

		return m
	}

	clone := Map(originalMap).Clone()

	assert.Equal(t, expectedCopy, transformer(clone), "DeepCopy was not mutated.")
	assert.Equal(t, expectedOriginal, originalMap, "Original was mutated.")
}

type NestedMap map[string]interface{}

func (n NestedMap) DeepCopy() interface{} {
	clone := make(NestedMap, len(n))

	for k, v := range n {
		clone[k] = DeepCopy(v)
	}

	return clone
}

func TestDeepCopy_Copyable_NestedMap(t *testing.T) {
	originalMap := NestedMap{"code": "FR", "country": NestedMap{"city": "Paris"}}
	expectedOriginal := NestedMap{"code": "FR", "country": NestedMap{"city": "Paris"}}
	expectedCopy := NestedMap{"code": "IT", "country": NestedMap{"city": "Roma"}}
	transformer := func(m map[string]interface{}) NestedMap {
		m["code"] = "IT"
		m["country"].(NestedMap)["city"] = "Roma"

		return m
	}

	clone := Map(originalMap).Clone()

	assert.Equal(t, expectedCopy, transformer(clone), "DeepCopy was not mutated.")
	assert.Equal(t, expectedOriginal, originalMap, "Original was mutated.")
}

func BenchmarkCopyableMap(b *testing.B) {
	m := map[string]interface{}{
		"foo": []interface{}{
			false,
			456,
			6.28,
			complex(3, 4),
			"bar",
			map[string]interface{}{"key1": 3, "key2": 2},
			map[string]bool{"key1": true, "key2": false},
			[]bool{false, false},
			[]int{4, 2, 3},
			[]float32{4.4, 2.2, 3.3},
			[]string{"baz", "bar"},
		},
	}

	for n := 0; n < b.N; n++ {
		Map(m).DeepCopy()
	}
}

func BenchmarkDeepCopy_Map(b *testing.B) {
	m := map[string]interface{}{
		"foo": []interface{}{
			false,
			456,
			6.28,
			complex(3, 4),
			"bar",
			map[string]interface{}{"key1": 3, "key2": 2},
			map[string]bool{"key1": true, "key2": false},
			[]bool{false, false},
			[]int{4, 2, 3},
			[]float32{4.4, 2.2, 3.3},
			[]string{"baz", "bar"},
		},
	}

	for n := 0; n < b.N; n++ {
		Map(m).DeepCopy()
	}
}
