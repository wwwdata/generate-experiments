package main

import (
	"fmt"
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
}
