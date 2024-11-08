package go_serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
)

func TestUnsafeBinarySerializer_Marshal(t *testing.T) {
	t.Run("success TestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		intPtr := 7
		strPtr := "test-str-ptr"
		boolPtr := true
		testData := TestData{
			FieldStr:     "test-data",
			FieldInt:     127,
			FieldBool:    true,
			FieldIntPtr:  &intPtr,
			FieldStrPtr:  &strPtr,
			FieldBoolPtr: &boolPtr,

			SubTestData: SubTestData{
				FieldStr:   "test-sub-data",
				FieldInt32: 127567,
				FieldBool:  false,
				FieldInt64: math.MaxInt64,
				FieldInt:   0,
			},

			SubTestDataPtr: &SubTestData{
				FieldStr:   "test-sub-data-ptr",
				FieldInt32: 765432,
				FieldBool:  true,
				FieldInt64: math.MaxInt16,
				FieldInt:   5432,
			},
		}

		bs, err := serializer.Marshal(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Unmarshal(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		t.Log(*td.FieldStrPtr)
		t.Log(*td.SubTestDataPtr)
	})

	t.Run("success TestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		strPtr := "test-str-ptr"
		testData := TestData{
			FieldStr:  "test-data",
			FieldInt:  127,
			FieldBool: true,

			FieldStrPtr: &strPtr,

			SubTestData: SubTestData{
				FieldStr:   "test-sub-data",
				FieldInt32: 127567,
				FieldBool:  false,
				FieldInt64: math.MaxInt64,
				FieldInt:   0,
			},
		}
		bs, err := serializer.Marshal(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Unmarshal(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		t.Log(*td.FieldStrPtr)
	})

	t.Run("success TestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		testData := TestData{
			FieldStr:  "test-data",
			FieldInt:  127,
			FieldBool: true,

			SubTestData: SubTestData{
				FieldStr:   "test-sub-data",
				FieldInt32: 127567,
				FieldBool:  false,
				FieldInt64: math.MaxInt64,
				FieldInt:   0,
			},
		}
		bs, err := serializer.Marshal(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Unmarshal(bs, &td)
		require.NoError(t, err)

		t.Log(td)
	})

	t.Run("success", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		bs, err := serializer.Marshal("test-again#$çcçá")
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var str string
		err = serializer.Unmarshal(bs, &str)
		require.NoError(t, err)

		t.Log(str)
	})

	t.Run("success SliceTestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		testData := TestData{
			FieldStr:  "test-data",
			FieldInt:  127,
			FieldBool: true,

			SliceTestData: SliceTestData{
				IntList: []int{1, 2, 3, 7, 11, 9, 19, 4},
				IntIntList: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				ThreeDIntList: [][][]int{
					{
						{1, 4, 7},
						{2, 5, 7, 9},
						{8, 7, 6},
					},
					{
						{11, 24, 57},
						{82, 75, 77, 99},
						{58, 47, 36},
					},
				},
				StrList: []string{
					"item-1", "item-2", "item-3", "item-4", "item-5", "item-6", "item-7",
				},
				StrStrList: [][]string{
					{"item-1", "item-2", "item-3", "item-4", "item-5", "item-6"},
					{"another-item-1", "another-item-2", "another-item-3"},
				},
				StructList: []SliceItem{
					{
						Int:  100,
						Str:  "any-item",
						Bool: true,
					},
					{
						Int:  99,
						Str:  "any-other-item",
						Bool: false,
					},
				},
				PtrStructList: []*SliceItem{
					{
						Int:  55,
						Str:  "any-ptr-item",
						Bool: true,
					},
					{
						Int:  56,
						Str:  "any-other-ptr-item",
						Bool: false,
					},
				},
				PtrStructNilList: []*SliceItem{
					nil, nil, nil, {
						Int:  56,
						Str:  "any-other-ptr-item",
						Bool: false,
					},
				},
			},
		}
		bs, err := serializer.Serialize(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Deserialize(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		for _, psl := range td.SliceTestData.PtrStructList {
			if psl != nil {
				t.Log(*psl)
				continue
			}

			t.Log(psl)
		}
		for _, psl := range td.SliceTestData.PtrStructNilList {
			if psl != nil {
				t.Log(*psl)
				continue
			}

			t.Log(psl)
		}
	})

	t.Run("success SliceTestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		testData := TestData{
			SliceTestData: SliceTestData{
				StrList: []string{
					"item-1", "item-2", "item-3", "item-4", "item-5", "item-6", "item-7",
				},
				StrStrList: [][]string{
					{"item-1", "item-2", "item-3", "item-4", "item-5", "item-6"},
					{"another-item-1", "another-item-2", "another-item-3"},
				},
			},
		}
		bs, err := serializer.Serialize(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Deserialize(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		for _, psl := range td.SliceTestData.PtrStructList {
			if psl != nil {
				t.Log(*psl)
				continue
			}

			t.Log(psl)
		}
		for _, psl := range td.SliceTestData.PtrStructNilList {
			if psl != nil {
				t.Log(*psl)
				continue
			}

			t.Log(psl)
		}
	})

	t.Run("success MapTestData", func(t *testing.T) {
		serializer := NewRawBinarySerializer()

		testData := TestData{
			FieldStr:  "test-data",
			FieldInt:  127,
			FieldBool: true,

			MapTestData: MapTestData{
				Int64KeyMapInt64Value: map[int64]int64{
					0: 4,
					2: 8,
					5: 7,
				},
				StrKeyMapStrValue: map[string]string{
					"key":         "value",
					"another-key": "another-value",
				},
			},
		}
		bs, err := serializer.Serialize(&testData)
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var td TestData
		err = serializer.Deserialize(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		for key, value := range td.MapTestData.Int64KeyMapInt64Value {
			t.Logf("%v: %v\n", key, value)
		}
		for key, value := range td.MapTestData.StrKeyMapStrValue {
			t.Logf("%v: %v\n", key, value)
		}
	})

	t.Run("Test Benchmark Data", func(t *testing.T) {
		Test_UnsafeBinary_Benchmark_Data(t)
	})
}

func Test_UnsafeBinary_Benchmark_Data(t *testing.T) {
	t.Run("main benchmark test data", func(t *testing.T) {
		msg := item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := NewRawBinarySerializer()

		bs, err := serializer.Serialize(&msg)
		require.NoError(t, err)

		var td item_models.Item
		err = serializer.Deserialize(bs, &td)
		require.NoError(t, err)

		t.Log(td)
		t.Log(td.SubItem)
	})

	t.Run("slice serialization", func(t *testing.T) {
		t.Run("slice of int", func(t *testing.T) {
			msg := &ProtoTypeSliceTestData{
				IntList:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				UintList: []uint64{math.MaxInt64, 2, 3, 4, 5, 6, 7, 8, 9, math.MaxUint64},
				StrList:  []string{"first-item", "second-item", "third-item", "fourth-item"},
				BytesBytesList: [][]byte{
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{},
				},
				BytesList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
			}
			serializer := NewRawBinarySerializer()

			var target ProtoTypeSliceTestData
			bs, err := serializer.Serialize(msg)
			require.NoError(t, err)
			err = serializer.Deserialize(bs, &target)
			require.NoError(t, err)
			t.Log(target)
		})
	})

	t.Run("success MapTestData", func(t *testing.T) {
		t.Run("map of int to int", func(t *testing.T) {
			serializer := NewRawBinarySerializer()

			msg := MapTestData{
				Int64KeyMapInt64Value: map[int64]int64{
					0:     100,
					7:     2,
					2:     8,
					8:     4,
					4:     16,
					100:   200,
					1_000: math.MaxInt64,
				},
			}
			bs, err := serializer.Serialize(&msg)
			require.NoError(t, err)

			//t.Log(string(bs), bs)

			var td MapTestData
			err = serializer.Deserialize(bs, &td)
			require.NoError(t, err)

			t.Log(td)
			for key, value := range td.Int64KeyMapInt64Value {
				t.Logf("%v: %v\n", key, value)
			}
			for key, value := range td.StrKeyMapStrValue {
				t.Logf("%v: %v\n", key, value)
			}
		})

		t.Run("map of string to string", func(t *testing.T) {
			serializer := NewRawBinarySerializer()

			msg := MapTestData{
				StrKeyMapStrValue: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
				},
			}
			bs, err := serializer.Serialize(&msg)
			require.NoError(t, err)

			//t.Log(string(bs), bs)

			var td MapTestData
			err = serializer.Deserialize(bs, &td)
			require.NoError(t, err)

			t.Log(td)
			for key, value := range td.Int64KeyMapInt64Value {
				t.Logf("%v: %v\n", key, value)
			}
			for key, value := range td.StrKeyMapStrValue {
				t.Logf("%v: %v\n", key, value)
			}
		})
	})
}
