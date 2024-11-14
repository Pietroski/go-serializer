package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/serializerx"
)

func Benchmark_BinaryXSerializer(b *testing.B) {
	msg := &testmodels.Item{
		Id:     "any-item",
		ItemId: 100,
		Number: 5_000_000_000,
		SubItem: &testmodels.SubItem{
			Date:     time.Now().Unix(),
			Amount:   1_000_000_000,
			ItemCode: "code-status",
		},
	}

	b.Run("encoding", func(b *testing.B) {
		s := serializerx.NewBinarySerializer()

		var err error
		for i := 0; i < b.N; i++ {
			_, err = s.Serialize(msg)
		}
		require.NoError(b, err)
	})

	b.Run("decoding", func(b *testing.B) {
		s := serializerx.NewBinarySerializer()

		bs, err := s.Serialize(msg)
		require.NoError(b, err)

		var target testmodels.Item
		for i := 0; i < b.N; i++ {
			err = s.Deserialize(bs, &target)
		}
		require.NoError(b, err)
	})

	b.Run("encoding - decoding", func(b *testing.B) {
		s := serializerx.NewBinarySerializer()

		var target testmodels.Item
		for i := 0; i < b.N; i++ {
			bs, _ := s.Serialize(msg)
			_ = s.Deserialize(bs, &target)
		}
	})
}

func BenchmarkType_BinaryXSerializer(b *testing.B) {
	b.Run("string serialization", func(b *testing.B) {
		s := serializerx.NewBinarySerializer()

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
		s := serializerx.NewBinarySerializer()

		msg := math.MaxInt64

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

	b.Run("uint serialization", func(b *testing.B) {
		s := serializerx.NewBinarySerializer()

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
			msg := &IntSliceTestData{
				IntList: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target IntSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target IntSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target IntSliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
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
			s := serializerx.NewBinarySerializer()

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
			s := serializerx.NewBinarySerializer()

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

		b.Run("slice of uint", func(b *testing.B) {
			msg := &ProtoTypeSliceTestData{
				UintList: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ProtoTypeSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ProtoTypeSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ProtoTypeSliceTestData
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("slice of string", func(b *testing.B) {
			msg := &ProtoTypeSliceTestData{
				StrList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ProtoTypeSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ProtoTypeSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ProtoTypeSliceTestData
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
			msg := &ProtoTypeSliceTestData{
				BytesBytesList: [][]byte{
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
					{255, 0, 4, 8, 16},
				},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ProtoTypeSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ProtoTypeSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ProtoTypeSliceTestData
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
			msg := &ProtoTypeSliceTestData{
				BytesBytesList: [][]byte{
					{},
					{},
					{},
					{},
					{},
				},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ProtoTypeSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ProtoTypeSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ProtoTypeSliceTestData
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
			msg := &ByteSliceTestData{
				ByteList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
			}
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ByteSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ByteSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ByteSliceTestData
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
			msg := &ByteSliceTestData{
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
			s := serializerx.NewBinarySerializer()

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target ByteSliceTestData
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target ByteSliceTestData
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encode - decode", func(b *testing.B) {
				var target ByteSliceTestData
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
			msg := &MapTestData{
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
			s := serializerx.NewBinarySerializer()

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
			s := serializerx.NewBinarySerializer()

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
