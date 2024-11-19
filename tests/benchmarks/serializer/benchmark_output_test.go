//go:build benchmark

package serializer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"gitlab.com/pietroski-software-company/devex/golang/serializer/internal/strconvx"
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

	//bs, err := json.MarshalIndent(seriesData, "", "  ")
	//require.NoError(t, err)
	//t.Log(string(bs))

	//debuggerFile, err := os.OpenFile(
	//	fmt.Sprintf("results/charts/debugger_results.json"),
	//	os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	//require.NoError(t, err)

	fullPageReport := components.NewPage()
	fullPageReportWithGob := components.NewPage()
	defer func() {
		fullChartReportFile, err := os.OpenFile(
			fmt.Sprintf("results/charts/benchmark_chart_results.html"),
			os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		require.NoError(t, err)
		defer func(file *os.File) { _ = file.Close() }(fullChartReportFile)

		fullChartReportFileWithGob, err := os.OpenFile(
			fmt.Sprintf("results/with_gob/charts/benchmark_chart_results.html"),
			os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		require.NoError(t, err)
		defer func(file *os.File) { _ = file.Close() }(fullChartReportFileWithGob)

		err = fullPageReport.Render(fullChartReportFile)
		require.NoError(t, err)
		err = fullPageReportWithGob.Render(fullChartReportFileWithGob)
		require.NoError(t, err)
	}()

	for quantifierName, quantifierValue := range seriesData {
		for dataTypeName, dataTypeResult := range quantifierValue {
			err := os.MkdirAll(fmt.Sprintf("results/charts/%s", dataTypeName), 0750)
			require.NoError(t, err)
			err = os.MkdirAll(fmt.Sprintf("results/with_gob/charts/%s", dataTypeName), 0750)
			require.NoError(t, err)

			chartFile, err := os.OpenFile(
				fmt.Sprintf("results/charts/%s/benchmark_chart_results_%s.html", dataTypeName, quantifierName),
				os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			require.NoError(t, err)
			chartFileWithGob, err := os.OpenFile(
				fmt.Sprintf("results/with_gob/charts/%s/benchmark_chart_results_%s.html",
					dataTypeName, quantifierName),
				os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			require.NoError(t, err)

			page := components.NewPage()
			pageWithGob := components.NewPage()

			for caseName, caseResult := range dataTypeResult {

				bar := charts.NewBar()
				barWithGob := charts.NewBar()
				chartGlobalOpts := []charts.GlobalOpts{
					charts.WithTitleOpts(opts.Title{
						Title: fmt.Sprintf(
							"Benchmark Chart - Serializer - %s - %s - %s",
							dataTypeName, quantifierName, caseName),
						Subtitle: "serializer types benchmark results",
					}),
					charts.WithLegendOpts(opts.Legend{
						Top: "50",
					}),
				}
				bar.SetGlobalOptions(chartGlobalOpts...)
				xAxis := bar.SetXAxis([]string{"encoding", "decoding", "rebind"})
				barWithGob.SetGlobalOptions(chartGlobalOpts...)
				xAxisWithGob := barWithGob.SetXAxis([]string{"encoding", "decoding", "rebind"})

				for serializerName, serializerResult := range caseResult {
					name := fmt.Sprintf("%s_%s_%s", serializerName, dataTypeName, caseName)

					seriesOpts := charts.WithItemStyleOpts(opts.ItemStyle{})
					if serializerName == "proto" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(84,112,198)"})
					}
					if serializerName == "binary" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(145,204,117)"})
					}
					if serializerName == "raw_binary" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(242, 202, 107)"})
					}
					if serializerName == "x_binary" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(238, 102, 102)"})
					}
					if serializerName == "msgpack" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(236, 138, 93)"})
					}
					if serializerName == "json" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(133, 190, 219)"})
					}
					if serializerName == "gob" {
						seriesOpts = charts.WithItemStyleOpts(opts.ItemStyle{Color: "rgb(89, 160, 118)"})
						xAxisWithGob.AddSeries(name, serializerResult, seriesOpts)
						continue
					}

					xAxisWithGob.AddSeries(name, serializerResult, seriesOpts)
					xAxis.AddSeries(name, serializerResult, seriesOpts)
				}

				page.AddCharts(bar)
				fullPageReport.AddCharts(bar)
				pageWithGob.AddCharts(barWithGob)
				fullPageReportWithGob.AddCharts(barWithGob)
			}

			err = page.Render(chartFile)
			require.NoError(t, err)
			require.NoError(t, chartFile.Close())

			err = pageWithGob.Render(chartFileWithGob)
			require.NoError(t, err)
			require.NoError(t, chartFileWithGob.Close())
		}
	}
}
