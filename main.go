package main

import (
	"FuncGo/lib"
	"go/token"
	"go/parser"
	"go/ast"
)

func main() {

	//makeFiler, err := Func.NewMakeFilerByBasicType(Func.TypeT, "/code")

	path := lib.NewPath("/Func/ast1.go")

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, path.GetPathByRoot(), nil, parser.ParseComments)


	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == "CompareTF" {
			for _, field := range fn.Type.Params.List {
				for _, name := range field.Names {
					name.Obj.Name = "aa"
				}
			}
		}
	}

	ast.Print(fset, f)
}
