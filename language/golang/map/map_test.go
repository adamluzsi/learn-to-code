package main

import (
	"testing"

	"go.llib.dev/testcase/assert"
	"go.llib.dev/testcase/random"
)

var rnd = random.New(random.CryptoSeed{})

func TestExerciseMakeEmptyMap(t *testing.T) {
	m := ExerciseMakeEmptyMap()

	assert.NotNil(t, m)
	assert.Empty(t, m)
}

func TestExerciseMakeNonEmptyMap(t *testing.T) {
	m := ExerciseMakeNonEmptyMap()
	assert.NotNil(t, m)
	assert.NotEmpty(t, m)
	assert.Equal(t, m, map[int]string{42: "The Answer"})
}

func TestExerciseAddKey(t *testing.T) {
	m := make(map[int]string)
	key := rnd.Int()
	value := rnd.String()
	ExerciseAddKey(m, key, value)
	assert.NotNil(t, m)
	assert.NotEmpty(t, m)
	assert.Equal(t, m[key], value)
}

func TestExerciseDeleteKey(t *testing.T) {
	key := rnd.Int()
	m := map[int]string{key: rnd.String()}
	ExerciseDeleteKey(m, key)
	assert.NotNil(t, m)
	assert.Empty(t, m)
}

func TestExerciseUpdateValue(t *testing.T) {
	key := rnd.Int()
	newValue := rnd.String()
	m := map[int]string{key: rnd.String()}
	ExerciseUpdateValue(m, key, newValue)
	assert.NotNil(t, m)
	assert.NotEmpty(t, m)
	assert.Equal(t, m[key], newValue)
}

func TestExerciseCheckKeyPresent(t *testing.T) {
	key := rnd.Int()
	m := map[int]string{key: rnd.String()}
	present := ExerciseCheckKeyPresent(m, key)
	assert.NotNil(t, present)
	assert.True(t, present)

	absentKey := rnd.Int()
	if absentKey == key {
		absentKey = rnd.Int()
	}
	present = ExerciseCheckKeyPresent(m, absentKey)
	assert.NotNil(t, present)
	assert.False(t, present)
}

func TestExerciseGetLength(t *testing.T) {
	m := map[int]struct{}{}

	var i int
	expLen := rnd.Repeat(3, 12, func() {
		i++
		m[i] = struct{}{}
	})

	assert.Equal(t, expLen, ExerciseGetLength(m))
}
