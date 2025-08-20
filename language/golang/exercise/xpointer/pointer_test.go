package xpointer_test

import (
	"testing"

	"golang-exercise/xpointer"

	"go.llib.dev/testcase/assert"
)

func TestCreatePointer(t *testing.T) {
	got := xpointer.CreatePointer()
	assert.NotNil(t, got, "the function was suppose to return a non nil pointer value")
	assert.Equal(t, *got, 42, "the function should return a pointer to the int variable with value 42")
}

func TestAccessValueThroughPointer(t *testing.T) {
	var x int = 42
	p := &x
	assert.Equal(t, 42, xpointer.AccessValueThroughPointer(p),
		"the function should return the value stored at the address pointed by p")
}

func TestModifyVariableViaPointer(t *testing.T) {
	var x int = 42
	p := &x
	modified := xpointer.ModifyVariableViaPointer(p)
	assert.NotNil(t, modified,
		"don't forget to return a non-nil pointer value")
	assert.Equal(t, 5, *modified,
		"the function should modify the variable's value to 5")
}

func TestCheckForNilPointer(t *testing.T) {
	assert.Equal(t, true, xpointer.CheckForNilPointer(nil),
		"the function should return true if p is nil")
}

func TestCopyPointerIntoNewPointer(t *testing.T) {
	var x int = 42
	p := &x
	copied := xpointer.CopyPointerIntoNewPointer(p)
	assert.NotNil(t, copied)
	assert.Equal(t, 42, *copied,
		"the function should copy the pointer into a new pointer with its own value")
	t.Log("given the original pointer's value is being modified")
	*p = 24
	assert.Equal(t, 42, *copied,
		"the function was expected to return a pointer which is an independent pointer that points to a copy of a value")
}
