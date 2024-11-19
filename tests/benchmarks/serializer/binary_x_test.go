//go:build benchmark

package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/serializerx"
)

func BenchmarkXBinarySerializer(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		msg := "test-again#$çcçá"

		s := serializerx.NewBinarySerializer()
		bs, err := s.Serialize(msg)
		require.NoError(b, err)
		var target string
		err = s.Deserialize(bs, &target)
		require.NoError(b, err)
		require.Equal(b, msg, target)

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

		b.Run("rebind", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
			}
		})
	})

	b.Run("number", func(b *testing.B) {
		b.Run("int64", func(b *testing.B) {
			msg := int64(math.MaxInt64)

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target int64
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.Equal(b, msg, target)

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

			b.Run("rebind", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("uint64", func(b *testing.B) {
			msg := uint64(math.MaxUint64)

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target uint64
			err = s.Deserialize(bs, &target)
			require.NoError(b, err)
			require.Equal(b, msg, target)

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

			b.Run("rebind", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})
	})

	b.Run("struct", func(b *testing.B) {
		b.Run("item sample", func(b *testing.B) {
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

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.Item
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("item sample - nil sub item", func(b *testing.B) {
			msg := &testmodels.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.Item
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("simplified special struct test data", func(b *testing.B) {
			msg := &testmodels.SimplifiedSpecialStructTestData{
				Bool:    true,
				String:  "any-string",
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

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.SimplifiedSpecialStructTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("string struct only", func(b *testing.B) {
			msg := &testmodels.StringStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.StringStruct
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("int64 struct only", func(b *testing.B) {
			msg := &testmodels.Int64Struct{
				FirstInt64:  math.MaxInt64,
				SecondInt64: -math.MaxInt64,
				ThirdInt64:  math.MaxInt64,
				FourthInt64: -math.MaxInt64,
				FifthInt64:  0,
				SixthInt64:  -0,
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.Int64Struct
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

			b.Run("encode - decode", func(b *testing.B) {
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
			msg := &testmodels.Int64SliceTestData{
				Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.Int64SliceTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]uint64", func(b *testing.B) {
			msg := &testmodels.Uint64SliceTestData{
				Uint64List: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.Uint64SliceTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]string", func(b *testing.B) {
			msg := &testmodels.StringSliceTestData{
				StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.StringSliceTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[]byte", func(b *testing.B) {
			msg := &testmodels.ByteSliceTestData{
				ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.ByteSliceTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})

		b.Run("[][]byte", func(b *testing.B) {
			msg := &testmodels.ByteByteSliceTestData{
				ByteByteList: [][]byte{
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
					{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
				},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.ByteByteSliceTestData
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

			b.Run("encode - decode", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})

			b.Log()
			b.Log(target)
		})
	})

	b.Run("map", func(b *testing.B) {
		b.Run("map[string]string", func(b *testing.B) {
			msg := &testmodels.MapStringStringTestData{
				MapStringString: map[string]string{
					"any-key":       "any-value",
					"any-other-key": "any-other-value",
					"another-key":   "another-value",
				},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.MapStringStringTestData
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
			msg := &testmodels.MapInt64Int64TestData{
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

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.MapInt64Int64TestData
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

		b.Run("map[int64]*testmodels.StructTestData", func(b *testing.B) {
			msg := &testmodels.MapInt64StructPointerTestData{
				MapInt64StructPointer: map[int64]*testmodels.StructTestData{
					0: {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					2: {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					4: {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.MapInt64StructPointerTestData
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

		b.Run("map[string]*testmodels.StructTestData", func(b *testing.B) {
			msg := &testmodels.MapStringStructPointerTestData{
				MapStringStructPointer: map[string]*testmodels.StructTestData{
					"any-key": {
						Bool:   true,
						String: "any-string",
						Int64:  math.MaxInt64,
					},
					"any-other-key": {
						Bool:   false,
						String: "any-other-string",
						Int64:  -math.MaxInt64,
					},
					"another-key": {
						Bool:   false,
						String: "",
						Int64:  0,
					},
				},
			}

			s := serializerx.NewBinarySerializer()
			bs, err := s.Serialize(msg)
			require.NoError(b, err)
			var target testmodels.MapStringStructPointerTestData
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
				msg := &testmodels.MapTestData{
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
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target testmodels.MapTestData
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
				msg := &testmodels.MapTestData{
					StrKeyMapStrValue: map[string]string{
						"any-key":       "any-value",
						"any-other-key": "any-other-value",
					},
				}

				s := serializerx.NewBinarySerializer()
				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				var target testmodels.MapTestData
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
