package xcore_test

import (
	"testing"

	"golang-exercise/xcore"

	"go.llib.dev/testcase/assert"
)

func TestVariableUsage(t *testing.T) {
	assert.Equal(t, "foo", xcore.VariableUsage("foo"),
		"in order to achieve this test, you need to use the VariableUsage function's argument as a return value")
}

func TestSumTwoIntegers(t *testing.T) {
	assert.Equal(t, 5, xcore.SumTwoIntegers(2, 3),
		"the function should return the sum of a and b")
}

func TestMultiplyByTwoWithShortDeclaration(t *testing.T) {
	assert.Equal(t, 6, xcore.MultiplyByTwoWithShortDeclaration(3),
		"the function should return input multiplied by two")
}

func TestConcatStrings(t *testing.T) {
	assert.Equal(t, "HelloGo", xcore.ConcatStrings("Hello", "Go"),
		"the function should concatenate the two strings")
}

func TestSumThreeIntegers(t *testing.T) {
	assert.Equal(t, 6, xcore.SumThreeIntegers(1, 2, 3),
		"the function should return the sum of a, b, and c")
}

func TestMultiplyByPi(t *testing.T) {
	assert.Equal(t, 2*3, xcore.MultiplyBy3(2),
		"the function should return input multiplied by Pi")
}
