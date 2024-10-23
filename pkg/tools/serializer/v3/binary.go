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
	BinarySerializer struct {
		byteOrder binary.ByteOrder
	}
)

func NewBinarySerializer() *BinarySerializer {
	bs := &BinarySerializer{
		byteOrder: binary.LittleEndian,
	}

	return bs
}

func (s *BinarySerializer) Serialize(data interface{}) ([]byte, error) {
	bbf := bytes.NewBuffer(make([]byte, 0, 1<<8))

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
			if field.IsNil() {
				if err := binary.Write(bbf, binary.LittleEndian, byte(1)); err != nil {
					return nil, err
				}

				continue
			}

			if err := binary.Write(bbf, binary.LittleEndian, byte(0)); err != nil {
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

		if err := s.deserializePrimitive(bbf, &f); err != nil {
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
