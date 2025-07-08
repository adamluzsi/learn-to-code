package xcontrol_test

import (
	"testing"

	"golang-exercise/xcontrol"

	"go.llib.dev/testcase/assert"
)

func TestIfEvenOdd(t *testing.T) {
	assert.Equal(t, "even", xcontrol.IfEvenOdd(2),
		"the function should return 'even' for even numbers")
	assert.Equal(t, "odd", xcontrol.IfEvenOdd(3),
		"the function should return 'odd' for odd numbers")
}

func TestElseIfCheck(t *testing.T) {
	assert.Equal(t, "divisible by 3", xcontrol.ElseIfCheck(6),
		"the function should return 'divisible by 3' for numbers divisible by 3")
	assert.Equal(t, "even", xcontrol.ElseIfCheck(4),
		"the function should return 'even' for even numbers")
	assert.Equal(t, "odd", xcontrol.ElseIfCheck(5),
		"the function should return 'odd' for odd numbers")
}

func TestSwitchWithExpression(t *testing.T) {
	assert.Equal(t, "Weekend", xcontrol.SwitchWithExpression(6),
		"the function should return 'Weekend' for Saturday/Sunday")
	assert.Equal(t, "Start of week", xcontrol.SwitchWithExpression(0),
		"the function should return 'Start of week' for Monday")
	assert.Equal(t, "Midweek", xcontrol.SwitchWithExpression(2),
		"the function should return 'Midweek' for other days")
}

func TestSwitchWithoutExpression(t *testing.T) {
	assert.Equal(t, "negative", xcontrol.SwitchWithoutExpression(-1),
		"the function should return 'negative' for negative numbers")
	assert.Equal(t, "zero", xcontrol.SwitchWithoutExpression(0),
		"the function should return 'zero' for zero")
	assert.Equal(t, "positive", xcontrol.SwitchWithoutExpression(1),
		"the function should return 'positive' for positive numbers")
}
