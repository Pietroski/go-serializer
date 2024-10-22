package go_serializer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
)

type (
	BinarySerializer struct {
		stringSerializer  stringSerializer
		booleanSerializer boolSerializer
		//numericSerializer *numericSerializerTypes
		ptrSerializer ptrSerializer
	}

	//numericSerializerTypes struct {
	//	is     numericSerializer[int64]
	//	i8s    numericSerializer[int8]
	//	i16s   numericSerializer[int16]
	//	i32s   numericSerializer[int32]
	//	i64s   numericSerializer[int64]
	//	uis    numericSerializer[uint64]
	//	ui8s   numericSerializer[uint8]
	//	ui16s  numericSerializer[uint16]
	//	ui32s  numericSerializer[uint32]
	//	ui64s  numericSerializer[uint64]
	//	f32s   numericSerializer[float32]
	//	f64s   numericSerializer[float64]
	//	c64s   numericSerializer[complex64]
	//	c128s  numericSerializer[complex128]
	//	uiptrs numericSerializer[uintptr]
	//}
)

func NewBinarySerializer() *BinarySerializer {
	bs := &BinarySerializer{
		stringSerializer:  &binaryRuneSerializer{},
		booleanSerializer: &binaryBoolSerializer{},
		//numericSerializer: &numericSerializerTypes{
		//	is:     &binaryNumericSerializer[int64]{},
		//	i8s:    &binaryNumericSerializer[int8]{},
		//	i16s:   &binaryNumericSerializer[int16]{},
		//	i32s:   &binaryNumericSerializer[int32]{},
		//	i64s:   &binaryNumericSerializer[int64]{},
		//	uis:    &binaryNumericSerializer[uint64]{},
		//	ui8s:   &binaryNumericSerializer[uint8]{},
		//	ui16s:  &binaryNumericSerializer[uint16]{},
		//	ui32s:  &binaryNumericSerializer[uint32]{},
		//	ui64s:  &binaryNumericSerializer[uint64]{},
		//	f32s:   &binaryNumericSerializer[float32]{},
		//	f64s:   &binaryNumericSerializer[float64]{},
		//	c64s:   &binaryNumericSerializer[complex64]{},
		//	c128s:  &binaryNumericSerializer[complex128]{},
		//	uiptrs: &binaryNumericSerializer[uintptr]{},
		//},
		ptrSerializer: &binaryPtrSerializer{},
	}

	return bs
}

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	return s.Marshal(data)
}

func (s *BinarySerializer) Deserialize(data []byte, target interface{}) error {
	return s.Unmarshal(data, target)
}

func (s *BinarySerializer) Marshal(data interface{}) ([]byte, error) {
	return s.encode(data)
}

func (s *BinarySerializer) Unmarshal(data []byte, target interface{}) error {
	bbf := bytes.NewBuffer(data)
	return s.decode(bbf, target)
}

func (s *BinarySerializer) encode(data interface{}) ([]byte, error) {
	if isPrimitive(data) {
		eBs, err := s.serializePrimitive(data)
		if err != nil {
			return nil, err
		}

		return eBs, nil
	}

	var bs []byte
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
				eBs, err := s.ptrSerializer.preEncode(true)
				if err != nil {
					return nil, err
				}

				bs = append(bs, eBs...)
				continue
			}

			eBs, err := s.ptrSerializer.preEncode(false)
			if err != nil {
				return nil, err
			}

			bs = append(bs, eBs...)

			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			eBs, err := s.encode(field.Interface())
			if err != nil {
				return nil, err
			}

			bs = append(bs, eBs...)
		}

		eBs, err := s.serializePrimitive(field.Interface())
		if err != nil {
			return nil, err
		}

		bs = append(bs, eBs...)
	}

	return bs, nil
}

func (s *BinarySerializer) serializePrimitive(data interface{}) ([]byte, error) {
	var bs []byte
	var err error

	switch v := data.(type) {
	case bool:
		bs, err = s.booleanSerializer.encode(v)
	case string:
		bs, err = s.stringSerializer.encode(v)
	case int:
		bs, err = binary.Append(nil, binary.BigEndian, int64(v)) // s.numericSerializer.is.encode(int64(v))
	case int8:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.i8s.encode(v)
	case int16:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.i16s.encode(v)
	case int32:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.i32s.encode(v)
	case int64:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.i64s.encode(v)
	case uint:
		bs, err = binary.Append(nil, binary.BigEndian, uint64(v)) // s.numericSerializer.uis.encode(uint64(v))
	case uint8:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.ui8s.encode(v)
	case uint16:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.ui16s.encode(v)
	case uint32:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.ui32s.encode(v)
	case uint64:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.ui64s.encode(v)
	case float32:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.f32s.encode(v)
	case float64:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.f64s.encode(v)
	case complex64:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.c64s.encode(v)
	case complex128:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.c128s.encode(v)
	case uintptr:
		bs, err = binary.Append(nil, binary.BigEndian, v) // s.numericSerializer.uiptrs.encode(v)
	}
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (s *BinarySerializer) decode(bbf *bytes.Buffer, target interface{}) error {
	voft := reflect.ValueOf(target)
	if voft.Kind() == reflect.Ptr {
		voft = voft.Elem()
	}

	if isPrimitive(target) {
		if err := s.decodePrimitiveType(bbf, &voft); err != nil {
			return err
		}

		return nil
	}

	limit := voft.NumField()
	for idx := 0; idx < limit; idx++ {
		field := voft.Field(idx)
		if field.Kind() == reflect.Chan {
			return fmt.Errorf("invalid type %v", field.Kind())
		}

		if field.Kind() == reflect.Ptr {
			isNull, err := s.ptrSerializer.preDecode(bbf, &field)
			if err != nil {
				return err
			}
			if isNull {
				continue
			} else {
				field.Set(reflect.New(field.Type().Elem()))
			}

			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			err := s.structDecode(bbf, &field)
			if err != nil {
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
			err := s.structDecode(bbf, &f)
			if err != nil {
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
	if field.Kind() == reflect.String {
		var str string
		err := s.stringSerializer.decode(bbf, &str)
		if err != nil {
			return err
		}

		field.SetString(str)
		return nil
	}

	if field.Kind() == reflect.Bool {
		var b bool
		err := s.booleanSerializer.decode(bbf, &b)
		if err != nil {
			return err
		}

		field.SetBool(b)
		return nil
	}

	if err := s.numericDeserializer(bbf, field); err != nil {
		return err
	}

	return nil
}

func numericIntDecoder[N numericalInt](
	bbf *bytes.Buffer, field *reflect.Value, // , ds numericSerializer[N],
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i) //binary.Read(bbf, binary.BigEndian, target)
	//err := ds.decode(bbf, &i)
	if err != nil {
		return err
	}

	field.SetInt(toInt64(i))

	return nil
}

func numericUintDecoder[N numericalUint](
	bbf *bytes.Buffer, field *reflect.Value, // , ds numericSerializer[N],
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i) // err := ds.decode(bbf, &i)
	if err != nil {
		return err
	}

	field.SetUint(toUint64(i))

	return nil
}

func numericFloatDecoder[N numericFloat](
	bbf *bytes.Buffer, field *reflect.Value, // , ds numericSerializer[N],
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i) // err := ds.decode(bbf, &i)
	if err != nil {
		return err
	}

	field.SetFloat(toFloat64(i))

	return nil
}

func numericComplexDecoder[N numericComplex](
	bbf *bytes.Buffer, field *reflect.Value, // , ds numericSerializer[N],
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i) // err := ds.decode(bbf, &i)
	if err != nil {
		return err
	}

	field.SetComplex(toComplex128(i))

	return nil
}

func numericUintPtrDecoder[N numericUintPtr](
	bbf *bytes.Buffer, field *reflect.Value, // , ds numericSerializer[N],
) error {
	var i N
	err := binary.Read(bbf, binary.BigEndian, &i) // err := ds.decode(bbf, &i)
	if err != nil {
		return err
	}

	field.SetInt(toInt64(i))

	return nil
}

func (s *BinarySerializer) numericDeserializer(bbf *bytes.Buffer, field *reflect.Value) (err error) {
	if field.Kind() == reflect.Int {
		err = numericIntDecoder[int64](bbf, field) //, s.numericSerializer.is)
	}

	if field.Kind() == reflect.Int8 {
		err = numericIntDecoder[int8](bbf, field) //, s.numericSerializer.i8s)
	}

	if field.Kind() == reflect.Int16 {
		err = numericIntDecoder[int16](bbf, field) //, s.numericSerializer.i16s)
	}

	if field.Kind() == reflect.Int32 {
		err = numericIntDecoder[int32](bbf, field) //, s.numericSerializer.i32s)
	}

	if field.Kind() == reflect.Int64 {
		err = numericIntDecoder[int64](bbf, field) //, s.numericSerializer.i64s)
	}

	if field.Kind() == reflect.Uint {
		err = numericUintDecoder[uint64](bbf, field) //, s.numericSerializer.uis)
	}

	if field.Kind() == reflect.Uint8 {
		err = numericUintDecoder[uint8](bbf, field) //, s.numericSerializer.ui8s)
	}

	if field.Kind() == reflect.Uint16 {
		err = numericUintDecoder[uint16](bbf, field) //, s.numericSerializer.ui16s)
	}

	if field.Kind() == reflect.Uint32 {
		err = numericUintDecoder[uint32](bbf, field) //, s.numericSerializer.ui32s)
	}

	if field.Kind() == reflect.Uint64 {
		err = numericUintDecoder[uint64](bbf, field) //, s.numericSerializer.ui64s)
	}

	if field.Kind() == reflect.Float32 {
		err = numericFloatDecoder[float32](bbf, field) //, s.numericSerializer.f32s)
	}

	if field.Kind() == reflect.Float64 {
		err = numericFloatDecoder[float64](bbf, field) //, s.numericSerializer.f64s)
	}

	if field.Kind() == reflect.Complex64 {
		err = numericComplexDecoder[complex64](bbf, field) //, s.numericSerializer.c64s)
	}

	if field.Kind() == reflect.Complex128 {
		err = numericComplexDecoder[complex128](bbf, field) //, s.numericSerializer.c128s)
	}

	if field.Kind() == reflect.Uintptr {
		err = numericUintPtrDecoder[uintptr](bbf, field) //, s.numericSerializer.uiptrs)
	}

	if err != nil {
		return err
	}

	return nil
}

type (
	binaryBytesSerializer struct{} // stringToBytes | binaryFromBytesToString
	binaryRuneSerializer  struct{} // stringToRune | binaryFromRuneToString

	stringSerializer interface {
		encode(str string) ([]byte, error)
		decode(bbf *bytes.Buffer, target *string) error
	}
)

func (s *binaryBytesSerializer) encode(str string) ([]byte, error) {
	var bs []byte

	for _, c := range str {
		bs = append(bs, byte(c))
	}

	buf := make([]byte, 0)
	var err error
	buf, err = binary.Append(buf, binary.BigEndian, uint64(len(bs)))
	if err != nil {
		return nil, err
	}

	for _, b := range bs {
		buf, err = binary.Append(buf, binary.BigEndian, b)
		if err != nil {
			return buf, nil
		}
	}

	return buf, nil
}

func (s *binaryBytesSerializer) decode(bbf *bytes.Buffer, target interface{}) error {
	var length uint64
	err := binary.Read(bbf, binary.BigEndian, &length)
	if err != nil {
		return err
	}

	var bs []byte
	for i := uint64(0); i < length; i++ {
		var b byte
		err = binary.Read(bbf, binary.BigEndian, &b)
		if err != nil {
			return err
		}

		bs = append(bs, b)
	}

	var str string
	for _, b := range bs {
		str += string(b)
	}

	ts, ok := target.(*string)
	if !ok {
		return fmt.Errorf("target is not a string pointer")
	}
	*ts = str

	return nil
}

func (s *binaryRuneSerializer) encode(str string) ([]byte, error) {
	rs := []rune(str)

	buf := make([]byte, 0, cap(rs)+8)
	var err error
	buf, err = binary.Append(buf, binary.BigEndian, uint64(len(rs)))
	if err != nil {
		return nil, err
	}

	for _, r := range rs {
		buf, err = binary.Append(buf, binary.BigEndian, r)
		if err != nil {
			return buf, nil
		}
	}

	return buf, nil
}

func (s *binaryRuneSerializer) decode(bbf *bytes.Buffer, target *string) error {
	var length uint64
	err := binary.Read(bbf, binary.BigEndian, &length)
	if err != nil {
		return err
	}

	rs := make([]rune, length)
	for i := uint64(0); i < length; i++ {
		err = binary.Read(bbf, binary.BigEndian, &rs[i])
		if err != nil {
			return err
		}
	}

	strBuilder := &strings.Builder{}
	for _, r := range rs {
		strBuilder.WriteString(string(r))
	}

	*target = strBuilder.String()

	return nil
}

type (
	binaryBoolSerializer struct{}

	boolSerializer interface {
		encode(b bool) ([]byte, error)
		decode(bbf *bytes.Buffer, target *bool) error
	}
)

func (s *binaryBoolSerializer) encode(b bool) ([]byte, error) {
	num := uint8(0)
	if b {
		num = 1
	}

	var bs []byte
	bs, err := binary.Append(bs, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (s *binaryBoolSerializer) decode(bbf *bytes.Buffer, target *bool) error {
	var boolean bool

	err := binary.Read(bbf, binary.BigEndian, &boolean)
	if err != nil {
		return err
	}

	*target = boolean

	return nil
}

type (
	binaryPtrSerializer struct{}

	ptrSerializer interface {
		preEncode(isNull bool) ([]byte, error)
		preDecode(bbf *bytes.Buffer, field *reflect.Value) (isNull bool, err error)
	}
)

func (s *binaryPtrSerializer) preEncode(isNull bool) ([]byte, error) {
	var bs []byte
	var err error
	if isNull {
		bs, err = binary.Append(bs, binary.BigEndian, byte(1))
	} else {
		bs, err = binary.Append(bs, binary.BigEndian, byte(0))
	}
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (s *binaryPtrSerializer) preDecode(bbf *bytes.Buffer, field *reflect.Value) (isNull bool, err error) {
	var isItNull byte
	err = binary.Read(bbf, binary.BigEndian, &isItNull)
	if err != nil {
		return
	}

	if isItNull == 1 {
		field.Set(reflect.Zero(field.Type()))

		isNull = true
		return
	}

	return
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

	//numerical interface {
	//	//int | int8 | int16 | int32 | int64 |
	//	//	uint | uint8 | uint16 | uint32 | uint64 |
	//	//	float32 | float64 |
	//	//	complex64 | complex128 |
	//	//	uintptr
	//
	//	numericalInt | numericalUint | numericFloat | numericComplex | numericUintPtr
	//}

	//binaryNumericSerializer[N numerical] struct{}

	//numericSerializer[N numerical] interface {
	//	encode(b N) ([]byte, error)
	//	decode(bbf *bytes.Buffer, target *N) error
	//}
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

//func (s *binaryNumericSerializer[N]) encode(n N) ([]byte, error) {
//	return binary.Append(nil, binary.BigEndian, n)
//}
//
//func (s *binaryNumericSerializer[N]) decode(bbf *bytes.Buffer, target *N) error {
//	return binary.Read(bbf, binary.BigEndian, target)
//}
