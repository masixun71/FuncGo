package main

import (
	"FuncGo/Func"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFiler("/Func/ast.go", "MaxT", new(Func.T))
	//makeFiler.MakeMethod(new(Func.ValueS))
	makeFiler.MakeFunc()
}
