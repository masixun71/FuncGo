package main

import (
	"reflect"
	"fmt"
	"FuncGo/Func"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	makeFiler := Func.NewMakeFiler("/Func/AST.go", "CompareT")
	makeFiler.MakeCode()

}

type YourT2 struct {
}

func (y *YourT2) MethodFoo(i int, oo string) {
	fmt.Println("MethodFoo called", i, oo)
}

func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}
