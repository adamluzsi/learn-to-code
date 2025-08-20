package xstruct_test

import (
	"reflect"
	"testing"

	"golang-exercise/xstruct"

	"go.llib.dev/testcase/assert"
)

func TestExercise1(t *testing.T) {
	p := xstruct.CreateStruct()
	assert.Equal(t, xstruct.Person{"Alice", 30}, p,
		"the function should return a Person struct with name 'Alice' and age 30")
}

func TestExercise2(t *testing.T) {
	p := xstruct.NewStruct()
	assert.NotNil(t, p)
	assert.Equal(t, xstruct.Person{"Zoe", 24}, *p,
		"the function should return a *Person with name 'Zoe' and age 24")
}

func TestExercise3(t *testing.T) {
	T := reflect.TypeOf(xstruct.Person{})
	sf, ok := T.FieldByName("MyField")

	assert.True(t, ok,
		`Add a new field called "MyField" to the Person struct.`,
		"The type is up to you.")

	t.Logf("hmmm, %s type? Good choice!", sf.Type.String())
}

func TestExercise4(t *testing.T) {
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

func TestExercise5(t *testing.T) {
	v := xstruct.Person{"Alice", 30}

	type WA interface {
		UpdateAge(age int)
	}

	wa, ok := any(&v).(WA)
	assert.True(t, ok)
	wa.UpdateAge(24)

	assert.Equal(t, v, xstruct.Person{Name: "Alice", Age: 24}, "expected our Person is updated.",
		"If this is not the case, then the method is using a value receiver type instead of a pointer receiver,",
		"and the receiver argument is a Pass By Value Copy on each method call.")
}

func TestExercise6(t *testing.T) {
	v := xstruct.ShowYourOwnType()
	assert.NotNil(t, v, "Create your very own struct using the `type` and `struct` keywords.")

	rv := reflect.ValueOf(v)
	assert.Equal(t, rv.Type().Kind(), reflect.Struct, "The type you returned in ShowYourOwnType is not a struct.")

	personType := reflect.TypeOf((*xstruct.Person)(nil)).Elem()
	assert.NotEqual(t, personType, rv.Type(),
		"No cheating, don't return back the Person type, but make your own struct!")

	assert.Equal(t, rv.Type().PkgPath(), personType.PkgPath(),
		"Please create your struct type within the xstruct package.")
}
