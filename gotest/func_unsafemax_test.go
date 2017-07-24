package gotest

import (
	"testing"
	"FuncGo/Func"
	"unsafe"
)

func Test_max_bool(t *testing.T) {

	val, err := Func.UnsafeMax(true, false)
	if err != nil {
		t.Log("test ok")
	} else {
		t.Error("test fail", val)
	}

}

func Test_max_int(t *testing.T)  {

	val, err := Func.UnsafeMax(2,3,123213214124124,21312312,-121)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 123213214124124 {
		t.Error("result is error", val)
	}
}

func Test_max_int8(t *testing.T) {

	val, err := Func.UnsafeMax(int8(1), int8(21), int8(-12), int8(0))

	if err != nil {
		t.Error("test fail", err)
	}

	if val != int8(21) {
		t.Error("result is error", val)
	}
}

func Test_max_int16(t *testing.T) {
	val, err := Func.UnsafeMax(int16(1), int16(256), int16(-255), int16(0))

	if err != nil {
		t.Error("test fail", err)
	}

	if val != int16(256) {
		t.Error("result is error", val)
	}
}

func Test_max_int32(t *testing.T)  {
	val, err := Func.UnsafeMax(int32(2),int32(3),int32(12321321),int32(21312312),int32(-121))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != int32(21312312) {
		t.Error("result is error", val)
	}
}

func Test_max_int64(t *testing.T)  {
	val, err := Func.UnsafeMax(int64(2),int64(3),int64(124124212321321),int64(21231231312312),int64(-1211241412))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != int64(124124212321321) {
		t.Error("result is error", val)
	}
}

func Test_max_uint8(t *testing.T)  {
	val, err := Func.UnsafeMax(uint8(2),uint8(3),uint8(12),uint8(2),uint8(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uint8(12) {
		t.Error("result is error", val)
	}
}


func Test_max_uint16(t *testing.T)  {
	val, err := Func.UnsafeMax(uint16(2),uint16(3232),uint16(12),uint16(2),uint16(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uint16(3232) {
		t.Error("result is error", val)
	}
}

func Test_max_uint32(t *testing.T)  {
	val, err := Func.UnsafeMax(uint32(2),uint32(323232),uint32(12),uint32(2),uint32(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uint32(323232) {
		t.Error("result is error", val)
	}
}

func Test_max_uint(t *testing.T)  {
	val, err := Func.UnsafeMax(uint(2),uint(323221332),uint(12),uint(2),uint(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uint(323221332) {
		t.Error("result is error", val)
	}
}

func Test_max_uint64(t *testing.T)  {
	val, err := Func.UnsafeMax(uint64(2),uint64(32321221332),uint64(12),uint64(2),uint64(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uint64(32321221332) {
		t.Error("result is error", val)
	}
}

func Test_max_uinptr(t *testing.T)  {
	val, err := Func.UnsafeMax(uintptr(2),uintptr(32321221332),uintptr(12),uintptr(2),uintptr(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != uintptr(32321221332) {
		t.Error("result is error", val)
	}
}

func Test_max_float32(t *testing.T)  {
	val, err := Func.UnsafeMax(float32(2.1),float32(32321221332.2),float32(12.1),float32(-222.3),float32(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != float32(32321221332.2) {
		t.Error("result is error", val)
	}
}

func Test_max_float64(t *testing.T)  {
	val, err := Func.UnsafeMax(float64(2.1),float64(32321221332.2),float64(12.1),float64(-222.3),float64(0))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != float64(32321221332.2) {
		t.Error("result is error", val)
	}
}

func Test_max_string(t *testing.T)  {
	val, err := Func.UnsafeMax(string("abc"),string("abd"),string("abcd"),string("abcd"),string("abcde"))
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != string("abd") {
		t.Error("result is error", val)
	}
}


func Test_max_unsafePoint(t *testing.T)  {
	var a,b,c,d,e int

	val, err := Func.UnsafeMax(unsafe.Pointer(&a),unsafe.Pointer(&b),unsafe.Pointer(&c),unsafe.Pointer(&d),unsafe.Pointer(&e))
	if err != nil {
		t.Log("test ok")
	} else {
		t.Error("test fail", val)
	}
}

func Test_max_unsafeArray(t *testing.T) {

	a := [5]int{1,2,3,4,5}
	b := [6]int{1,2,3,4,5,6}
	_, err := Func.UnsafeMax(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.UnsafeMax(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 5 {
		t.Error("result is error", val)
	}

}


func Test_max_unsafeArrayStruct(t *testing.T) {

	a := foo{a: 1}
	b := foo{a: 3}
	c := foo{a: 2}


	ff := [3]foo{a,b,c}

	val, err := Func.UnsafeMax(ff)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != b {
		t.Error("result is error", val)
	}

}


func Test_max_unsafeSlice(t *testing.T) {

	a := []int{1,2,3,4,5}
	b := []int{1,2,3,4,5,6}
	_, err := Func.UnsafeMax(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.UnsafeMax(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 5 {
		t.Error("result is error", val)
	}

}

func Test_max_unsafeMap(t *testing.T)  {
	a := make(map[int]int, 0)
	a[0] = 1
	a[1] = 2
	b := make(map[int]int, 0)
	b[0] = 1
	b[1] = 2

	_, err := Func.UnsafeMax(a, b)
	if err != nil {
		t.Log("ok", err)
	}

	val, err := Func.UnsafeMax(a)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != 2 {
		t.Error("result is error", val)
	}

}

func Test_max_unsafeMapStruct(t *testing.T) {

	a := foo{a: 1}
	b := foo{a: 3}
	c := foo{a: 2}

	ff := make(map[int]foo, 0)
	ff[0] = a
	ff[1] = b
	ff[2] = c

	val, err := Func.UnsafeMax(ff)
	t.Log(val)
	if err != nil {
		t.Error("test fail", err)
	}

	if val != b {
		t.Error("result is error", val)
	}

}


type foo struct {
	a int
	Func.Comparer
}
type foo2 struct {
	b string
}

func (f foo) Compare(value interface{}) bool {
	val, _ := value.(foo)

	if f.a >= val.a {
		return true
	} else {
		return false
	}

}

func Test_max_struct(t *testing.T)  {

	a := foo{a: 1}
	b := foo2{b: "123"}
	c := foo{a: 2}

	val, err := Func.UnsafeMax(a, b)
	t.Log(val)
	if err != nil {
		t.Log("test fail", err)
	}

	val, err = Func.UnsafeMax(a, c)
	t.Log(val)
	if err != nil {
		t.Log("test fail", err)
	}

	if val != c {
		t.Error("result is error", val)
	}
}

