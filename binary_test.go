package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
)

func TestBinarySerializer(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		msg := "test-again#$çcçá"

		s := NewBinarySerializer()

		bs, err := s.Serialize(msg)
		assert.NoError(t, err)
		assert.NotNil(t, bs)

		var target string
		err = s.Deserialize(bs, &target)
		assert.NoError(t, err)
		assert.Equal(t, msg, target)
		t.Log(target)
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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

					s := NewBinarySerializer()

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

			s := NewBinarySerializer()

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

		t.Run("string struct only", func(t *testing.T) {
			msg := &testmodels.StringsStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			s := NewBinarySerializer()

			bs, err := s.Serialize(msg)
			assert.NoError(t, err)
			assert.NotNil(t, bs)

			var target testmodels.StringsStruct
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

			s := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

			var target testmodels.Uint64SliceTestData
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
			serializer := NewBinarySerializer()

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
			serializer := NewBinarySerializer()

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
		//
	})
}
