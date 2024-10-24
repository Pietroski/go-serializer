package go_serializer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	TestData struct {
		FieldStr  string
		FieldInt  int8
		FieldBool bool

		FieldStrPtr    *string
		FieldIntPtr    *int
		FieldBoolPtr   *bool
		SubTestData    SubTestData
		SubTestDataPtr *SubTestData
		SliceTestData  SliceTestData
		MapTestData    MapTestData
	}

	SubTestData struct {
		FieldStr   string
		FieldInt32 int32
		FieldBool  bool
		FieldInt64 int64
		FieldInt   int
	}

	SliceTestData struct {
		IntList          []int
		IntIntList       [][]int
		ThreeDIntList    [][][]int
		StrList          []string
		StrStrList       [][]string
		StructList       []SliceItem
		PtrStructList    []*SliceItem
		PtrStructNilList []*SliceItem
	}

	MapTestData struct {
		Int64KeyMapInt64Value map[int64]int64
		StrKeyMapStrValue     map[string]string
	}

	SliceItem struct {
		Int  int
		Str  string
		Bool bool
	}
)

func TestBinarySerializer_Marshal(t *testing.T) {
	t.Run("success TestData", func(t *testing.T) {
		serializer := NewBinarySerializer()

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
		serializer := NewBinarySerializer()

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
		serializer := NewBinarySerializer()

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
		serializer := NewBinarySerializer()

		bs, err := serializer.Marshal("test-again#$çcçá")
		require.NoError(t, err)

		//t.Log(string(bs), bs)

		var str string
		err = serializer.Unmarshal(bs, &str)
		require.NoError(t, err)

		t.Log(str)
	})

	t.Run("success SliceTestData", func(t *testing.T) {
		serializer := NewBinarySerializer()

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

	t.Run("success MapTestData", func(t *testing.T) {
		serializer := NewBinarySerializer()

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
}
