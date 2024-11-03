package go_serializer

import (
	"math"
	"reflect"
	"unsafe"

	error_builder "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/builder/errors"
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
	bs, err := s.Serialize(payload)
	if err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

	if err = s.Deserialize(bs, target); err != nil {
		return error_builder.Err(RebinderErrMsg, err)
	}

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
	bbw := newBytesWriter(make([]byte, 1<<5))

	if isPrimitive(data) {
		s.serializePrimitive(bbw, data)

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

func (s *RawBinarySerializer) decode(data []byte, target interface{}) int {
	bbr := newBytesReader(data)

	value := reflect.ValueOf(target)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if isReflectPrimitive(value.Kind()) {
		s.deserializePrimitive(bbr, &value)

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

func (s *RawBinarySerializer) serializePrimitive(bbw *bytesWriter, data interface{}) {
	switch v := data.(type) {
	case bool:
		if v {
			bbw.put(1)
		} else {
			bbw.put(0)
		}
	case string:
		s.encodeUnsafeString(bbw, v)
	case int:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v))
		bbw.write(bs)
	case int8:
		bbw.put(byte(v))
	case int16:
		bs := make([]byte, 2)
		PutUint16(bs, uint16(v))
		bbw.write(bs)
	case int32:
		bs := make([]byte, 4)
		PutUint32(bs, uint32(v))
		bbw.write(bs)
	case int64:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v))
		bbw.write(bs)
	case uint:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v))
		bbw.write(bs)
	case uint8:
		bbw.put(v)
	case uint16:
		bs := make([]byte, 2)
		PutUint16(bs, v)
		bbw.write(bs)
	case uint32:
		bs := make([]byte, 4)
		PutUint32(bs, v)
		bbw.write(bs)
	case uint64:
		bs := make([]byte, 8)
		PutUint64(bs, v)
		bbw.write(bs)
	case float32:
		bs := make([]byte, 4)
		PutUint32(bs, math.Float32bits(v))
		bbw.write(bs)
	case float64:
		bs := make([]byte, 8)
		PutUint64(bs, math.Float64bits(v))
		bbw.write(bs)
	case complex64:
		bs := make([]byte, 4)
		PutUint32(bs, math.Float32bits(real(v)))
		bbw.write(bs)

		bs = make([]byte, 4)
		PutUint32(bs, math.Float32bits(imag(v)))
		bbw.write(bs)
	case complex128:
		bs := make([]byte, 8)
		PutUint64(bs, math.Float64bits(real(v)))
		bbw.write(bs)

		bs = make([]byte, 8)
		PutUint64(bs, math.Float64bits(imag(v)))
		bbw.write(bs)
	case uintptr:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v))
		bbw.write(bs)
	}
}

func (s *RawBinarySerializer) serializeReflectPrimitive(bbw *bytesWriter, v *reflect.Value) {
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
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v.Int()))
		bbw.write(bs)
	case reflect.Int8:
		bbw.put(byte(v.Int()))
	case reflect.Int16:
		bs := make([]byte, 2)
		PutUint16(bs, uint16(v.Int()))
		bbw.write(bs)
	case reflect.Int32:
		bs := make([]byte, 4)
		PutUint32(bs, uint32(v.Int()))
		bbw.write(bs)
	case reflect.Int64:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v.Int()))
		bbw.write(bs)
	case reflect.Uint:
		bs := make([]byte, 8)
		PutUint64(bs, v.Uint())
		bbw.write(bs)
	case reflect.Uint8:
		bbw.put(byte(v.Uint()))
	case reflect.Uint16:
		bs := make([]byte, 2)
		PutUint16(bs, uint16(v.Uint()))
		bbw.write(bs)
	case reflect.Uint32:
		bs := make([]byte, 4)
		PutUint32(bs, uint32(v.Uint()))
		bbw.write(bs)
	case reflect.Uint64:
		bs := make([]byte, 8)
		PutUint64(bs, v.Uint())
		bbw.write(bs)
	case reflect.Float32:
		bs := make([]byte, 4)
		PutUint32(bs, math.Float32bits(float32(v.Float())))
		bbw.write(bs)
	case reflect.Float64:
		bs := make([]byte, 8)
		PutUint64(bs, math.Float64bits(v.Float()))
		bbw.write(bs)
	case reflect.Complex64:
		bs := make([]byte, 4)
		PutUint32(bs, math.Float32bits(real(complex64(v.Complex()))))
		bbw.write(bs)

		bs = make([]byte, 4)
		PutUint32(bs, math.Float32bits(imag(complex64(v.Complex()))))
		bbw.write(bs)
	case reflect.Complex128:
		bs := make([]byte, 8)
		PutUint64(bs, math.Float64bits(real(v.Complex())))
		bbw.write(bs)

		bs = make([]byte, 8)
		PutUint64(bs, math.Float64bits(imag(v.Complex())))
		bbw.write(bs)
	case reflect.Uintptr:
		bs := make([]byte, 8)
		PutUint64(bs, uint64(v.Int()))
		bbw.write(bs)
	default:
	}
}

func (s *RawBinarySerializer) deserializePrimitive(br *bytesReader, field *reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(s.decodeUnsafeString(br))
	case reflect.Bool:
		field.SetBool(br.next() == 1)
	case reflect.Int:
		field.SetInt(int64(Uint64(br.read(8))))
	case reflect.Int8:
		field.SetInt(int64(br.next()))
	case reflect.Int16:
		field.SetInt(int64(Uint16(br.read(2))))
	case reflect.Int32:
		field.SetInt(int64(Uint32(br.read(4))))
	case reflect.Int64:
		field.SetInt(int64(Uint64(br.read(8))))
	case reflect.Uint:
		field.SetUint(Uint64(br.read(8)))
	case reflect.Uint8:
		field.SetUint(uint64(br.next()))
	case reflect.Uint16:
		field.SetUint(uint64(Uint16(br.read(2))))
	case reflect.Uint32:
		field.SetUint(uint64(Uint32(br.read(4))))
	case reflect.Uint64:
		field.SetUint(Uint64(br.read(8)))
	case reflect.Float32:
		field.SetFloat(float64(math.Float32frombits(Uint32(br.read(4)))))
	case reflect.Float64:
		field.SetFloat(math.Float64frombits(Uint64(br.read(8))))
	case reflect.Complex64:
		field.SetComplex(complex(
			float64(math.Float32frombits(Uint32(br.read(4)))),
			float64(math.Float32frombits(Uint32(br.read(4)))),
		))
	case reflect.Complex128:
		field.SetComplex(complex(
			math.Float64frombits(Uint64(br.read(8))),
			math.Float64frombits(Uint64(br.read(8))),
		))
	case reflect.Uintptr:
		field.SetInt(int64(Uint64(br.read(8))))
	default:
	}
}

// ################################################################################################################## \\
// struct encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) structEncode(bbw *bytesWriter, field *reflect.Value) {
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

func (s *RawBinarySerializer) structDecode(bbr *bytesReader, field *reflect.Value) {
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

func (s *RawBinarySerializer) sliceArrayEncode(bbw *bytesWriter, field *reflect.Value) {
	fLen := field.Len()

	bs := make([]byte, 4)
	PutUint32(bs, uint32(fLen))
	bbw.write(bs)

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
			// if it is an slice or array
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

func (s *RawBinarySerializer) sliceArrayDecode(bbr *bytesReader, field *reflect.Value) {
	length := Uint32(bbr.read(4))

	field.Set(reflect.MakeSlice(field.Type(), int(length), int(length)))

	for i := uint32(0); i < length; i++ {
		f := field.Index(int(i))

		if isReflectPrimitive(f.Kind()) {
			s.deserializePrimitive(bbr, &f)
			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			bbr.skip(s.decode(bbr.bytesFromCursor(), f.Addr().Interface()))
			continue
		}

		if f.Kind() == reflect.Ptr {
			// isItNil
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

func (s *RawBinarySerializer) mapEncode(bbw *bytesWriter, field *reflect.Value) {
	// map's length
	fLen := field.Len()
	bs := make([]byte, 4)
	PutUint32(bs, uint32(fLen))
	bbw.write(bs)

	switch rawFieldValue := field.Interface().(type) {
	case map[int]int:
		for k, v := range rawFieldValue {
			bs = make([]byte, 8)
			PutUint64(bs, uint64(k))
			bbw.write(bs)

			bs = make([]byte, 8)
			PutUint64(bs, uint64(v))
			bbw.write(bs)
		}

		return
	case map[int]interface{}:
		for k, v := range rawFieldValue {
			bs = make([]byte, 8)
			PutUint64(bs, uint64(k))
			bbw.write(bs)

			s.serializePrimitive(bbw, &v)
		}

		return
	case map[int64]int64:
		for k, v := range rawFieldValue {
			bs = make([]byte, 8)
			PutUint64(bs, uint64(k))
			bbw.write(bs)

			bs = make([]byte, 8)
			PutUint64(bs, uint64(v))
			bbw.write(bs)
		}

		return
	case map[int64]interface{}:
		for k, v := range rawFieldValue {
			bs = make([]byte, 8)
			PutUint64(bs, uint64(k))
			bbw.write(bs)

			s.serializePrimitive(bbw, &v)
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
		for _, key := range field.MapKeys() {
			// key
			bbw.write(s.encode(key.Interface()))

			// value type
			value := field.MapIndex(key)
			// value
			bbw.write(s.encode(value.Interface()))
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

func (s *RawBinarySerializer) mapDecode(bbr *bytesReader, field *reflect.Value) {
	length := Uint32(bbr.read(4))

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
		// temporary map to decode
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
		// temporary map to decode
		tmtd := make(map[interface{}]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfcKey interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcKey))
			var itrfcType interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfcType))
			tmtd[itrfcKey] = itrfcType
		}
		field.Set(reflect.ValueOf(tmtd))
	}
}

// ################################################################################################################## \\
// string unsafe encoder
// ################################################################################################################## \\

func (s *RawBinarySerializer) encodeUnsafeString(bbw *bytesWriter, str string) {
	bs := make([]byte, 4)
	PutUint32(bs, uint32(len(str)))
	bbw.write(bs)

	bbw.write(unsafe.Slice(unsafe.StringData(str), len(str))) // []byte(str)
}

func (s *RawBinarySerializer) decodeUnsafeString(bbr *bytesReader) string {
	bs := bbr.read(int(Uint32(bbr.read(4))))
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
