package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	grpc_item "gitlab.com/pietroski-software-company/devex/golang/serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/models/item"
	go_serializer "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/tools/serializer"
)

func Benchmark_JsonSerializer(b *testing.B) {
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
		serializer := go_serializer.NewJsonSerializer()

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
		serializer := go_serializer.NewJsonSerializer()
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
		serializer := go_serializer.NewJsonSerializer()

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
		serializer := go_serializer.NewJsonSerializer()

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
		serializer := go_serializer.NewJsonSerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
		}
	})
}

func Benchmark_ProtoJsonSerializer(b *testing.B) {
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
		serializer := go_serializer.NewJsonSerializer()
		for i := 0; i < b.N; i++ {
			_, err := serializer.Serialize(msg)
			require.NoError(b, err)
		}
	})
}

func validateStructMsgAndTarget(b *testing.B, msg, target *item_models.Item) {
	require.Equal(b, msg.Id, target.Id)
	require.Equal(b, msg.ItemId, target.ItemId)
	require.Equal(b, msg.Number, target.Number)
	require.Equal(b, msg.SubItem.Date, target.SubItem.Date)
	require.Equal(b, msg.SubItem.Amount, target.SubItem.Amount)
	require.Equal(b, msg.SubItem.ItemCode, target.SubItem.ItemCode)
}

func BenchmarkType_JsonSerializer(b *testing.B) {
	b.Run("slice serialization", func(b *testing.B) {
		b.Run("slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			serializer := go_serializer.NewJsonSerializer()

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

		b.Run("slice of slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntIntList: [][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				},
			}
			serializer := go_serializer.NewJsonSerializer()

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
			serializer := go_serializer.NewJsonSerializer()

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
	})

	b.Run("map serialization", func(b *testing.B) {
		b.Run("map of int to int", func(b *testing.B) {
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
			serializer := go_serializer.NewJsonSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target MapTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target MapTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target MapTestData
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
			msg := MapTestData{
				StrKeyMapStrValue: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
				},
			}
			serializer := go_serializer.NewJsonSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = serializer.Serialize(msg)
				}

				var target MapTestData
				_ = serializer.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := serializer.Serialize(msg)

				var target MapTestData
				for i := 0; i < b.N; i++ {
					_ = serializer.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target MapTestData
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
