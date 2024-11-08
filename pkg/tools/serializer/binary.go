package go_serializer

import (
	"math"
	"reflect"
	"unsafe"
)

type BinarySerializer struct{}

func NewBinarySerializer() *BinarySerializer {
	return &BinarySerializer{}
}

// ################################################################################################################## \\
// serializer interface implementation
// ################################################################################################################## \\

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	return s.encode(data), nil
}

func (s *BinarySerializer) Deserialize(data []byte, target interface{}) error {
	s.decode(data, target)
	return nil
}

func (s *BinarySerializer) DataRebind(payload interface{}, target interface{}) error {
	s.decode(s.encode(payload), target)
	return nil
}

// ################################################################################################################## \\
// encoding interface implementation
// ################################################################################################################## \\

func (s *BinarySerializer) Marshal(data interface{}) ([]byte, error) {
	return s.Serialize(data)
}

func (s *BinarySerializer) Unmarshal(data []byte, target interface{}) error {
	return s.Deserialize(data, target)
}

// ################################################################################################################## \\
// private encoder implementation
// ################################################################################################################## \\

func (s *BinarySerializer) encode(data interface{}) []byte {
	bbw := newBytesWriter(make([]byte, 1<<6))

	if s.serializePrimitive(bbw, data) {
		return bbw.bytes()
	}

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		s.structEncode(bbw, &value)
		return bbw.bytes()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayEncode(bbw, &value)
		return bbw.bytes()
	}

	if value.Kind() == reflect.Map {
		s.mapEncode(bbw, &value)
		return bbw.bytes()
	}

	if value.Kind() == reflect.Chan {
		return nil
	}

	return bbw.bytes()
}

func (s *BinarySerializer) decode(data []byte, target interface{}) int {
	bbr := newBytesReader(data)

	value := reflect.ValueOf(target)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if s.deserializePrimitive(bbr, &value) {
		return bbr.yield()
	}

	if value.Kind() == reflect.Struct {
		s.structDecode(bbr, &value)
		return bbr.yield()
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayDecode(bbr, &value)
		return bbr.yield()
	}

	if value.Kind() == reflect.Map {
		s.mapDecode(bbr, &value)
		return bbr.yield()
	}

	return bbr.yield()
}

// ################################################################################################################## \\
// primitive encoder
// ################################################################################################################## \\

func (s *BinarySerializer) serializePrimitive(bbw *bytesWriter, data interface{}) bool {
	switch v := data.(type) {
	case bool:
		if v {
			bbw.put(1)
		} else {
			bbw.put(0)
		}

		return true
	case string:
		s.encodeUnsafeString(bbw, v)
		return true
	case int:
		bbw.write(AddUint64(uint64(v)))
		return true
	case int8:
		bbw.put(byte(v))
		return true
	case int16:
		bbw.write(AddUint16(uint16(v)))
		return true
	case int32:
		bbw.write(AddUint32(uint32(v)))
		return true
	case int64:
		bbw.write(AddUint64(uint64(v)))
		return true
	case uint:
		bbw.write(AddUint64(uint64(v)))
		return true
	case uint8:
		bbw.put(v)
		return true
	case uint16:
		bbw.write(AddUint16(v))
		return true
	case uint32:
		bbw.write(AddUint32(v))
		return true
	case uint64:
		bbw.write(AddUint64(v))
		return true
	case float32:
		bbw.write(AddUint32(math.Float32bits(v)))
		return true
	case float64:
		bbw.write(AddUint64(math.Float64bits(v)))
		return true
	case complex64:
		bbw.write(AddUint32(math.Float32bits(real(v))))
		bbw.write(AddUint32(math.Float32bits(imag(v))))
		return true
	case complex128:
		bbw.write(AddUint64(math.Float64bits(real(v))))
		bbw.write(AddUint64(math.Float64bits(imag(v))))
		return true
	case uintptr:
		bbw.write(AddUint64(uint64(v)))
		return true
	}

	return false
}

func (s *BinarySerializer) serializeReflectPrimitive(bbw *bytesWriter, v *reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			bbw.put(1)
		} else {
			bbw.put(0)
		}
	case reflect.String:
		s.encodeUnsafeString(bbw, v.String())
	case reflect.Int:
		bbw.write(AddUint64(uint64(v.Int())))
	case reflect.Int8:
		bbw.put(byte(v.Int()))
	case reflect.Int16:
		bbw.write(AddUint16(uint16(v.Int())))
	case reflect.Int32:
		bbw.write(AddUint32(uint32(v.Int())))
	case reflect.Int64:
		bbw.write(AddUint64(uint64(v.Int())))
	case reflect.Uint:
		bbw.write(AddUint64(v.Uint()))
	case reflect.Uint8:
		bbw.put(byte(v.Uint()))
	case reflect.Uint16:
		bbw.write(AddUint16(uint16(v.Uint())))
	case reflect.Uint32:
		bbw.write(AddUint32(uint32(v.Uint())))
	case reflect.Uint64:
		bbw.write(AddUint64(v.Uint()))
	case reflect.Float32:
		bbw.write(AddUint32(math.Float32bits(float32(v.Float()))))
	case reflect.Float64:
		bbw.write(AddUint64(math.Float64bits(v.Float())))
	case reflect.Complex64:
		bbw.write(AddUint32(math.Float32bits(real(complex64(v.Complex())))))
		bbw.write(AddUint32(math.Float32bits(imag(complex64(v.Complex())))))
	case reflect.Complex128:
		bbw.write(AddUint64(math.Float64bits(real(v.Complex()))))
		bbw.write(AddUint64(math.Float64bits(imag(v.Complex()))))
	case reflect.Uintptr:
		bbw.write(AddUint64(uint64(v.Int())))
	default:
	}
}

func (s *BinarySerializer) deserializePrimitive(bbr *bytesReader, field *reflect.Value) bool {
	switch field.Kind() {
	case reflect.String:
		field.SetString(s.decodeUnsafeString(bbr))
	case reflect.Bool:
		field.SetBool(bbr.next() == 1)
		return true
	case reflect.Int:
		field.SetInt(int64(Uint64(bbr.read(8))))
		return true
	case reflect.Int8:
		field.SetInt(int64(bbr.next()))
		return true
	case reflect.Int16:
		field.SetInt(int64(Uint16(bbr.read(2))))
		return true
	case reflect.Int32:
		field.SetInt(int64(Uint32(bbr.read(4))))
		return true
	case reflect.Int64:
		field.SetInt(int64(Uint64(bbr.read(8))))
		return true
	case reflect.Uint:
		field.SetUint(Uint64(bbr.read(8)))
		return true
	case reflect.Uint8:
		field.SetUint(uint64(bbr.next()))
		return true
	case reflect.Uint16:
		field.SetUint(uint64(Uint16(bbr.read(2))))
		return true
	case reflect.Uint32:
		field.SetUint(uint64(Uint32(bbr.read(4))))
		return true
	case reflect.Uint64:
		field.SetUint(Uint64(bbr.read(8)))
		return true
	case reflect.Float32:
		field.SetFloat(float64(math.Float32frombits(Uint32(bbr.read(4)))))
		return true
	case reflect.Float64:
		field.SetFloat(math.Float64frombits(Uint64(bbr.read(8))))
		return true
	case reflect.Complex64:
		field.SetComplex(complex(
			float64(math.Float32frombits(Uint32(bbr.read(4)))),
			float64(math.Float32frombits(Uint32(bbr.read(4)))),
		))
		return true
	case reflect.Complex128:
		field.SetComplex(complex(
			math.Float64frombits(Uint64(bbr.read(8))),
			math.Float64frombits(Uint64(bbr.read(8))),
		))
		return true
	case reflect.Uintptr:
		field.SetInt(int64(Uint64(bbr.read(8))))
		return true
	default:
		return false
	}

	return false
}

func (s *BinarySerializer) serializePrimitiveSliceArray(bbw *bytesWriter, data interface{}) bool {
	switch v := data.(type) {
	case []bool:
		for _, b := range v {
			if b {
				bbw.put(1)
			} else {
				bbw.put(0)
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
			bbw.write(AddUint64(uint64(n)))
		}

		return true
	case []int8:
		for _, n := range v {
			bbw.put(byte(n))
		}

		return true
	case []int16:
		for _, n := range v {
			bbw.write(AddUint16(uint16(n)))
		}

		return true
	case []int32:
		for _, n := range v {
			bbw.write(AddUint32(uint32(n)))
		}

		return true
	case []int64:
		for _, n := range v {
			bbw.write(AddUint64(uint64(n)))
		}

		return true
	case []uint:
		for _, n := range v {
			bbw.write(AddUint64(uint64(n)))
		}

		return true
	case []uint8:
		bbw.write(v)
		return true
	case []uint16:
		for _, n := range v {
			bbw.write(AddUint16(n))
		}

		return true
	case []uint32:
		for _, n := range v {
			bbw.write(AddUint32(n))
		}

		return true
	case []uint64:
		for _, n := range v {
			bbw.write(AddUint64(n))
		}

		return true
	case []float32:
		for _, n := range v {
			bbw.write(AddUint32(math.Float32bits(n)))
		}

		return true
	case []float64:
		for _, n := range v {
			bbw.write(AddUint64(math.Float64bits(n)))
		}

		return true
	case []complex64:
		for _, n := range v {
			bbw.write(AddUint32(math.Float32bits(real(n))))
			bbw.write(AddUint32(math.Float32bits(imag(n))))
		}

		return true
	case []complex128:
		for _, n := range v {
			bbw.write(AddUint64(math.Float64bits(real(n))))
			bbw.write(AddUint64(math.Float64bits(imag(n))))
		}

		return true
	case []uintptr:
		for _, n := range v {
			bbw.write(AddUint64(uint64(n)))
		}

		return true
	case [][]byte:
		for _, bs := range v {
			size := len(bs)
			bbw.write(AddUint32(uint32(size)))
			if size == 0 {
				continue
			}

			bbw.write(bs)
		}

		return true
	}

	return false
}

func (s *BinarySerializer) serializeReflectPrimitiveSliceArray(
	bbw *bytesWriter, field *reflect.Value, length int,
) bool {
	switch field.Type().String() {
	case "[]bool":
		for i := 0; i < length; i++ {
			if field.Index(i).Bool() {
				bbw.put(1)
			} else {
				bbw.put(0)
			}
		}

		return true
	case "[]string":
		for i := 0; i < length; i++ {
			s.encodeUnsafeString(bbw, field.Index(i).String())
		}

		return true
	case "[]int":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(uint64(field.Index(i).Int())))
		}

		return true
	case "[]int8":
		for i := 0; i < length; i++ {
			bbw.put(byte(field.Index(i).Int()))
		}

		return true
	case "[]int16":
		for i := 0; i < length; i++ {
			bbw.write(AddUint16(uint16(field.Index(i).Int())))
		}

		return true
	case "[]int32":
		for i := 0; i < length; i++ {
			bbw.write(AddUint32(uint32(field.Index(i).Int())))
		}

		return true
	case "[]int64":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(uint64(field.Index(i).Int())))
		}

		return true
	case "[]uint":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(field.Index(i).Uint()))
		}

		return true
	case "[]uint8":
		bbw.write(field.Bytes())
		return true
	case "[]uint16":
		for i := 0; i < length; i++ {
			bbw.write(AddUint16(uint16(field.Index(i).Uint())))
		}

		return true
	case "[]uint32":
		for i := 0; i < length; i++ {
			bbw.write(AddUint32(uint32(field.Index(i).Uint())))
		}

		return true
	case "[]uint64":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(field.Index(i).Uint()))
		}

		return true
	case "[]float32":
		for i := 0; i < length; i++ {
			bbw.write(AddUint32(math.Float32bits(float32(field.Index(i).Float()))))
		}

		return true
	case "[]float64":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(math.Float64bits(field.Index(i).Float())))
		}

		return true
	case "[]complex64":
		for i := 0; i < length; i++ {
			bbw.write(AddUint32(math.Float32bits(real(complex64(field.Index(i).Complex())))))
			bbw.write(AddUint32(math.Float32bits(imag(complex64(field.Index(i).Complex())))))
		}

		return true
	case "[]complex128":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(math.Float64bits(real(field.Index(i).Complex()))))
			bbw.write(AddUint64(math.Float64bits(imag(field.Index(i).Complex()))))
		}

		return true
	case "[]uintptr":
		for i := 0; i < length; i++ {
			bbw.write(AddUint64(uint64(field.Index(i).Int())))
		}

		return true
	case "[][]uint8":
		for i := 0; i < length; i++ {
			f := field.Index(i)
			size := f.Len()
			bbw.write(AddUint32(uint32(size)))
			if size == 0 {
				continue
			}

			bbw.write(f.Bytes())
		}

		return true
	}

	return false
}

func (s *BinarySerializer) deserializeReflectPrimitiveSliceArray(
	bbr *bytesReader, field *reflect.Value, length uint32,
) bool {
	switch field.Type().String() {
	case "[]string":
		ss := make([]string, length)
		for i := range ss {
			ss[i] = s.decodeUnsafeString(bbr)
		}

		field.Set(reflect.ValueOf(ss))
		return true
	case "[]bool":
		bb := make([]bool, length)
		for i := range bb {
			bb[i] = bbr.next() == 1
		}

		field.Set(reflect.ValueOf(bb))
		return true
	case "[]int":
		ii := make([]int, length)
		for i := range ii {
			ii[i] = int(Uint64(bbr.read(8)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]int8":
		ii := make([]int8, length)
		for i := range ii {
			ii[i] = int8(bbr.next())
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]int16":
		ii := make([]int16, length)
		for i := range ii {
			ii[i] = int16(Uint64(bbr.read(2)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]int32":
		ii := make([]int32, length)
		for i := range ii {
			ii[i] = int32(Uint64(bbr.read(4)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]int64":
		ii := make([]int64, length)
		for i := range ii {
			ii[i] = int64(Uint64(bbr.read(8)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]uint":
		ii := make([]uint, length)
		for i := range ii {
			ii[i] = uint(Uint64(bbr.read(8)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]uint8":
		field.Set(reflect.ValueOf(bbr.read(int(length))))
		return true
	case "[]uint16":
		ii := make([]uint16, length)
		for i := range ii {
			ii[i] = uint16(Uint64(bbr.read(2)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]uint32":
		ii := make([]uint32, length)
		for i := range ii {
			ii[i] = uint32(Uint64(bbr.read(4)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]uint64":
		ii := make([]uint64, length)
		for i := range ii {
			ii[i] = Uint64(bbr.read(8))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[][]uint8":
		ii := make([][]byte, length)
		for i := range ii {
			l := Uint32(bbr.read(4))
			if l == 0 {
				continue
			}

			ii[i] = bbr.read(int(l))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	}

	return false
}

// ################################################################################################################## \\
// struct encoder
// ################################################################################################################## \\

func (s *BinarySerializer) structEncode(bbw *bytesWriter, field *reflect.Value) {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)

		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				bbw.put(1)

				continue
			}

			bbw.put(0)
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			bbw.write(s.encode(f.Interface()))
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

func (s *BinarySerializer) structDecode(bbr *bytesReader, field *reflect.Value) {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)

		if f.Kind() == reflect.Ptr {
			// isItNil?
			if bbr.next() == 1 {
				continue
			}

			f.Set(reflect.New(f.Type().Elem()))
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			bbr.skip(s.decode(bbr.bytesFromCursor(), f.Addr().Interface()))
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

func (s *BinarySerializer) sliceArrayEncode(bbw *bytesWriter, field *reflect.Value) {
	fLen := field.Len()
	bbw.write(AddUint32(uint32(fLen)))
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
				bbw.put(1)
				continue
			}

			bbw.put(0)
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			s.structEncode(bbw, &f)
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			bbw.write(s.encode(f.Interface()))
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

func (s *BinarySerializer) sliceArrayDecode(bbr *bytesReader, field *reflect.Value) {
	length := Uint32(bbr.read(4))
	if length == 0 {
		return
	}

	if s.deserializeReflectPrimitiveSliceArray(bbr, field, length) {
		return
	}

	field.Set(reflect.MakeSlice(field.Type(), int(length), int(length)))
	for i := uint32(0); i < length; i++ {
		f := field.Index(int(i))

		if s.deserializePrimitive(bbr, &f) {
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			bbr.skip(s.decode(bbr.bytesFromCursor(), f.Addr().Interface()))
			continue
		}

		if f.Kind() == reflect.Ptr {
			if bbr.next() == 1 {
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
	}
}

// ################################################################################################################## \\
// map encoder
// ################################################################################################################## \\

func (s *BinarySerializer) mapEncode(bbw *bytesWriter, field *reflect.Value) {
	fLen := field.Len()
	bbw.write(AddUint32(uint32(fLen)))

	if fLen == 0 {
		return
	}

	switch rawFieldValue := field.Interface().(type) {
	case map[int]int:
		for k, v := range rawFieldValue {
			bbw.write(AddUint64(uint64(k)))
			bbw.write(AddUint64(uint64(v)))
		}

		return
	case map[int]interface{}:
		for k, v := range rawFieldValue {
			bbw.write(AddUint64(uint64(k)))
			bbw.write(s.encode(v))
		}

		return
	case map[int64]int64:
		for k, v := range rawFieldValue {
			bbw.write(AddUint64(uint64(k)))
			bbw.write(AddUint64(uint64(v)))
		}

		return
	case map[int64]interface{}:
		for k, v := range rawFieldValue {
			bbw.write(AddUint64(uint64(k)))
			bbw.write(s.encode(v))
		}

		return
	case map[string]string:
		for k, v := range rawFieldValue {
			s.encodeUnsafeString(bbw, k)
			s.encodeUnsafeString(bbw, v)
		}

		return
	case map[string]interface{}:
		for k, v := range rawFieldValue {
			s.encodeUnsafeString(bbw, k)
			bbw.write(s.encode(v))
		}

		return
	case map[interface{}]interface{}:
		for k, v := range rawFieldValue {
			bbw.write(s.encode(k))
			bbw.write(s.encode(v))
		}
	default:
		for _, key := range field.MapKeys() {
			// key
			bbw.write(s.encode(key.Interface()))

			// value type
			value := field.MapIndex(key)
			// value
			bbw.write(s.encode(value.Interface()))
		}
	}
}

func (s *BinarySerializer) mapDecode(bbr *bytesReader, field *reflect.Value) {
	length := Uint32(bbr.read(4))
	if length == 0 {
		return
	}

	switch field.Interface().(type) {
	case map[int]int:
		tmtd := make(map[int]int, length)
		for i := uint32(0); i < length; i++ {
			tmtd[int(Uint64(bbr.read(8)))] = int(Uint64(bbr.read(8)))
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[int]interface{}:
		tmtd := make(map[int]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfc interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfc))
			tmtd[int(Uint64(bbr.read(8)))] = itrfc
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[int64]int64:
		tmtd := make(map[int64]int64, length)
		for i := uint32(0); i < length; i++ {
			tmtd[int64(Uint64(bbr.read(8)))] = int64(Uint64(bbr.read(8)))
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[int64]interface{}:
		tmtd := make(map[int64]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfc interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfc))
			tmtd[int64(Uint64(bbr.read(8)))] = itrfc
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[string]string:
		tmtd := make(map[string]string, length)
		for i := uint32(0); i < length; i++ {
			tmtd[s.decodeUnsafeString(bbr)] = s.decodeUnsafeString(bbr)
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[string]interface{}:
		tmtd := make(map[string]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfc interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfc))
			tmtd[s.decodeUnsafeString(bbr)] = itrfc
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[interface{}]interface{}:
		tmtd := make(map[interface{}]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfcKey interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcKey))
			var itrfcType interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcType))
			tmtd[itrfcKey] = itrfcType
		}
		field.Set(reflect.ValueOf(tmtd))
	default:
		field.Set(reflect.MakeMapWithSize(field.Type(), int(length)))
		mapKeys := field.MapKeys()
		for _, mapKey := range mapKeys {
			var itrfcKey interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcKey))
			mapKey.Set(reflect.ValueOf(itrfcKey))
			mapValue := field.MapIndex(mapKey)
			var itrfcType interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcType))
			mapValue.Set(reflect.ValueOf(itrfcType))
		}
	}
}

// ################################################################################################################## \\
// string unsafe encoder
// ################################################################################################################## \\

func (s *BinarySerializer) encodeUnsafeString(bbw *bytesWriter, str string) {
	strLen := len(str)
	bbw.write(AddUint32(uint32(strLen)))
	bbw.write(unsafe.Slice(unsafe.StringData(str), strLen))
}

func (s *BinarySerializer) decodeUnsafeString(bbr *bytesReader) string {
	bs := bbr.read(int(Uint32(bbr.read(4))))
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
