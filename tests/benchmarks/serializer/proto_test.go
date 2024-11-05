package serializer

import (
	"math"
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
		var err error
		for i := 0; i < b.N; i++ {
			_, err = serializer.Serialize(msg)
		}
		require.NoError(b, err)
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
		}
		require.NoError(b, err)
		validateMsgAndTarget(b, msg, &target)
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

func BenchmarkType_ProtoSerializer(b *testing.B) {
	b.Run("slice serialization", func(b *testing.B) {
		b.Run("slice of int", func(b *testing.B) {
			msg := &grpc_item.SliceTestData{
				IntList: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			serializer := go_serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target grpc_item.SliceTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target grpc_item.SliceTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.SliceTestData
				bs, err := serializer.Serialize(msg)
				require.NoError(b, err)
				err = serializer.Deserialize(bs, &target)
				require.NoError(b, err)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
					_ = serializer.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntIntList: [][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				},
			}
			serializer := go_serializer.NewBinarySerializer()

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

		b.Run("slice of slice of slice of int", func(b *testing.B) {
			msg := SliceTestData{
				ThreeDIntList: [][][]int{
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
						{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
					},
					{
						{12, 22, 32, 42, 52, 62, 72, 82, 92, 10},
						{102, 92, 82, 72, 62, 52, 42, 32, 22, 1},
					},
				},
			}
			serializer := go_serializer.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target SliceTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target SliceTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
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
	})

	b.Run("map serialization", func(b *testing.B) {
		b.Run("map of int to int", func(b *testing.B) {
			msg := &grpc_item.MapTestData{
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
			serializer := go_serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target grpc_item.MapTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target grpc_item.MapTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target grpc_item.MapTestData
				bs, _ := serializer.Serialize(msg)
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
					_ = serializer.Deserialize(bs, &target)
				}
			})
		})

		b.Run("map of string to string", func(b *testing.B) {
			msg := &grpc_item.MapTestData{
				StrKeyMapStrValue: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
				},
			}
			serializer := go_serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target grpc_item.MapTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target grpc_item.MapTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target grpc_item.MapTestData
				bs, _ := serializer.Serialize(msg)
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
					_ = serializer.Deserialize(bs, &target)
				}
			})
		})
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
