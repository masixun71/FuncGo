package main

import (
	"FuncGo/Func"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFilerSimple("/Func/ast.go", "MaxTF", Func.TypeT)
	//makeFiler.MakeMethod(new(Func.ValueS))
	makeFiler.MakeFunc()
}
