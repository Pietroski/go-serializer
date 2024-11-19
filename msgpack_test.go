//go:build unit

package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
)

func TestMsgPackSerializer(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		msg := "test-again#$çcçá"

		s := NewMsgPackSerializer()

		bs, err := s.Serialize(msg)
		assert.NoError(t, err)
		assert.NotNil(t, bs)

		var target string
		err = s.Deserialize(bs, &target)
		assert.NoError(t, err)
		assert.Equal(t, msg, target)
		t.Log(target)

		t.Run("DataRebind", func(t *testing.T) {
			s = NewMsgPackSerializer()
			err = s.DataRebind(msg, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
		})
	})

	t.Run("number", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			testCases := map[string]struct {
				msg int
			}{
				"max": {
					msg: math.MaxInt64,
				},
				"min": {
					msg: -math.MaxInt64,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target int
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("int8", func(t *testing.T) {
			testCases := map[string]struct {
				msg int8
			}{
				"max": {
					msg: math.MaxInt8,
				},
				"min": {
					msg: -math.MaxInt8,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target int8
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("int16", func(t *testing.T) {
			testCases := map[string]struct {
				msg int16
			}{
				"max": {
					msg: math.MaxInt16,
				},
				"min": {
					msg: -math.MaxInt16,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target int16
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("int32", func(t *testing.T) {
			testCases := map[string]struct {
				msg int32
			}{
				"max": {
					msg: math.MaxInt32,
				},
				"min": {
					msg: -math.MaxInt32,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target int32
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("int64", func(t *testing.T) {
			testCases := map[string]struct {
				msg int64
			}{
				"max": {
					msg: math.MaxInt64,
				},
				"min": {
					msg: -math.MaxInt64,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target int64
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("uint", func(t *testing.T) {
			testCases := map[string]struct {
				msg uint
			}{
				"max": {
					msg: math.MaxUint64,
				},
				"min": {
					msg: -0,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target uint
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("uint8", func(t *testing.T) {
			testCases := map[string]struct {
				msg uint8
			}{
				"max": {
					msg: math.MaxUint8,
				},
				"min": {
					msg: -0,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target uint8
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("uint16", func(t *testing.T) {
			testCases := map[string]struct {
				msg uint16
			}{
				"max": {
					msg: math.MaxUint16,
				},
				"min": {
					msg: -0,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target uint16
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("uint32", func(t *testing.T) {
			testCases := map[string]struct {
				msg uint32
			}{
				"max": {
					msg: math.MaxUint32,
				},
				"min": {
					msg: -0,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target uint32
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("uint64", func(t *testing.T) {
			testCases := map[string]struct {
				msg uint64
			}{
				"max": {
					msg: math.MaxUint64,
				},
				"min": {
					msg: -0,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target uint64
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("float32", func(t *testing.T) {
			testCases := map[string]struct {
				msg float32
			}{
				"max": {
					msg: math.MaxFloat32,
				},
				"max-smaller": {
					msg: 3.4028234663852886e38,
				},
				"max-bigger": {
					msg: 3.40282346638528859811704183484516925440e+38,
				},
				"min": {
					msg: -math.MaxFloat32,
				},
				"min-smaller": {
					msg: -3.4028234663852886e38,
				},
				"min-bigger": {
					msg: -3.40282346638528859811704183484516925440e+38,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target float32
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})

		t.Run("float64", func(t *testing.T) {
			testCases := map[string]struct {
				msg float64
			}{
				"max": {
					msg: math.MaxFloat64,
				},
				"max-smaller": {
					msg: 1.7976931348623157e308,
				},
				"max-bigger": {
					msg: 1.79769313486231570814527423731704356798070e+308,
				},
				"min": {
					msg: -math.MaxFloat64,
				},
				"min-smaller": {
					msg: -1.7976931348623157e308,
				},
				"min-bigger": {
					msg: -1.79769313486231570814527423731704356798070e+308,
				},
				"zero": {
					msg: 0,
				},
			}

			for testName, testCase := range testCases {
				t.Run(testName, func(t *testing.T) {
					msg := testCase.msg

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target float64
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		t.Run("default benchmark", func(t *testing.T) {
			msg := &testmodels.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
				SubItem: &testmodels.SubItem{
					Date:     time.Now().Unix(),
					Amount:   1_000_000_000,
					ItemCode: "code-status",
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.Item
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, &target)
			t.Log(target)
			t.Log(target.SubItem)
		})

		t.Run("default benchmark - nil sub item", func(t *testing.T) {
			msg := &testmodels.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.Item
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, &target)
			t.Log(target)
			t.Log(target.SubItem)
		})

		t.Run("simplified special struct test data", func(t *testing.T) {
			msg := &testmodels.SimplifiedSpecialStructTestData{
				Bool:    true,
				String:  "any-string",
				Int32:   math.MaxInt32,
				Int64:   math.MaxInt64,
				Uint32:  math.MaxUint32,
				Uint64:  math.MaxUint64,
				Float32: math.MaxFloat32,
				Float64: math.MaxFloat64,
				Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
				RepeatedBytes: [][]byte{
					{-0, 0, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, math.MaxInt8, 255, 0, -0},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.SimplifiedSpecialStructTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, *msg, target)
			t.Log(target)
		})

		t.Run("string struct only", func(t *testing.T) {
			msg := &testmodels.StringStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.StringStruct
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, &target)
			t.Log(target)
		})

		t.Run("int64 struct only", func(t *testing.T) {
			msg := &testmodels.Int64Struct{
				FirstInt64:  math.MaxInt64,
				SecondInt64: -math.MaxInt64,
				ThirdInt64:  math.MaxInt64,
				FourthInt64: -math.MaxInt64,
				FifthInt64:  0,
				SixthInt64:  -0,
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.Int64Struct
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, &target)
			t.Log(target)
		})
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("[]int", func(t *testing.T) {
			msg := &testmodels.IntSliceTestData{
				IntList: []int{-math.MaxInt64, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, math.MaxInt64},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.IntSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]int8", func(t *testing.T) {
			msg := &testmodels.Int8SliceTestData{
				Int8List: []int8{-math.MaxInt8, -128, 0, 4, 5, 100, 8, 127, math.MaxInt8},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Int8SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]int16", func(t *testing.T) {
			msg := &testmodels.Int16SliceTestData{
				Int16List: []int16{-math.MaxInt16, -32768, -128, 0, 4, 5, 100, 8, 127, 32767, math.MaxInt16},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Int16SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]int32", func(t *testing.T) {
			msg := &testmodels.Int32SliceTestData{
				Int32List: []int32{
					-math.MaxInt32, -2147483648, -32768, -128, 0, 4, 5, 100, 8, 127, 32767, 2147483647, math.MaxInt32,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Int32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]int64", func(t *testing.T) {
			msg := &testmodels.Int64SliceTestData{
				Int64List: []int64{
					-math.MaxInt64, -9223372036854775808, -0, 0, 2, 12345678, 4, 5, 5170, 10, 8,
					87654321, 9223372036854775807, math.MaxInt64,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Int64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint", func(t *testing.T) {
			msg := &testmodels.UintSliceTestData{
				UintList: []uint{math.MaxUint64, math.MaxInt64, -0, 0, 2, 12345678, 4, 5, 5170, 10, 8,
					87654321, 18446744073709551615, math.MaxUint64,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.UintSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint8", func(t *testing.T) {
			msg := &testmodels.Uint8SliceTestData{
				Uint8List: []uint8{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Uint8SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint16", func(t *testing.T) {
			msg := &testmodels.Uint16SliceTestData{
				Uint16List: []uint16{math.MaxUint16, -0, 0, 4, 5, 8, 127, 32767, 65535, math.MaxInt16, math.MaxUint16},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Uint16SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint32", func(t *testing.T) {
			msg := &testmodels.Uint32SliceTestData{
				Uint32List: []uint32{
					-0, 0, 4, 5, 100, 8, 127, 32767, 2147483647, 4294967295, math.MaxInt32, math.MaxUint32,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Uint32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]uint64", func(t *testing.T) {
			msg := &testmodels.Uint64SliceTestData{
				Uint64List: []uint64{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
					9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Uint64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]float32", func(t *testing.T) {
			msg := &testmodels.Float32SliceTestData{
				Float32List: []float32{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, 9223372036854775807, 18446744073709551615,
					3.40282346638528859811704183484516925440e+38, 3.4028234663852886e38, 0.00000001,
					-math.MaxFloat32, math.MaxFloat32,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Float32SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]float64", func(t *testing.T) {
			msg := &testmodels.Float64SliceTestData{
				Float64List: []float64{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, 9223372036854775807, 18446744073709551615,
					3.40282346638528859811704183484516925440e+38, 3.4028234663852886e38, 0.00000001,
					1.7976931348623157e308, -1.7976931348623157e308,
					1.79769313486231570814527423731704356798070e+308,
					-1.79769313486231570814527423731704356798070e+308,
					-math.MaxFloat32, math.MaxFloat32, -math.MaxFloat64, math.MaxFloat64,
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.Float64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]byte", func(t *testing.T) {
			msg := &testmodels.ByteSliceTestData{
				ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.ByteSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[][]byte", func(t *testing.T) {
			msg := &testmodels.ByteByteSliceTestData{
				ByteByteList: [][]byte{
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.ByteByteSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]string", func(t *testing.T) {
			msg := &testmodels.StringSliceTestData{
				StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.StringSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]SimplifiedSpecialStructTestData", func(t *testing.T) {
			msg := &testmodels.SimplifiedSpecialStructSliceTestData{
				SimplifiedSpecialStructSliceTestData: []testmodels.SimplifiedSpecialStructTestData{
					{
						Bool:    true,
						String:  "any-string",
						Int32:   math.MaxInt32,
						Int64:   math.MaxInt64,
						Uint32:  math.MaxUint32,
						Uint64:  math.MaxUint64,
						Float32: math.MaxFloat32,
						Float64: math.MaxFloat64,
						Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
						RepeatedBytes: [][]byte{
							{-0, 0, 255, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, 255, 0, -0},
						},
					},
					{
						Bool:    false,
						String:  "any-other-string",
						Int32:   -math.MaxInt32,
						Int64:   -math.MaxInt64,
						Uint32:  math.MaxUint32,
						Uint64:  math.MaxUint64,
						Float32: -math.MaxFloat32,
						Float64: -math.MaxFloat64,
						Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
						RepeatedBytes: [][]byte{
							{-0, 0, 255, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, 255, 0, -0},
						},
					},
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.SimplifiedSpecialStructSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("[]*SimplifiedSpecialStructTestData", func(t *testing.T) {
			msg := &testmodels.SimplifiedSpecialStructPointerSliceTestData{
				SimplifiedSpecialStructPointerSliceTestData: []*testmodels.SimplifiedSpecialStructTestData{
					{
						Bool:    true,
						String:  "any-string",
						Int32:   math.MaxInt32,
						Int64:   math.MaxInt64,
						Uint32:  math.MaxUint32,
						Uint64:  math.MaxUint64,
						Float32: math.MaxFloat32,
						Float64: math.MaxFloat64,
						Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
						RepeatedBytes: [][]byte{
							{-0, 0, 255, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, 255, 0, -0},
						},
					},
					{
						Bool:    false,
						String:  "any-other-string",
						Int32:   -math.MaxInt32,
						Int64:   -math.MaxInt64,
						Uint32:  math.MaxUint32,
						Uint64:  math.MaxUint64,
						Float32: -math.MaxFloat32,
						Float64: -math.MaxFloat64,
						Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
						RepeatedBytes: [][]byte{
							{-0, 0, 255, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, math.MaxUint8},
							{math.MaxUint8, math.MaxInt8, 255, 0, -0},
						},
					},
				},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.SimplifiedSpecialStructPointerSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("various types together", func(t *testing.T) {
			msg := &testmodels.ProtoTypeSliceTestData{
				IntList: []int64{
					-math.MaxInt64, -math.MaxInt8, -math.MaxInt16, -math.MaxInt32,
					-0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10,
					math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64,
				},
				UintList: []uint64{
					-0, 0, 2, 4, 5, 7, 8, 10, math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64,
				},
				StrList: []string{
					"first-item", "second-item", "third-item", "fourth-item",
				},
				BytesBytesList: [][]byte{
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					nil,
				},
				BytesList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.ProtoTypeSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)

			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			assert.Equal(t, msg, &target)

			t.Log(target)
			t.Log(target.IntList)
			t.Log(target.UintList)
			t.Log(target.StrList)
			t.Log(target.BytesBytesList)
			t.Log(target.BytesList)
		})

		t.Run("various types and dimensions together", func(t *testing.T) {
			msg := &testmodels.SliceTestData{
				IntList: []int{
					-math.MaxInt64, -math.MaxInt8, -math.MaxInt16, -math.MaxInt32,
					-0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10,
					math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64,
				},
				IntIntList: [][]int{
					{
						-math.MaxInt64, -math.MaxInt8, -math.MaxInt16, -math.MaxInt32,
						-0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10,
						math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64,
					},
					{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
				},
				ThreeDIntList: [][][]int{
					{
						{
							-math.MaxInt64, -math.MaxInt8, -math.MaxInt16, -math.MaxInt32,
							-0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10,
							math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64,
						},
						{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
						{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
					},
					{
						{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
						{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
					},
					{
						{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
						{222, 7, 5, 88, 77, 2357},
					},
				},
				StrList: []string{
					"first-item", "second-item", "third-item", "fourth-item",
				},
				StrStrList: [][]string{
					{"first-item", "second-item", "third-item", "fourth-item", "fifth-item"},
					{"unordered item list", "fifth-item", "fourth-item", "first-item", "second-item", "third-item"},
				},
				StructList: []testmodels.SliceItem{
					{
						Int:  100,
						Str:  "any string",
						Bool: true,
					},
					{
						Int:  500,
						Str:  "any other string",
						Bool: false,
					},
					{
						Int: 700,
						Str: "another any other string",
					},
				},
				PtrStructList: []*testmodels.SliceItem{
					{
						Int:  100,
						Str:  "any string",
						Bool: true,
					},
					{
						Int:  500,
						Str:  "any other string",
						Bool: false,
					},
					{
						Int: 700,
						Str: "another any other string",
					},
				},
				PtrStructNilList: []*testmodels.SliceItem{nil, nil, nil, nil, nil},
			}
			serializer := NewMsgPackSerializer()

			var target testmodels.SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)

			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			assert.Equal(t, msg, &target)

			t.Log(target)
			t.Log(target.IntList)
			t.Log(target.IntIntList)
			t.Log(target.ThreeDIntList)
			t.Log(target.StrList)
			t.Log(target.StrStrList)
			t.Log(target.StructList)
			t.Log(target.PtrStructList)
			for _, ptrStruct := range target.PtrStructList {
				t.Log(ptrStruct)
			}
			t.Log(target.PtrStructNilList)
		})
	})

	t.Run("map", func(t *testing.T) {
		t.Run("map[string]string", func(t *testing.T) {
			msg := testmodels.MapStringStringTestData{
				MapStringString: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
					"another-key":   "another-value",
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapStringStringTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int]int", func(t *testing.T) {
			msg := testmodels.MapIntIntTestData{
				MapIntInt: map[int]int{
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

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapIntIntTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int64]int64", func(t *testing.T) {
			msg := testmodels.MapInt64Int64TestData{
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

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapInt64Int64TestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int]StructTestData", func(t *testing.T) {
			msg := testmodels.MapIntStructTestData{
				MapIntStruct: map[int]testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapIntStructTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int]MapIntStructPointerTestData", func(t *testing.T) {
			msg := testmodels.MapIntStructPointerTestData{
				MapIntStructPointer: map[int]*testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapIntStructPointerTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int64]MapInt64StructTestData", func(t *testing.T) {
			msg := testmodels.MapInt64StructTestData{
				MapInt64Struct: map[int64]testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapInt64StructTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[int64]MapInt64StructPointerTestData", func(t *testing.T) {
			msg := testmodels.MapInt64StructPointerTestData{
				MapInt64StructPointer: map[int64]*testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapInt64StructPointerTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[string]StructTestData", func(t *testing.T) {
			msg := testmodels.MapStringStructTestData{
				MapStringStruct: map[string]testmodels.StructTestData{
					"any-key": {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					"any-other-key": {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					"another-key": {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapStringStructTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map[string]StructPointerTestData", func(t *testing.T) {
			msg := testmodels.MapStringStructPointerTestData{
				MapStringStructPointer: map[string]*testmodels.StructTestData{
					"any-key": {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					"any-other-key": {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					"another-key": {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := NewMsgPackSerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.MapStringStructPointerTestData
			err = s.Deserialize(bs, &target)
			assert.NoError(t, err)
			assert.Equal(t, msg, target)
			t.Log(target)
		})

		t.Run("map only", func(t *testing.T) {
			t.Run("map[string]string", func(t *testing.T) {
				msg := map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
					"another-key":   "another-value",
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[string]string
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int]int", func(t *testing.T) {
				msg := map[int]int{
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
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int]int
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int64]int64", func(t *testing.T) {
				msg := map[int64]int64{
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
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int64]int64
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int]StructTestData", func(t *testing.T) {
				msg := map[int]testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int]testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int]StructPointerTestData", func(t *testing.T) {
				msg := map[int]*testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int]*testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int64]StructTestData", func(t *testing.T) {
				msg := map[int64]testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int64]testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[int64]StructPointerTestData", func(t *testing.T) {
				msg := map[int64]*testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[int64]*testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[string]StructTestData", func(t *testing.T) {
				msg := map[string]testmodels.StructTestData{
					"any-key": {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					"any-other-key": {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					"another-key": {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[string]testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("map[string]StructPointerTestData", func(t *testing.T) {
				msg := map[string]*testmodels.StructTestData{
					"any-key": {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					"any-other-key": {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					"another-key": {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				}

				s := NewMsgPackSerializer()

				bs, err := s.Serialize(msg)
				assert.NoError(t, err)
				assert.NotNil(t, bs)

				var target map[string]*testmodels.StructTestData
				err = s.Deserialize(bs, &target)
				assert.NoError(t, err)
				assert.Equal(t, msg, target)
				t.Log(target)
			})

			t.Run("empty map", func(t *testing.T) {
				t.Run("map[int]int", func(t *testing.T) {
					var msg map[int]int

					s := NewMsgPackSerializer()

					bs, err := s.Serialize(msg)
					assert.NoError(t, err)
					assert.NotNil(t, bs)

					var target map[int]int
					err = s.Deserialize(bs, &target)
					assert.NoError(t, err)
					assert.Equal(t, msg, target)
					t.Log(target)
				})
			})
		})
	})
}
