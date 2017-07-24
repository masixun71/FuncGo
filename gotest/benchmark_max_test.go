package gotest

import (
	"testing"
	"FuncGo/Func"
)
func Benchmark_max_int(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(2,3,123213214124124,21312312,-121)
	}
}

func Benchmark_max_unsafeint(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(2,3,123213214124124,21312312,-121)
	}
}


func Benchmark_max_realInt(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxInt(2,3,123213214124124,21312312,-121)
	}
}

func Benchmark_max_int8(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(int8(1), int8(21), int8(-12), int8(0))
	}
}

func Benchmark_max_unsafeint8(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(int8(1), int8(21), int8(-12), int8(0))
	}
}


func Benchmark_max_realInt8(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxInt8(int8(1), int8(21), int8(-12), int8(0))
	}
}

func Benchmark_max_int16(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(int16(1), int16(256), int16(-255), int16(0))
	}
}

func Benchmark_max_unsafeint16(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(int16(1), int16(256), int16(-255), int16(0))
	}
}


func Benchmark_max_realInt16(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxInt16(int16(1), int16(256), int16(-255), int16(0))
	}
}

func Benchmark_max_int32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(int32(2),int32(3),int32(12321321),int32(21312312),int32(-121))
	}
}

func Benchmark_max_unsafeint32(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(int32(2),int32(3),int32(12321321),int32(21312312),int32(-121))
	}
}


func Benchmark_max_realInt32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxInt32(int32(2),int32(3),int32(12321321),int32(21312312),int32(-121))
	}
}

func Benchmark_max_int64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(int64(2),int64(3),int64(124124212321321),int64(21231231312312),int64(-1211241412))
	}
}

func Benchmark_max_unsafeint64(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(int64(2),int64(3),int64(124124212321321),int64(21231231312312),int64(-1211241412))
	}
}


func Benchmark_max_realInt64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxInt64(int64(2),int64(3),int64(124124212321321),int64(21231231312312),int64(-1211241412))
	}
}


func Benchmark_max_Uint(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uint(2),uint(323221332),uint(12),uint(2),uint(0))
	}
}

func Benchmark_max_unsafeUint(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uint(2),uint(323221332),uint(12),uint(2),uint(0))
	}
}


func Benchmark_max_realUint(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUint(uint(2),uint(323221332),uint(12),uint(2),uint(0))
	}
}

func Benchmark_max_Uint8(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uint8(2),uint8(3),uint8(12),uint8(2),uint8(0))
	}
}

func Benchmark_max_unsafeUint8(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uint8(2),uint8(3),uint8(12),uint8(2),uint8(0))
	}
}


func Benchmark_max_realUint8(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUint8(uint8(2),uint8(3),uint8(12),uint8(2),uint8(0))
	}
}

func Benchmark_max_Uint16(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uint16(2),uint16(3232),uint16(12),uint16(2),uint16(0))
	}
}

func Benchmark_max_unsafeUint16(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uint16(2),uint16(3232),uint16(12),uint16(2),uint16(0))
	}
}


func Benchmark_max_realUint16(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUint16(uint16(2),uint16(3232),uint16(12),uint16(2),uint16(0))
	}
}

func Benchmark_max_Uint32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uint32(2),uint32(323232),uint32(12),uint32(2),uint32(0))
	}
}

func Benchmark_max_unsafeUint32(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uint32(2),uint32(323232),uint32(12),uint32(2),uint32(0))
	}
}


func Benchmark_max_realUint32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUint32(uint32(2),uint32(323232),uint32(12),uint32(2),uint32(0))
	}
}

func Benchmark_max_Uint64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uint64(2),uint64(32321221332),uint64(12),uint64(2),uint64(0))
	}
}

func Benchmark_max_unsafeUint64(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uint64(2),uint64(32321221332),uint64(12),uint64(2),uint64(0))
	}
}


func Benchmark_max_realUint64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUint64(uint64(2),uint64(32321221332),uint64(12),uint64(2),uint64(0))
	}
}

func Benchmark_max_UintPtr(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(uintptr(2),uintptr(32321221332),uintptr(12),uintptr(2),uintptr(0))
	}
}

func Benchmark_max_unsafeUintPtr(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(uintptr(2),uintptr(32321221332),uintptr(12),uintptr(2),uintptr(0))
	}
}


func Benchmark_max_realUintPtr(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxUintPtr(uintptr(2),uintptr(32321221332),uintptr(12),uintptr(2),uintptr(0))
	}
}

func Benchmark_max_Float32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(float32(2.1),float32(32321221332.2),float32(12.1),float32(-222.3),float32(0))
	}
}

func Benchmark_max_unsafeFloat32(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(float32(2.1),float32(32321221332.2),float32(12.1),float32(-222.3),float32(0))
	}
}


func Benchmark_max_realFloat32(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxFloat32(float32(2.1),float32(32321221332.2),float32(12.1),float32(-222.3),float32(0))
	}
}

func Benchmark_max_Float64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(float64(2.1),float64(32321221332.2),float64(12.1),float64(-222.3),float64(0))
	}
}

func Benchmark_max_unsafeFloat64(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(float64(2.1),float64(32321221332.2),float64(12.1),float64(-222.3),float64(0))
	}
}


func Benchmark_max_realFloat64(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxFloat64(float64(2.1),float64(32321221332.2),float64(12.1),float64(-222.3),float64(0))
	}
}

func Benchmark_max_String(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.Max(string("abc"),string("abd"),string("abcd"),string("abcd"),string("abcde"))
	}
}

func Benchmark_max_unsafeString(b *testing.B)  {
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(string("abc"),string("abd"),string("abcd"),string("abcd"),string("abcde"))
	}
}


func Benchmark_max_realString(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i ++ {
		Func.MaxString(string("abc"),string("abd"),string("abcd"),string("abcd"),string("abcde"))
	}
}



func Benchmark_max_Struct(b *testing.B)  {
	//b.StopTimer()
	//b.StartTimer()
	a := foo{a: 1}
	c := foo{a: 2}
	for i := 0; i < b.N; i ++ {
		Func.Max(a, c)
	}
}

func Benchmark_max_unsafeStruct(b *testing.B)  {
	a := foo{a: 1}
	c := foo{a: 2}
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(a, c)
	}
}


func Benchmark_max_unsafeArray(b *testing.B) {
	a := [5]int{1,2,3,4,5}
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(a)
	}

}

func Benchmark_max_Array(b *testing.B) {
	a := [5]int{1,2,3,4,5}
	for i := 0; i < b.N; i ++ {
		Func.Max(a)
	}

}

func Benchmark_max_unsafeSlice(b *testing.B) {
	a := []int{1,2,3,4,5}
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(a)
	}

}

func Benchmark_max_Slice(b *testing.B) {
	a := []int{1,2,3,4,5}
	for i := 0; i < b.N; i ++ {
		Func.Max(a)
	}

}

func Benchmark_max_unsafeMap(b *testing.B) {
	a := make(map[int]int, 0)
	a[0] = 1
	a[1] = 2
	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(a)
	}

}

func Benchmark_max_Map(b *testing.B) {
	a := make(map[int]int, 0)
	a[0] = 1
	a[1] = 2
	for i := 0; i < b.N; i ++ {
		Func.Max(a)
	}

}

func Benchmark_max_unsafeArrayStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}


	ff := [3]foo{a,bb,c}

	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(ff)
	}

}

func Benchmark_max_ArrayStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}


	ff := [3]foo{a,bb,c}

	for i := 0; i < b.N; i ++ {
		Func.Max(ff)
	}

}

func Benchmark_max_unsafeSliceStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}


	ff := []foo{a,bb,c}

	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(ff)
	}

}

func Benchmark_max_SliceStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}


	ff := []foo{a,bb,c}

	for i := 0; i < b.N; i ++ {
		Func.Max(ff)
	}

}


func Benchmark_max_unsafeMapStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}

	ff := make(map[int]foo, 0)
	ff[0] = a
	ff[1] = bb
	ff[2] = c

	for i := 0; i < b.N; i ++ {
		Func.UnsafeMax(ff)
	}

}

func Benchmark_max_MapStruct(b *testing.B) {
	a := foo{a: 1}
	bb := foo{a: 3}
	c := foo{a: 2}

	ff := make(map[int]foo, 0)
	ff[0] = a
	ff[1] = bb
	ff[2] = c

	for i := 0; i < b.N; i ++ {
		Func.Max(ff)
	}

}


