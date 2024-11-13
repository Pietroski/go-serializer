package serializer

import (
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/generated/go/pkg/item"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
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
		s := serializer.NewProtoSerializer()
		var err error
		for i := 0; i < b.N; i++ {
			_, err = s.Serialize(msg)
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
		s := serializer.NewProtoSerializer()
		bs, err := s.Serialize(msg)
		require.NoError(b, err)

		var target grpc_item.Item
		for i := 0; i < b.N; i++ {
			err = s.Deserialize(bs, &target)
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
		s := serializer.NewProtoSerializer()

		var target grpc_item.Item
		for i := 0; i < b.N; i++ {
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
		}
	})
}

func BenchmarkType_ProtoSerializer(b *testing.B) {
	b.Run("slice serialization", func(b *testing.B) {
		b.Run("slice of int", func(b *testing.B) {
			msg := &grpc_item.IntSliceTestData{
				IntList: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.IntSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.IntSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.IntSliceTestData
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of uint", func(b *testing.B) {
			msg := &grpc_item.SliceTestData{
				UintList: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.SliceTestData
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of string", func(b *testing.B) {
			msg := &grpc_item.SliceTestData{
				StrList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.SliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of slice of bytes", func(b *testing.B) {
			msg := &grpc_item.SliceTestData{
				BytesBytesList: [][]byte{
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
				},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.SliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of slice of bytes", func(b *testing.B) {
			msg := &grpc_item.SliceTestData{
				BytesBytesList: [][]byte{
					{},
					{},
					{},
					{},
					{},
				},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.SliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of bytes", func(b *testing.B) {
			msg := &grpc_item.ByteSliceTestData{
				ByteList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.ByteSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.ByteSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.ByteSliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of bytes", func(b *testing.B) {
			msg := &grpc_item.ByteSliceTestData{
				ByteList: []byte{math.MaxUint8,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
					255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
				},
			}
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.ByteSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.ByteSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target grpc_item.ByteSliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
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
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.MapTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.MapTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target grpc_item.MapTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
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
			s := serializer.NewProtoSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target grpc_item.MapTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target grpc_item.MapTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target grpc_item.MapTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
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
