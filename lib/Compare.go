package lib

import "reflect"

type Comparer interface {

	compareInt() int
	compareInt8() int
	compareInt16() int
	compareInt32() int
	compareInt64() int
	compareUint() int
	compareUint8() int
	compareUint16() int
	compareUint32() int
	compareUint64() int
	compareFloat32() int
	compareFloat64() int
	compareString() int
}

const (
	GT = 1 //>
	EQ = 0 //==
	LT = -1 //<
)


type Compare struct {
	first interface{}
	second interface{}
}

func compareInt(first, second interface{}) int {
	tFirst,_ := first.(int)
	tSecond,_ := second.(int)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareInt8(first, second interface{}) int {
	tFirst,_ := first.(int8)
	tSecond,_ := second.(int8)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareInt16(first, second interface{}) int {

	tFirst,_ := first.(int16)
	tSecond,_ := second.(int16)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareInt32(first, second interface{}) int {

	tFirst,_ := first.(int32)
	tSecond,_ := second.(int32)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareInt64(first, second interface{}) int {

	tFirst,_ := first.(int64)
	tSecond,_ := second.(int64)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareUint(first, second interface{}) int {
	tFirst,_ := first.(uint)
	tSecond,_ := second.(uint)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareUint8(first, second interface{}) int {

	tFirst,_ := first.(uint8)
	tSecond,_ := second.(uint8)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareUint16(first, second interface{}) int {

	tFirst,_ := first.(uint16)
	tSecond,_ := second.(uint16)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareUint32(first, second interface{}) int {
	tFirst,_ := first.(uint32)
	tSecond,_ := second.(uint32)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareUint64(first, second interface{}) int {
	tFirst,_ := first.(uint64)
	tSecond,_ := second.(uint64)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareFloat32(first, second interface{}) int {
	tFirst,_ := first.(float32)
	tSecond,_ := second.(float32)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareFloat64(first, second interface{}) int {
	tFirst,_ := first.(float64)
	tSecond,_ := second.(float64)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func compareString(first, second interface{}) int {
	tFirst,_ := first.(string)
	tSecond,_ := second.(string)

	if tFirst > tSecond {
		return GT
	}else if tFirst < tSecond{
		return LT
	} else {
		return EQ
	}
}

func CompareBetween(first, second interface{}) int {

	types := reflect.TypeOf(first).Kind()
	if types != reflect.TypeOf(second).Kind() {
		panic("compare type must equal")
	}

	switch reflect.TypeOf(first).Kind() {
	case reflect.Int:
		return compareInt(first, second)
	case reflect.Int8:
		return compareInt8(first, second)
	case reflect.Int16:
		return compareInt16(first, second)
	case reflect.Int32:
		return compareInt32(first, second)
	case reflect.Int64:
		return compareInt64(first, second)
	case reflect.Uint:
		return compareUint(first, second)
	case reflect.Uint8:
		return compareUint8(first, second)
	case reflect.Uint16:
		return compareUint16(first, second)
	case reflect.Uint32:
		return compareUint32(first, second)
	case reflect.Uint64:
		return compareUint64(first, second)
	case reflect.Float32:
		return compareFloat32(first, second)
	case reflect.Float64:
		return compareFloat64(first, second)
	case reflect.String:
		return compareString(first, second)
	default:
		panic("未识别的类型")
	}
}


