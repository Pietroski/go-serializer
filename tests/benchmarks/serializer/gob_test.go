package serializer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	grpc_item "gitlab.com/pietroski-software-company/devex/golang/serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/models/item"
)

func Benchmark_GobSerializer(b *testing.B) {
	b.Run("proto serialization", func(b *testing.B) {
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
		s := serializer.NewGobSerializer()
		for i := 0; i < b.N; i++ {
			_, err := s.Serialize(msg)
			require.NoError(b, err)
		}
	})

	b.Run("gob/json serialization and deserialization - clean - no validation", func(b *testing.B) {
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
		s := serializer.NewGobSerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
		}
	})
}

func Benchmark_ProtoGobSerializer(b *testing.B) {
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
		s := serializer.NewGobSerializer()
		for i := 0; i < b.N; i++ {
			_, err := s.Serialize(msg)
			require.NoError(b, err)
		}
	})
}
