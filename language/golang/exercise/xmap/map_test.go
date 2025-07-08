package xmap_test

import (
	"testing"

	"golang-exercise/xmap"

	"go.llib.dev/testcase/assert"
)

func TestCreateMap(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, xmap.CreateMap(), "the function should return the map with given key-value pairs")
}

func TestAccessValueByKey(t *testing.T) {
	assert.Equal(t, 1, xmap.AccessValueByKey(map[string]int{"a": 1, "b": 2}, "a"),
		"the function should return the value for key 'a'")
}

func TestModifyValueByKey(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2}
	modified := xmap.ModifyValueByKey(original, "a", 5)
	assert.Equal(t, map[string]int{"a": 5, "b": 2}, modified,
		"the function should modify the value for key 'a' to 5")
}

func TestIterateOverMap(t *testing.T) {
	assert.Equal(t, []int{1, 2}, xmap.IterateOverMap(map[string]int{"a": 1, "b": 2}),
		"the function should return the slice of values from the map")
}

func TestCopyMapIntoNewMap(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2}
	copied := xmap.CopyMapIntoNewMap(original)
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, copied,
		"the function should copy the map into a new map with its own backing")
}
