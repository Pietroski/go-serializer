package serializer

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
)

func BenchmarkBinarySerializer(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		msg := "test-again#$çcçá"

		s := serializer.NewBinarySerializer()
		bs, err := s.Serialize(msg)
		require.NoError(b, err)
		var target string
		err = s.Deserialize(bs, &target)
		require.NoError(b, err)
		require.EqualExportedValues(b, msg, &target)

		b.Run("encoding", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.Serialize(msg)
			}
		})

		b.Run("decoding", func(b *testing.B) {
			_ = s.Deserialize(bs, &target)
		})

		b.Run("rebind", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
			}
		})

		b.Log()
		b.Log(target)
	})

	b.Run("number", func(b *testing.B) {
		b.Run("int", func(b *testing.B) {
			s := serializer.NewBinarySerializer()

			msg := int64(math.MaxInt64)

			b.Run("encoding", func(b *testing.B) {
				var bs []byte
				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
				}

				var target int64
				_ = s.Deserialize(bs, &target)
				b.Log(target)
			})

			b.Run("decoding", func(b *testing.B) {
				bs, _ := s.Serialize(msg)

				var target int64
				for i := 0; i < b.N; i++ {
					_ = s.Deserialize(bs, &target)
				}
				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				var target int64
				bs, _ := s.Serialize(msg)
				_ = s.Deserialize(bs, &target)
				b.Log(target)

				for i := 0; i < b.N; i++ {
					bs, _ = s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("uint64", func(b *testing.B) {
			s := serializer.NewBinarySerializer()

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

			b.Run("encoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var err error
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
				require.NoError(b, err)
			})

			b.Run("decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				bs, err := s.Serialize(msg)
				require.NoError(b, err)

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					err = s.Deserialize(bs, &target)
				}
				require.NoError(b, err)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					bs, _ := s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})

		b.Run("item sample - nil sub item", func(b *testing.B) {
			msg := &testmodels.Item{
				Id:     "any-item",
				ItemId: 100,
				Number: 5_000_000_000,
			}

			b.Run("encoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var err error
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
				require.NoError(b, err)
			})

			b.Run("decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				bs, err := s.Serialize(msg)
				require.NoError(b, err)

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					err = s.Deserialize(bs, &target)
				}
				require.NoError(b, err)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					bs, _ := s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
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

			b.Run("encoding", func(b *testing.B) {
				s := serializer.NewProtoSerializer()

				var err error
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
				require.NoError(b, err)
			})

			b.Run("decoding", func(b *testing.B) {
				s := serializer.NewProtoSerializer()

				bs, err := s.Serialize(msg)
				require.NoError(b, err)
				require.NotNil(b, bs)

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					err = s.Deserialize(bs, &target)
				}
				require.NoError(b, err)

				b.Log(target)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				s := serializer.NewProtoSerializer()

				var target testmodels.Item
				for i := 0; i < b.N; i++ {
					bs, _ := s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}

				b.Log(target)
			})
		})

		b.Run("string struct only", func(b *testing.B) {
			msg := &testmodels.StringStruct{
				FirstString:  "first string value",
				SecondString: "second string value",
				ThirdString:  "third string value",
				FourthString: "fourth string value",
				FifthString:  "fifth string value",
			}

			b.Run("encoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var err error
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
				require.NoError(b, err)
			})

			b.Run("decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				bs, err := s.Serialize(msg)
				require.NoError(b, err)

				var target testmodels.StringStruct
				for i := 0; i < b.N; i++ {
					err = s.Deserialize(bs, &target)
				}
				require.NoError(b, err)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var target testmodels.StringStruct
				for i := 0; i < b.N; i++ {
					bs, _ := s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
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

			b.Run("encoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var err error
				for i := 0; i < b.N; i++ {
					_, err = s.Serialize(msg)
				}
				require.NoError(b, err)
			})

			b.Run("decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				bs, err := s.Serialize(msg)
				require.NoError(b, err)

				var target testmodels.Int64Struct
				for i := 0; i < b.N; i++ {
					err = s.Deserialize(bs, &target)
				}
				require.NoError(b, err)
			})

			b.Run("encoding - decoding", func(b *testing.B) {
				s := serializer.NewBinarySerializer()

				var target testmodels.Int64Struct
				for i := 0; i < b.N; i++ {
					bs, _ := s.Serialize(msg)
					_ = s.Deserialize(bs, &target)
				}
			})
		})
	})

	b.Run("slice", func(b *testing.B) {
		b.Run("[]int64", func(b *testing.B) {
			msg := &testmodels.Int64SliceTestData{
				Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			}

			s := serializer.NewBinarySerializer()
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
				_ = s.Deserialize(bs, &target)
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

		b.Run("extra cases", func(b *testing.B) {
			//b.Run("slice of slice of int", func(b *testing.B) {
			//	msg := testmodels.SliceTestData{
			//		IntIntList: [][]int{
			//			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			//			{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			//		},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target SliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target SliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encoding - decoding", func(b *testing.B) {
			//		var target SliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of slice of slice of int", func(b *testing.B) {
			//	msg := SliceTestData{
			//		ThreeDIntList: [][][]int{
			//			{
			//				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			//				{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			//			},
			//			{
			//				{12, 22, 32, 42, 52, 62, 72, 82, 92, 10},
			//				{102, 92, 82, 72, 62, 52, 42, 32, 22, 1},
			//			},
			//		},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target SliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target SliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encoding - decoding", func(b *testing.B) {
			//		var target SliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of uint", func(b *testing.B) {
			//	msg := &ProtoTypeSliceTestData{
			//		UintList: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ProtoTypeSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ProtoTypeSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ProtoTypeSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of string", func(b *testing.B) {
			//	msg := &ProtoTypeSliceTestData{
			//		StrList: []string{"first-item", "second-item", "third-item", "fourth-item"},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ProtoTypeSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ProtoTypeSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ProtoTypeSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of slice of bytes", func(b *testing.B) {
			//	msg := &ProtoTypeSliceTestData{
			//		BytesBytesList: [][]byte{
			//			{255, 0, 4, 8, 16},
			//			{255, 0, 4, 8, 16},
			//			{255, 0, 4, 8, 16},
			//			{255, 0, 4, 8, 16},
			//			{255, 0, 4, 8, 16},
			//		},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ProtoTypeSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ProtoTypeSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ProtoTypeSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of slice of bytes", func(b *testing.B) {
			//	msg := &ProtoTypeSliceTestData{
			//		BytesBytesList: [][]byte{
			//			{},
			//			{},
			//			{},
			//			{},
			//			{},
			//		},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ProtoTypeSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ProtoTypeSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ProtoTypeSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of bytes", func(b *testing.B) {
			//	msg := &ByteSliceTestData{
			//		ByteList: []byte{255, 0, 4, 8, 16, 48, 56, 32, 44, 200},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ByteSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ByteSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ByteSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("slice of bytes", func(b *testing.B) {
			//	msg := &ByteSliceTestData{
			//		ByteList: []byte{math.MaxUint8,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//			255, 0, 4, 8, 16, 48, 56, 32, 44, 200,
			//		},
			//	}
			//	s := serializer.NewBinarySerializer()
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		var bs []byte
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//		}
			//
			//		var target ByteSliceTestData
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		bs, _ := s.Serialize(msg)
			//
			//		var target ByteSliceTestData
			//		for i := 0; i < b.N; i++ {
			//			_ = s.Deserialize(bs, &target)
			//		}
			//		b.Log(target)
			//	})
			//
			//	b.Run("encode - decode", func(b *testing.B) {
			//		var target ByteSliceTestData
			//		bs, _ := s.Serialize(msg)
			//		_ = s.Deserialize(bs, &target)
			//		b.Log(target)
			//
			//		for i := 0; i < b.N; i++ {
			//			bs, _ = s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//	})
			//})
			//
			//b.Run("simplified special struct test data", func(b *testing.B) {
			//	msg := &testmodels.ProtoTypeSliceTestData{
			//		Bool:    true,
			//		String:  "any-string",
			//		Int32:   math.MaxInt32,
			//		Int64:   math.MaxInt64,
			//		Uint32:  math.MaxUint32,
			//		Uint64:  math.MaxUint64,
			//		Float32: math.MaxFloat32,
			//		Float64: math.MaxFloat64,
			//		Bytes:   []byte{-0, 0, 255, math.MaxInt8, math.MaxUint8},
			//		RepeatedBytes: [][]byte{
			//			{-0, 0, 255, math.MaxInt8, math.MaxUint8},
			//			{math.MaxUint8, math.MaxInt8, math.MaxUint8},
			//			{math.MaxUint8, math.MaxInt8, 255, 0, -0},
			//		},
			//	}
			//
			//	b.Run("encoding", func(b *testing.B) {
			//		s := serializer.NewProtoSerializer()
			//
			//		var err error
			//		for i := 0; i < b.N; i++ {
			//			_, err = s.Serialize(msg)
			//		}
			//		require.NoError(b, err)
			//	})
			//
			//	b.Run("decoding", func(b *testing.B) {
			//		s := serializer.NewProtoSerializer()
			//
			//		bs, err := s.Serialize(msg)
			//		require.NoError(b, err)
			//		require.NotNil(b, bs)
			//
			//		var target testmodels.Item
			//		for i := 0; i < b.N; i++ {
			//			err = s.Deserialize(bs, &target)
			//		}
			//		require.NoError(b, err)
			//
			//		b.Log(target)
			//	})
			//
			//	b.Run("encoding - decoding", func(b *testing.B) {
			//		s := serializer.NewProtoSerializer()
			//
			//		var target testmodels.Item
			//		for i := 0; i < b.N; i++ {
			//			bs, _ := s.Serialize(msg)
			//			_ = s.Deserialize(bs, &target)
			//		}
			//
			//		b.Log(target)
			//	})
			//})
		})
	})

	b.Run("map", func(b *testing.B) {
		//b.Run("map[string]string", func(b *testing.B) {
		//	msg := MapTestData{
		//		StrKeyMapStrValue: map[string]string{
		//			"any-key":       "any-value",
		//			"any-other-key": "any-other-value",
		//		},
		//	}
		//	s := serializer.NewBinarySerializer()
		//
		//	b.Run("encoding", func(b *testing.B) {
		//		var bs []byte
		//		for i := 0; i < b.N; i++ {
		//			bs, _ = s.Serialize(msg)
		//		}
		//
		//		var target MapTestData
		//		_ = s.Deserialize(bs, &target)
		//		b.Log(target)
		//	})
		//
		//	b.Run("decoding", func(b *testing.B) {
		//		bs, _ := s.Serialize(msg)
		//
		//		var target MapTestData
		//		for i := 0; i < b.N; i++ {
		//			_ = s.Deserialize(bs, &target)
		//		}
		//		b.Log(target)
		//	})
		//
		//	b.Run("encoding - decoding", func(b *testing.B) {
		//		var target MapTestData
		//		bs, _ := s.Serialize(msg)
		//		_ = s.Deserialize(bs, &target)
		//		b.Log(target)
		//
		//		for i := 0; i < b.N; i++ {
		//			bs, _ = s.Serialize(msg)
		//			_ = s.Deserialize(bs, &target)
		//		}
		//	})
		//})
		//
		//b.Run("map[int]int", func(b *testing.B) {
		//	msg := &MapTestData{
		//		Int64KeyMapInt64Value: map[int64]int64{
		//			0:     100,
		//			7:     2,
		//			2:     8,
		//			8:     4,
		//			4:     16,
		//			100:   200,
		//			1_000: math.MaxInt64,
		//		},
		//	}
		//	s := serializer.NewBinarySerializer()
		//
		//	b.Run("encoding", func(b *testing.B) {
		//		var bs []byte
		//		for i := 0; i < b.N; i++ {
		//			bs, _ = s.Serialize(msg)
		//		}
		//
		//		var target MapTestData
		//		_ = s.Deserialize(bs, &target)
		//		b.Log(target)
		//	})
		//
		//	b.Run("decoding", func(b *testing.B) {
		//		bs, _ := s.Serialize(msg)
		//
		//		var target MapTestData
		//		for i := 0; i < b.N; i++ {
		//			_ = s.Deserialize(bs, &target)
		//		}
		//		b.Log(target)
		//	})
		//
		//	b.Run("encoding - decoding", func(b *testing.B) {
		//		var target MapTestData
		//		bs, _ := s.Serialize(msg)
		//		_ = s.Deserialize(bs, &target)
		//		b.Log(target)
		//
		//		for i := 0; i < b.N; i++ {
		//			bs, _ = s.Serialize(msg)
		//			_ = s.Deserialize(bs, &target)
		//		}
		//	})
		//})
	})
}
