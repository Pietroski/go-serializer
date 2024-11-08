## Go Serializer

### Profile to Identify Bottlenecks

It might help to profile your code to pinpoint where the exact performance bottleneck lies. The Go profiler can reveal
if the primary cost is in reflect.Index, AddUint64, or bbw.write. Here’s a simple way to profile:

```bash
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

This can provide more insight, allowing you to concentrate on the specific bottleneck.
