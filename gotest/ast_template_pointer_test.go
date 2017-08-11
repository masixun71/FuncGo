package gotest

import (
	"testing"
	"FuncGo/Func"
	"os"
	"FuncGo/lib"
)



func Test_simple_pointer_string(t *testing.T) {

	makeFiler,_ := Func.NewMakeFilerByBasicType(Func.TypeT, "/code", "**")

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




func Test_simple_pointer_file(t *testing.T) {

	makeFiler,_ := Func.NewMakeFilerByBasicType(Func.TypeT, "/code", "**")

	file, err := os.OpenFile(lib.GetRoot()+"/Func/ast.go", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}
	_, err = makeFiler.MakeFuncSourceWithFile(file)

	if  err != nil {
		t.Error("test fail", err)
	}
}




func Test_simple_pointer_func(t *testing.T) {
	makeFiler,err := Func.NewMakeFilerByBasicType(Func.TypeT, "/code", "***")

	if err != nil {
		t.Error(err)
	}
	path := lib.NewPath("/Func/ast.go")

	_, err = makeFiler.MakeFuncSourceWithFunc(path, "MaxTF")
	if  err != nil {
		t.Error("test fail", err)
	}
}




func Test_pointer_string(t *testing.T) {

	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", false, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}

	_,err = makeFiler.MakeFuncSourceWithString(`
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



func Test_pointer_file(t *testing.T) {

	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", false, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}

	file, err := os.OpenFile(lib.GetRoot()+"/Func/ast.go", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}
	_, err = makeFiler.MakeFuncSourceWithFile(file)

	if  err != nil {
		t.Error("test fail", err)
	}
}



func Test_pointer_func(t *testing.T) {
	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", false, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}

	path := lib.NewPath("/Func/ast.go")

	_, err = makeFiler.MakeFuncSourceWithFunc(path, "MaxTF")
	if  err != nil {
		t.Error("test fail", err)
	}
}



func Test_method_pointer_func(t *testing.T) {
	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", true, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}


	path := lib.NewPath("/Func/ast.go")

	_, err = makeFiler.MakeMethodSourceWithFunc(new(lib.Pather), path, "MaxTF")
	if err != nil {
		t.Error(err)
	}
}




func Test_method_pointer_file(t *testing.T) {
	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", true, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}

	file, err := os.OpenFile(lib.GetRoot()+"/Func/ast.go", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}

	_, err = makeFiler.MakeMethodSourceWithFile(new(lib.Pather), file)
	if err != nil {
		t.Error(err)
	}
}




func Test_method_pointer_string(t *testing.T) {
	makeFiler,err := Func.NewMakeFiler(&Func.BuildType{FuncString:"TF", TypeString:"T"}, "/code", true, "**", Func.TmapByBasicType(), 0)

	if err != nil {
		t.Error(err)
	}

	_, err = makeFiler.MakeMethodSourceWithString(new(lib.Pather), `
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
	if err != nil {
		t.Error(err)
	}
}