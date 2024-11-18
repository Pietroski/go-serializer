package serializer

import (
	"math"
	"time"

	"gitlab.com/pietroski-software-company/devex/golang/serializer"
	grpc_item "gitlab.com/pietroski-software-company/devex/golang/serializer/internal/generated/go/pkg/item"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/testmodels"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/serializerx"
)

type (
	BenchmarkRunner struct {
		Benchmarks Benchmarks `json:"benchmarks,omitempty" yaml:"benchmarks,omitempty"`
	}

	Benchmarks struct {
		Proto     Runner `json:"proto,omitempty" yaml:"proto,omitempty"`
		Binary    Runner `json:"binary,omitempty" yaml:"binary,omitempty"`
		RawBinary Runner `json:"raw_binary,omitempty" yaml:"raw_binary,omitempty"`
		XBinary   Runner `json:"x_binary,omitempty" yaml:"x_binary,omitempty"`
		MsgPack   Runner `json:"msgpack,omitempty" yaml:"msgpack,omitempty"`
		Json      Runner `json:"json,omitempty" yaml:"json,omitempty"`
		Gob       Runner `json:"gob,omitempty" yaml:"gob,omitempty"`
	}

	Runner struct {
		Serializer models.Serializer `json:"-" yaml:"-"`
		DataType   DataType          `json:"data_type,omitempty" yaml:"data_type,omitempty"`
		ExtraCases DataType          `json:"extra_cases,omitempty" yaml:"extra_cases,omitempty"`
	}

	DataType struct {
		String []TestCase `json:"string,omitempty" yaml:"string,omitempty"`
		Number []TestCase `json:"number,omitempty" yaml:"number,omitempty"`
		Struct []TestCase `json:"struct,omitempty" yaml:"struct,omitempty"`
		Slice  []TestCase `json:"slice,omitempty" yaml:"slice,omitempty"`
		Map    []TestCase `json:"map,omitempty" yaml:"map,omitempty"`
	}

	TestCase struct {
		CaseName    string      `json:"case_name,omitempty" yaml:"case_name,omitempty"`
		TestResults TestResults `json:"test_results,omitempty" yaml:"test_results,omitempty"`
		TestData    TestData    `json:"-" yaml:"-"`
	}

	TestResults struct {
		Encoding   Result `json:"encoding,omitempty" yaml:"encoding,omitempty"`
		Decoding   Result `json:"decoding,omitempty" yaml:"decoding,omitempty"`
		DataRebind Result `json:"data_rebind,omitempty" yaml:"data_rebind,omitempty"`
	}

	Result struct {
		OpsCount  uint64 `json:"ops_count,omitempty" yaml:"ops_count,omitempty"`
		AvgOpTime string `json:"avg_op_time,omitempty" yaml:"avg_op_time,omitempty"`
		AllocSize string `json:"alloc_size,omitempty" yaml:"alloc_size,omitempty"`
		Allocs    string `json:"allocs,omitempty" yaml:"allocs,omitempty"`
	}

	TestData struct {
		Msg    interface{} `json:"-" yaml:"-"`
		Target interface{} `json:"-" yaml:"-"`
	}
)

var BenchmarkData = BenchmarkRunner{
	Benchmarks: Benchmarks{
		Proto: Runner{
			Serializer: serializer.NewProtoSerializer(),
			DataType: DataType{
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &grpc_item.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &grpc_item.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(grpc_item.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &grpc_item.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(grpc_item.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &grpc_item.SimplifiedSpecialStructTestData{
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
							},
							Target: new(grpc_item.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &grpc_item.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(grpc_item.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &grpc_item.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(grpc_item.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &grpc_item.Int64SliceTestData{
								Int64List: []int64{
									-math.MaxInt64, -9223372036854775808, -0, 0, 2, 12345678, 4, 5, 5170, 10, 8,
									87654321, 9223372036854775807, math.MaxInt64,
								},
							},
							Target: new(grpc_item.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &grpc_item.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(grpc_item.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &grpc_item.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(grpc_item.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &grpc_item.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(grpc_item.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &grpc_item.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(grpc_item.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &grpc_item.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(grpc_item.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &grpc_item.MapInt64Int64TestData{
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
							},
							Target: new(grpc_item.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &grpc_item.MapInt64StructPointerTestData{
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
							},
							Target: new(grpc_item.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &grpc_item.MapStringStructPointerTestData{
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
							},
							Target: new(grpc_item.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		Binary: Runner{
			Serializer: serializer.NewBinarySerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		RawBinary: Runner{
			Serializer: serializer.NewRawBinarySerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		XBinary: Runner{
			Serializer: serializerx.NewBinarySerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		MsgPack: Runner{
			Serializer: serializer.NewMsgPackSerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		Json: Runner{
			Serializer: serializer.NewJsonSerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
		Gob: Runner{
			Serializer: serializer.NewGobSerializer(),
			DataType: DataType{
				String: []TestCase{
					{
						CaseName: "complex string",
						TestData: TestData{
							Msg:    "test-again#$çcçá",
							Target: new(string),
						},
					},
				},
				Number: []TestCase{
					{
						CaseName: "int",
						TestData: TestData{
							Msg:    int(1_500_700_429),
							Target: new(int),
						},
					},
					{
						CaseName: "int64",
						TestData: TestData{
							Msg:    int64(1_500_700_095),
							Target: new(int64),
						},
					},
					{
						CaseName: "uint",
						TestData: TestData{
							Msg:    uint(1_500_700_070),
							Target: new(uint),
						},
					},
					{
						CaseName: "uint64",
						TestData: TestData{
							Msg:    uint64(1_500_700_787),
							Target: new(uint64),
						},
					},
				},
				Struct: []TestCase{
					{
						CaseName: "item sample",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
								SubItem: &testmodels.SubItem{
									Date:     time.Now().Unix(),
									Amount:   1_000_000_000,
									ItemCode: "code-status",
								},
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "item sample - nil sub item",
						TestData: TestData{
							Msg: &testmodels.Item{
								Id:     "any-item",
								ItemId: 100,
								Number: 5_000_000_000,
							},
							Target: new(testmodels.Item),
						},
					},
					{
						CaseName: "simplified special struct test data",
						TestData: TestData{
							Msg: &testmodels.SimplifiedSpecialStructTestData{
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
							},
							Target: new(testmodels.SimplifiedSpecialStructTestData),
						},
					},
					{
						CaseName: "string struct only",
						TestData: TestData{
							Msg: &testmodels.StringStruct{
								FirstString:  "first string value",
								SecondString: "second string value",
								ThirdString:  "third string value",
								FourthString: "fourth string value",
								FifthString:  "fifth string value",
							},
							Target: new(testmodels.StringStruct),
						},
					},
					{
						CaseName: "int64 struct only",
						TestData: TestData{
							Msg: &testmodels.Int64Struct{
								FirstInt64:  math.MaxInt64,
								SecondInt64: -math.MaxInt64,
								ThirdInt64:  math.MaxInt64,
								FourthInt64: -math.MaxInt64,
								FifthInt64:  0,
								SixthInt64:  -0,
							},
							Target: new(testmodels.Int64Struct),
						},
					},
				},
				Slice: []TestCase{
					{
						CaseName: "[]int64",
						TestData: TestData{
							Msg: &testmodels.Int64SliceTestData{
								Int64List: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
							},
							Target: new(testmodels.Int64SliceTestData),
						},
					},
					{
						CaseName: "[]uint64",
						TestData: TestData{
							Msg: &testmodels.Uint64SliceTestData{
								Uint64List: []uint64{
									-0, 0, 2, 12345678, 4, 5, 5170, 10, 8, 87654321,
									9223372036854775807, 18446744073709551615, math.MaxInt64, math.MaxUint64,
								},
							},
							Target: new(testmodels.Uint64SliceTestData),
						},
					},
					{
						CaseName: "[]string",
						TestData: TestData{
							Msg: &testmodels.StringSliceTestData{
								StringList: []string{"first-item", "second-item", "third-item", "fourth-item"},
							},
							Target: new(testmodels.StringSliceTestData),
						},
					},
					{
						CaseName: "[]byte",
						TestData: TestData{
							Msg: &testmodels.ByteSliceTestData{
								ByteList: []byte{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
							},
							Target: new(testmodels.ByteSliceTestData),
						},
					},
					{
						CaseName: "[][]byte",
						TestData: TestData{
							Msg: &testmodels.ByteByteSliceTestData{
								ByteByteList: [][]byte{
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
									{math.MaxUint8, -0, 0, 4, 5, 100, 8, 127, 255, math.MaxInt8, math.MaxUint8},
								},
							},
							Target: new(testmodels.ByteByteSliceTestData),
						},
					},
				},
				Map: []TestCase{
					{
						CaseName: "map[string]string",
						TestData: TestData{
							Msg: &testmodels.MapStringStringTestData{
								MapStringString: map[string]string{
									"any-key":       "any-value",
									"any-other-key": "any-other-value",
									"another-key":   "another-value",
								},
							},
							Target: new(testmodels.MapStringStringTestData),
						},
					},
					{
						CaseName: "map[int64]int64",
						TestData: TestData{
							Msg: &testmodels.MapInt64Int64TestData{
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
							},
							Target: new(testmodels.MapInt64Int64TestData),
						},
					},
					{
						CaseName: "map[int64]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapInt64StructPointerTestData{
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
							},
							Target: new(testmodels.MapInt64StructPointerTestData),
						},
					},
					{
						CaseName: "map[string]*StructTestData",
						TestData: TestData{
							Msg: &testmodels.MapStringStructPointerTestData{
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
							},
							Target: new(testmodels.MapStringStructPointerTestData),
						},
					},
				},
			},
		},
	},
}
