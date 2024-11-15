package serializer

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/strconvx"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/require"
)

func TestWriteBenchmarkResults(t *testing.T) {
	writeBenchmarkResults(t)
}

func writeBenchmarkResults(t *testing.T) {
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

						caseName := fmt.Sprintf("%s/%s/%s/%s", "BenchmarkAll",
							strings.Split(
								reflect.TypeOf(benchmarks.Interface()).Field(j).Tag.Get("json"),
								",")[0],
							strings.Split(
								reflect.TypeOf(dataType.Interface()).Field(l).Tag.Get("json"),
								",")[0],
							strings.ReplaceAll(testCase.FieldByName("CaseName").String(), " ", "_"),
						)

						testResults := testCase.FieldByName("TestResults")
						encoding := testResults.FieldByName("Encoding")
						decoding := testResults.FieldByName("Decoding")
						dataRebind := testResults.FieldByName("DataRebind")

						parsedResults := parseBenchmarkResults[*testing.T](t)
						for _, pr := range parsedResults {
							// t.Log(pr.TestName, caseName, pr.TestName == caseName, strings.Contains(pr.TestName, caseName))

							if strings.Contains(pr.TestName, caseName+"/rebind-12") {
								dataRebind.FieldByName("OpsCount").SetUint(pr.OpsCount)
								dataRebind.FieldByName("AvgOpTime").SetString(pr.AvgOpTime)
								dataRebind.FieldByName("AllocSize").SetString(pr.AllocSize)
								dataRebind.FieldByName("Allocs").SetString(pr.AllocCount)
							} else if strings.Contains(pr.TestName, caseName+"/encoding-12") {
								encoding.FieldByName("OpsCount").SetUint(pr.OpsCount)
								encoding.FieldByName("AvgOpTime").SetString(pr.AvgOpTime)
								encoding.FieldByName("AllocSize").SetString(pr.AllocSize)
								encoding.FieldByName("Allocs").SetString(pr.AllocCount)
							} else if strings.Contains(pr.TestName, caseName+"/decoding-12") {
								decoding.FieldByName("OpsCount").SetUint(pr.OpsCount)
								decoding.FieldByName("AvgOpTime").SetString(pr.AvgOpTime)
								decoding.FieldByName("AllocSize").SetString(pr.AllocSize)
								decoding.FieldByName("Allocs").SetString(pr.AllocCount)
							}
						}
					}
				}
			}
		}
	}

	yamlFile, err := os.OpenFile("results/benchmark_results.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer func() { _ = yamlFile.Close() }()

	err = yaml.NewEncoder(yamlFile).Encode(BenchmarkData)
	require.NoError(t, err)

	jsonFile, err := os.OpenFile("results/benchmark_results.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer func() { _ = jsonFile.Close() }()

	err = json.NewEncoder(jsonFile).Encode(BenchmarkData)
	require.NoError(t, err)

	beautifulJsonFile, err :=
		os.OpenFile("results/beautiful_benchmark_results.json",
			os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer func() { _ = jsonFile.Close() }()

	bs, err := json.MarshalIndent(BenchmarkData, "", "  ")
	require.NoError(t, err)

	_, err = beautifulJsonFile.Write(bs)
	require.NoError(t, err)

	time.Sleep(time.Second * 1)
	createBenchmarkChartResults(t)
}

func createBenchmarkChartResults(t *testing.T) {
	seriesData := make(map[string]map[string]map[string]map[string][]opts.BarData)

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

						serializerName := strings.Split(
							reflect.TypeOf(benchmarks.Interface()).Field(j).Tag.Get("json"),
							",")[0]

						dataTypeName := strings.Split(
							reflect.TypeOf(dataType.Interface()).Field(l).Tag.Get("json"),
							",")[0]

						caseName := testCase.FieldByName("CaseName").String()

						casePathName := fmt.Sprintf("%s/%s/%s/%s", "BenchmarkAll",
							strings.Split(
								reflect.TypeOf(benchmarks.Interface()).Field(j).Tag.Get("json"),
								",")[0],
							strings.Split(
								reflect.TypeOf(dataType.Interface()).Field(l).Tag.Get("json"),
								",")[0],
							strings.ReplaceAll(testCase.FieldByName("CaseName").String(), " ", "_"),
						)

						// t.Log(caseName)

						parsedResults := parseBenchmarkResults[*testing.T](t)
						for _, pr := range parsedResults {
							opsCount := pr.OpsCount
							avgOpTime, err := strconvx.StrTimeToFloat64(pr.AvgOpTime)
							require.NoError(t, err)
							allocSize, err := strconvx.StrSizeToFloat64(pr.AllocSize)
							require.NoError(t, err)
							allocCount, err := strconvx.StrAllocCountToFloat64(pr.AllocCount)
							require.NoError(t, err)

							{
								// #####################################################################################

								{
									_, ok := seriesData["opsCount"]
									if !ok {
										seriesData["opsCount"] = map[string]map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["opsCount"][dataTypeName]
									if !ok {
										seriesData["opsCount"][dataTypeName] = map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["opsCount"][dataTypeName][caseName]
									if !ok {
										seriesData["opsCount"][dataTypeName][caseName] = map[string][]opts.BarData{}
									}

									_, ok = seriesData["opsCount"][dataTypeName][caseName][serializerName]
									if !ok {
										seriesData["opsCount"][dataTypeName][caseName][serializerName] =
											make([]opts.BarData, 3)
									}
								}

								// #####################################################################################

								{
									_, ok := seriesData["avgOpTime"]
									if !ok {
										seriesData["avgOpTime"] = map[string]map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["avgOpTime"][dataTypeName]
									if !ok {
										seriesData["avgOpTime"][dataTypeName] = map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["avgOpTime"][dataTypeName][caseName]
									if !ok {
										seriesData["avgOpTime"][dataTypeName][caseName] = map[string][]opts.BarData{}
									}

									_, ok = seriesData["avgOpTime"][dataTypeName][caseName][serializerName]
									if !ok {
										seriesData["avgOpTime"][dataTypeName][caseName][serializerName] =
											make([]opts.BarData, 3)
									}
								}

								// #####################################################################################

								{
									_, ok := seriesData["allocSize"]
									if !ok {
										seriesData["allocSize"] = map[string]map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocSize"][dataTypeName]
									if !ok {
										seriesData["allocSize"][dataTypeName] = map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocSize"][dataTypeName][caseName]
									if !ok {
										seriesData["allocSize"][dataTypeName][caseName] = map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocSize"][dataTypeName][caseName][serializerName]
									if !ok {
										seriesData["allocSize"][dataTypeName][caseName][serializerName] =
											make([]opts.BarData, 3)
									}
								}

								// #####################################################################################

								{
									_, ok := seriesData["allocCount"]
									if !ok {
										seriesData["allocCount"] = map[string]map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocCount"][dataTypeName]
									if !ok {
										seriesData["allocCount"][dataTypeName] = map[string]map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocCount"][dataTypeName][caseName]
									if !ok {
										seriesData["allocCount"][dataTypeName][caseName] = map[string][]opts.BarData{}
									}

									_, ok = seriesData["allocCount"][dataTypeName][caseName][serializerName]
									if !ok {
										seriesData["allocCount"][dataTypeName][caseName][serializerName] =
											make([]opts.BarData, 3)
									}
								}

								// #####################################################################################
							}

							{
								//t.Log(strings.Contains(pr.TestName, caseName+"/rebind-12"), pr.TestName, caseName+"/rebind-12")
								if strings.Contains(pr.TestName, casePathName+"/rebind-12") {
									seriesData["opsCount"][dataTypeName][caseName][serializerName][2] =
										opts.BarData{Value: opsCount}
									seriesData["avgOpTime"][dataTypeName][caseName][serializerName][2] =
										opts.BarData{Value: avgOpTime}
									seriesData["allocSize"][dataTypeName][caseName][serializerName][2] =
										opts.BarData{Value: allocSize}
									seriesData["allocCount"][dataTypeName][caseName][serializerName][2] =
										opts.BarData{Value: allocCount}
								} else if strings.Contains(pr.TestName, casePathName+"/encoding-12") {
									seriesData["opsCount"][dataTypeName][caseName][serializerName][1] =
										opts.BarData{Value: opsCount}
									seriesData["avgOpTime"][dataTypeName][caseName][serializerName][1] =
										opts.BarData{Value: avgOpTime}
									seriesData["allocSize"][dataTypeName][caseName][serializerName][1] =
										opts.BarData{Value: allocSize}
									seriesData["allocCount"][dataTypeName][caseName][serializerName][1] =
										opts.BarData{Value: allocCount}
								} else if strings.Contains(pr.TestName, casePathName+"/decoding-12") {
									seriesData["opsCount"][dataTypeName][caseName][serializerName][0] =
										opts.BarData{Value: opsCount}
									seriesData["avgOpTime"][dataTypeName][caseName][serializerName][0] =
										opts.BarData{Value: avgOpTime}
									seriesData["allocSize"][dataTypeName][caseName][serializerName][0] =
										opts.BarData{Value: allocSize}
									seriesData["allocCount"][dataTypeName][caseName][serializerName][0] =
										opts.BarData{Value: allocCount}
								}
							}
						}
					}
				}
			}
		}
	}

	//t.Log(seriesData)

	bs, err := json.MarshalIndent(seriesData, "", "  ")
	require.NoError(t, err)
	t.Log(string(bs))

	chartFile, err := os.OpenFile(
		"results/charts/benchmark_chart_results.html",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer func() { _ = chartFile.Close() }()

	page := components.NewPage()

	for quantifierName, quantifierValue := range seriesData {
		for dataTypeName, dataTypeResult := range quantifierValue {
			for caseName, caseResult := range dataTypeResult {
				bar := charts.NewBar()
				bar.SetGlobalOptions(
					charts.WithTitleOpts(opts.Title{
						Title: fmt.Sprintf(
							"Benchmark Chart - Serializer - %s - %s - %s",
							dataTypeName, quantifierName, caseName),
						Subtitle: "serializer types benchmark results",
						//Top:         "50",
					}),
					charts.WithLegendOpts(opts.Legend{
						Top: "50",
						//Bottom:        "",
					}),
				)
				xAxis := bar.SetXAxis([]string{"encoding", "decoding", "rebind"}) // *1

				//xAxis := bar.SetXAxis(map[string][]string{
				//	"opsCount":   {"encoding", "decoding", "rebind"},
				//	"avgOpTime":  {"encoding", "decoding", "rebind"},
				//	"allocSize":  {"encoding", "decoding", "rebind"},
				//	"allocCount": {"encoding", "decoding", "rebind"},
				//}) // *1

				for serializerName, serializerResult := range caseResult {
					name := fmt.Sprintf("%s_%s_%s", serializerName, dataTypeName, caseName)

					//seriesOpts := charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(0, 0, 0)"})
					//if serializerName == "x_binary" {
					//	seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(238, 102, 102)"})
					//}
					//if serializerName == "proto" {
					//	seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(84, 112, 198)"})
					//} // 145,204,117 (green) |

					xAxis.AddSeries(name, serializerResult) // seriesOpts
				}

				page.AddCharts(bar)
			}
		}
	}

	err = page.Render(chartFile)
	require.NoError(t, err)
}
