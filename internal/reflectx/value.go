package reflectx

import (
	"reflect"
	"unsafe"
)

type Value struct {
	*reflect.Value

	// Store internal fields we need access to
	ptr unsafe.Pointer
	typ reflect.Type
}

// ValueOf creates a new ValueExt wrapper
func ValueOf(v *reflect.Value) Value {
	// Get the pointer to the value's data
	//var ptr unsafe.Pointer
	//if v.CanAddr() {
	//	ptr = unsafe.Pointer(v.UnsafeAddr())
	//} else {
	//	// For non-addressable values, we need to get the pointer differently
	//	ptr = unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&v))[1])
	//}

	ptr := unsafe.Pointer(v.UnsafeAddr())
	return Value{
		Value: v,
		ptr:   ptr,
		typ:   v.Type(),
	}
}

//func (v Value) Elem() Value {
//	return ValueOf(v.Value.Elem())
//}
//
//func (v Value) Index(i int) Value {
//	return ValueOf(v.Value.Index(i))
//}
//
//func (v Value) Field(i int) Value {
//	return ValueOf(v.Value.Field(i))
//}

func (v Value) StringToBytes() []byte {
	// it is known to be a string
	//if v.Kind() != reflect.String {
	//	//return []byte(fmt.Sprintf("not a string - %s", v.Kind()))
	//	panic("reflect_ext: StringToBytes of non-string value")
	//}

	str := *(*string)(v.ptr)
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func Bytefy(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func Stringify(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

func (v Value) SetStringFromBytes(b []byte) {
	if !v.CanSet() {
		panic("reflect_ext: SetStringFromBytes of unaddressable value")
	}
	if v.Kind() != reflect.String {
		panic("reflect_ext: SetStringFromBytes of non-string value")
	}

	// Create a new string from the bytes without copying
	*(*string)(v.ptr) = unsafe.String(unsafe.SliceData(b), len(b))
}

// Int64Slice returns v's underlying value as []int64
func (v Value) Int64Slice() []int64 {
	if v.Kind() == reflect.Slice {
		return *(*[]int64)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*int64)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Int64Slice of unaddressable array")
	}

	panic("reflect_ext: Int64Slice of non-slice nor non-int64-slice value")
}

// Int64SliceToBytes converts []int64 to []byte
func (v Value) Int64SliceToBytes() []byte {
	slice := v.Int64Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*8)
}

// SetInt64Slice sets v's underlying value from []int64
func (v Value) SetInt64Slice(x []int64) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt64Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int64 {
		panic("reflect_ext: SetInt64Slice of non-slice nor non-int64-slice value")
	}

	*(*[]int64)(v.ptr) = x
}

// SetBytesIntoInt64Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoInt64Slice(x []byte) {
	if !v.CanSet() && len(x)%8 != 0 {
		panic("reflect_ext: SetBytesIntoInt64Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int64 {
		panic("reflect_ext: SetBytesIntoInt64Slice of non-slice nor non-int64-slice value")
	}

	*(*[]int64)(v.ptr) = unsafe.Slice((*int64)(unsafe.Pointer(&x[0])), len(x)/8)
}

// Int32Slice returns v's underlying value as []int32
func (v Value) Int32Slice() []int32 {
	if v.Kind() == reflect.Slice {
		return *(*[]int32)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*int32)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Int32Slice of unaddressable array")
	}

	panic("reflect_ext: Int32Slice of non-slice nor non-int32-slice value")
}

// Int32SliceToBytes converts []int32 to []byte
func (v Value) Int32SliceToBytes() []byte {
	slice := v.Int32Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*4)
}

// SetInt32Slice sets v's underlying value from []int32
func (v Value) SetInt32Slice(x []int32) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt32Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int32 {
		panic("reflect_ext: SetInt32Slice of non-slice nor non-int32-slice value")
	}

	*(*[]int32)(v.ptr) = x
}

// SetBytesIntoInt32Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoInt32Slice(x []byte) {
	if !v.CanSet() && len(x)%4 != 0 {
		panic("reflect_ext: SetBytesIntoInt32Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int32 {
		panic("reflect_ext: SetBytesIntoInt32Slice of non-slice nor non-int32-slice value")
	}

	*(*[]int32)(v.ptr) = unsafe.Slice((*int32)(unsafe.Pointer(&x[0])), len(x)/4)
}

// Int16Slice returns v's underlying value as []int16
func (v Value) Int16Slice() []int16 {
	if v.Kind() == reflect.Slice {
		return *(*[]int16)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*int16)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Int16Slice of unaddressable array")
	}

	panic("reflect_ext: Int16Slice of non-slice nor non-int16-slice value")
}

// Int16SliceToBytes converts []int16 to []byte
func (v Value) Int16SliceToBytes() []byte {
	slice := v.Int16Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*2)
}

// SetInt16Slice sets v's underlying value from []int16
func (v Value) SetInt16Slice(x []int16) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt16Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int16 {
		panic("reflect_ext: SetInt16Slice of non-slice nor non-int16-slice value")
	}

	*(*[]int16)(v.ptr) = x
}

// SetBytesIntoInt16Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoInt16Slice(x []byte) {
	if !v.CanSet() && len(x)%2 != 0 {
		panic("reflect_ext: SetBytesIntoInt16Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int16 {
		panic("reflect_ext: SetBytesIntoInt16Slice of non-slice nor non-int16-slice value")
	}

	*(*[]int16)(v.ptr) = unsafe.Slice((*int16)(unsafe.Pointer(&x[0])), len(x)/2)
}

// Int8Slice returns v's underlying value as []int8
func (v Value) Int8Slice() []int8 {
	if v.Kind() == reflect.Slice {
		return *(*[]int8)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*int8)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Int8Slice of unaddressable array")
	}

	panic("reflect_ext: Int8Slice of non-slice nor non-int8-slice value")
}

// Int8SliceToBytes converts []int8 to []byte
func (v Value) Int8SliceToBytes() []byte {
	slice := v.Int8Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice))
}

// SetInt8Slice sets v's underlying value from []int8
func (v Value) SetInt8Slice(x []int8) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt8Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int8 {
		panic("reflect_ext: SetInt8Slice of non-slice nor non-int8-slice value")
	}

	*(*[]int8)(v.ptr) = x
}

// SetBytesIntoInt8Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoInt8Slice(x []byte) {
	if !v.CanSet() {
		panic("reflect_ext: SetBytesIntoInt8Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Int8 {
		panic("reflect_ext: SetBytesIntoInt8Slice of non-slice nor non-int8-slice value")
	}

	*(*[]int8)(v.ptr) = unsafe.Slice((*int8)(unsafe.Pointer(&x[0])), len(x))
}

// Uint64Slice returns v's underlying value as []uint64
func (v Value) Uint64Slice() []uint64 {
	if v.Kind() == reflect.Slice {
		return *(*[]uint64)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*uint64)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Uint64Slice of unaddressable array")
	}

	panic("reflect_ext: Uint64Slice of non-slice nor non-uint64-slice value")
}

// Uint64SliceToBytes converts []uint64 to []byte
func (v Value) Uint64SliceToBytes() []byte {
	slice := v.Uint64Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*8)
}

// SetUint64Slice sets v's underlying value from []int64
func (v Value) SetUint64Slice(x []uint64) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt64Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint64 {
		panic("reflect_ext: SetUint64Slice of non-slice nor non-uint64-slice value")
	}

	*(*[]uint64)(v.ptr) = x
}

// SetBytesIntoUint64Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoUint64Slice(x []byte) {
	if !v.CanSet() && len(x)%8 != 0 {
		panic("reflect_ext: SetBytesIntoInt64Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint64 {
		panic("reflect_ext: SetBytesIntoUint64Slice of non-slice nor non-uint64-slice value")
	}

	*(*[]uint64)(v.ptr) = unsafe.Slice((*uint64)(unsafe.Pointer(&x[0])), len(x)/8)
}

// Uint32Slice returns v's underlying value as []uint32
func (v Value) Uint32Slice() []uint32 {
	if v.Kind() == reflect.Slice {
		return *(*[]uint32)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*uint32)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Uint32Slice of unaddressable array")
	}

	panic("reflect_ext: Uint32Slice of non-slice nor non-uint32-slice value")
}

// Uint32SliceToBytes converts []uint32 to []byte
func (v Value) Uint32SliceToBytes() []byte {
	slice := v.Uint32Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*4)
}

// SetUint32Slice sets v's underlying value from []int32
func (v Value) SetUint32Slice(x []uint32) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt32Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint32 {
		panic("reflect_ext: SetUint32Slice of non-slice nor non-uint32-slice value")
	}

	*(*[]uint32)(v.ptr) = x
}

// SetBytesIntoUint32Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoUint32Slice(x []byte) {
	if !v.CanSet() && len(x)%4 != 0 {
		panic("reflect_ext: SetBytesIntoInt32Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint32 {
		panic("reflect_ext: SetBytesIntoUint32Slice of non-slice nor non-uint32-slice value")
	}

	*(*[]uint32)(v.ptr) = unsafe.Slice((*uint32)(unsafe.Pointer(&x[0])), len(x)/4)
}

// Uint16Slice returns v's underlying value as []uint16
func (v Value) Uint16Slice() []uint16 {
	if v.Kind() == reflect.Slice {
		return *(*[]uint16)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*uint16)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Uint16Slice of unaddressable array")
	}

	panic("reflect_ext: Uint16Slice of non-slice nor non-uint16-slice value")
}

// Uint16SliceToBytes converts []uint16 to []byte
func (v Value) Uint16SliceToBytes() []byte {
	slice := v.Uint16Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*2)
}

// SetUint16Slice sets v's underlying value from []int16
func (v Value) SetUint16Slice(x []uint16) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt16Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint16 {
		panic("reflect_ext: SetUint16Slice of non-slice nor non-uint16-slice value")
	}

	*(*[]uint16)(v.ptr) = x
}

// SetBytesIntoUint16Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoUint16Slice(x []byte) {
	if !v.CanSet() && len(x)%2 != 0 {
		panic("reflect_ext: SetBytesIntoInt16Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint16 {
		panic("reflect_ext: SetBytesIntoUint16Slice of non-slice nor non-uint16-slice value")
	}

	*(*[]uint16)(v.ptr) = unsafe.Slice((*uint16)(unsafe.Pointer(&x[0])), len(x)/2)
}

// Uint8Slice returns v's underlying value as []uint8
func (v Value) Uint8Slice() []uint8 {
	if v.Kind() == reflect.Slice {
		return *(*[]uint8)(v.ptr)
	}

	if v.Kind() == reflect.Array {
		p := (*uint8)(v.ptr)
		return unsafe.Slice(p, v.Len())
	}

	if !v.CanAddr() {
		panic("reflect_ext: Uint8Slice of unaddressable array")
	}

	panic("reflect_ext: Uint8Slice of non-slice nor non-uint8-slice value")
}

// Uint8SliceToBytes converts []uint8 to []byte
func (v Value) Uint8SliceToBytes() []byte {
	slice := v.Uint8Slice()
	if len(slice) == 0 {
		return nil
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice))
}

// SetUint8Slice sets v's underlying value from []int8
func (v Value) SetUint8Slice(x []uint8) {
	if !v.CanSet() {
		panic("reflect_ext: SetInt8Slice of unaddressable value")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 {
		panic("reflect_ext: SetUint8Slice of non-slice nor non-uint8-slice value")
	}

	*(*[]uint8)(v.ptr) = x
}

// SetBytesIntoUint8Slice sets v's underlying value from []byte
func (v Value) SetBytesIntoUint8Slice(x []byte) {
	if !v.CanSet() && len(x) != 0 {
		panic("reflect_ext: SetBytesIntoInt8Slice of unaddressable value or does not have the correct size")
	}
	if v.Kind() != reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 {
		panic("reflect_ext: SetBytesIntoUint8Slice of non-slice nor non-uint8-slice value")
	}

	*(*[]uint8)(v.ptr) = unsafe.Slice((*uint8)(unsafe.Pointer(&x[0])), len(x))
}
