package main

import (
	"FuncGo/Func"
)

func main() {

	makeFiler := Func.NewMakeFiler("/Func/AST.go", "CompareT")
	makeFiler.MakeCode()

}
