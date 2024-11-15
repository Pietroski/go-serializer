## Go Serializer

### Profile to Identify Bottlenecks

It might help to profile your code to pinpoint where the exact performance bottleneck lies. The Go profiler can reveal
if the primary cost is in reflect.Index, AddUint64, or bbw.write. Here’s a simple way to profile:

```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

This can provide more insight, allowing you to concentrate on the specific bottleneck.

For more advanced analysis, you can use Go's built-in profiling tools:

Generate profile: go test -bench=. -cpuprofile=cpu.prof
Analyze with pprof: go tool pprof cpu.prof

To run the benchmark and get results:

Save the file as main_test.go
Run in terminal: go test -bench=.

You can get more detailed results with these flags:

Memory allocation stats: go test -bench=. -benchmem
CPU profile: go test -bench=. -cpuprofile=cpu.out
Memory profile: go test -bench=. -memprofile=mem.out
Count (run multiple times): go test -bench=. -count=5

## Benchmark results

- struct

    - item sample

        - proto serialzier:
            - encoding:
                - ops count: 10_515_591
                - avg op time: 112.0 ns/op
            - decoding:
                - ops count: 7_653_538
                - avg op time: 154.3 ns/op
            - encoding - decoding:
                - ops count: 4_319_391
                - avg op time: 274.5 ns/op

        - binary serializer:
            - encoding:
                - ops count: 12_578_024
                - avg op time: 94.77 ns/op
            - decoding:
                - ops count: 10_895_841
                - avg op time: 108.9 ns/op
            - encoding - decoding:
                - ops count: 5_446_920
                - avg op time: 220.0 ns/op

    - item sample - nil sub item

        - proto serialzier:
            - encoding:
                - ops count: 13963114
                - avg op time: 72.22 ns/op
            - decoding:
                - ops count: 16776402
                - avg op time: 69.48 ns/op
            - encoding - decoding:
                - ops count: 8330821
                - avg op time: 142.1 ns/op

        - binary serializer:
            - encoding:
                - ops count: 24927391
                - avg op time: 48.42 ns/op
            - decoding:
                - ops count: 31225570
                - avg op time: 37.83 ns/op
            - encoding - decoding:
                - ops count: 13172132
                - avg op time: 90.08 ns/op
