package serializer

import (
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBenchResults(t *testing.T) {

	matches := parseBenchmarkResults(t)
	_ = matches
	t.Log(matches)

	//t.Log("benchmarkedFunction,ops_count,avg_op_time,alloc_size,allocs")
	//for _, m := range matches {
	//	t.Logf("%s,%s,%s,%s,%s\n", m[1], m[2], m[3], m[4], m[5])
	//}
}

type (
	ParsedBenchmarkResults []ParsedBenchmarkResult

	ParsedBenchmarkResult struct {
		TestName   string
		OpsCount   uint64
		AvgOpTime  string
		AllocSize  string
		AllocCount string
	}
)

type BenchTest interface {
	require.TestingT
	*testing.T | *testing.B
}

func parseBenchmarkResults[T BenchTest](tb T) ParsedBenchmarkResults {
	b, err := os.ReadFile("results/BenchmarkAll.log")
	require.NoError(tb, err)

	benchmarkResult := string(b)
	regexBench := regexp.MustCompile(
		`([a-zA-Z0-9/_\-\]\[]+)+\s+(\d+)+\s+(\d+.?\d+?\s+[a-z]+/op)+\s+(\d+\s+[a-zA-Z]/op)+\s+(\d+\s+allocs/op)`)
	matches := regexBench.FindAllStringSubmatch(benchmarkResult, -1)

	parsedBenchmarkResults := make(ParsedBenchmarkResults, len(matches))
	for idx, match := range matches {
		i64, err := strconv.ParseUint(match[2], 10, 64)
		assert.NoError(tb, err)
		parsedBenchmarkResults[idx] = ParsedBenchmarkResult{
			TestName:   match[1],
			OpsCount:   i64,
			AvgOpTime:  match[3],
			AllocSize:  match[4],
			AllocCount: match[5],
		}
	}

	return parsedBenchmarkResults
}

func BenchmarkTestTime(b *testing.B) {
	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			time.Sleep(time.Second)
		}
	})

	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			time.Sleep(time.Millisecond)
		}
	})

	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			time.Sleep(time.Microsecond)
		}
	})

	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			time.Sleep(time.Nanosecond)
		}
	})
}

func TestTime(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Log(time.Second)
	})

	t.Run("", func(t *testing.T) {
		t.Log(time.Millisecond)
	})

	t.Run("", func(t *testing.T) {
		t.Log(time.Microsecond)
	})

	t.Run("", func(t *testing.T) {
		t.Log(time.Nanosecond)
	})
}
