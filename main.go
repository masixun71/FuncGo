package main

import (
	"FuncGo/lib"
	"go/token"
	"go/parser"
	"go/ast"
)

func main() {

	fset := token.NewFileSet()

	path := lib.NewPath("/code/MaxTF.go")

	f, err := parser.ParseFile(fset, path.GetPathByRoot(), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)

}
