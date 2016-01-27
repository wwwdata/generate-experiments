package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// how to parse consts of a type?
func main() {
	// example go file that we want to parse all the Unicorn constants out of
	src := `
		package foo

		type Unicorn int

		const (
			Pink Unicorn = iota
			Fluffy, Puffy
			Rainbow
			MagicNumber int = 42
		)
	`

	file, err := parser.ParseFile(token.NewFileSet(), "unicorns.go", src, 0)
	if err != nil {
		panic(err)
	}

	var values []string
	// we now hard code the type for the sake of simplicity. The go:generator tool would read the type via a
	// command line parameter like the Stringer oder Jsonenum tool
	typeName := "Unicorn"

	ast.Inspect(file, func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			// we need consts
			return true
		}

		// current type of the const we are iterating over
		typ := ""

		for _, spec := range decl.Specs {
			vspec := spec.(*ast.ValueSpec) // must work, because we are in a const

			// we have seen that the first const had the type and the others are untyped, but we will assume that they
			// all have the same type, like in the Stringer example. So we will need to remember it and check if a new
			// type occurs.
			if vspec.Type != nil {
				ident, ok := vspec.Type.(*ast.Ident)
				if !ok {
					continue
				}
				typ = ident.Name
			}

			if typ != typeName {
				continue
			}

			// outputs one line of code
			fmt.Println(vspec.Names)
		}

		return false
	})

	fmt.Println("just to compile it")
	fmt.Println(values, typeName)
}
