package main

import (
	"FuncGo/Func"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFiler("/Func/ast.go", "MaxT", Func.TypeT)
	//makeFiler.MakeMethod(new(Func.ValueS))
	makeFiler.MakeFunc()
}
