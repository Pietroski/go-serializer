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
	bbf := bytes.NewBuffer(make([]byte, 0, 1<<8))

	if isPrimitive(data) {
		if err := s.serializePrimitive(bbf, data); err != nil {
			return nil, err
		}

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
			if field.IsNil() {
				if err := binary.Write(bbf, binary.BigEndian, byte(1)); err != nil {
					return nil, err
				}

				continue
			}

			if err := binary.Write(bbf, binary.BigEndian, byte(0)); err != nil {
				return nil, err
			}

			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			eBs, err := s.Marshal(field.Interface())
			if err != nil {
				return nil, err
			}

			bbf.Write(eBs)

			continue
		}

		if err := s.serializePrimitive(bbf, field.Interface()); err != nil {
			return nil, err
		}
	}

	return bbf.Bytes(), nil
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
			if err := binary.Read(bbf, binary.BigEndian, &isItNull); err != nil {
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

func (s *BinarySerializer) serializePrimitive(bbf *bytes.Buffer, data interface{}) error {
	switch v := data.(type) {
	case bool:
		return binary.Write(bbf, binary.BigEndian, v)
	case string:
		return encodeRune(bbf, v)
	case int:
		return binary.Write(bbf, binary.BigEndian, int64(v))
	case int8:
		return binary.Write(bbf, binary.BigEndian, v)
	case int16:
		return binary.Write(bbf, binary.BigEndian, v)
	case int32:
		return binary.Write(bbf, binary.BigEndian, v)
	case int64:
		return binary.Write(bbf, binary.BigEndian, v)
	case uint:
		return binary.Write(bbf, binary.BigEndian, uint64(v))
	case uint8:
		return binary.Write(bbf, binary.BigEndian, v)
	case uint16:
		return binary.Write(bbf, binary.BigEndian, v)
	case uint32:
		return binary.Write(bbf, binary.BigEndian, v)
	case uint64:
		return binary.Write(bbf, binary.BigEndian, v)
	case float32:
		return binary.Write(bbf, binary.BigEndian, v)
	case float64:
		return binary.Write(bbf, binary.BigEndian, v)
	case complex64:
		return binary.Write(bbf, binary.BigEndian, v)
	case complex128:
		return binary.Write(bbf, binary.BigEndian, v)
	case uintptr:
		return binary.Write(bbf, binary.BigEndian, v)
	}

	return fmt.Errorf("invalid type %v - type is not a primitive", reflect.TypeOf(data))
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
		if err := binary.Read(bbf, binary.BigEndian, &b); err != nil {
			return err
		}

		field.SetBool(b)
		return nil
	case reflect.Int:
		return numericIntDecoder[int64](bbf, field)
	case reflect.Int8:
		return numericIntDecoder[int8](bbf, field)
	case reflect.Int16:
		return numericIntDecoder[int16](bbf, field)
	case reflect.Int32:
		return numericIntDecoder[int32](bbf, field)
	case reflect.Int64:
		return numericIntDecoder[int64](bbf, field)
	case reflect.Uint:
		return numericUintDecoder[uint64](bbf, field)
	case reflect.Uint8:
		return numericUintDecoder[uint8](bbf, field)
	case reflect.Uint16:
		return numericUintDecoder[uint16](bbf, field)
	case reflect.Uint32:
		return numericUintDecoder[uint32](bbf, field)
	case reflect.Uint64:
		return numericUintDecoder[uint64](bbf, field)
	case reflect.Float32:
		return numericFloatDecoder[float32](bbf, field)
	case reflect.Float64:
		return numericFloatDecoder[float64](bbf, field)
	case reflect.Complex64:
		return numericComplexDecoder[complex64](bbf, field)
	case reflect.Complex128:
		return numericComplexDecoder[complex128](bbf, field)
	case reflect.Uintptr:
		return numericUintPtrDecoder[uintptr](bbf, field)
	default:
		return fmt.Errorf("unsupported type %s - not numerical", field.Kind())
	}
}

func numericIntDecoder[N numericalInt](
	bbf *bytes.Buffer, field *reflect.Value,
) error {
	var i N
	if err := binary.Read(bbf, binary.BigEndian, &i); err != nil {
		return err
	}

	field.SetInt(toInt64(i))

	return nil
}

func numericUintDecoder[N numericalUint](
	bbf *bytes.Buffer, field *reflect.Value,
) error {
	var i N
	if err := binary.Read(bbf, binary.BigEndian, &i); err != nil {
		return err
	}

	field.SetUint(toUint64(i))

	return nil
}

func numericFloatDecoder[N numericFloat](
	bbf *bytes.Buffer, field *reflect.Value,
) error {
	var i N
	if err := binary.Read(bbf, binary.BigEndian, &i); err != nil {
		return err
	}

	field.SetFloat(toFloat64(i))

	return nil
}

func numericComplexDecoder[N numericComplex](
	bbf *bytes.Buffer, field *reflect.Value,
) error {
	var i N
	if err := binary.Read(bbf, binary.BigEndian, &i); err != nil {
		return err
	}

	field.SetComplex(toComplex128(i))

	return nil
}

func numericUintPtrDecoder[N numericUintPtr](
	bbf *bytes.Buffer, field *reflect.Value,
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i)
	if err != nil {
		return err
	}

	field.SetInt(toInt64(i))

	return nil
}

func encodeRune(bbf *bytes.Buffer, str string) error {
	rs := []rune(str)

	if err := binary.Write(bbf, binary.BigEndian, uint64(len(rs))); err != nil {
		return err
	}

	if err := binary.Write(bbf, binary.BigEndian, rs); err != nil {
		return err
	}

	return nil
}

func decodeRune(bbf *bytes.Buffer, target *string) error {
	var length uint64
	if err := binary.Read(bbf, binary.BigEndian, &length); err != nil {
		return err
	}

	rs := make([]rune, length)
	if err := binary.Read(bbf, binary.BigEndian, &rs); err != nil {
		return err
	}

	strBuilder := &strings.Builder{}
	for _, r := range rs {
		strBuilder.WriteRune(r)
	}

	*target = strBuilder.String()

	return nil
}

type (
	numericalInt interface {
		int | int8 | int16 | int32 | int64
	}

	numericalUint interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	numericFloat interface {
		float32 | float64
	}

	numericComplex interface {
		complex64 | complex128
	}

	numericUintPtr interface {
		uintptr
	}
)

func toInt64(v interface{}) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		return int64(n)
	case float32:
		return int64(n)
	case float64:
		return int64(n)
	case uintptr:
		return int64(n)
	}

	return 0
}

func toUint64(v interface{}) uint64 {
	switch n := v.(type) {
	case int:
		return uint64(n)
	case int8:
		return uint64(n)
	case int16:
		return uint64(n)
	case int32:
		return uint64(n)
	case int64:
		return uint64(n)
	case uint:
		return uint64(n)
	case uint8:
		return uint64(n)
	case uint16:
		return uint64(n)
	case uint32:
		return uint64(n)
	case uint64:
		return n
	case float32:
		return uint64(n)
	case float64:
		return uint64(n)
	}

	return 0
}

func toFloat64(v interface{}) float64 {
	switch n := v.(type) {
	case int:
		return float64(n)
	case int8:
		return float64(n)
	case int16:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	case uint:
		return float64(n)
	case uint8:
		return float64(n)
	case uint16:
		return float64(n)
	case uint32:
		return float64(n)
	case uint64:
		return float64(n)
	case float32:
		return float64(n)
	case float64:
		return n
	}

	return 0
}

func toComplex128(v interface{}) complex128 {
	switch n := v.(type) {
	case complex64:
		return complex128(n)
	case complex128:
		return n
	}

	return 0
}
