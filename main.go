package main

import (
	"FuncGo/Func"
	"FuncGo/code"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFiler("/Func/ast.go", "MaxT", new(Func.T))
	makeFiler.MakeMethod(new(code.ValueS))
}
