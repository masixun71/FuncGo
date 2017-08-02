package main

import (
	"FuncGo/Func"
	"FuncGo/lib"
)

func main() {

	//InvokeObjectMethod(new(YourT2), "MethodFoo", 10, "abc")

	//var file *os.File
	//
	//var ss interface{}
	//ss = file
	//
	//
	//switch s := ss.(type) {
	//case *os.File:
	//	fmt.Println(2,s)
	//case io.Reader:
	//	fmt.Println(1,s)
	//}

	makeFiler,_ := Func.NewMakeFilerSimple(Func.TypeT)

	//file, err := os.OpenFile(lib.GetRoot()+"/Func/ast.go", os.O_RDWR|os.O_CREATE, 0777)
	//
	//if err != nil {
	//	panic(err)
	//}

	//makeFiler.MakeFuncSourceWithFile(file)

	path := lib.NewPath("/Func/ast.go")

	makeFiler.MakeFuncSourceWithFunc(path, "MaxTF")

}
