package serializer

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/models"
)

func Benchmark(b *testing.B) {
	testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{
		Name: "TestBenchmarkAll",
		F:    TestBenchmarkAll,
	}})

	bs, err := exec.Command(
		"pwd",
	).Output()
	require.NoError(b, err)
	b.Log(string(bs))

	err = exec.Command(
		"cd", "../../..", "&&", "go", "test", "-bench", "BenchmarkAll",
		"-benchmem", "./...", "&>", "./tests/benchmarks/serializer/results/BenchmarkAll.log",
	).Run()
	require.NoError(b, err)
}

func BenchmarkAll(b *testing.B) {
	testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{
		Name: "TestBenchmarkAll",
		F:    TestBenchmarkAll,
	}})

	benchmarkData := reflect.ValueOf(BenchmarkData)

	benchmarkDataLimit := benchmarkData.NumField()
	for i := 0; i < benchmarkDataLimit; i++ {
		benchmarks := benchmarkData.Field(i)

		benchmarksLimit := benchmarks.NumField()
		for j := 0; j < benchmarksLimit; j++ {
			runner := benchmarks.Field(j)

			runnerLimit := runner.NumField()
			for k := 1; k < runnerLimit; k++ {
				dataType := runner.Field(k)

				dataTypeLimit := dataType.NumField()
				for l := 0; l < dataTypeLimit; l++ {
					testCases := dataType.Field(l)

					testCasesLimit := testCases.Len()
					for m := 0; m < testCasesLimit; m++ {
						testCase := testCases.Index(m)

						testData := testCase.FieldByName("TestData")
						msg := testData.FieldByName("Msg").Interface()
						target := testData.FieldByName("Target").Interface()

						caseName := fmt.Sprintf("%s/%s/%s",
							strings.Split(
								reflect.TypeOf(benchmarks.Interface()).Field(j).Tag.Get("json"),
								",")[0],
							strings.Split(
								reflect.TypeOf(dataType.Interface()).Field(l).Tag.Get("json"),
								",")[0],
							testCase.FieldByName("CaseName").String(),
						)

						b.Run(caseName, RunBenchmarkTestCase(runner, msg, target))
					}
				}
			}
		}
	}

	time.Sleep(time.Second * 1)
	testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{{
		Name: "TestWriteBenchmarkResults",
		F:    TestWriteBenchmarkResults,
	}})
}

func RunBenchmarkTestCase(
	runner reflect.Value, msg, target interface{},
) func(b *testing.B) {
	return func(b *testing.B) {
		sValue := runner.FieldByName("Serializer")
		if sValue.IsNil() {
			return
		}

		s := sValue.Interface().(models.Serializer)

		bs, err := s.Serialize(msg)
		require.NoError(b, err)
		err = s.Deserialize(bs, target)
		require.NoError(b, err)
		//require.EqualExportedValues(b, msg, target)

		switch msg.(type) {
		case bool, string,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			tgt := reflect.Indirect(reflect.ValueOf(target))
			assert.EqualValues(b, msg, tgt.Interface())
		default:
			assert.EqualExportedValues(b, msg, target)
		}

		b.Run("encoding", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.Serialize(msg)
			}
		})

		b.Run("decoding", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = s.Deserialize(bs, target)
			}
		})

		b.Run("rebind", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = s.Serialize(msg)
				_ = s.Deserialize(bs, target)
			}
		})
	}
}

func TestBenchmarkAll(t *testing.T) {
	benchmarkData := reflect.ValueOf(BenchmarkData)

	benchmarkDataLimit := benchmarkData.NumField()
	for i := 0; i < benchmarkDataLimit; i++ {
		benchmarks := benchmarkData.Field(i)

		benchmarksLimit := benchmarks.NumField()
		for j := 0; j < benchmarksLimit; j++ {
			runner := benchmarks.Field(j)

			runnerLimit := runner.NumField()
			for k := 1; k < runnerLimit; k++ {
				dataType := runner.Field(k)

				dataTypeLimit := dataType.NumField()
				for l := 0; l < dataTypeLimit; l++ {
					testCases := dataType.Field(l)

					testCasesLimit := testCases.Len()
					for m := 0; m < testCasesLimit; m++ {
						testCase := testCases.Index(m)

						testData := testCase.FieldByName("TestData")
						msg := testData.FieldByName("Msg").Interface()
						target := testData.FieldByName("Target").Interface()

						caseName := fmt.Sprintf("%s/%s/%s",
							strings.Split(
								reflect.TypeOf(benchmarks.Interface()).Field(j).Tag.Get("json"),
								",")[0],
							strings.Split(
								reflect.TypeOf(dataType.Interface()).Field(l).Tag.Get("json"),
								",")[0],
							testCase.FieldByName("CaseName").String(),
						)

						t.Run(caseName, RunTestBenchmarkTestCase(runner, msg, target))
					}
				}
			}
		}
	}
}

func RunTestBenchmarkTestCase(
	runner reflect.Value, msg, target interface{},
) func(t *testing.T) {
	return func(t *testing.T) {
		sValue := runner.FieldByName("Serializer")
		if sValue.IsNil() {
			return
		}

		s := sValue.Interface().(models.Serializer)

		bs, err := s.Serialize(msg)
		require.NoError(t, err)
		err = s.Deserialize(bs, target)
		require.NoError(t, err)

		switch msg.(type) {
		case bool, string,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			tgt := reflect.Indirect(reflect.ValueOf(target))
			assert.EqualValues(t, msg, tgt.Interface())
		default:
			assert.EqualExportedValues(t, msg, target)
		}
	}
}
