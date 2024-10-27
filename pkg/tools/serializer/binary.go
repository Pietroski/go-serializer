package go_serializer

import (
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
		encodeRune(bbw, v)
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

func (s *BinarySerializer) deserializePrimitive(br *bytesReader, field *reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(decodeRune(br))
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

		s.serializePrimitive(bbw, f.Interface())
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

		// this is always a primitive
		s.serializePrimitive(bbw, f.Interface())
		continue
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
			encodeRune(bbw, k)
			encodeRune(bbw, v)
		}

		return
	case map[string]interface{}:
		for k, v := range rawFieldValue {
			encodeRune(bbw, k)
			s.serializePrimitive(bbw, &v)
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
			tmtd[decodeRune(bbr)] = decodeRune(bbr)
		}
		field.Set(reflect.ValueOf(tmtd))
		return
	case map[string]interface{}:
		tmtd := make(map[string]interface{}, length)
		for i := uint32(0); i < length; i++ {
			var itrfc interface{}
			bbr.skip(s.decode(bbr.bytesFromCursor(), &itrfc))
			tmtd[decodeRune(bbr)] = itrfc
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

func encodeRune(bbw *bytesWriter, str string) {
	bs := make([]byte, 4)
	PutUint32(bs, uint32(len(str)))
	bbw.write(bs)

	bbw.write([]byte(str))
}

func decodeRune(bbr *bytesReader) string {
	return string(bbr.read(int(Uint32(bbr.read(4)))))
}

// ################################################################################################################## \\
// map type mappings
// ################################################################################################################## \\

const (
	reflectMapType = iota
	intIntMapType
	int64Int64MapType
	intInterfaceMapType
	int64InterfaceMapType
	stringStringMapType
	stringInterfaceMapType
)

func getMapType(field *reflect.Value) uint8 {
	switch field.Interface().(type) {
	case map[int]int:
		return intIntMapType
	case map[int64]int64:
		return int64Int64MapType
	case map[int]interface{}:
		return intInterfaceMapType
	case map[int64]interface{}:
		return int64InterfaceMapType
	case map[string]string:
		return stringStringMapType
	case map[string]interface{}:
		return stringInterfaceMapType
	}

	return reflectMapType
}

func fromMapType(mapType byte, length int) reflect.Value {
	switch mapType {
	case intIntMapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[int]int{}), length)
	case int64Int64MapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[int64]int64{}), length)
	case intInterfaceMapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[int]interface{}{}), length)
	case int64InterfaceMapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[int64]interface{}{}), length)
	case stringStringMapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[string]string{}), length)
	case stringInterfaceMapType:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[string]interface{}{}), length)
	default:
		return reflect.MakeMapWithSize(reflect.TypeOf(map[interface{}]interface{}{}), length)
	}
}

func keyType(keyTypeMapping map[reflect.Kind]uint8, value reflect.Value) uint8 {
	vk := value.Kind()
	kt, ok := keyTypeMapping[vk]
	if ok {
		return kt
	}

	switch vk {
	case reflect.String:
		return 1
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return 2
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return 3
	case reflect.Float32, reflect.Float64:
		return 4
	case reflect.Complex64, reflect.Complex128:
		return 5
	default:
		return 0
	}

	//switch value {
	//case reflect.String:
	//	return stringKey
	//case reflect.Int:
	//	return intKey
	//case reflect.Int8:
	//	return int8Key
	//case reflect.Int16:
	//	return int16Key
	//case reflect.Int32:
	//	return int32Key
	//case reflect.Int64:
	//	return int64Key
	//case reflect.Uint:
	//	return uintKey
	//case reflect.Uint8:
	//	return uint8Key
	//case reflect.Uint16:
	//	return uint16Key
	//case reflect.Uint32:
	//	return uint32Key
	//case reflect.Uint64:
	//	return uint64Key
	//case reflect.Float32:
	//	return float32Key
	//case reflect.Float64:
	//	return float64Key
	//}
}

func fromKeyType(keyTypeMapping map[uint8]reflect.Kind, kt uint8) reflect.Kind {
	kind, ok := keyTypeMapping[kt]
	if ok {
		return kind
	}

	switch kt {
	case 0:
		return reflect.Kind(0)
	case 1:
		return reflect.String
	case 2:
		return reflect.Int64
	case 3:
		return reflect.Uint64
	case 4:
		return reflect.Float64
	case 5:
		return reflect.Complex128
	}

	return reflect.Kind(0)
}

const (
	stringKey uint8 = 1 + iota
	intKey
	int8Key
	int16Key
	int32Key
	int64Key
	uintKey
	uint8Key
	uint16Key
	uint32Key
	uint64Key
	float32Key
	float64Key
)

const (
	stringValue uint16 = 1 + iota
	intValue
	int8Value
	int16Value
	int32Value
	int64Value
	uintValue
	uint8Value
	uint16Value
	uint32Value
	uint64Value
	float32Value
	float64Value
	complex64Value
	complex128Value
	uintptrValue
	intSliceValue
	int8SliceValue
	int16SliceValue
	int32SliceValue
	int64SliceValue
	uintSliceValue
	uint8SliceValue
	uint16SliceValue
	uint32SliceValue
	uint64SliceValue
	float32SliceValue
	float64SliceValue
	complex64SliceValue
	complex128SliceValue
	uintptrSliceValue
	structValue
	mapValue
)

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

func AppendUint16(b []byte, v uint16) []byte {
	return append(b,
		byte(v),
		byte(v>>8),
	)
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

func AppendUint32(b []byte, v uint32) []byte {
	return append(b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)
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
