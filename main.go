package main

import (
	"go/token"
	"go/parser"
	"go/ast"
	"fmt"
	"strconv"
	"FuncGo/lib"
)

func main() {

	values := []int{1, 2, 3}

	vars := make(map[string]interface{}, 0)

	intVars := make(map[string]int, 0)


	var cycleName string
	var lenName string
	var maxName string
	var tmpValueName string


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
							intVars[x.Name] = len(values)
							lenName = x.Name
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
								maxName = x.Name
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
								intVars[x.Name] = yindex
								cycleName = x.Name
							}
						}
					}
				}

				cond, ok := forStmt.Cond.(*ast.BinaryExpr)
				if ok {
					x, ok := cond.X.(*ast.Ident)
					if ok {
						_, ok := intVars[x.Name]
						if !ok {
							panic("变量不存在" + x.Name)
						}

						if cycleName != x.Name {
							panic("for 循环变量有误")
						}
					}

					if cond.Op != token.LSS {
						panic("for 循环变量有误")
					}

					y, ok := cond.Y.(*ast.Ident)
					if ok {
						if lenName != y.Name{
							panic("for 循环变量有误")
						}

					}


				}

				post, ok := forStmt.Post.(*ast.IncDecStmt)
				if ok {
					x, ok := cond.X.(*ast.Ident)
					if ok {
						if x.Name != cycleName {
							panic("for 循环变量有误")
						}
					}
					if post.Tok != token.INC {
						panic("for 循环变量有误")

					}
				}

				ass, ok := forStmt.Body.List[0].(*ast.AssignStmt)
				if ok {
					x, ok := ass.Lhs[0].(*ast.Ident)
					if ok {
						y, ok := ass.Rhs[0].(*ast.IndexExpr)
						if ok {
							_, ok := y.Index.(*ast.Ident)
							if ok {
								tmpValueName = x.Name
							}
						}
					}
				}

				ifStmt, ok := forStmt.Body.List[1].(*ast.IfStmt)
				if ok {
					cond, ok := ifStmt.Cond.(*ast.BinaryExpr)
					if ok {
						x, ok := cond.X.(*ast.Ident)
						if ok {
							if x.Name != maxName{
								panic("if 判断 条件有误")
							}
						}

						y, ok := cond.Y.(*ast.Ident)
						if ok {
							if y.Name != tmpValueName {
								panic("if 判断 条件有误")
							}
						}
					}

					ass, ok := ifStmt.Body.List[0].(*ast.AssignStmt)
					if ok {
						x, ok := ass.Lhs[0].(*ast.Ident)
						if ok {
							if x.Name != maxName {
								panic("if 内容有误")
							}
						}

						if ass.Tok != token.ASSIGN {
							panic("if 内容有误")
						}

						y, ok := ass.Rhs[0].(*ast.Ident)
						if ok {
							if y.Name != tmpValueName {
								panic("if 内容有误")
							}
						}
					}

				}

			}

		}
	}


	for cycle := intVars[cycleName]; cycle < intVars[lenName]; cycle++ {
		vars[tmpValueName] = values[cycle]

		if lib.CompareBetween(vars[maxName], vars[tmpValueName]) == lib.LT {
			vars[maxName] = vars[tmpValueName]
		}
	}

}
