package serializer

import (
	"fmt"
	"math"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"

	item_models "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/models/item"
	go_serializer "gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer"
)

func Benchmark_UnsafeBinarySerializer(b *testing.B) {
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
		serializer := go_serializer.NewRawBinarySerializer()
		var err error
		for i := 0; i < b.N; i++ {
			_, err = serializer.Serialize(msg)
		}
		require.NoError(b, err)
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
		serializer := go_serializer.NewRawBinarySerializer()
		bs, err := serializer.Serialize(msg)
		require.NoError(b, err)

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			err = serializer.Deserialize(bs, &target)
		}
		require.NoError(b, err)
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
		serializer := go_serializer.NewRawBinarySerializer()

		var target item_models.Item
		for i := 0; i < b.N; i++ {
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
		}
	})
}

func BenchmarkType_UnsafeBinarySerializer(b *testing.B) {
	b.Run("string serialization", func(b *testing.B) {
		serializer := go_serializer.NewRawBinarySerializer()

		msg := "test-again#$çcçá"

		b.Run("encoding", func(b *testing.B) {
			var bs []byte
			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
			}

			var target string
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)
		})

		b.Run("decoding", func(b *testing.B) {
			bs, _ := serializer.Serialize(msg)

			var target string
			for i := 0; i < b.N; i++ {
				_ = serializer.Deserialize(bs, &target)
			}
			b.Log(target)
		})

		b.Run("encoding - decoding", func(b *testing.B) {
			var target string
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)

			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
				_ = serializer.Deserialize(bs, &target)
			}
		})
	})

	b.Run("int serialization", func(b *testing.B) {
		serializer := go_serializer.NewRawBinarySerializer()

		msg := uint64(math.MaxUint64)

		b.Run("encoding", func(b *testing.B) {
			var bs []byte
			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
			}

			var target uint64
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)
		})

		b.Run("decoding", func(b *testing.B) {
			bs, _ := serializer.Serialize(msg)

			var target uint64
			for i := 0; i < b.N; i++ {
				_ = serializer.Deserialize(bs, &target)
			}
			b.Log(target)
		})

		b.Run("encoding - decoding", func(b *testing.B) {
			var target uint64
			bs, _ := serializer.Serialize(msg)
			_ = serializer.Deserialize(bs, &target)
			b.Log(target)

			for i := 0; i < b.N; i++ {
				bs, _ = serializer.Serialize(msg)
				_ = serializer.Deserialize(bs, &target)
			}
		})
	})

	b.Run("slice serialization", func(b *testing.B) {
		b.Run("slice of int", func(b *testing.B) {
			msg := SliceTestData{
				IntList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}
			serializer := go_serializer.NewRawBinarySerializer()

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

			b.Run("encode - decode", func(b *testing.B) {
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
			serializer := go_serializer.NewRawBinarySerializer()

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
			serializer := go_serializer.NewRawBinarySerializer()

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
			serializer := go_serializer.NewRawBinarySerializer()

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
			serializer := go_serializer.NewRawBinarySerializer()

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

func Benchmark_Operation(b *testing.B) {
	b.Run("", func(b *testing.B) {
		var abs [8]byte
		i64 := int64(100)

		var i64Value int64
		for i := 0; i < b.N; i++ {
			*(*int64)(unsafe.Pointer(&abs)) = i64
			i64Value = *(*int64)(unsafe.Pointer(&abs[0]))
		}

		fmt.Println(i64Value)
	})

	b.Run("", func(b *testing.B) {
		bs := make([]byte, 8)
		i64 := int64(100)

		var i64Value int64
		for i := 0; i < b.N; i++ {
			go_serializer.PutUint64(bs, uint64(i64))
			i64Value = int64(go_serializer.Uint64(bs))
		}

		fmt.Println(i64Value)
	})

	b.Run("", func(b *testing.B) {
		var abs [8]byte
		i64 := uint64(100)

		var i64Value uint64
		for i := 0; i < b.N; i++ {
			*(*uint64)(unsafe.Pointer(&abs)) = i64
			i64Value = *(*uint64)(unsafe.Pointer(&abs[0]))
		}

		fmt.Println(i64Value)
	})

	b.Run("", func(b *testing.B) {
		bs := make([]byte, 8)
		i64 := uint64(100)

		var i64Value uint64
		for i := 0; i < b.N; i++ {
			go_serializer.PutUint64(bs, i64)
			i64Value = go_serializer.Uint64(bs)
		}

		fmt.Println(i64Value)
	})
}
