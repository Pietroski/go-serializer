package go_serializer

import (
	"math"
	"reflect"
	"unsafe"
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

func (s *RawBinarySerializer) decode(data []byte, target interface{}) int {
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

func (s *RawBinarySerializer) serializePrimitive(bbw *bytesWriter, data interface{}) bool {
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

func (s *RawBinarySerializer) deserializePrimitive(bbr *bytesReader, field *reflect.Value) bool {
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

func (s *RawBinarySerializer) serializePrimitiveSliceArray(bbw *bytesWriter, data interface{}) bool {
	switch v := data.(type) {
	case []int64:
		for _, n := range v {
			bbw.write(AddUint64(uint64(n)))
		}

		return true
	case []string:
		for _, str := range v {
			s.encodeUnsafeString(bbw, str)
		}

		return true
	}

	return false
}

func (s *RawBinarySerializer) deserializeReflectPrimitiveSliceArray(
	bbr *bytesReader, field *reflect.Value, length uint32,
) bool {
	switch field.Type().String() {
	case "[]int64":
		ii := make([]int64, length)
		for i := range ii {
			ii[i] = int64(Uint64(bbr.read(8)))
		}

		field.Set(reflect.ValueOf(ii))
		return true
	case "[]string":
		ii := make([]string, length)
		for i := range ii {
			ii[i] = s.decodeUnsafeString(bbr)
		}

		field.Set(reflect.ValueOf(ii))
		return true
	}

	return false
}

func (s *RawBinarySerializer) sliceArrayEncode(bbw *bytesWriter, field *reflect.Value) {
	fLen := field.Len()
	bbw.write(AddUint32(uint32(fLen)))

	if fLen == 0 {
		return
	}

	if s.serializePrimitiveSliceArray(bbw, field.Interface()) {
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

func (s *RawBinarySerializer) mapEncode(bbw *bytesWriter, field *reflect.Value) {
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

func (s *RawBinarySerializer) mapDecode(bbr *bytesReader, field *reflect.Value) {
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
	strLen := len(str)
	bbw.write(AddUint32(uint32(strLen)))
	bbw.write(unsafe.Slice(unsafe.StringData(str), strLen))
}

func (s *RawBinarySerializer) decodeUnsafeString(bbr *bytesReader) string {
	bs := bbr.read(int(Uint32(bbr.read(4))))
	return unsafe.String(unsafe.SliceData(bs), len(bs))
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
	capacity := cap(data)
	length := len(data)
	bbw := &bytesWriter{
		data:    data,
		freeCap: capacity - length,
	}
	if bbw.freeCap == 0 {
		bbw.freeCap = capacity
	}
	if length == 0 {
		bbw.data = bbw.data[:capacity]
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
	bsLen := len(bs)
	dataLimit := len(bbw.data)

	if bsLen > bbw.freeCap {
		newDataCap := cap(bbw.data) << 1
		currentMaxSize := dataLimit + bsLen - bbw.freeCap
		for currentMaxSize > newDataCap {
			newDataCap <<= 1
		}

		nbs := make([]byte, newDataCap)
		copy(nbs, bbw.data)
		bbw.data = nbs
		bbw.freeCap = newDataCap - bbw.cursor
	}

	copy(bbw.data[bbw.cursor:], bs)
	bbw.cursor += bsLen
	bbw.freeCap -= bsLen
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

// ################################################################################################################## \\
// addon to binary little endian functions
// ################################################################################################################## \\

func AddUint16(v uint16) []byte {
	b := [2]byte{}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	bs := b[:]
	return bs
}

func AddUint32(v uint32) []byte {
	b := [4]byte{}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	bs := b[:]
	return bs
}

func AddUint64(v uint64) []byte {
	b := [8]byte{}
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	bs := b[:]
	return bs
}
