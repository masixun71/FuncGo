package gotest

import (
	"testing"
	"FuncGo/Func"
)

func Test_max_Array(t *testing.T) {

	a := [5]int{1,2,3,4,5}
	b := [6]int{1,2,3,4,5,6}
	_, err := Func.Max(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.Max(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 5 {
		t.Error("result is error", val)
	}

}


func Test_max_Slice(t *testing.T) {

	a := []int{1,2,3,4,5}
	b := []int{1,2,3,4,5,6}
	_, err := Func.Max(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.Max(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 5 {
		t.Error("result is error", val)
	}

}

func Test_max_Map(t *testing.T)  {
	a := make(map[int]int, 0)
	a[0] = 1
	a[1] = 2
	b := make(map[int]int, 0)
	b[0] = 1
	b[1] = 2

	_, err := Func.Max(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.Max(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 2 {
		t.Error("result is error", val)
	}

}

func Test_max_SliceStruct(t *testing.T) {

	a := foo{a: 1}
	b := foo{a: 3}
	c := foo{a: 2}


	ff := []foo{a,b,c}

	val, err := Func.Max(ff)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != b {
		t.Error("result is error", val)
	}

}

func Test_max_ArrayStruct(t *testing.T) {

	a := foo{a: 1}
	b := foo{a: 3}
	c := foo{a: 2}


	ff := [3]foo{a,b,c}

	val, err := Func.Max(ff)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != b {
		t.Error("result is error", val)
	}

}

func Test_max_MapStruct(t *testing.T) {

	a := foo{a: 1}
	b := foo{a: 3}
	c := foo{a: 2}

	ff := make(map[int]foo, 0)
	ff[0] = a
	ff[1] = b
	ff[2] = c

	val, err := Func.Max(ff)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != b {
		t.Error("result is error", val)
	}

}
