package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/generated/go/pkg/item"
)

func BenchmarkProtoSerializer(b *testing.B) {
	b.Run("struct", func(b *testing.B) {
		b.Run("item sample", func(b *testing.B) {
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
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
			b.Log(target.SubItem)
		})

		b.Run("item sample - nil sub item", func(b *testing.B) {
			msg := &grpc_item.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.Item
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
			b.Log(target.SubItem)
		})

		b.Run("simplified special struct test data", func(b *testing.B) {
			msg := &grpc_item.SimplifiedSpecialStructTestData{
				Bool:    true,
				Str:     "any-string",
				Int32:   math.MaxInt32,
				Int64:   math.MaxInt64,
				Uint32:  math.MaxUint32,
				Uint64:  math.MaxUint64,
				Float32: math.MaxFloat32,
				Float64: math.MaxFloat64,
				Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
				RepeatedBytes: [][]byte{
					{-0, 0, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, math.MaxInt8, 255, 0, -0},
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.SimplifiedSpecialStructTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("string struct only", func(b *testing.B) {
			msg := &grpc_item.StringStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.StringStruct
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("int64 struct only", func(b *testing.B) {
			msg := &grpc_item.Int64Struct{
				FirstInt64:  math.MaxInt64,
				SecondInt64: -math.MaxInt64,
				ThirdInt64:  math.MaxInt64,
				FourthInt64: -math.MaxInt64,
				FifthInt64:  0,
				SixthInt64:  -0,
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.Int64Struct
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})
	})

	b.Run("slice", func(b *testing.B) {
		b.Run("[]int64", func(b *testing.B) {
			msg := &grpc_item.Int64SliceTestData{
				Int64List: []int64{
					-math.MaxInt64, -9223372036854775808, -0, 0, 2, 12345678, 4, 5, 5170, 10, 8,
					87654321, 9223372036854775807, math.MaxInt64,
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.Int64SliceTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]uint64", func(b *testing.B) {
			msg := &grpc_item.Uint64SliceTestData{
				Uint64List: []uint64{
					-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
					9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.Uint64SliceTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]string", func(b *testing.B) {
			msg := &grpc_item.StringSliceTestData{
				StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.StringSliceTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]byte", func(b *testing.B) {
			msg := &grpc_item.ByteSliceTestData{
				ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.ByteSliceTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[][]byte", func(b *testing.B) {
			msg := &grpc_item.ByteByteSliceTestData{
				ByteByteList: [][]byte{
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.ByteByteSliceTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("extra cases", func(b *testing.B) {
			b.Run("[]int64", func(b *testing.B) {
				msg := &grpc_item.SliceTestData{
					Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.SliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[]uint64", func(b *testing.B) {
				msg := &grpc_item.SliceTestData{
					Uint64List: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.SliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[]string", func(b *testing.B) {
				msg := &grpc_item.SliceTestData{
					StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.SliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[][]bytes", func(b *testing.B) {
				msg := &grpc_item.SliceTestData{
					RepeatedBytesList: [][]byte{
						{255, 0, 4, 8, 16},
						{255, 0, 4, 8, 16},
						{255, 0, 4, 8, 16},
						{255, 0, 4, 8, 16},
						{255, 0, 4, 8, 16},
					},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.SliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[][]bytes(nil)", func(b *testing.B) {
				msg := &grpc_item.SliceTestData{
					RepeatedBytesList: [][]byte{
						{},
						{},
						{},
						{},
						{},
					},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.SliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[]bytes", func(b *testing.B) {
				msg := &grpc_item.ByteSliceTestData{
					ByteList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.ByteSliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("[]bytes - huge", func(b *testing.B) {
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
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.ByteSliceTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})
		})
	})

	b.Run("map", func(b *testing.B) {
		b.Run("map[string]string", func(b *testing.B) {
			msg := &grpc_item.MapStringStringTestData{
				MapStringString: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
					"another-key":   "another-value",
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.MapStringStringTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("map[int64]int64", func(b *testing.B) {
			msg := &grpc_item.MapInt64Int64TestData{
				MapInt64Int64: map[int64]int64{
					0:              math.MaxInt64,
					1:              math.MaxInt8,
					2:              math.MaxInt16,
					3:              math.MaxInt32,
					4:              math.MaxInt64,
					math.MaxInt64:  0,
					math.MaxInt8:   1,
					math.MaxInt16:  2,
					math.MaxInt32:  3,
					-math.MaxInt64: 4,
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.MapInt64Int64TestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("map[int64]*grpc_item.StructTestData", func(b *testing.B) {
			msg := &grpc_item.MapInt64StructPointerTestData{
				MapInt64StructPointerTestData: map[int64]*grpc_item.StructTestData{
					0: {
						Bool:  true,
						Str:   "any-string",
						Int64: math.MaxInt64,
					},
					2: {
						Bool:  false,
						Str:   "any-other-string",
						Int64: -math.MaxInt64,
					},
					4: {
						Bool:  false,
						Str:   "",
						Int64: 0,
					},
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.MapInt64StructPointerTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("map[string]*grpc_item.StructTestData", func(b *testing.B) {
			msg := &grpc_item.MapStringStructPointerTestData{
				MapStringStructPointerTestData: map[string]*grpc_item.StructTestData{
					"any-key": {
						Bool:  true,
						Str:   "any-string",
						Int64: math.MaxInt64,
					},
					"any-other-key": {
						Bool:  false,
						Str:   "any-other-string",
						Int64: -math.MaxInt64,
					},
					"another-key": {
						Bool:  false,
						Str:   "",
						Int64: 0,
					},
				},
			}

			s := serializer.NewProtoSerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target grpc_item.MapStringStructPointerTestData
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.EqualExportedValues(b, msg, &target)

			b.Run("encoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
				}
			})

			b.Run("decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("extra cases", func(b *testing.B) {
			b.Run("map[int64]int64", func(b *testing.B) {
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
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.MapTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})

			b.Run("map[string]string", func(b *testing.B) {
				msg := &grpc_item.MapTestData{
					StrKeyMapStrValue: map[string]string{
						"any-key":       "any-value",
						"any-other-key": "any-other-value",
					},
				}

				s := serializer.NewProtoSerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target grpc_item.MapTestData
				err = s.Deserialize(bs, &target)
				require.NoError(b, err)
				require.EqualExportedValues(b, msg, &target)

				b.Run("encoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
					}
				})

				b.Run("decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Run("encoding - decoding", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_, _ = s.Serialize(msg)
						_ = s.Deserialize(bs, &target)
					}
				})

				b.Log()
				b.Log(target)
			})
		})
	})
}
