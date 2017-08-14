package gotest

import (
	"testing"
	"FuncGo/Func"
)

func Benchmark_simple_string(b *testing.B) {

	makeFiler,_ := Func.NewMakeFilerByBasicType(Func.TypeT, "/code", "")
	for i := 0; i < b.N; i ++ {

		makeFiler.MakeFuncSourceWithString(`
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
	}
}