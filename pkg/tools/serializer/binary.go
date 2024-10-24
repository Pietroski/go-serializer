package go_serializer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"reflect"
	"unsafe"
)

type (
	BinarySerializer struct{}
)

func NewBinarySerializer() *BinarySerializer {
	bs := &BinarySerializer{}

	return bs
}

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	bbf := &bytes.Buffer{}

	if isPrimitive(data) {
		s.serializePrimitive(bbf, data)

		return bbf.Bytes(), nil
	}

	vofd := reflect.ValueOf(data)
	if vofd.Kind() == reflect.Ptr {
		vofd = vofd.Elem()
	}

	limit := vofd.NumField()
	for idx := 0; idx < limit; idx++ {
		field := vofd.Field(idx)
		if field.Kind() == reflect.Chan {
			return nil, fmt.Errorf("invalid type %v", field.Kind())
		}

		if field.Kind() == reflect.Ptr {
			bs := make([]byte, 1)
			if field.IsNil() {
				bs[0] = 1
				bbf.Write(bs)

				continue
			}

			bs[0] = 0
			bbf.Write(bs)

			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			eBs, err := s.Serialize(field.Interface())
			if err != nil {
				return nil, err
			}

			bbf.Write(eBs)

			continue
		}

		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			if err := s.sliceArrayEncoder(bbf, &field); err != nil {
				return nil, err
			}

			continue
		}

		if field.Kind() == reflect.Map {
			if err := s.mapEncoder(bbf, &field); err != nil {
				return nil, err
			}

			continue
		}

		s.serializePrimitive(bbf, field.Interface())
	}

	return bbf.Bytes(), nil
}

func (s *BinarySerializer) Deserialize(data []byte, target interface{}) error {
	bbf := bytes.NewBuffer(data)

	voft := reflect.ValueOf(target)
	if voft.Kind() == reflect.Ptr {
		voft = voft.Elem()
	}

	if isPrimitive(target) {
		return s.deserializePrimitive(bbf, &voft)
	}

	limit := voft.NumField()
	for idx := 0; idx < limit; idx++ {
		field := voft.Field(idx)
		if field.Kind() == reflect.Chan {
			return fmt.Errorf("invalid type %v", field.Kind())
		}

		if field.Kind() == reflect.Ptr {
			isItNull := make([]byte, 1)
			if _, err := io.ReadFull(bbf, isItNull); err != nil {
				return err
			}
			if isItNull[0] == 1 {
				continue
			}

			field.Set(reflect.New(field.Type().Elem()))
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			if err := s.structDecode(bbf, &field); err != nil {
				return err
			}

			continue
		}

		if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
			if err := s.sliceArrayDecoder(bbf, &field); err != nil {
				return err
			}

			continue
		}

		if err := s.deserializePrimitive(bbf, &field); err != nil {
			return err
		}
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

func (s *BinarySerializer) deserializePrimitive(bbf *bytes.Buffer, field *reflect.Value) error {
	switch field.Kind() {
	case reflect.String:
		var str string
		if err := decodeRune(bbf, &str); err != nil {
			return err
		}

		field.SetString(str)
		return nil
	case reflect.Bool:
		bs := make([]byte, 1)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetBool(bs[0] == 1)
	case reflect.Int:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetInt(int64(binary.LittleEndian.Uint64(bs)))
	case reflect.Int8:
		bs := make([]byte, 1)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetInt(int64(bs[0]))
	case reflect.Int16:
		bs := make([]byte, 2)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetInt(int64(binary.LittleEndian.Uint16(bs)))
	case reflect.Int32:
		bs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetInt(int64(binary.LittleEndian.Uint32(bs)))
	case reflect.Int64:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetInt(int64(binary.LittleEndian.Uint64(bs)))
	case reflect.Uint:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetUint(binary.LittleEndian.Uint64(bs))
	case reflect.Uint8:
		bs := make([]byte, 1)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetUint(uint64(bs[0]))
	case reflect.Uint16:
		bs := make([]byte, 2)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetUint(uint64(binary.LittleEndian.Uint16(bs)))
	case reflect.Uint32:
		bs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetUint(uint64(binary.LittleEndian.Uint32(bs)))
	case reflect.Uint64:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetUint(binary.LittleEndian.Uint64(bs))
	case reflect.Float32:
		bs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(bs))))
	case reflect.Float64:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		field.SetFloat(math.Float64frombits(binary.LittleEndian.Uint64(bs)))
	case reflect.Complex64:
		rbs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, rbs); err != nil {
			return err
		}

		ibs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, ibs); err != nil {
			return err
		}

		field.SetComplex(complex(
			float64(math.Float32frombits(binary.LittleEndian.Uint32(rbs))),
			float64(math.Float32frombits(binary.LittleEndian.Uint32(ibs))),
		))
	case reflect.Complex128:
		rbs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, rbs); err != nil {
			return err
		}

		ibs := make([]byte, 4)
		if _, err := io.ReadFull(bbf, ibs); err != nil {
			return err
		}

		field.SetComplex(complex(
			math.Float64frombits(binary.LittleEndian.Uint64(rbs)),
			math.Float64frombits(binary.LittleEndian.Uint64(ibs)),
		))
	case reflect.Uintptr:
		bs := make([]byte, 8)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		iPtr := int(binary.LittleEndian.Uint64(bs))
		field.SetPointer(unsafe.Pointer(&iPtr))
	default:
		return nil
	}

	return nil
}

func (s *BinarySerializer) structDecode(bbf *bytes.Buffer, field *reflect.Value) error {
	limit := field.NumField()
	for idx := 0; idx < limit; idx++ {
		f := field.Field(idx)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}

		if f.Kind() == reflect.Chan {
			return fmt.Errorf("invalid type %v", f.Kind())
		}

		if f.Kind() == reflect.Struct {
			if err := s.structDecode(bbf, &f); err != nil {
				return err
			}

			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			if err := s.sliceArrayDecoder(bbf, &f); err != nil {
				return err
			}

			continue
		}

		if f.Kind() == reflect.Map {
			if err := s.mapDecoder(bbf, &f); err != nil {
				return err
			}
		}

		if err := s.deserializePrimitive(bbf, &f); err != nil {
			return err
		}
	}

	return nil
}

func (s *BinarySerializer) sliceArrayEncoder(bbf *bytes.Buffer, field *reflect.Value) error {
	fLen := field.Len()

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(fLen))
	bbf.Write(bs)

	for i := 0; i < fLen; i++ {
		f := field.Index(i)
		if f.Type().Kind() != reflect.Slice && f.Type().Kind() != reflect.Array {
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

			eBs, err := s.Serialize(f.Interface())
			if err != nil {
				return err
			}

			bbf.Write(eBs)
			continue
		}

		if err := s.sliceArrayEncoder(bbf, &f); err != nil {
			return err
		}
	}

	return nil
}

func (s *BinarySerializer) sliceArrayDecoder(bbf *bytes.Buffer, field *reflect.Value) error {
	bs := make([]byte, 4)
	if _, err := io.ReadFull(bbf, bs); err != nil {
		return err
	}

	length := binary.LittleEndian.Uint32(bs)

	field.Set(reflect.MakeSlice(field.Type(), int(length), int(length)))

	for i := uint32(0); i < length; i++ {
		f := field.Index(int(i))

		if isPrimitive(f.Interface()) {
			if err := s.deserializePrimitive(bbf, &f); err != nil {
				return err
			}

			continue
		}

		if f.Kind() == reflect.Slice || f.Kind() == reflect.Array {
			if err := s.sliceArrayDecoder(bbf, &f); err != nil {
				return err
			}

			continue
		}

		if f.Kind() == reflect.Ptr {
			isItNull := make([]byte, 1)
			if _, err := io.ReadFull(bbf, isItNull); err != nil {
				return err
			}
			if isItNull[0] == 1 {
				continue
			}

			f.Set(reflect.New(f.Type().Elem()))
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {
			if err := s.structDecode(bbf, &f); err != nil {
				return err
			}

			continue
		}
	}

	return nil
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
		//	if err = s.sliceArrayEncoder(bbf, &value); err != nil {
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

func (s *BinarySerializer) mapDecoder(bbf *bytes.Buffer, field *reflect.Value) error {
	bs := make([]byte, 4)
	if _, err := io.ReadFull(bbf, bs); err != nil {
		return err
	}

	length := binary.LittleEndian.Uint32(bs)

	field.Set(reflect.MakeMapWithSize(field.Type(), int(length)))

	keyTypeMapping := map[uint8]reflect.Kind{}

	for i := uint32(0); i < length; i++ {
		// key
		bs = make([]byte, 1)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		kind := fromKeyType(keyTypeMapping, bs[0])
		fKey := reflect.New(field.Type().Key()).Elem()

		if isReflectPrimitive(kind) {
			if err := s.deserializePrimitive(bbf, &fKey); err != nil {
				return err
			}
		}

		// value
		bs = make([]byte, 1)
		if _, err := io.ReadFull(bbf, bs); err != nil {
			return err
		}

		kind = fromKeyType(keyTypeMapping, bs[0])
		fValue := reflect.New(field.Type().Elem()).Elem()

		if isReflectPrimitive(kind) {
			if err := s.deserializePrimitive(bbf, &fValue); err != nil {
				return err
			}
		}

		field.SetMapIndex(fKey, fValue)
	}

	return nil
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

func decodeRune(bbf *bytes.Buffer, target *string) error {
	bs := make([]byte, 4)
	if _, err := io.ReadFull(bbf, bs); err != nil {
		return err
	}

	length := binary.LittleEndian.Uint32(bs)

	bs = make([]byte, length<<2)
	if _, err := io.ReadFull(bbf, bs); err != nil {
		return err
	}

	bf := &bytes.Buffer{}
	for i := uint32(0); i < length; i++ {
		bf.WriteRune(int32(binary.LittleEndian.Uint32(bs[4*i:])))
	}

	*target = bf.String()

	return nil
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
