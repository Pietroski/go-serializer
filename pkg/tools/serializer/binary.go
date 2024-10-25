package go_serializer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

type (
	BinarySerializer struct{}
)

func NewBinarySerializer() *BinarySerializer {
	return &BinarySerializer{}
}

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	bbf := &bytes.Buffer{}

	if isPrimitive(data) {
		s.serializePrimitive(bbf, data)

		return bbf.Bytes(), nil
	}

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		if err := s.structEncode(bbf, &value); err != nil {
			return nil, err
		}

		return bbf.Bytes(), nil
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		if err := s.sliceArrayEncode(bbf, &value); err != nil {
			return nil, err
		}

		return bbf.Bytes(), nil
	}

	if value.Kind() == reflect.Chan {
		return nil, fmt.Errorf("invalid type %v", value.Kind())
	}

	return bbf.Bytes(), nil
}

func (s *BinarySerializer) Deserialize(data []byte, target interface{}) error {
	bbr := newBytesReader(data)

	value := reflect.ValueOf(target)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if isReflectPrimitive(value.Kind()) {
		s.deserializePrimitive(bbr, &value)
	}

	if value.Kind() == reflect.Struct {
		//s.structDecode(bbr, &value)
		limit := value.NumField()
		for idx := 0; idx < limit; idx++ {
			f := value.Field(idx)
			if f.Kind() == reflect.Ptr {
				// isItNil?
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

			if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
				s.sliceArrayDecode(bbr, &f)

				continue
			}

			if f.Kind() == reflect.Map {
				s.mapDecoder(bbr, &f)

				continue
			}

			s.deserializePrimitive(bbr, &f)
		}

		return nil
	}

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		s.sliceArrayDecode(bbr, &value)

		return nil
	}

	return nil
}

func (s *BinarySerializer) Marshal(data interface{}) ([]byte, error) {
	return s.Serialize(data)
}

func (s *BinarySerializer) Unmarshal(data []byte, target interface{}) error {
	return s.Deserialize(data, target)
}

func (s *BinarySerializer) serializePrimitive(bbf *bytes.Buffer, data interface{}) {
	switch v := data.(type) {
	case bool:
		bs := make([]byte, 1)
		if v {
			bs[0] = 1
			bbf.Write(bs)
			return
		}

		bs[0] = 0
		bbf.Write(bs)
	case string:
		encodeRune(bbf, v)
	case int:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		bbf.Write(bs)
	case int8:
		bs := make([]byte, 1)
		bs[0] = byte(v)
		bbf.Write(bs)
	case int16:
		bs := make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, uint16(v))
		bbf.Write(bs)
	case int32:
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(v))
		bbf.Write(bs)
	case int64:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		bbf.Write(bs)
	case uint:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		bbf.Write(bs)
	case uint8:
		bs := make([]byte, 1)
		bs[0] = v
		bbf.Write(bs)
	case uint16:
		bs := make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, v)
		bbf.Write(bs)
	case uint32:
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, v)
		bbf.Write(bs)
	case uint64:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, v)
		bbf.Write(bs)
	case float32:
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, math.Float32bits(v))
		bbf.Write(bs)
	case float64:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, math.Float64bits(v))
		bbf.Write(bs)
	case complex64:
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, math.Float32bits(real(v)))
		bbf.Write(bs)

		bs = make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, math.Float32bits(imag(v)))
		bbf.Write(bs)
	case complex128:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, math.Float64bits(real(v)))
		bbf.Write(bs)

		bs = make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, math.Float64bits(imag(v)))
		bbf.Write(bs)
	case uintptr:
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		bbf.Write(bs)
	}
}

func (s *BinarySerializer) deserializePrimitive(br *bytesReader, field *reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(decodeRune(br))
	case reflect.Bool:
		field.SetBool(br.next() == 1)
	case reflect.Int:
		field.SetInt(int64(binary.LittleEndian.Uint64(br.readBytes(8))))
	case reflect.Int8:
		field.SetInt(int64(br.next()))
	case reflect.Int16:
		field.SetInt(int64(binary.LittleEndian.Uint16(br.readBytes(2))))
	case reflect.Int32:
		field.SetInt(int64(binary.LittleEndian.Uint32(br.readBytes(4))))
	case reflect.Int64:
		field.SetInt(int64(binary.LittleEndian.Uint64(br.readBytes(8))))
	case reflect.Uint:
		field.SetUint(binary.LittleEndian.Uint64(br.readBytes(8)))
	case reflect.Uint8:
		field.SetUint(uint64(br.next()))
	case reflect.Uint16:
		field.SetUint(uint64(binary.LittleEndian.Uint16(br.readBytes(2))))
	case reflect.Uint32:
		field.SetUint(uint64(binary.LittleEndian.Uint32(br.readBytes(4))))
	case reflect.Uint64:
		field.SetUint(binary.LittleEndian.Uint64(br.readBytes(8)))
	case reflect.Float32:
		field.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(br.readBytes(4)))))
	case reflect.Float64:
		field.SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(br.readBytes(8))))
	case reflect.Complex64:
		field.SetComplex(complex(
			float64(math.Float32frombits(binary.LittleEndian.Uint32(br.readBytes(4)))),
			float64(math.Float32frombits(binary.LittleEndian.Uint32(br.readBytes(4)))),
		))
	case reflect.Complex128:
		field.SetComplex(complex(
			math.Float64frombits(binary.LittleEndian.Uint64(br.readBytes(8))),
			math.Float64frombits(binary.LittleEndian.Uint64(br.readBytes(8))),
		))
	case reflect.Uintptr:
		iPtr := int(binary.LittleEndian.Uint64(br.readBytes(8)))
		field.SetPointer(unsafe.Pointer(&iPtr))
	default:
	}
}

func (s *BinarySerializer) structEncode(bbf *bytes.Buffer, field *reflect.Value) error {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)
		if f.Kind() == reflect.Ptr {
			bs := make([]byte, 1)
			if f.IsNil() {
				bs[0] = 1
				bbf.Write(bs)

				continue
			}

			bs[0] = 0
			bbf.Write(bs)

			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			eBs, err := s.Serialize(f.Interface())
			if err != nil {
				return err
			}

			bbf.Write(eBs)

			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			if err := s.sliceArrayEncode(bbf, &f); err != nil {
				return err
			}

			continue
		}

		if f.Kind() == reflect.Map {
			if err := s.mapEncoder(bbf, &f); err != nil {
				return err
			}

			continue
		}

		s.serializePrimitive(bbf, f.Interface())
	}

	return nil
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
			s.structDecode(bbr, &f)

			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			s.sliceArrayDecode(bbr, &f)

			continue
		}

		if f.Kind() == reflect.Map {
			s.mapDecoder(bbr, &f)

			continue
		}

		s.deserializePrimitive(bbr, &f)
	}
}

func (s *BinarySerializer) sliceArrayEncode(bbf *bytes.Buffer, field *reflect.Value) error {
	fLen := field.Len()

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(fLen))
	bbf.Write(bs)

	for i := 0; i < fLen; i++ {
		f := field.Index(i)
		if f.Type().Kind() != reflect.Slice && f.Type().Kind() != reflect.Array {
			if f.Kind() == reflect.Ptr {
				bs = make([]byte, 1)
				if f.IsNil() {
					bs[0] = 1
					bbf.Write(bs)

					continue
				}

				bs[0] = 0
				bbf.Write(bs)

				f = f.Elem()
			}

			if f.Kind() == reflect.Struct {
				eBs, err := s.Serialize(f.Interface())
				if err != nil {
					return err
				}

				bbf.Write(eBs)

				continue
			}

			if isPrimitive(f.Interface()) {
				s.serializePrimitive(bbf, f.Interface())

				continue
			}
		}

		eBs, err := s.Serialize(f.Interface())
		if err != nil {
			return err
		}

		bbf.Write(eBs)
	}

	return nil
}

func (s *BinarySerializer) sliceArrayDecode(bbr *bytesReader, field *reflect.Value) {
	length := binary.LittleEndian.Uint32(bbr.readBytes(4))

	field.Set(reflect.MakeSlice(field.Type(), int(length), int(length)))

	for i := uint32(0); i < length; i++ {
		f := field.Index(int(i))

		if isPrimitive(f.Interface()) {
			s.deserializePrimitive(bbr, &f)

			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			s.sliceArrayDecode(bbr, &f)

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
	}
}

func (s *BinarySerializer) mapEncoder(bbf *bytes.Buffer, field *reflect.Value) error {
	fLen := field.Len()

	// map's length
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(fLen))
	bbf.Write(bs)

	keyTypeMapping := map[reflect.Kind]uint8{}

	for _, key := range field.MapKeys() {
		value := field.MapIndex(key)

		// key type
		kt := keyType(keyTypeMapping, key)
		bs = make([]byte, 1)
		bs[0] = kt
		bbf.Write(bs)

		// key
		eBs, err := s.Serialize(key.Interface())
		if err != nil {
			return err
		}
		bbf.Write(eBs)

		// value
		//if value.Kind() == reflect.Ptr {
		//	bs = make([]byte, 1)
		//	if value.IsNil() {
		//		bs[0] = 1
		//		bbf.Write(bs)
		//
		//		continue
		//	}
		//
		//	bs[0] = 0
		//	bbf.Write(bs)
		//
		//	value = value.Elem()
		//}
		//
		//if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		//	if err = s.sliceArrayEncode(bbf, &value); err != nil {
		//		return err
		//	}
		//
		//	continue
		//}

		// value type
		kt = keyType(keyTypeMapping, key)
		bs = make([]byte, 1)
		bs[0] = kt
		bbf.Write(bs)

		// value
		eBs, err = s.Serialize(value.Interface())
		if err != nil {
			return err
		}
		bbf.Write(eBs)
	}

	return nil
}

func (s *BinarySerializer) mapDecoder(bbr *bytesReader, field *reflect.Value) {
	length := binary.LittleEndian.Uint32(bbr.readBytes(4))

	field.Set(reflect.MakeMapWithSize(field.Type(), int(length)))

	keyTypeMapping := map[uint8]reflect.Kind{}

	for i := uint32(0); i < length; i++ {
		// key
		kind := fromKeyType(keyTypeMapping, bbr.next())
		fKey := reflect.New(field.Type().Key()).Elem()

		if isReflectPrimitive(kind) {
			s.deserializePrimitive(bbr, &fKey)
		}

		// value
		kind = fromKeyType(keyTypeMapping, bbr.next())
		fValue := reflect.New(field.Type().Elem()).Elem()

		if isReflectPrimitive(kind) {
			s.deserializePrimitive(bbr, &fValue)
		}

		field.SetMapIndex(fKey, fValue)
	}
}

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
	case nil:
		return true
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

func encodeRune(bbf *bytes.Buffer, str string) {
	rs := []rune(str)

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(len(rs)))
	bbf.Write(bs)

	bs = make([]byte, len(rs)<<2)
	for i, r := range rs {
		binary.LittleEndian.PutUint32(bs[4*i:], uint32(r))
	}

	bbf.Write(bs)
}

// TODO: bench it against returning
func decodeRune(br *bytesReader) string {
	length := binary.LittleEndian.Uint32(br.readBytes(4))
	bs := br.readBytes(int(length << 2))

	rb := make([]rune, int(length))
	for i := uint32(0); i < length; i++ {
		rb[i] = int32(binary.LittleEndian.Uint32(bs[4*i:]))
	}

	return string(rb)
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

func (bbr *bytesReader) readBytes(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = bbr.data[bbr.cursor+i]
	}

	bbr.cursor += n
	return bs
}
