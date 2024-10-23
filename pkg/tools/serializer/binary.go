package go_serializer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
)

type (
	BinarySerializer struct{}
)

func NewBinarySerializer() *BinarySerializer {
	bs := &BinarySerializer{}

	return bs
}

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	return s.Marshal(data)
}

func (s *BinarySerializer) Deserialize(data []byte, target interface{}) error {
	return s.Unmarshal(data, target)
}

func (s *BinarySerializer) Marshal(data interface{}) ([]byte, error) {
	if isPrimitive(data) {
		return s.serializePrimitive(data)
	}

	vofd := reflect.ValueOf(data)
	if vofd.Kind() == reflect.Ptr {
		vofd = vofd.Elem()
	}

	limit := vofd.NumField()
	bb := &bytes.Buffer{}
	for idx := 0; idx < limit; idx++ {
		field := vofd.Field(idx)
		if field.Kind() == reflect.Chan {
			return nil, fmt.Errorf("invalid type %v", field.Kind())
		}

		if field.Kind() == reflect.Ptr {
			bs := make([]byte, 0, 8)
			if field.IsNil() {
				eBs, err := binary.Append(bs, binary.LittleEndian, byte(1))
				if err != nil {
					return nil, err
				}

				bb.Write(eBs)
				continue
			}

			eBs, err := binary.Append(bs, binary.LittleEndian, byte(0))
			if err != nil {
				return nil, err
			}

			bb.Write(eBs)

			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			eBs, err := s.Marshal(field.Interface())
			if err != nil {
				return nil, err
			}

			bb.Write(eBs)

			continue
		}

		eBs, err := s.serializePrimitive(field.Interface())
		if err != nil {
			return nil, err
		}

		bb.Write(eBs)
	}

	return bb.Bytes(), nil
}

func (s *BinarySerializer) Unmarshal(data []byte, target interface{}) error {
	bbf := bytes.NewBuffer(data)

	voft := reflect.ValueOf(target)
	if voft.Kind() == reflect.Ptr {
		voft = voft.Elem()
	}

	if isPrimitive(target) {
		return s.decodePrimitiveType(bbf, &voft)
	}

	limit := voft.NumField()
	for idx := 0; idx < limit; idx++ {
		field := voft.Field(idx)
		if field.Kind() == reflect.Chan {
			return fmt.Errorf("invalid type %v", field.Kind())
		}

		if field.Kind() == reflect.Ptr {
			var isItNull byte
			if err := binary.Read(bbf, binary.LittleEndian, &isItNull); err != nil {
				return err
			}
			if isItNull == 1 {
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

		if err := s.decodePrimitiveType(bbf, &field); err != nil {
			return err
		}
	}

	return nil
}

func (s *BinarySerializer) serializePrimitive(data interface{}) ([]byte, error) {
	bs := make([]byte, 0, 8)
	switch v := data.(type) {
	case bool:
		return binary.Append(bs, binary.LittleEndian, v)
	case string:
		return encodeRune(v)
	case int:
		return binary.Append(bs, binary.LittleEndian, int64(v))
	case int8:
		return binary.Append(bs, binary.LittleEndian, v)
	case int16:
		return binary.Append(bs, binary.LittleEndian, v)
	case int32:
		return binary.Append(bs, binary.LittleEndian, v)
	case int64:
		return binary.Append(bs, binary.LittleEndian, v)
	case uint:
		return binary.Append(bs, binary.LittleEndian, uint64(v))
	case uint8:
		return binary.Append(bs, binary.LittleEndian, v)
	case uint16:
		return binary.Append(bs, binary.LittleEndian, v)
	case uint32:
		return binary.Append(bs, binary.LittleEndian, v)
	case uint64:
		return binary.Append(bs, binary.LittleEndian, v)
	case float32:
		return binary.Append(bs, binary.LittleEndian, v)
	case float64:
		return binary.Append(bs, binary.LittleEndian, v)
	case complex64:
		return binary.Append(bs, binary.LittleEndian, v)
	case complex128:
		return binary.Append(bs, binary.LittleEndian, v)
	case uintptr:
		return binary.Append(bs, binary.LittleEndian, v)
	}

	return nil, fmt.Errorf("invalid type %v - type is not a primitive", reflect.TypeOf(data))
}

func (s *BinarySerializer) decodePrimitiveType(bbf *bytes.Buffer, field *reflect.Value) error {
	switch field.Kind() {
	case reflect.String:
		var str string
		if err := decodeRune(bbf, &str); err != nil {
			return err
		}

		field.SetString(str)
		return nil
	case reflect.Bool:
		var b bool
		if err := binary.Read(bbf, binary.LittleEndian, &b); err != nil {
			return err
		}

		field.SetBool(b)
		return nil
	case reflect.Int:
		var i int64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(i)

		return nil
	case reflect.Int8:
		var i int8
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(int64(i))

		return nil
	case reflect.Int16:
		var i int16
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(int64(i))

		return nil
	case reflect.Int32:
		var i int32
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(int64(i))

		return nil
	case reflect.Int64:
		var i int64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(i)

		return nil
	case reflect.Uint:
		var i uint64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetUint(i)

		return nil
	case reflect.Uint8:
		var i uint8
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetUint(uint64(i))

		return nil
	case reflect.Uint16:
		var i uint16
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetUint(uint64(i))

		return nil
	case reflect.Uint32:
		var i uint32
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetUint(uint64(i))

		return nil
	case reflect.Uint64:
		var i uint64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetUint(i)

		return nil
	case reflect.Float32:
		var i float32
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetFloat(float64(i))

		return nil
	case reflect.Float64:
		var i float64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetFloat(i)

		return nil
	case reflect.Complex64:
		var i complex64
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetComplex(complex128(i))

		return nil
	case reflect.Complex128:
		var i complex128
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetComplex(i)

		return nil
	case reflect.Uintptr:
		var i uintptr
		if err := binary.Read(bbf, binary.LittleEndian, &i); err != nil {
			return err
		}

		field.SetInt(int64(i))

		return nil
	default:
		return fmt.Errorf("unsupported type %s - not numerical", field.Kind())
	}
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

		if err := s.decodePrimitiveType(bbf, &f); err != nil {
			return err
		}
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

func encodeRune(str string) ([]byte, error) {
	rs := []rune(str)

	buf := make([]byte, 0, cap(rs)<<1)
	buf, err := binary.Append(buf, binary.LittleEndian, uint64(len(rs)))
	if err != nil {
		return nil, err
	}

	if buf, err = binary.Append(buf, binary.LittleEndian, rs); err != nil {
		return nil, nil
	}

	return buf, nil
}

func decodeRune(bbf *bytes.Buffer, target *string) error {
	var length uint64
	if err := binary.Read(bbf, binary.LittleEndian, &length); err != nil {
		return err
	}

	rs := make([]rune, length)
	for i := uint64(0); i < length; i++ {
		if err := binary.Read(bbf, binary.LittleEndian, &rs[i]); err != nil {
			return err
		}
	}

	strBuilder := &strings.Builder{}
	for _, r := range rs {
		strBuilder.WriteRune(r)
	}

	*target = strBuilder.String()

	return nil
}
