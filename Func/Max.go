package Func

import (
	"reflect"
	"errors"
	"unsafe"
)


func Max(values ...interface{}) (interface{}, error) {
	firstType := reflect.TypeOf(values[0])
	res, err := checkValue(values, firstType)
	if !res {
		return nil, err
	}

	return maxSwitch(firstType)(values...)
}

func UnsafeMax(values ...interface{}) (interface{}, error) {

	firstType := reflect.TypeOf(values[0])
	res, err := checkValue(values, firstType)
	if !res {
		return nil, err
	}

	return unsafeMaxSwitch(firstType)(values...)
}

func checkValue(values []interface{}, firstType reflect.Type) (bool, error) {
	if len(values) == 0 {
		return false, nil
	}
	for _, value := range values {
		tmpType := reflect.TypeOf(value)
		if firstType != tmpType {
			return false, errors.New("The max function requires the same type of argument to compare")
		}
	}

	return true, nil
}

func maxSwitch(p reflect.Type) func(...interface{}) (interface{}, error) {
	switch p.Kind() {
	case reflect.Invalid:
		return maxInvalid
	case reflect.Bool:
		return maxBool
	case reflect.Int:
		return maxInt
	case reflect.Int8:
		return maxInt8
	case reflect.Int16:
		return maxInt16
	case reflect.Int32:
		return maxInt32
	case reflect.Int64:
		return maxInt64
	case reflect.Uint:
		return maxUint
	case reflect.Uint8:
		return maxUint8
	case reflect.Uint16:
		return maxUint16
	case reflect.Uint32:
		return maxUint32
	case reflect.Uint64:
		return maxUint64
	case reflect.Uintptr:
		return maxUintPtr
	case reflect.Float32:
		return maxFloat32
	case reflect.Float64:
		return maxFloat64
	case reflect.Complex128:
		return unsafeMaxComplex128
	case reflect.String:
		return maxString
	case reflect.UnsafePointer:
		return unsafeMaxPointer
	case reflect.Ptr:
		return maxUintPtr
	case reflect.Array:
		return maxArray
	case reflect.Chan:
		return maxChan
	case reflect.Func:
		return maxFunc
	case reflect.Interface:
		return maxInterface
	case reflect.Map:
		return maxMap
	case reflect.Slice:
		return maxSlice
	case reflect.Struct:
		return maxStruct
	}

	return maxInvalid
}

func unsafeMaxSwitch(p reflect.Type) func(...interface{}) (interface{}, error) {
	switch p.Kind() {
	case reflect.Invalid:
		return maxInvalid
	case reflect.Bool:
		return maxBool
	case reflect.Int:
		return unsafeMaxInt
	case reflect.Int8:
		return unsafeMaxInt8
	case reflect.Int16:
		return unsafeMaxInt16
	case reflect.Int32:
		return unsafeMaxInt32
	case reflect.Int64:
		return unsafeMaxInt64
	case reflect.Uint:
		return unsafeMaxUint
	case reflect.Uint8:
		return unsafeMaxUint8
	case reflect.Uint16:
		return unsafeMaxUint16
	case reflect.Uint32:
		return unsafeMaxUint32
	case reflect.Uint64:
		return unsafeMaxUint64
	case reflect.Uintptr:
		return unsafeMaxUintptr
	case reflect.Float32:
		return unsafeMaxFloat32
	case reflect.Float64:
		return unsafeMaxFloat64
	case reflect.Complex128:
		return unsafeMaxComplex128
	case reflect.String:
		return unsafeMaxString
	case reflect.UnsafePointer:
		return unsafeMaxPointer
	case reflect.Ptr:
		return unsafeMaxUintptr
	case reflect.Array:
		return unsafeMaxArray
	case reflect.Chan:
		return maxChan
	case reflect.Func:
		return maxFunc
	case reflect.Interface:
		return maxInterface
	case reflect.Map:
		return unsafeMaxMap
	case reflect.Slice:
		return unsafeMaxSlice
	case reflect.Struct:
		return maxStruct
	}

	return maxInvalid
}

func maxInvalid(values ...interface{}) (interface{}, error) {

	return nil, errors.New("Can not be resolved")

}

func maxBool(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Type Bool Can not be resolved")
}

func maxInt(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(int)
	}

	return MaxInt(arr...)
}

func MaxInt(values ...int) (int, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxInt(values ...interface{}) (interface{}, error) {

	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*int)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*int)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxInt8(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]int8, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(int8)
	}

	return MaxInt8(arr...)
}

func MaxInt8(values ...int8) (int8, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxInt8(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*int8)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*int8)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxInt16(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]int16, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(int16)
	}

	return MaxInt16(arr...)
}

func MaxInt16(values ...int16) (int16, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxInt16(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*int16)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*int16)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxInt32(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]int32, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(int32)
	}

	return MaxInt32(arr...)
}

func MaxInt32(values ...int32) (int32, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxInt32(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*int32)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*int32)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxInt64(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]int64, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(int64)
	}

	return MaxInt64(arr...)
}

func MaxInt64(values ...int64) (int64, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxInt64(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*int64)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*int64)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUint8(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uint8, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uint8)
	}

	return MaxUint8(arr...)
}

func MaxUint8(values ...uint8) (uint8, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUint8(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uint8)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uint8)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUint(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uint, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uint)
	}

	return MaxUint(arr...)
}

func MaxUint(values ...uint) (uint, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUint(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uint)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uint)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUint16(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uint16, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uint16)
	}

	return MaxUint16(arr...)
}

func MaxUint16(values ...uint16) (uint16, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUint16(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uint16)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uint16)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUint32(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uint32, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uint32)
	}

	return MaxUint32(arr...)
}

func MaxUint32(values ...uint32) (uint32, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUint32(values ...interface{}) (interface{}, error) {

	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uint32)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uint32)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUint64(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uint64, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uint64)
	}

	return MaxUint64(arr...)
}

func MaxUint64(values ...uint64) (uint64, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUint64(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uint64)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uint64)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxUintPtr(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]uintptr, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(uintptr)
	}

	return MaxUintPtr(arr...)
}

func MaxUintPtr(values ...uintptr) (uintptr, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxUintptr(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*uintptr)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*uintptr)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxFloat32(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]float32, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(float32)
	}

	return MaxFloat32(arr...)
}

func MaxFloat32(values ...float32) (float32, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxFloat32(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*float32)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*float32)(unsafe.Pointer(uintptr(tmpinterfacePtr))))
		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func maxFloat64(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]float64, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(float64)
	}

	return MaxFloat64(arr...)
}

func MaxFloat64(values ...float64) (float64, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxFloat64(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*float64)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*float64)(unsafe.Pointer(uintptr(tmpinterfacePtr))))
		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxComplex128(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Not supported yet")
}

func maxString(values ...interface{}) (interface{}, error) {

	len := len(values)
	arr := make([]string, len)
	for i := 0; i < len; i++ {
		arr[i] = values[i].(string)
	}

	return MaxString(arr...)
}

func MaxString(values ...string) (string, error) {
	len := len(values)

	max := values[0]
	for i := 1; i < len; i++ {
		tmpvalue := values[i]

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxString(values ...interface{}) (interface{}, error) {
	interfacePtr := (*emptyInterface)(unsafe.Pointer(&values[0])).word
	max := *((*string)(unsafe.Pointer(uintptr(interfacePtr))))

	for i := 1; i < len(values); i++ {
		tmpinterfacePtr := (*emptyInterface)(unsafe.Pointer(&values[i])).word
		tmpvalue := *((*string)(unsafe.Pointer(uintptr(tmpinterfacePtr))))

		if max < tmpvalue {
			max = tmpvalue
		}
	}

	return max, nil
}

func unsafeMaxPointer(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Not supported yet")
}

func maxChan(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Not supported yet")
}

func maxFunc(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Not supported yet")
}

func maxInterface(values ...interface{}) (interface{}, error) {
	return nil, errors.New("Not supported yet")
}

func maxMapDo(values ...interface{}) ([]interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an map, only one value can be passed")
	}

	typ := reflect.ValueOf(values[0])
	newIn := make([]interface{}, typ.Len())

	mapkeys := typ.MapKeys()
	for index, key := range mapkeys {
		newIn[index] = typ.MapIndex(key).Interface()
	}

	return newIn, nil
}

func maxMap(values ...interface{}) (interface{}, error) {

	newIn, err := maxMapDo(values...)
	if err != nil {
		return nil, err
	}
	value, _ := Max(newIn...)

	return value, nil
}

func unsafeMaxMap(values ...interface{}) (interface{}, error) {

	newIn, err := maxMapDo(values...)
	if err != nil {
		return nil, err
	}
	value, _ := UnsafeMax(newIn...)

	return value, nil
}

func unsafeMaxArray(values ...interface{}) (interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an array, only one value can be passed")
	}
	newIn, err := toSliceInterface(reflect.ValueOf(values[0]))
	if err != nil {
		return nil, err
	}
	value, _ := UnsafeMax(newIn...)

	return value, nil
}

func unsafeMaxSlice(values ...interface{}) (interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an slice, only one value can be passed")
	}
	newIn, _ := toSliceInterface(reflect.ValueOf(values[0]))
	value, _ := UnsafeMax(newIn...)

	return value, nil
}

func maxSliceDo(values ...interface{}) (interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an slice, only one value can be passed")
	}
	typ := reflect.TypeOf(values[0]).Elem()
	switch typ.Kind() {
	case reflect.Invalid:
		return nil, errors.New("The value of map can not be defined")
	case reflect.Bool:
		return nil, errors.New("The value of map is a boolean value")
	case reflect.Int:
		arr := values[0].([]int)
		return MaxInt(arr...)
	case reflect.Int8:
		arr := values[0].([]int8)
		return MaxInt8(arr...)
	case reflect.Int16:
		arr := values[0].([]int16)
		return MaxInt16(arr...)
	case reflect.Int32:
		arr := values[0].([]int32)
		return MaxInt32(arr...)
	case reflect.Int64:
		arr := values[0].([]int64)
		return MaxInt64(arr...)
	case reflect.Uint:
		arr := values[0].([]uint)
		return MaxUint(arr...)
	case reflect.Uint8:
		arr := values[0].([]uint8)
		return MaxUint8(arr...)
	case reflect.Uint16:
		arr := values[0].([]uint16)
		return MaxUint16(arr...)
	case reflect.Uint32:
		arr := values[0].([]uint32)
		return MaxUint32(arr...)
	case reflect.Uint64:
		arr := values[0].([]uint64)
		return MaxUint64(arr...)
	case reflect.Uintptr:
		arr := values[0].([]uintptr)
		return MaxUintPtr(arr...)
	case reflect.Float32:
		arr := values[0].([]float32)
		return MaxFloat32(arr...)
	case reflect.Float64:
		arr := values[0].([]float64)
		return MaxFloat64(arr...)
	case reflect.Complex128:
		return unsafeMaxComplex128(values[0])
	case reflect.String:
		arr := values[0].([]string)
		return MaxString(arr...)
	case reflect.UnsafePointer:
		return unsafeMaxPointer(values[0])
	case reflect.Ptr:
		return maxUintPtr(values[0])
	case reflect.Array:
		return nil, errors.New("Multidimensional arrays are not supported")
	case reflect.Chan:
		return nil, errors.New("Do not support array values for chan")
	case reflect.Func:
		return nil, errors.New("Do not support array values for func")
	case reflect.Interface:
		return nil, errors.New("Do not support array values for interface")
	case reflect.Map:
		return nil, errors.New("Do not support array values for map")
	case reflect.Slice:
		return nil, errors.New("Do not support array values for slice")
	case reflect.Struct:
		arr, _ := toSliceInterface(reflect.ValueOf(values[0]))
		return maxStruct(arr...)
	}

	return nil, nil
}

func maxArray(values ...interface{}) (interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an array, only one value can be passed")
	}
	arr, _ := toSliceInterface(reflect.ValueOf(values[0]))

	return Max(arr...)
}

func maxSlice(values ...interface{}) (interface{}, error) {
	if len(values) > 1 {
		return nil, errors.New("If the type is an slice, only one value can be passed")
	}
	return maxSliceDo(values[0])
}

func toSliceInterface(typ reflect.Value) ([]interface{}, error) {

	newIn := make([]interface{}, typ.Len())

	for i := 0; i < typ.Len(); i++ {
		newIn[i] = typ.Index(i).Interface()
	}

	return newIn, nil
}

func maxStruct(values ...interface{}) (interface{}, error) {
	max, ok := values[0].(Comparer)
	if !ok {
		return nil, errors.New("The struct needs to implement the Comparer interface ")
	}
	for i := 1; i < len(values); i++ {
		if !max.Compare(values[i]) {
			tmpValue, _ := values[i].(Comparer)
			max = tmpValue
		}
	}

	return max, nil
}
