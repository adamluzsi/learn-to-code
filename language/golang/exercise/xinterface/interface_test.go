package xinterface_test

import (
	"testing"

	"golang-exercise/xinterface"

	"go.llib.dev/testcase/assert"
)

func TestDefineInterface(t *testing.T) {
	// Interface is defined with method Describe() string, no test needed here.
}

func TestImplementInterface(t *testing.T) {
	v := xinterface.T{Desc: "foo bar"}
	d, ok := any(v).(xinterface.Describer)

	assert.True(t, ok,
		"You need to implement the `Describer` interface on `T`.",
		"You can do that by creating all the methods listed in `Describer` on T.",
		"For example, if Describer would have a `Foo()` function, then you would need `func (v T) Foo() {}` to implement it.",
		"In Go, interfaces are implemented implicitly, so you don't need to have `implements` or similar keywords to express that a type implements it.",
		"It is more akin to how duck typing works in other high-level languages.",
	)

	assert.Equal(t, v.Desc, d.Describe(), "Your `Describe() string` function should return the T.Desc field.")
}

type D struct{ d string }

func (d D) Describe() string { return d.d }

func TestUseInterfaceMethod(t *testing.T) {
	var (
		exp    = "bar baz"
		v      = D{d: exp}
		result = xinterface.UseInterfaceMethod(v)
	)
	assert.Equal(t, exp, result,
		"the function should call the Describe method on the interface and return its result")
}

func TestInterfaceImplementation(t *testing.T) {
	d := xinterface.TestInterfaceImplementation()
	assert.NotNil(t, d, "expected that you return a non-nil Describer value (T{})")
	_, ok := d.(xinterface.T)

	assert.True(t, ok,
		"The goal of this exercise to see that you can return a concrete type value for an expected interface return result type.",
		"Try returning a T value",
	)

	assert.Equal(t, "Person: Alice, Age 30", d.Describe(),
		"the function should return a Person struct that implements the Describer interface")
}
