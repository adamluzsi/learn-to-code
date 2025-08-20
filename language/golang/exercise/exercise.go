package exercise

import (
	"embed"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"strings"
	"testing"

	"go.llib.dev/testcase/assert"
	"go.llib.dev/testcase/random"
)

var rnd = random.New(random.CryptoSeed{})

func Random() *random.Random { return rnd }

//go:embed all:*
var files embed.FS

func findFile(tb testing.TB, filename string) fs.File {
	tb.Helper()

	var paths []string
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if d != nil && d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, filename) {
			paths = append(paths, path)
		}
		return nil
	})

	if len(paths) == 0 {
		tb.Fatalf("%s filename is not found", filename)
		return nil
	}

	if len(paths) != 1 {
		tb.Log("Too many matching filename is found, could you improve your filename identifier?")
		for _, p := range paths {
			tb.Log(p)
		}
		tb.FailNow()
		return nil
	}

	file, err := files.Open(paths[0])
	assert.NoError(tb, err)
	return file

}

func GetFunc(tb testing.TB, filename, funcName string) *ast.FuncDecl {
	tb.Helper()

	file := findFile(tb, filename)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", file, parser.ParseComments)
	if err != nil {
		assert.NoError(tb, err)
		log.Fatal(err)
	}
	var fn *ast.FuncDecl
	ast.Inspect(f, func(node ast.Node) bool {
		if fn != nil {
			return false
		}
		if decl, ok := node.(*ast.FuncDecl); ok && decl.Name.Name == funcName {
			fn = decl
		}
		return true
	})
	assert.NotNil(tb, fn,
		assert.MessageF("it is unexpected that func %s is missing, did you renamed it or remove id by any chance?", funcName))
	return fn
}

func Description(tb testing.TB, logs ...string) {
	tb.Helper()
	var msg string
	msg = strings.Join(logs, "\n")
	msg = strings.TrimSpace(msg)
	tb.Logf("\n\n%s\n\n", msg)
}
