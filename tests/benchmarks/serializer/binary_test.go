package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
)

type (
	SliceTestData struct {
		IntList          []int        `json:"int_list,omitempty"`
		IntIntList       [][]int      `json:"int_int_list,omitempty"`
		ThreeDIntList    [][][]int    `json:"three_d_int_list,omitempty"`
		StrList          []string     `json:"str_list,omitempty"`
		StrStrList       [][]string   `json:"str_str_list,omitempty"`
		StructList       []SliceItem  `json:"struct_list,omitempty"`
		PtrStructList    []*SliceItem `json:"ptr_struct_list,omitempty"`
		PtrStructNilList []*SliceItem `json:"ptr_struct_nil_list,omitempty"`
	}

	SliceItem struct {
		Int  int    `json:"int,omitempty"`
		Str  string `json:"str,omitempty"`
		Bool bool   `json:"bool,omitempty"`
	}
)

func Benchmark_BinarySerializer(b *testing.B) {
	b.Run("binary serialization", func(b *testing.B) {
		msg := &item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := go_serializer.NewBinarySerializer()
		for i := 0; i < b.N; i++ {
			_, err := serializer.Serialize(msg)
			require.NoError(b, err)
		}
	})

	b.Run("binary deserialization", func(b *testing.B) {
		msg := &item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := go_serializer.NewBinarySerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(b, err)

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err := serializer.Deserialize(bs, &target)
			require.NoError(b, err)
		}
		validateStructMsgAndTarget(b, msg, &target)
	})

	b.Run("binary serialization and deserialization", func(b *testing.B) {
		msg := &item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := go_serializer.NewBinarySerializer()

		var bs []byte
		var err error
		for i := 0; i < b.N; i++ {
			bs, err = serializer.Serialize(msg)
			require.NoError(b, err)
		}

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err := serializer.Deserialize(bs, &target)
			require.NoError(b, err)
		}
		validateStructMsgAndTarget(b, msg, &target)
	})

	b.Run("binary serialization and deserialization", func(b *testing.B) {
		msg := &item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := go_serializer.NewBinarySerializer()

		for i := 0; i < b.N; i++ {
			bs, err := serializer.Serialize(msg)
			require.NoError(b, err)

			var target item_models.Item
			err = serializer.Deserialize(bs, &target)
			require.NoError(b, err)
			validateStructMsgAndTarget(b, msg, &target)
		}
	})

	b.Run("binary serialization and deserialization - clean - no validation", func(b *testing.B) {
		msg := &item_models.Item{
			Id:     "any-item",
			ItemId: 100,
			Number: 5_000_000_000,
			SubItem: &item_models.SubItem{
				Date:     time.Now().Unix(),
				Amount:   1_000_000_000,
				ItemCode: "code-status",
			},
		}
		serializer := go_serializer.NewBinarySerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
		}
	})
}

func BenchmarkType_BinarySerializer(b *testing.B) {
	b.Run("slice serialization", func(b *testing.B) {
		b.Run("int slice", func(b *testing.B) {
			msg := SliceTestData{
				IntList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			serializer := go_serializer.NewBinarySerializer()

			var target SliceTestData
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)

			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
				_ = serializer.Deserialize(bs, &target)
			}
		})
	})
}
