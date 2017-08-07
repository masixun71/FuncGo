package Func

import (
	"unsafe"
	"reflect"
	"strings"
)

type T int

type TInterface interface{}

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

const (
	TypeT int = 1
	TypeInterface int = 2
)


const (
	IntPointer     reflect.Kind = iota
	Int8Pointer
	Int16Pointer
	Int32Pointer
	Int64Pointer
	UintPointer
	Uint8Pointer
	Uint16Pointer
	Uint32Pointer
	Uint64Pointer
	Float32Pointer
	Float64Pointer
	StringPointer
)

type Tmap map[reflect.Kind]string

var TYPE_STRING Tmap = Tmap{
	reflect.Int:     "int",
	reflect.Int8:    "int8",
	reflect.Int16:   "int16",
	reflect.Int32:   "int32",
	reflect.Int64:   "int64",
	reflect.Uint:    "uint",
	reflect.Uint8:   "uint8",
	reflect.Uint16:  "uint16",
	reflect.Uint32:  "uint32",
	reflect.Uint64:  "uint64",
	reflect.Float32: "float32",
	reflect.Float64: "float64",
	reflect.String:  "string",
}

var TypeDependentByBasicType = TypeDependenter{Dependents: &TYPE_STRING}

var TYPE_POINTER_STRING Tmap = Tmap{
	IntPointer:     "*int",
	Int8Pointer:    "*int8",
	Int16Pointer:   "*int16",
	Int32Pointer:   "*int32",
	Int64Pointer:   "*int64",
	UintPointer:    "*uint",
	Uint8Pointer:   "*uint8",
	Uint16Pointer:  "*uint16",
	Uint32Pointer:  "*uint32",
	Uint64Pointer:  "*uint64",
	Float32Pointer: "*float32",
	Float64Pointer: "*float64",
	StringPointer:  "*string",
}

type ToString interface {
	String() string
}

type TypeDependent interface {
	GetMap() *Tmap
}


type TypeDependenter struct {
	Dependents *Tmap
}

func (T *TypeDependenter) GetMap() *Tmap {
	return T.Dependents
}



func NewTmapByString(values ...string) TypeDependent {
	var tmap = make(Tmap, 0)


	for index,value := range values {
		tmap[reflect.Kind(index)] = value
	}

	return &TypeDependenter{Dependents:&tmap}
}



func TmapByBasicType() TypeDependent {
	return &TypeDependentByBasicType
}


func NewTmapByInterface(values ...interface{}) TypeDependent {
	var tmap = make(Tmap, 0)


	for index,value := range values {
		if vs, ok := value.(ToString); ok {
			tmap[reflect.Kind(index)] = vs.String()
		} else {
			tmap[reflect.Kind(index)] = reflect.TypeOf(value).String()
		}
	}

	return &TypeDependenter{Dependents:&tmap}
}




func TypeToFuncName(str string) string {
	return strings.Title(str)
}

func PointerTypeToFuncName(str string) string {
	return "Pointer" + TypeToFuncName(str)
}
