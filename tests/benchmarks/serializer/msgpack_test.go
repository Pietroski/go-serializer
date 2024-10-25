package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	grpc_item "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
)

func Benchmark_MsgPackSerializer(b *testing.B) {
	b.Run("json serialization", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()

		var err error
		for i := 0; i < b.N; i++ {
			_, err = serializer.Serialize(msg)
		}
		require.NoError(b, err)
	})

	b.Run("json deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(b, err)

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err = serializer.Deserialize(bs, &target)
		}
		require.NoError(b, err)
		validateStructMsgAndTarget(b, msg, &target)
	})

	b.Run("json serialization and deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()

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

	b.Run("json serialization and deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()

		for i := 0; i < b.N; i++ {
			bs, err := serializer.Serialize(msg)
			require.NoError(b, err)

			var target item_models.Item
			err = serializer.Deserialize(bs, &target)
			require.NoError(b, err)
			validateStructMsgAndTarget(b, msg, &target)
		}
	})

	b.Run("json serialization and deserialization - clean - no validation", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
		}
	})
}

func Benchmark_ProtoMsgPackSerializerSerializer(b *testing.B) {
	b.Run("proto serialization", func(b *testing.B) {
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
		serializer := go_serializer.NewMsgPackSerializer()
		for i := 0; i < b.N; i++ {
			_, err := serializer.Serialize(msg)
			require.NoError(b, err)
		}
	})
}

func BenchmarkType_MsgPackSerializer(b *testing.B) {
	b.Run("slice serialization", func(b *testing.B) {
		msg := SliceTestData{
			IntList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}
		serializer := go_serializer.NewMsgPackSerializer()

		b.Run("int slice - encoding", func(b *testing.B) {
			var bs []byte
			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
			}

			var target SliceTestData
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)
		})

		b.Run("int slice - decoding", func(b *testing.B) {
			bs, _ := serializer.Serialize(msg)

			var target SliceTestData
			for i := 0; i < b.N; i++ {
				_ = serializer.Deserialize(bs, &target)
			}
			b.Log(target)
		})

		b.Run("int slice", func(b *testing.B) {
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
