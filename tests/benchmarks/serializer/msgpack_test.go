package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	grpc_item "gitlab.com/pietroski-software-company/devex/golang/serializer/generated/go/pkg/item"
	item_models "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/models/item"
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
		s := serializer.NewMsgPackSerializer()

		var err error
		for i := 0; i < b.N; i++ {
			_, err = s.Serialize(msg)
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
		s := serializer.NewMsgPackSerializer()
		bs, err := s.Serialize(msg)
		require.NoError(b, err)

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err = s.Deserialize(bs, &target)
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
		s := serializer.NewMsgPackSerializer()

		var bs []byte
		var err error
		for i := 0; i < b.N; i++ {
			bs, err = s.Serialize(msg)
			require.NoError(b, err)
		}

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err := s.Deserialize(bs, &target)
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
		s := serializer.NewMsgPackSerializer()

		for i := 0; i < b.N; i++ {
			bs, err := s.Serialize(msg)
			require.NoError(b, err)

			var target item_models.Item
			err = s.Deserialize(bs, &target)
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
		s := serializer.NewMsgPackSerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
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
		s := serializer.NewMsgPackSerializer()
		for i := 0; i < b.N; i++ {
			_, err := s.Serialize(msg)
			require.NoError(b, err)
		}
	})
}

func BenchmarkType_MsgPackSerializer(b *testing.B) {
	b.Run("string serialization", func(b *testing.B) {
		s := serializer.NewMsgPackSerializer()

		msg := "test-again#$çcçá"

		b.Run("encoding", func(b *testing.B) {
			var bs []byte
			for i := 0; i < b.N; i++ {
				bs, _ = s.Serialize(msg)
			}

			var target string
			_ = s.Deserialize(bs, &target)
			b.Log(target)
		})

		b.Run("decoding", func(b *testing.B) {
			bs, _ := s.Serialize(msg)

			var target string
			for i := 0; i < b.N; i++ {
				_ = s.Deserialize(bs, &target)
			}
			b.Log(target)
		})

		b.Run("encoding - decoding", func(b *testing.B) {
			var target string
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
			b.Log(target)

			for i := 0; i < b.N; i++ {
				bs, _ = s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
			}
		})
	})

	b.Run("int serialization", func(b *testing.B) {
		s := serializer.NewMsgPackSerializer()

		msg := uint64(math.MaxUint64)

		b.Run("encoding", func(b *testing.B) {
			var bs []byte
			for i := 0; i < b.N; i++ {
				bs, _ = s.Serialize(msg)
			}

			var target uint64
			_ = s.Deserialize(bs, &target)
			b.Log(target)
		})

		b.Run("decoding", func(b *testing.B) {
			bs, _ := s.Serialize(msg)

			var target uint64
			for i := 0; i < b.N; i++ {
				_ = s.Deserialize(bs, &target)
			}
			b.Log(target)
		})

		b.Run("encoding - decoding", func(b *testing.B) {
			var target uint64
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
			b.Log(target)

			for i := 0; i < b.N; i++ {
				bs, _ = s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
			}
		})
	})

	b.Run("slice serialization", func(b *testing.B) {
		b.Run("slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			s := serializer.NewMsgPackSerializer()

			b.Run("int slice - encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("int slice - decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("int slice", func(b *testing.B) {
				var target SliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of of slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntIntList: [][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
					{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
				},
			}
			s := serializer.NewMsgPackSerializer()

			b.Run("int slice - encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("int slice - decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("int slice", func(b *testing.B) {
				var target SliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
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
			s := serializer.NewMsgPackSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target SliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target SliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target SliceTestData
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
			s := serializer.NewMsgPackSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target MapTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target MapTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target MapTestData
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
			msg := MapTestData{
				StrKeyMapStrValue: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
				},
			}
			s := serializer.NewMsgPackSerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target MapTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target MapTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target MapTestData
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
