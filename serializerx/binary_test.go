package serializerx

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
		t.Run("default benchmark data", func(t *testing.T) {
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
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("[]int64", func(t *testing.T) {
			msg := &testmodels.Int64SliceTestData{
				Int64List: []int64{-math.MaxInt64, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321, math.MaxInt64},
			}
			serializer := NewBinarySerializer()

			var target testmodels.Int64SliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})

		t.Run("various types", func(t *testing.T) {
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
					{},
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
	})

	t.Run("map", func(t *testing.T) {
		//
	})
}
