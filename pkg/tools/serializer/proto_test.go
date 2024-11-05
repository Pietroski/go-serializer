//go:build unit

package go_serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	grpc_item "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
)

func Test_ProtoSerializer(t *testing.T) {
	t.Run("proto serialization", func(t *testing.T) {
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
		serializer := NewProtoSerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(t, err)
		require.NotEmpty(t, bs)
	})

	t.Run("proto serialization error", func(t *testing.T) {
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
		serializer := NewProtoSerializer()
		bs, err := serializer.Serialize(msg)
		require.Error(t, err)
		require.Len(t, bs, 0)
	})

	t.Run("proto deserialization", func(t *testing.T) {
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
		serializer := NewProtoSerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(t, err)
		require.NotEmpty(t, bs)

		var target grpc_item.Item
		err = serializer.Deserialize(bs, &target)
		require.NoError(t, err)
		validateMsgAndTarget(t, msg, &target)
	})

	t.Run("proto deserialization error", func(t *testing.T) {
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
		s := NewJsonSerializer()
		bs, err := s.Serialize(msg)
		require.NoError(t, err)
		require.Greater(t, len(bs), 0)

		serializer := NewProtoSerializer()
		var target grpc_item.Item
		err = serializer.Deserialize(bs, &target)
		require.Error(t, err)
		t.Log(err)
	})

	t.Run("proto data rebind", func(t *testing.T) {
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

		serializer := NewProtoSerializer()

		var target grpc_item.Item
		err := serializer.DataRebind(msg, &target)
		require.NoError(t, err)
		validateMsgAndTarget(t, msg, &target)
	})

	t.Run("proto data rebind", func(t *testing.T) {
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

		serializer := NewProtoSerializer()

		var target grpc_item.AnotherItem
		err := serializer.DataRebind(msg, &target)
		require.NoError(t, err)
		require.Equal(t, msg.Id, target.Id)
		require.Equal(t, msg.ItemId, target.ItemId)
		require.Equal(t, msg.Number, target.Number)
		require.Equal(t, msg.SubItem.Date, target.SubItem.Date)
		require.Equal(t, msg.SubItem.Amount, target.SubItem.Amount)
		require.Equal(t, msg.SubItem.ItemCode, target.SubItem.ItemCode)
	})

	t.Run("proto data rebind error on serializing", func(t *testing.T) {
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

		serializer := NewProtoSerializer()

		var target grpc_item.Item
		err := serializer.DataRebind(msg, &target)
		require.Error(t, err)
		require.Empty(t, &target)
	})
}

func Test_ProtoSerializer_Benchmark_Data(t *testing.T) {
	t.Run("success MapTestData", func(t *testing.T) {
		t.Run("map of int to int", func(t *testing.T) {
			serializer := NewProtoSerializer()

			msg := grpc_item.MapTestData{
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

			var td grpc_item.MapTestData
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
			serializer := NewProtoSerializer()

			msg := grpc_item.MapTestData{
				StrKeyMapStrValue: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
				},
			}
			bs, err := serializer.Serialize(&msg)
			require.NoError(t, err)

			//t.Log(string(bs), bs)

			var td grpc_item.MapTestData
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

func validateMsgAndTarget(t *testing.T, msg, target *grpc_item.Item) {
	require.Equal(t, msg.Id, target.Id)
	require.Equal(t, msg.ItemId, target.ItemId)
	require.Equal(t, msg.Number, target.Number)
	require.Equal(t, msg.SubItem.Date, target.SubItem.Date)
	require.Equal(t, msg.SubItem.Amount, target.SubItem.Amount)
	require.Equal(t, msg.SubItem.ItemCode, target.SubItem.ItemCode)
}
