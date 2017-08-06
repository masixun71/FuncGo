package main

import (
	"FuncGo/lib"
	"go/token"
	"go/parser"
	"go/ast"
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {

	//makeFiler, err := Func.NewMakeFilerByBasicType(Func.TypeT, "/code")

	path := lib.NewPath("/Func/ast1.go")

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, path.GetPathByRoot(), nil, parser.ParseComments)

	b, _ := ioutil.ReadFile(path.GetPathByRoot())

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == "CompareTF" {

			str := string(b[fn.Type.Func-1: fn.Body.Rbrace])
			//var buffer bytes.Buffer
			for _, field := range fn.Type.Params.List {
				for _, name := range field.Names {
					//varLen := len(name.Name)
					index := strings.Index(str, name.Name)
					str = str[0:index] + "p" + str[index:len(str)]

				}
			}

			fmt.Println(str)

		}
	}

}
