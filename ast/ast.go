package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
)

// how to parse consts of a type?
func main() {
	// example go file that we want to parse all the Unicorn constants out of
	src := `
		package foo

		type Unicorn int

		const (
			Pink Unicorn = iota
			Fluffy
			Rainbow
			MagicNumber int = 42
		)
	`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "unicorns.go", src, 0)
	if err != nil {
		panic(err)
	}

	// type-check for the package
	defs := make(map[*ast.Ident]types.Object)
	config := types.Config{}
	info := &types.Info{Defs: defs}
	_, err = config.Check("unicorn.go", fset, []*ast.File{file}, info)
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

			for _, name := range vspec.Names {
				values = append(values, name.Name)
				// hm cool, but how do we get the actual integer number out of them?
				// now we have defs and can extract the actual values out of there
				obj, ok := defs[name]
				if !ok {
					fmt.Println("no value for constant ", name)
				}

				fmt.Printf("value %s %d\n", name, obj.(*types.Const).Val())
			}
		}

		return false
	})

	fmt.Println("all values", values, "type", typeName)
}
