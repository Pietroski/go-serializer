package serializer

import (
	grpc_item "gitlab.com/pietroski-software-company/devex/golang/serializer/internal/generated/go/pkg/item"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtoSerializer(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		t.Run("default benchmark", func(t *testing.T) {
			msg := &grpc_item.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
				SubItem: &grpc_item.SubItem{
					Date:     time.Now().Unix(),
					Amount:   1_000_000_000,
					ItemCode: "code-status",
				},
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.Item
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			//assert.Equal(t, *msg, target)
			assert.Equal(t, msg.Id, target.Id)
			assert.Equal(t, msg.ItemId, target.ItemId)
			assert.Equal(t, msg.Number, target.Number)
			assert.Equal(t, msg.SubItem.Date, target.SubItem.Date)
			assert.Equal(t, msg.SubItem.Amount, target.SubItem.Amount)
			assert.Equal(t, msg.SubItem.ItemCode, target.SubItem.ItemCode)
			t.Log(target)
			t.Log(target.SubItem)
		})

		t.Run("default benchmark - nil sub item", func(t *testing.T) {
			msg := &grpc_item.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.Item
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			//assert.Equal(t, *msg, target)
			assert.Equal(t, msg.Id, target.Id)
			assert.Equal(t, msg.ItemId, target.ItemId)
			assert.Equal(t, msg.Number, target.Number)
			assert.Nil(t, target.SubItem)
			t.Log(target)
			t.Log(target.SubItem)
		})

		t.Run("string struct only", func(t *testing.T) {
			msg := &grpc_item.StringStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.StringStruct
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg.FirstString, target.FirstString)
			assert.Equal(t, msg.SecondString, target.SecondString)
			assert.Equal(t, msg.ThirdString, target.ThirdString)
			assert.Equal(t, msg.FourthString, target.FourthString)
			assert.Equal(t, msg.FifthString, target.FifthString)
			t.Log(target)
		})

		t.Run("int64 struct only", func(t *testing.T) {
			msg := &grpc_item.Int64Struct{
				FirstInt64:  math.MaxInt64,
				SecondInt64: -math.MaxInt64,
				ThirdInt64:  math.MaxInt64,
				FourthInt64: -math.MaxInt64,
				FifthInt64:  0,
				SixthInt64:  -0,
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.Int64Struct
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg.FirstInt64, target.FirstInt64)
			assert.Equal(t, msg.SecondInt64, target.SecondInt64)
			assert.Equal(t, msg.ThirdInt64, target.ThirdInt64)
			assert.Equal(t, msg.FourthInt64, target.FourthInt64)
			assert.Equal(t, msg.FifthInt64, target.FifthInt64)
			assert.Equal(t, msg.SixthInt64, target.SixthInt64)
			t.Log(target)
		})
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("[]int32", func(t *testing.T) {
			msg := &grpc_item.Int32SliceTestData{
				Int32List: []int32{
					-math.MaxInt32, -2147483648, -32768, -128, 0, 4, 5, 100, 8, 127, 32767, 2147483647, math.MaxInt32,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Int32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]int64", func(t *testing.T) {
			msg := &grpc_item.Int64SliceTestData{
				Int64List: []int64{
					-math.MaxInt64, -9223372036854775808, -0, 0, 2, 12345678, 4, 5, 5170, 10, 8,
					87654321, 9223372036854775807, math.MaxInt64,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Int64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint32", func(t *testing.T) {
			msg := &grpc_item.Uint32SliceTestData{
				Uint32List: []uint32{
					-0, 0, 4, 5, 100, 8, 127, 32767, 2147483647, 4294967295, math.MaxInt32, math.MaxUint32,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Uint32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint64", func(t *testing.T) {
			msg := &grpc_item.Uint64SliceTestData{
				Uint64List: []uint64{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
					9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Uint64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]float32", func(t *testing.T) {
			msg := &grpc_item.Float32SliceTestData{
				Float32List: []float32{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, 9223372036854775807, 18446744073709551615,
					3.40282346638528859811704183484516925440e+38, 3.4028234663852886e38, 0.00000001,
					-math.MaxFloat32, math.MaxFloat32,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Float32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]float64", func(t *testing.T) {
			msg := &grpc_item.Float64SliceTestData{
				Float64List: []float64{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, 9223372036854775807, 18446744073709551615,
					3.40282346638528859811704183484516925440e+38, 3.4028234663852886e38, 0.00000001,
					1.7976931348623157e308, -1.7976931348623157e308,
					1.79769313486231570814527423731704356798070e+308,
					-1.79769313486231570814527423731704356798070e+308,
					-math.MaxFloat32, math.MaxFloat32, -math.MaxFloat64, math.MaxFloat64,
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.Float64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]byte", func(t *testing.T) {
			msg := &grpc_item.ByteSliceTestData{
				ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.ByteSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[][]byte", func(t *testing.T) {
			msg := &grpc_item.ByteByteSliceTestData{
				ByteByteList: [][]byte{
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
				},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.ByteByteSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]string", func(t *testing.T) {
			msg := &grpc_item.StringSliceTestData{
				StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}
			serializer := NewProtoSerializer()

			var target grpc_item.StringSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})
	})

	t.Run("map", func(t *testing.T) {
		t.Run("map[string]string", func(t *testing.T) {
			msg := &grpc_item.MapStringStringTestData{
				MapStringString: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
					"another-key":   "another-value",
				},
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.MapStringStringTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg.GetMapStringString(), target.GetMapStringString())
			t.Log(target)
		})

		t.Run("map[int64]int64", func(t *testing.T) {
			msg := &grpc_item.MapInt64Int64TestData{
				MapInt64Int64: map[int64]int64{
					0:              math.MaxInt64,
					1:              math.MaxInt8,
					2:              math.MaxInt16,
					3:              math.MaxInt32,
					4:              math.MaxInt64,
					math.MaxInt64:  0,
					math.MaxInt8:   1,
					math.MaxInt16:  2,
					math.MaxInt32:  3,
					-math.MaxInt64: 4,
				},
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.MapInt64Int64TestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg.GetMapInt64Int64(), target.GetMapInt64Int64())
			t.Log(target)
		})

		t.Run("map[int64]MapInt64StructPointerTestData", func(t *testing.T) {
			msg := &grpc_item.MapInt64StructPointerTestData{
				MapInt64StructPointerTestData: map[int64]*grpc_item.StructTestData{
					0: {
						Bool:  true,
						Str:   "any-string",
						Int64: math.MaxInt64,
					},
					2: {
						Bool:  false,
						Str:   "any-other-string",
						Int64: -math.MaxInt64,
					},
					4: {
						Bool:  false,
						Str:   "",
						Int64: 0,
					},
				},
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.MapInt64StructPointerTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			for key, value := range msg.MapInt64StructPointerTestData {
				assert.Equal(t, value.Int64, target.MapInt64StructPointerTestData[key].Int64)
				assert.Equal(t, value.Str, target.MapInt64StructPointerTestData[key].Str)
				assert.Equal(t, value.Bool, target.MapInt64StructPointerTestData[key].Bool)
			}
			for key, value := range msg.GetMapInt64StructPointerTestData() {
				assert.Equal(t, value.GetStr(), target.GetMapInt64StructPointerTestData()[key].GetStr())
				assert.Equal(t, value.GetBool(), target.GetMapInt64StructPointerTestData()[key].GetBool())
				assert.Equal(t, value.GetInt64(), target.GetMapInt64StructPointerTestData()[key].GetInt64())
			}
			t.Log(target)
		})

		t.Run("map[string]MapStringStructPointerTestData", func(t *testing.T) {
			msg := &grpc_item.MapStringStructPointerTestData{
				MapStringStructPointerTestData: map[string]*grpc_item.StructTestData{
					"any-key": {
						Bool:  true,
						Str:   "any-string",
						Int64: math.MaxInt64,
					},
					"any-other-key": {
						Bool:  false,
						Str:   "any-other-string",
						Int64: -math.MaxInt64,
					},
					"another-key": {
						Bool:  false,
						Str:   "",
						Int64: 0,
					},
				},
			}

			s := NewProtoSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target grpc_item.MapStringStructPointerTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			for key, value := range msg.MapStringStructPointerTestData {
				assert.Equal(t, value.Int64, target.MapStringStructPointerTestData[key].Int64)
				assert.Equal(t, value.Str, target.MapStringStructPointerTestData[key].Str)
				assert.Equal(t, value.Bool, target.MapStringStructPointerTestData[key].Bool)
			}
			for key, value := range msg.GetMapStringStructPointerTestData() {
				assert.Equal(t, value.GetStr(), target.GetMapStringStructPointerTestData()[key].GetStr())
				assert.Equal(t, value.GetBool(), target.GetMapStringStructPointerTestData()[key].GetBool())
				assert.Equal(t, value.GetInt64(), target.GetMapStringStructPointerTestData()[key].GetInt64())
			}
			t.Log(target)
		})
	})
}
