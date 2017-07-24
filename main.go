package main

import (
	"unsafe"
	"fmt"
	"reflect"
	"FuncGo/Func"
)

type TestObj struct {
	field1 string
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func ff(values ...interface{}) {

	fmt.Println(reflect.TypeOf(values))

}

func Max(values ...interface{}) interface{} {

	fmt.Println(reflect.TypeOf(values))

	ff(values[0])

	return nil
}

var cache = map[*uintptr]map[string]uintptr{}

func getType(val interface{}) string {

	typ := reflect.TypeOf(val)
	switch typ.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int:

	}

	itab := *(**uintptr)(unsafe.Pointer(&val))

	m, ok := cache[itab]
	if !ok {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	return "1"
}

type foo struct {
	A int
	Func.Comparer
}
type foo2 struct {
	b string
}

func (f foo) Compare(value interface{}) bool {
	val, _ := value.(foo)

	if f.A >= val.A {
		return true
	} else {
		return false
	}

}



func main() {
	a := [5]int{1,2,3,4,5}
	//b := [6]int{1,2,3,4,5,6}
	//_, err := Func.Max(a, b)
	//if err != nil {
	//	fmt.Println("ok")
	//}

	val, err := Func.Max(a)
	fmt.Println(val)
	if err != nil {
		fmt.Println("test fail", err)
	}

	if val != 5 {
		fmt.Println("result is error", val)
	}
}

func fff(slice interface{}) {


	val,ok := slice.([]string)
	fmt.Println(val)
	fmt.Println(ok)

}
