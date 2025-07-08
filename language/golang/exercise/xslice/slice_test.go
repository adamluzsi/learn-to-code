package xslice_test

import (
	"testing"

	"golang-exercise/xslice"

	"go.llib.dev/testcase/assert"
)

func TestCreateSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, xslice.CreateSlice(), "the function should return the Slice with given elements")
}

func TestAccessElementAtZero(t *testing.T) {
	assert.Equal(t, 1, xslice.AccessElementAtZero([]int{1, 2, 3}),
		"the function should return the element at index 0")
}

func TestModifyElementAtIndexOne(t *testing.T) {
	original := []int{1, 2, 3}
	modified := xslice.ModifyElementAtIndexOne(original)
	assert.Equal(t, []int{1, 5, 3}, modified,
		"the function should modify the element at index 1 to 5")
}

func TestAppendToSlice(t *testing.T) {
	original := []int{1, 2}
	n := xslice.AppendToSlice(original, 3)
	assert.Equal(t, []int{1, 2, 3}, n,
		"the function should append the value and return the new Slice")
}

func TestCopySliceIntoNewSlice(t *testing.T) {
	original := []int{1, 2, 3}
	copied := xslice.CopySliceIntoNewSlice(original)
	assert.Equal(t, []int{1, 2, 3}, copied,
		"the function should copy the slice into a new Slice with its own backing")
}
