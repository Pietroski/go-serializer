package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	grpc_item "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/generated/go/pkg/item"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
)

func Benchmark_ProtoSerializer(b *testing.B) {
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
		serializer := go_serializer.NewProtoSerializer()
		for i := 0; i < b.N; i++ {
			_, err := serializer.Serialize(msg)
			require.NoError(b, err)
		}
	})

	b.Run("proto deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewProtoSerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(b, err)

		var target grpc_item.Item
		for i := 0; i < b.N; i++ {
			err = serializer.Deserialize(bs, &target)
			require.NoError(b, err)
		}
		validateMsgAndTarget(b, msg, &target)
	})
}

func Benchmark_ProtoSerializerAndDeserialization(b *testing.B) {
	b.Run("proto serialization and deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewProtoSerializer()
		var bs []byte
		var err error
		for i := 0; i < b.N; i++ {
			bs, err = serializer.Serialize(msg)
			require.NoError(b, err)
		}

		var target grpc_item.Item
		for i := 0; i < b.N; i++ {
			err = serializer.Deserialize(bs, &target)
			require.NoError(b, err)
		}
		validateMsgAndTarget(b, msg, &target)
	})

	b.Run("proto serialization and deserialization", func(b *testing.B) {
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
		serializer := go_serializer.NewProtoSerializer()

		for i := 0; i < b.N; i++ {
			bs, err := serializer.Serialize(msg)
			require.NoError(b, err)

			var target grpc_item.Item
			err = serializer.Deserialize(bs, &target)
			require.NoError(b, err)
			validateMsgAndTarget(b, msg, &target)
		}
	})

	b.Run("proto serialization and deserialization - clean - no validation", func(b *testing.B) {
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
		serializer := go_serializer.NewProtoSerializer()

		var target grpc_item.Item
		for i := 0; i < b.N; i++ {
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
		}
	})
}

func validateMsgAndTarget(b *testing.B, msg, target *grpc_item.Item) {
	require.Equal(b, msg.Id, target.Id)
	require.Equal(b, msg.ItemId, target.ItemId)
	require.Equal(b, msg.Number, target.Number)
	require.Equal(b, msg.SubItem.Date, target.SubItem.Date)
	require.Equal(b, msg.SubItem.Amount, target.SubItem.Amount)
	require.Equal(b, msg.SubItem.ItemCode, target.SubItem.ItemCode)
}
