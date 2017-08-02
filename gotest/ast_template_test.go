package gotest

import (
	"testing"
	"FuncGo/Func"
	"os"
	"FuncGo/lib"
)

func Test_simple_string(t *testing.T) {

	makeFiler,_ := Func.NewMakeFilerSimple(Func.TypeT)

	_,err := makeFiler.MakeFuncSourceWithString(`
		package Func
		func MaxTF(values ...T) (T, error) {

		len := len(values)

		max := values[0]
		for i := 1; i < len; i++ {
			tmpvalue := values[i]

			if max < tmpvalue {
				max = tmpvalue
			}
		}

		return max, nil
	}`)

	t.Log(makeFiler)
	if  err != nil {
		t.Error("test fail", err)
	}
}


func Test_simple_file(t *testing.T) {

	makeFiler,_ := Func.NewMakeFilerSimple(Func.TypeT)

	file, err := os.OpenFile(lib.GetRoot()+"/Func/ast.go", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}
	_, err = makeFiler.MakeFuncSourceWithFile(file)

	if  err != nil {
		t.Error("test fail", err)
	}
}

func Test_simple_func(t *testing.T) {
	makeFiler,_ := Func.NewMakeFilerSimple(Func.TypeT)

	path := lib.NewPath("/Func/ast.go")

	_, err := makeFiler.MakeFuncSourceWithFunc(path, "MaxTF")
	if  err != nil {
		t.Error("test fail", err)
	}
}

