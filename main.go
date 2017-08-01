package main

import (
	"FuncGo/Func"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFiler("/Func/AST.go", "StringTInterface", new(Func.TInterface))
	makeFiler.MakeCode()
}
