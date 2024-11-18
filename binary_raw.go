package serializer

import (
	"math"
	"reflect"
	"unsafe"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/bytesx"
)

type RawBinarySerializer struct{}

func NewRawBinarySerializer() *RawBinarySerializer {
	return &RawBinarySerializer{}
}

// ################################################################################################################## \\
// serializer interface implementation
// ################################################################################################################## \\

func (s *RawBinarySerializer) Serialize(data interface{}) ([]byte, error) {
	return s.encode(data), nil
}

func (s *RawBinarySerializer) Deserialize(data []byte, target interface{}) error {
	s.decode(data, target)
	return nil
}

func (s *RawBinarySerializer) DataRebind(payload interface{}, target interface{}) error {
	s.decode(s.encode(payload), target)
	return nil
}

// ################################################################################################################## \\
// encoding interface implementation
// ################################################################################################################## \\

func (s *RawBinarySerializer) Marshal(data interface{}) ([]byte, error) {
	return s.Serialize(data)
}

func (s *RawBinarySerializer) Unmarshal(data []byte, target interface{}) error {
	return s.Deserialize(data, target)
}

// ################################################################################################################## \\
// private encoder implementation
// ################################################################################################################## \\

func (s *RawBinarySerializer) encode(data interface{}) []byte {
	bbw := bytesx.NewWriter(make([]byte, 1<<6))

	if s.serializePrimitive(bbw, data) {
		return bbw.Bytes()
	}

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		s.structEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Map {
		s.mapEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Chan {
		return nil
	}

	return bbw.Bytes()
}

func (s *RawBinarySerializer) reflectEncode(value reflect.Value) []byte {
	bbw := bytesx.NewWriter(make([]byte, 1<<6))

	if s.serializeReflectPrimitive(bbw, &value) {
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			bbw.Put(1)

			return bbw.Bytes()
		}

		bbw.Put(0)
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		s.structEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Map {
		s.mapEncode(bbw, &value)
		return bbw.Bytes()
	}

	if value.Kind() == reflect.Chan {
		return nil
	}

	return bbw.Bytes()
}

func (s *RawBinarySerializer) decode(data []byte, target interface{}) int {
	bbr := bytesx.NewReader(data)

	value := reflect.ValueOf(target)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if s.deserializePrimitive(bbr, &value) {
		return bbr.Yield()
	}

	if value.Kind() == reflect.Struct {
		s.structDecode(bbr, &value)
		return bbr.Yield()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayDecode(bbr, &value)
		return bbr.Yield()
	}

	if value.Kind() == reflect.Map {
		s.mapDecode(bbr, &value)
		return bbr.Yield()
	}

	return bbr.Yield()
}

func (s *RawBinarySerializer) reflectDecode(data []byte, value reflect.Value) int {
	bbr := bytesx.NewReader(data)

	if value.Kind() == reflect.Ptr {
		if bbr.Next() == 1 {
			return bbr.Yield()
		}

		value.Set(reflect.New(value.Type().Elem()))
		value = value.Elem()
	}

	if s.deserializePrimitive(bbr, &value) {
		return bbr.Yield()
	}

	if value.Kind() == reflect.Struct {
		s.structDecode(bbr, &value)
		return bbr.Yield()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayDecode(bbr, &value)
		return bbr.Yield()
	}

	if value.Kind() == reflect.Map {
		s.mapDecode(bbr, &value)
		return bbr.Yield()
	}

	return bbr.Yield()
}

// ################################################################################################################## \\
// primitive encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) serializePrimitive(bbw *bytesx.Writer, data interface{}) bool {
	switch v := data.(type) {
	case bool:
		if v {
			bbw.Put(1)
		} else {
			bbw.Put(0)
		}

		return true
	case string:
		s.encodeUnsafeString(bbw, v)
		return true
	case int:
		bbw.Write(bytesx.AddUint64(uint64(v)))
		return true
	case int8:
		bbw.Put(byte(v))
		return true
	case int16:
		bbw.Write(bytesx.AddUint16(uint16(v)))
		return true
	case int32:
		bbw.Write(bytesx.AddUint32(uint32(v)))
		return true
	case int64:
		bbw.Write(unsafe.Slice((*byte)(unsafe.Pointer(&v)), 8))
		//bbw.Write(bytesx.AddUint64(uint64(v)))
		return true
	case uint:
		bbw.Write(bytesx.AddUint64(uint64(v)))
		return true
	case uint8:
		bbw.Put(v)
		return true
	case uint16:
		bbw.Write(bytesx.AddUint16(v))
		return true
	case uint32:
		bbw.Write(bytesx.AddUint32(v))
		return true
	case uint64:
		bbw.Write(bytesx.AddUint64(v))
		return true
	case float32:
		bbw.Write(bytesx.AddUint32(math.Float32bits(v)))
		return true
	case float64:
		bbw.Write(bytesx.AddUint64(math.Float64bits(v)))
		return true
	case complex64:
		bbw.Write(bytesx.AddUint32(math.Float32bits(real(v))))
		bbw.Write(bytesx.AddUint32(math.Float32bits(imag(v))))
		return true
	case complex128:
		bbw.Write(bytesx.AddUint64(math.Float64bits(real(v))))
		bbw.Write(bytesx.AddUint64(math.Float64bits(imag(v))))
		return true
	default:
		return false
	}
}

func (s *RawBinarySerializer) serializeReflectPrimitive(bbw *bytesx.Writer, v *reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			bbw.Put(1)
		} else {
			bbw.Put(0)
		}

		return true
	case reflect.String:
		s.encodeUnsafeString(bbw, v.String())
		return true
	case reflect.Int:
		bbw.Write(bytesx.AddUint64(uint64(v.Int())))
		return true
	case reflect.Int8:
		bbw.Put(byte(v.Int()))
		return true
	case reflect.Int16:
		bbw.Write(bytesx.AddUint16(uint16(v.Int())))
		return true
	case reflect.Int32:
		bbw.Write(bytesx.AddUint32(uint32(v.Int())))
		return true
	case reflect.Int64:
		bbw.Write(bytesx.AddUint64(uint64(v.Int()))) // bbw.Write(unsafe.Slice((*byte)(v.UnsafePointer()), 8)) // AddUint64(uint64(v.Int()))
		return true
	case reflect.Uint:
		bbw.Write(bytesx.AddUint64(v.Uint()))
		return true
	case reflect.Uint8:
		bbw.Put(byte(v.Uint()))
		return true
	case reflect.Uint16:
		bbw.Write(bytesx.AddUint16(uint16(v.Uint())))
		return true
	case reflect.Uint32:
		bbw.Write(bytesx.AddUint32(uint32(v.Uint())))
		return true
	case reflect.Uint64:
		bbw.Write(bytesx.AddUint64(v.Uint()))
		return true
	case reflect.Float32:
		bbw.Write(bytesx.AddUint32(math.Float32bits(float32(v.Float()))))
		return true
	case reflect.Float64:
		bbw.Write(bytesx.AddUint64(math.Float64bits(v.Float())))
		return true
	case reflect.Complex64:
		bbw.Write(bytesx.AddUint32(math.Float32bits(real(complex64(v.Complex())))))
		bbw.Write(bytesx.AddUint32(math.Float32bits(imag(complex64(v.Complex())))))
		return true
	case reflect.Complex128:
		bbw.Write(bytesx.AddUint64(math.Float64bits(real(v.Complex()))))
		bbw.Write(bytesx.AddUint64(math.Float64bits(imag(v.Complex()))))
		return true
	case reflect.Uintptr:
		bbw.Write(bytesx.AddUint64(uint64(v.Int())))
		return true
	default:
		return false
	}
}

func (s *RawBinarySerializer) deserializePrimitive(bbr *bytesx.Reader, field *reflect.Value) bool {
	switch field.Kind() {
	case reflect.Bool:
		field.SetBool(bbr.Next() == 1)
		return true
	case reflect.String:
		field.SetString(s.decodeUnsafeString(bbr))
		return true
	case reflect.Int:
		field.SetInt(int64(bytesx.Uint64(bbr.Read(8))))
		return true
	case reflect.Int8:
		field.SetInt(int64(bbr.Next()))
		return true
	case reflect.Int16:
		field.SetInt(int64(bytesx.Uint16(bbr.Read(2))))
		return true
	case reflect.Int32:
		field.SetInt(int64(bytesx.Uint32(bbr.Read(4))))
		return true
	case reflect.Int64:
		field.SetInt(int64(bytesx.Uint64(bbr.Read(8))))
		return true
	case reflect.Uint:
		field.SetUint(bytesx.Uint64(bbr.Read(8)))
		return true
	case reflect.Uint8:
		field.SetUint(uint64(bbr.Next()))
		return true
	case reflect.Uint16:
		field.SetUint(uint64(bytesx.Uint16(bbr.Read(2))))
		return true
	case reflect.Uint32:
		field.SetUint(uint64(bytesx.Uint32(bbr.Read(4))))
		return true
	case reflect.Uint64:
		field.SetUint(bytesx.Uint64(bbr.Read(8)))
		return true
	case reflect.Float32:
		field.SetFloat(float64(math.Float32frombits(bytesx.Uint32(bbr.Read(4)))))
		return true
	case reflect.Float64:
		field.SetFloat(math.Float64frombits(bytesx.Uint64(bbr.Read(8))))
		return true
	case reflect.Complex64:
		field.SetComplex(complex(
			float64(math.Float32frombits(bytesx.Uint32(bbr.Read(4)))),
			float64(math.Float32frombits(bytesx.Uint32(bbr.Read(4)))),
		))
		return true
	case reflect.Complex128:
		field.SetComplex(complex(
			math.Float64frombits(bytesx.Uint64(bbr.Read(8))),
			math.Float64frombits(bytesx.Uint64(bbr.Read(8))),
		))
		return true
	default:
		return false
	}
}

func (s *RawBinarySerializer) serializePrimitiveSliceArray(bbw *bytesx.Writer, data interface{}) bool {
	switch v := data.(type) {
	case []bool:
		for _, b := range v {
			if b {
				bbw.Put(1)
			} else {
				bbw.Put(0)
			}
		}

		return true
	case []string:
		for _, str := range v {
			s.encodeUnsafeString(bbw, str)
		}

		return true
	case []int:
		for _, n := range v {
			bbw.Write(bytesx.AddUint64(uint64(n)))
		}

		return true
	case []int8:
		for _, n := range v {
			bbw.Put(byte(n))
		}

		return true
	case []int16:
		for _, n := range v {
			bbw.Write(bytesx.AddUint16(uint16(n)))
		}

		return true
	case []int32:
		for _, n := range v {
			bbw.Write(bytesx.AddUint32(uint32(n)))
		}

		return true
	case []int64:
		//for _, n := range v {
		//	bbw.Write(bytesx.AddUint64(uint64(n)))
		//}
		//
		//return true

		bbw.Write(unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), len(v)*8))
		return true
	case []uint:
		for _, n := range v {
			bbw.Write(bytesx.AddUint64(uint64(n)))
		}

		return true
	case []uint8:
		bbw.Write(v)
		return true
	case []uint16:
		for _, n := range v {
			bbw.Write(bytesx.AddUint16(n))
		}

		return true
	case []uint32:
		for _, n := range v {
			bbw.Write(bytesx.AddUint32(n))
		}

		return true
	case []uint64:
		for _, n := range v {
			bbw.Write(bytesx.AddUint64(n))
		}

		return true
	case []float32:
		for _, n := range v {
			bbw.Write(bytesx.AddUint32(math.Float32bits(n)))
		}

		return true
	case []float64:
		for _, n := range v {
			bbw.Write(bytesx.AddUint64(math.Float64bits(n)))
		}

		return true
	case []complex64:
		for _, n := range v {
			bbw.Write(bytesx.AddUint32(math.Float32bits(real(n))))
			bbw.Write(bytesx.AddUint32(math.Float32bits(imag(n))))
		}

		return true
	case []complex128:
		for _, n := range v {
			bbw.Write(bytesx.AddUint64(math.Float64bits(real(n))))
			bbw.Write(bytesx.AddUint64(math.Float64bits(imag(n))))
		}

		return true
	case [][]byte:
		for _, bs := range v {
			size := len(bs)
			bbw.Write(bytesx.AddUint32(uint32(size)))
			if size == 0 {
				if bs == nil {
					bbw.Put(1)
				} else {
					bbw.Put(0)
				}

				continue
			}

			bbw.Write(bs)
		}

		return true
	}

	return false
}

func (s *RawBinarySerializer) serializeReflectPrimitiveSliceArray(
	bbw *bytesx.Writer, field *reflect.Value, length int,
) bool {
	switch field.Type().String() {
	case "[]bool":
		for i := 0; i < length; i++ {
			if field.Index(i).Bool() {
				bbw.Put(1)
			} else {
				bbw.Put(0)
			}
		}

		return true
	case "[]string":
		for i := 0; i < length; i++ {
			s.encodeUnsafeString(bbw, field.Index(i).String())
		}

		return true
	case "[]int":
		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*8))
		return true
	case "[]int8":
		//for i := 0; i < length; i++ {
		//	bbw.Put(byte(field.Index(i).Int()))
		//}

		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()))
		return true
	case "[]int16":
		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*2))
		return true
	case "[]int32":
		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*4))
		return true
	case "[]int64":
		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*8))
		return true
	case "[]uint":
		//for i := 0; i < length; i++ {
		//	bbw.Write(bytesx.AddUint64(field.Index(i).Uint()))
		//}

		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*8))
		return true
	case "[]uint8":
		bbw.Write(field.Bytes())
		return true
	case "[]uint16":
		//for i := 0; i < length; i++ {
		//	bbw.Write(bytesx.AddUint16(uint16(field.Index(i).Uint())))
		//}

		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*2))
		return true
	case "[]uint32":
		//for i := 0; i < length; i++ {
		//	bbw.Write(bytesx.AddUint32(uint32(field.Index(i).Uint())))
		//}

		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*4))
		return true
	case "[]uint64":
		//for i := 0; i < length; i++ {
		//	bbw.Write(bytesx.AddUint64(field.Index(i).Uint()))
		//}

		bbw.Write(unsafe.Slice((*byte)(field.UnsafePointer()), field.Len()*8))
		return true
	case "[]float32":
		for i := 0; i < length; i++ {
			bbw.Write(bytesx.AddUint32(math.Float32bits(float32(field.Index(i).Float()))))
		}

		return true
	case "[]float64":
		for i := 0; i < length; i++ {
			bbw.Write(bytesx.AddUint64(math.Float64bits(field.Index(i).Float())))
		}

		return true
	case "[]complex64":
		for i := 0; i < length; i++ {
			bbw.Write(bytesx.AddUint32(math.Float32bits(real(complex64(field.Index(i).Complex())))))
			bbw.Write(bytesx.AddUint32(math.Float32bits(imag(complex64(field.Index(i).Complex())))))
		}

		return true
	case "[]complex128":
		for i := 0; i < length; i++ {
			bbw.Write(bytesx.AddUint64(math.Float64bits(real(field.Index(i).Complex()))))
			bbw.Write(bytesx.AddUint64(math.Float64bits(imag(field.Index(i).Complex()))))
		}

		return true
	case "[][]uint8":
		for i := 0; i < length; i++ {
			f := field.Index(i)
			size := f.Len()
			bbw.Write(bytesx.AddUint32(uint32(size)))
			if size == 0 {
				continue
			}

			bbw.Write(f.Bytes())
		}

		return true
	default:
		return false
	}
}

func (s *RawBinarySerializer) deserializeReflectPrimitiveSliceArray(
	bbr *bytesx.Reader, field *reflect.Value, length int,
) bool {
	switch field.Type().String() {
	case "[]bool":
		bb := make([]bool, length)
		for i := range bb {
			bb[i] = bbr.Next() == 1
		}

		field.Set(reflect.ValueOf(bb))
		return true
	case "[]string":
		ss := make([]string, length)
		for i := range ss {
			ss[i] = s.decodeUnsafeString(bbr)
		}

		field.Set(reflect.ValueOf(ss))
		return true
	case "[]int":
		*(*[]int64)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*int64)(unsafe.Pointer(&bbr.Read(length * 8)[0])), length)
		return true
	case "[]int8":
		*(*[]int8)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*int8)(unsafe.Pointer(&bbr.Read(length)[0])), length)
		return true
	case "[]int16":
		*(*[]int16)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*int16)(unsafe.Pointer(&bbr.Read(length * 2)[0])), length)
		return true
	case "[]int32":
		*(*[]int32)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*int32)(unsafe.Pointer(&bbr.Read(length * 4)[0])), length)
		return true
	case "[]int64":
		*(*[]int64)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*int64)(unsafe.Pointer(&bbr.Read(length * 8)[0])), length)
		return true
	case "[]uint":
		*(*[]uint64)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*uint64)(unsafe.Pointer(&bbr.Read(length * 8)[0])), length)
		return true
	case "[]uint8":
		field.SetBytes(bbr.Read(length))
		return true
	case "[]uint16":
		*(*[]uint16)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*uint16)(unsafe.Pointer(&bbr.Read(length * 2)[0])), length)
		return true
	case "[]uint32":
		*(*[]uint32)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*uint32)(unsafe.Pointer(&bbr.Read(length * 4)[0])), length)
		return true
	case "[]uint64":
		*(*[]uint64)(unsafe.Pointer(field.UnsafeAddr())) =
			unsafe.Slice((*uint64)(unsafe.Pointer(&bbr.Read(length * 8)[0])), length)
		return true
	case "[][]uint8":
		ii := make([][]byte, length)
		for i := range ii {
			l := int(bytesx.Uint32(bbr.Read(4)))
			if l == 0 {
				continue
			}

			ii[i] = bbr.Read(l)
		}

		field.Set(reflect.ValueOf(ii))
		return true
	default:
		return false
	}
}

// ################################################################################################################## \\
// struct encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) structEncode(bbw *bytesx.Writer, field *reflect.Value) {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)

		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				bbw.Put(1)

				continue
			}

			bbw.Put(0)
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			//bbw.Write(s.encode(f.Interface()))
			bbw.Write(s.reflectEncode(f))
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			s.sliceArrayEncode(bbw, &f)
			continue
		}

		if f.Kind() == reflect.Map {
			s.mapEncode(bbw, &f)
			continue
		}

		s.serializeReflectPrimitive(bbw, &f)
	}
}

func (s *RawBinarySerializer) structDecode(bbr *bytesx.Reader, field *reflect.Value) {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)

		if f.Kind() == reflect.Ptr {
			if bbr.Next() == 1 {
				continue
			}

			f.Set(reflect.New(f.Type().Elem()))
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			bbr.Skip(s.reflectDecode(bbr.BytesFromCursor(), f))
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			s.sliceArrayDecode(bbr, &f)
			continue
		}

		if f.Kind() == reflect.Map {
			s.mapDecode(bbr, &f)
			continue
		}

		s.deserializePrimitive(bbr, &f)
	}
}

// ################################################################################################################## \\
// slice & array encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) sliceArrayEncode(bbw *bytesx.Writer, field *reflect.Value) {
	fLen := field.Len()
	bbw.Write(bytesx.AddUint32(uint32(fLen)))
	if fLen == 0 {
		return
	}

	//if s.serializePrimitiveSliceArray(bbw, field.Interface()) {
	//	return
	//}
	if s.serializeReflectPrimitiveSliceArray(bbw, field, fLen) {
		return
	}

	for i := 0; i < fLen; i++ {
		f := field.Index(i)

		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				bbw.Put(1)
				continue
			}

			bbw.Put(0)
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			s.structEncode(bbw, &f)
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			bbw.Write(s.reflectEncode(f))
			continue
		}

		if f.Kind() == reflect.Map {
			s.mapEncode(bbw, &f)
			continue
		}

		// this is always a primitive
		s.serializeReflectPrimitive(bbw, &f)
		continue
	}
}

func (s *RawBinarySerializer) sliceArrayDecode(bbr *bytesx.Reader, field *reflect.Value) {
	length := int(bytesx.Uint32(bbr.Read(4)))
	if length == 0 {
		// field.Set(reflect.MakeSlice(field.Type(), 0, length))
		return
	}

	if s.deserializeReflectPrimitiveSliceArray(bbr, field, length) {
		return
	}

	field.Set(reflect.MakeSlice(field.Type(), length, length))
	for i := 0; i < length; i++ {
		f := field.Index(i)

		if s.deserializePrimitive(bbr, &f) {
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			bbr.Skip(s.reflectDecode(bbr.BytesFromCursor(), f))
			continue
		}

		if f.Kind() == reflect.Ptr {
			if bbr.Next() == 1 {
				continue
			}

			f.Set(reflect.New(f.Type().Elem()))
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			s.structDecode(bbr, &f)
			continue
		}

		if f.Kind() == reflect.Map {
			s.mapDecode(bbr, &f)
			continue
		}

		// this is always a primitive
		s.deserializePrimitive(bbr, &f)
	}
}

// ################################################################################################################## \\
// map encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) mapEncode(bbw *bytesx.Writer, field *reflect.Value) {
	fLen := field.Len()
	bbw.Write(bytesx.AddUint32(uint32(fLen)))

	if fLen == 0 {
		return
	}

	switch rawFieldValue := field.Interface().(type) {
	case map[int]int:
		for k, v := range rawFieldValue {
			bbw.Write(bytesx.AddUint64(uint64(k)))
			bbw.Write(bytesx.AddUint64(uint64(v)))
		}

		return
	case map[int64]int64:
		for k, v := range rawFieldValue {
			bbw.Write(bytesx.AddUint64(uint64(k)))
			bbw.Write(bytesx.AddUint64(uint64(v)))
		}

		return
	case map[string]string:
		for k, v := range rawFieldValue {
			s.encodeUnsafeString(bbw, k)
			s.encodeUnsafeString(bbw, v)
		}

		return

	// TODO: implement these map types
	//case map[int]interface{}:
	//	for k, v := range rawFieldValue {
	//		bbw.Write(bytesx.AddUint64(uint64(k)))
	//		bbw.Write(s.encode(v))
	//	}
	//
	//	return
	//case map[int64]interface{}:
	//	for k, v := range rawFieldValue {
	//		bbw.Write(bytesx.AddUint64(uint64(k)))
	//		bbw.Write(s.encode(v))
	//	}
	//
	//	return
	//case map[string]interface{}:
	//	for k, v := range rawFieldValue {
	//		s.encodeUnsafeString(bbw, k)
	//		bbw.Write(s.encode(v))
	//	}
	//
	//	return
	//case map[interface{}]interface{}:
	//	for k, v := range rawFieldValue {
	//		bbw.Write(s.encode(k))
	//		bbw.Write(s.encode(v))
	//	}
	default:
		for _, key := range field.MapKeys() {
			// key
			bbw.Write(s.reflectEncode(key))

			// value type
			value := field.MapIndex(key)
			// value
			bbw.Write(s.reflectEncode(value))
		}
	}
}

func (s *RawBinarySerializer) mapDecode(bbr *bytesx.Reader, field *reflect.Value) {
	length := int(bytesx.Uint32(bbr.Read(4)))
	if length == 0 {
		return
	}

	switch field.Interface().(type) {
	case map[int]int:
		tmtd := make(map[int]int, length)
		for i := 0; i < length; i++ {
			tmtd[int(bytesx.Uint64(bbr.Read(8)))] = int(bytesx.Uint64(bbr.Read(8)))
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[int64]int64:
		tmtd := make(map[int64]int64, length)
		for i := 0; i < length; i++ {
			tmtd[int64(bytesx.Uint64(bbr.Read(8)))] = int64(bytesx.Uint64(bbr.Read(8)))
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[string]string:
		tmtd := make(map[string]string, length)
		for i := 0; i < length; i++ {
			tmtd[s.decodeUnsafeString(bbr)] = s.decodeUnsafeString(bbr)
		}
		field.Set(reflect.ValueOf(tmtd))
		return

	// TODO: implement these map types
	//case map[int]interface{}:
	//	tmtd := make(map[int]interface{}, length)
	//	for i := uint32(0); i < length; i++ {
	//		var itrfc interface{}
	//		bbr.Skip(s.decode(bbr.BytesFromCursor(), &itrfc))
	//		tmtd[int(bytesx.Uint64(bbr.Read(8)))] = itrfc
	//	}
	//	field.Set(reflect.ValueOf(tmtd))
	//	return
	//case map[int64]interface{}:
	//	tmtd := make(map[int64]interface{}, length)
	//	for i := uint32(0); i < length; i++ {
	//		var itrfc interface{}
	//		bbr.Skip(s.decode(bbr.BytesFromCursor(), &itrfc))
	//		tmtd[int64(bytesx.Uint64(bbr.Read(8)))] = itrfc
	//	}
	//	field.Set(reflect.ValueOf(tmtd))
	//	return
	//case map[string]interface{}:
	//	tmtd := make(map[string]interface{}, length)
	//	for i := uint32(0); i < length; i++ {
	//		var itrfc interface{}
	//		bbr.Skip(s.decode(bbr.BytesFromCursor(), &itrfc))
	//		tmtd[s.decodeUnsafeString(bbr)] = itrfc
	//	}
	//	field.Set(reflect.ValueOf(tmtd))
	//	return
	//case map[interface{}]interface{}:
	//	tmtd := make(map[interface{}]interface{}, length)
	//	for i := uint32(0); i < length; i++ {
	//		var itrfcKey interface{}
	//		bbr.Skip(s.decode(bbr.BytesFromCursor(), &itrfcKey))
	//		var itrfcType interface{}
	//		bbr.Skip(s.decode(bbr.BytesFromCursor(), &itrfcType))
	//		tmtd[itrfcKey] = itrfcType
	//	}
	//	field.Set(reflect.ValueOf(tmtd))
	default:
		field.Set(reflect.MakeMapWithSize(field.Type(), length))
		for i := 0; i < length; i++ {
			keyValue := reflect.New(field.Type().Key()).Elem()
			bbr.Skip(s.reflectDecode(bbr.BytesFromCursor(), keyValue))

			valueValue := reflect.New(field.Type().Elem()).Elem()
			bbr.Skip(s.reflectDecode(bbr.BytesFromCursor(), valueValue))

			field.SetMapIndex(keyValue, valueValue)
		}
	}
}

// ################################################################################################################## \\
// string unsafe encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) encodeUnsafeString(bbw *bytesx.Writer, str string) {
	strLen := len(str)
	bbw.Write(bytesx.AddUint32(uint32(strLen)))
	bbw.Write(unsafe.Slice(unsafe.StringData(str), strLen))
}

func (s *RawBinarySerializer) decodeUnsafeString(bbr *bytesx.Reader) string {
	bs := bbr.Read(int(bytesx.Uint32(bbr.Read(4))))
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
