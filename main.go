package main

import (
	"go/token"
	"go/parser"
	"go/ast"
	"fmt"
	"strconv"
)

func main() {

	values := []int{1, 2, 3}

	vars := make(map[string]interface{}, 0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./Func/AST.go", nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == "astMax" {

			assignStmt, ok := fn.Body.List[0].(*ast.AssignStmt)
			if ok {
				x, ok := assignStmt.Lhs[0].(*ast.Ident)
				if ok {
					y, ok := assignStmt.Rhs[0].(*ast.CallExpr)
					if ok {
						ident, ok := y.Fun.(*ast.Ident)
						if ok && ident.Name == "len" {
							vars[x.Name] = len(values)
						}
					}

				}

			}

			assignStmt, ok = fn.Body.List[1].(*ast.AssignStmt)
			if ok {
				x, ok := assignStmt.Lhs[0].(*ast.Ident)
				if ok {
					y, ok := assignStmt.Rhs[0].(*ast.IndexExpr)
					if ok {
						index, ok := y.Index.(*ast.BasicLit)
						if ok {
							if index.Kind == token.INT {
								yindex, _ := strconv.Atoi(index.Value)
								vars[x.Name] = values[yindex]
							}
						}
					}
				}

			}

			forStmt, ok := fn.Body.List[2].(*ast.ForStmt)
			if ok {
				init, ok := forStmt.Init.(*ast.AssignStmt)
				if ok {
					x, ok := init.Lhs[0].(*ast.Ident)
					if ok {
						initValue, ok := init.Rhs[0].(*ast.BasicLit)
						if ok {
							if initValue.Kind == token.INT {
								yindex, _ := strconv.Atoi(initValue.Value)
								vars[x.Name] = yindex
							}
						}
					}
				}

				cond, ok := forStmt.Cond.(*ast.BinaryExpr)
				if ok {
					x, ok := cond.X.(*ast.Ident)
					if ok {
						_, ok := vars[x.Name]
						if !ok {
							panic("变量不存在" + x.Name)
						}
					}

					if cond.Op != token.LSS {
						panic("for 循环变量有误")
					}

				}

				post, ok := forStmt.Post.(*ast.IncDecStmt)
				if ok {
					x, ok := cond.X.(*ast.Ident)
					if ok {
						if x.Name != "i" {
							panic("for 循环变量有误")
						}
					}
					if post.Tok != token.INC {
						panic("for 循环变量有误")

					}
				}

			}

		}
	}

	fmt.Println(vars)

}
