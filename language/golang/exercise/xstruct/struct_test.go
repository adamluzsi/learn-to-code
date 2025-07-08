package xstruct_test

import (
	"testing"

	"golang-exercise/xstruct"

	"go.llib.dev/testcase/assert"
)

func TestCreateStruct(t *testing.T) {
	p := xstruct.CreateStruct()
	assert.Equal(t, xstruct.Person{"Alice", 30}, p,
		"the function should return a Person struct with name 'Alice' and age 30")
}

func TestNewStruct(t *testing.T) {
	p := xstruct.NewStruct()
	assert.NotNil(t, p)
	assert.Equal(t, xstruct.Person{"Zoe", 24}, *p,
		"the function should return a *Person with name 'Zoe' and age 24")
}

func TestModifyAgeViaValueReceiver(t *testing.T) {
	p := xstruct.Person{"Alice", 30}

	type WA interface {
		WithAge(age int) xstruct.Person
	}

	wa, ok := any(&p).(WA)
	assert.True(t, ok)
	got := wa.WithAge(42)

	assert.Equal(t, got, xstruct.Person{Name: "Alice", Age: 42}, "expected that we got back a new Person value with the desired age")
	assert.Equal(t, p, xstruct.Person{"Alice", 30},
		"The original value should have not changed, this can only happens if unintentionally a pointer receiver were used.",
		"You should only use a pointer receiver, if you want to mutate the original value of your method.",
	)
}

func TestModifyAgeViaPointerReceiver(t *testing.T) {
	p := xstruct.Person{"Alice", 30}

	type WA interface {
		UpdateAge(age int)
	}

	wa, ok := any(&p).(WA)
	assert.True(t, ok)
	wa.UpdateAge(24)

	assert.Equal(t, p, xstruct.Person{Name: "Alice", Age: 24}, "expected our Person is updated.",
		"If this is not the case, then the method is using a value receiver type instead of a pointer receiver,",
		"and the receiver argument is a Pass By Value Copy on each method call.")
}

func TestCopyStructIntoNewStruct(t *testing.T) {
	p := xstruct.Person{"Alice", 30}
	copied := xstruct.CopyStructIntoNewStruct(p)
	assert.Equal(t, xstruct.Person{"Alice", 30}, copied,
		"the function should return a copy of the struct with same fields")
}

func TestGetAgeViaPointerReceiver(t *testing.T) {
	p := &xstruct.Person{"Alice", 30}
	assert.Equal(t, 30, xstruct.GetAgeViaPointerReceiver(p),
		"the function should return the age through the pointer")
}
