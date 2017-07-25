package gotest

import (
	"testing"
	"go/ast"
	"go/token"
	"strconv"
	"go/parser"
)

func Benchmark_ast_done(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", `
	package main
func main(){
	// comments
	x:=1
	println(x+1)
}`, parser.ParseComments)

		if err != nil {
			panic(err)
		}

		var vars = map[string]int{}

		for _, decl := range f.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if ok && fn.Name.Name == "main" {
				for _, stmt := range fn.Body.List {
					exprstmt, ok := stmt.(*ast.ExprStmt)
					if ok {
						call, ok := exprstmt.X.(*ast.CallExpr)
						if ok {
							ident, ok := call.Fun.(*ast.Ident)
							if ok && ident.Name == "println" {
								expr, ok := call.Args[0].(*ast.BinaryExpr)
								if ok && expr.Op == token.ADD {
									x := expr.X.(*ast.Ident)
									y := expr.Y.(*ast.BasicLit)
									yn, _ := strconv.ParseInt(y.Value, 10, 64)
									arg := vars[x.Name] + int(yn)
									println(arg)
								}
							}
						}
					}

					ass, ok := stmt.(*ast.AssignStmt)
					if ok {
						x, ok := ass.Lhs[0].(*ast.Ident)
						if ok {
							y, ok := ass.Rhs[0].(*ast.BasicLit)
							if ok {
								yn, err := strconv.ParseInt(y.Value, 10, 64)
								if err != nil {
									panic(err)
								}
								vars[x.Name] = int(yn)
							}
						}
					}
				}
			}

		}
	}
}

func Benchmark_normal_done(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 1
		println(x + 1)
	}
}
