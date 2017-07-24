package Func

import "unsafe"

type Comparer interface {
	Compare(value interface{}) bool
}

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
