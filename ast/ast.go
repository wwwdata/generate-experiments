package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	file, err := parser.ParseFile(token.NewFileSet(), "ast.go", nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	for _, i := range file.Imports {
		fmt.Println(i.Path.Value)
	}

	// parsing comments
	file, err = parser.ParseFile(token.NewFileSet(), "ast.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// such comment wow
	for _, i := range file.Comments {
		fmt.Println(i.Text())
	}

	// Parse very small example go program and inspect all the nodes
	fset := token.NewFileSet()
	src := `
	package foo
	const c = 1.0
	`
	f, err := parser.ParseFile(fset, "something.go", src, 0)
	ast.Inspect(f, func(n ast.Node) bool {
		fmt.Printf("%#v\n", n)
		return true
	})
}
