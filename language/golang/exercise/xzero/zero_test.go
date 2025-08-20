package xzero_test

import (
	_ "embed"
	"go/ast"
	"go/token"
	exercise "golang-exercise"
	"testing"

	"golang-exercise/xzero"

	"go.llib.dev/testcase/assert"
)

const srcfile = "xzero/zero.go"

func TestExercise1(t *testing.T) {
	ast.Inspect(exercise.GetFunc(t, srcfile, "ZeroInt"), func(n ast.Node) bool {
		if bl, ok := n.(*ast.BasicLit); ok {
			assert.False(t, bl.Kind == token.INT && bl.Value == `0`,
				"You need to use the default zero int value of a type by creating a string variable, without any value assigment")
		}
		return true
	})

	assert.Equal(t, 0, xzero.ZeroInt(),
		"the function should return the zero value of int via var declaration")
}

func TestExercise2(t *testing.T) {
	ast.Inspect(exercise.GetFunc(t, srcfile, "ZeroString"), func(n ast.Node) bool {
		if bl, ok := n.(*ast.BasicLit); ok {
			assert.False(t, bl.Kind == token.STRING && bl.Value == `""`,
				"You need to use the default zero string value of a type by creating a string variable, without any value assigment")
		}
		return true
	})

	assert.Equal(t, "", xzero.ZeroString(),
		"the function should return the zero value of string via var declaration")
}

func TestExercise3(t *testing.T) {
	ast.Inspect(exercise.GetFunc(t, srcfile, "ZeroBool"), func(n ast.Node) bool {
		indent, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		assert.False(t, indent.Name == "false",
			"You need to use the default zero bool value of a type by creating a string variable, without any value assigment")

		return true
	})

	assert.Equal(t, false, xzero.ZeroBool(),
		"the function should return the zero value of bool via var declaration")
}
