package gotest

import (
	"testing"
	"FuncGo/Func"
)

func Benchmark_ast_max_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Func.Max(2, 3, 123213214124124, 21312312, -121)
	}
}

func Benchmark_unsafe_ast_max_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Func.UnsafeMax(2, 3, 123213214124124, 21312312, -121)
	}
}
