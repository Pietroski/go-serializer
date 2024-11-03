package go_serializer

import (
	error_builder "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/builder/errors"
	"math"
	"reflect"
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

func (s *BinarySerializer) decode(data []byte, target interface{}) int {
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

func (s *BinarySerializer) serializePrimitive(bbw *bytesWriter, data interface{}) {
	switch v := data.(type) {
	case bool:
		if v {
			bbw.put(1)
		} else {
			bbw.put(0)
		}
	case string:
		s.encodeRune(bbw, v)
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

func (s *BinarySerializer) serializeReflectPrimitive(bbw *bytesWriter, v *reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			bbw.put(1)
		} else {
			bbw.put(0)
		}
	case reflect.String:
		s.encodeRune(bbw, v.String())
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

func (s *BinarySerializer) deserializePrimitive(br *bytesReader, field *reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(s.decodeRune(br))
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

		s.serializeReflectPrimitive(bbw, &f)
	}
}

func (s *BinarySerializer) sliceArrayDecode(bbr *bytesReader, field *reflect.Value) {
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

func (s *BinarySerializer) mapEncode(bbw *bytesWriter, field *reflect.Value) {
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
			s.encodeRune(bbw, k)
			s.encodeRune(bbw, v)
		}

		return
	case map[string]interface{}:
		for k, v := range rawFieldValue {
			s.encodeRune(bbw, k)
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
			return
		}
	default:
		for _, key := range field.MapKeys() {
			// key
			bbw.write(s.encode(key.Interface()))

			// value type
			value := field.MapIndex(key)
			// value
			bbw.write(s.encode(value.Interface()))
			return
		}
	}
}

func (s *BinarySerializer) mapDecode(bbr *bytesReader, field *reflect.Value) {
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
			tmtd[s.decodeRune(bbr)] = s.decodeRune(bbr)
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[string]interface{}:
		tmtd := make(map[string]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfc interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfc))
			tmtd[s.decodeRune(bbr)] = itrfc
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
		return
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
		return
	}
}

// ################################################################################################################## \\
// primitive & reflect primitive checks -- string included
// ################################################################################################################## \\

func isPrimitive(target interface{}) bool {
	switch target.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128,
		uintptr,
		*int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64,
		*float32, *float64,
		*complex64, *complex128,
		*uintptr,
		string, *string,
		bool, *bool:
		return true
	//case nil:
	//	return true
	default:
		return false
	}
}

func isReflectPrimitive(target reflect.Kind) bool {
	switch target {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Bool:
		return true
	default:
		return false
	}
}

// ################################################################################################################## \\
// rune encoder
// ################################################################################################################## \\

func (s *BinarySerializer) encodeRune(bbw *bytesWriter, str string) {
	bs := make([]byte, 4)
	PutUint32(bs, uint32(len(str)))
	bbw.write(bs)

	bbw.write([]byte(str))
}

func (s *BinarySerializer) decodeRune(bbr *bytesReader) string {
	return string(bbr.read(int(Uint32(bbr.read(4)))))
}

// ################################################################################################################## \\
// bytes reader & bytes writer
// ################################################################################################################## \\

type bytesReader struct {
	data   []byte
	cursor int
}

func newBytesReader(data []byte) *bytesReader {
	return &bytesReader{
		data: data,
	}
}

func (bbr *bytesReader) next() byte {
	bbr.cursor++
	return bbr.data[bbr.cursor-1]
}

func (bbr *bytesReader) read(n int) []byte {
	bs := bbr.data[bbr.cursor : bbr.cursor+n]

	bbr.cursor += n
	return bs
}

func (bbr *bytesReader) yield() int {
	return bbr.cursor
}

func (bbr *bytesReader) skip(n int) {
	bbr.cursor += n
}

func (bbr *bytesReader) bytes() []byte {
	return bbr.data[:bbr.cursor]
}

func (bbr *bytesReader) bytesFromCursor() []byte {
	return bbr.data[bbr.cursor:]
}

func (bbr *bytesReader) cutBytes() []byte {
	bbr.data = bbr.data[bbr.cursor:]
	bbr.cursor = 0
	return bbr.data
}

type bytesWriter struct {
	data   []byte
	cursor int

	freeCap int // cap(data) - len(data)
}

func newBytesWriter(data []byte) *bytesWriter {
	bbw := &bytesWriter{
		data:    data,
		freeCap: cap(data) - len(data),
	}
	if bbw.freeCap == 0 {
		bbw.freeCap = cap(data)
	}
	if len(data) == 0 {
		bbw.data = bbw.data[:cap(data)]
	}

	return bbw
}

func (bbw *bytesWriter) put(b byte) {
	if 1 >= bbw.freeCap {
		newDataCap := cap(bbw.data) << 1
		nbs := make([]byte, newDataCap)
		copy(nbs, bbw.data)
		bbw.data = nbs
		bbw.freeCap = newDataCap - bbw.cursor
	}

	bbw.data[bbw.cursor] = b
	bbw.cursor++
	bbw.freeCap--
}

func (bbw *bytesWriter) write(bs []byte) {
	limit := len(bs)
	dataLimit := len(bbw.data)
	dataCap := cap(bbw.data)

	if limit > bbw.freeCap {
		newDataCap := dataCap << 1
		for dataLimit+limit-bbw.freeCap > newDataCap {
			newDataCap <<= 1
		}

		nbs := make([]byte, newDataCap)
		copy(nbs, bbw.data)
		bbw.data = nbs
		bbw.freeCap = newDataCap - bbw.cursor
	}

	copy(bbw.data[bbw.cursor:], bs)
	bbw.cursor += limit
	bbw.freeCap -= limit
}

func (bbw *bytesWriter) bytes() []byte {
	return bbw.data[:bbw.cursor]
}

// ################################################################################################################## \\
// binary little endian functions
// ################################################################################################################## \\

func Uint16(b []byte) uint16 {
	_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
	return uint16(b[0]) | uint16(b[1])<<8
}

func PutUint16(b []byte, v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}

func Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func Uint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func PutUint64(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}
