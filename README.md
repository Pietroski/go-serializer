# Serializer

## Quick Warning

The commit history will be deleted at some point in time (because this library has been reorganised). The only reason
this repo's commit history will be kept for now is because I'll need it in order to go back in time and see some of the
changes that will be used to my academic paper.

## Intro

Serializer provides a way to serialize and deserialize data into bytes to it can be transferred over the internet or be
written and read into a files.

The icing on the top of this library is the Binary serializer. It implements the `Serializer` interface but also extends
its methods to the default `encoding` methods such as `Marshal` and `Unmarshal`. The binary serializer comes in three
different flavours (whose benchmarks can be visited; more details at [Benchmark results](#benchmark-results)),
`serializer.BinarySerializer`, `serializer.RawBinarySerializer` and the extra `serializerx.BinarySerializer`.

`serializer.BinarySerializer` is the Default one that stays 100% under the golang's safe helm. It is at least as fast as
the `proto` serializer with a plus that is does not use any `unsafe` operations compared to `proto`, `msgpack` and other
serializers. Yet, it is as least as fast as `proto`'s.

`serializerx.BinarySerializer` is similar to `serializer.BinarySerializer`, but under the hoods it does do `unsafe`
operations under the `reflectx` package. Per se, the `reflect` package heavily uses `unsafe` operations in a highly
safe, secure and controlled way and environment. `reflectx` behaves as a `reflect` extension that would not be
incorporated by the Go programming language as referred [here](https://github.com/golang/go/issues/70267). However, it
than extends the `reflect` interface adding specific methods that keeps all the `unsafe` operations constrained to a
single place and package; leaving it easier to be maintained although if minimal performance cost. In order to reduce
`reflectx`'s performance cost, it skips a few type checks that are done in the standard `reflect` package because and
only because all the necessary pre-checks and calls are done within the `serializerx.BinarySerializer` package and since
the `reflectx` package is internal and will remain so, there is no point on repeating the necessary type checks at this
point.

Finally, the `serializer.RawBinarySerializer` directly calls and do the `unsafe` operations `rawly` without the
`reflectx` encapsulation that although minimally, comes at a cost. These raw and `unsafe` calls provide about 0-3%
better performance than the code encapsulated by `reflectx` package.

The Binary Serializers expose the `Serializer` and standard `encoding` interface methods to give usage flexibility in
different case scenarios and also to keep standard compatibility with Go's `stdlib`.
As a plus, the `Serializer` lib encapsulates the stdlib interface into its own `Serializer` interface to keep package
consistency.

## Benchmark results

All The benchmark results can be found under `./tests/benchmarks/serializer/results`.
The results are available in a few
formats, [json](./tests/benchmarks/serializer/results/benchmark_results.yaml),
[yaml](./tests/benchmarks/serializer/results/beautiful_benchmark_results.json)
and from the json file, `html` results are displayed and generated in bar charts.

The benchmark results can be found
here [tests/benchmarks/serializer/results/index.html](./tests/benchmarks/serializer/results/index.html)

The benchmark results with the Gob serializer were omitted because of its horrible performance. But the results can
still be found
here [tests/benchmarks/serializer/results/with_gob/index.html](./tests/benchmarks/serializer/results/with_gob/index.html)

The charts can also be seen at -> https://go-serializer-benchmarks.tiiny.site/

### General benchmark tips and hints on profiling

Ways to benchmark you application

Generate profile:

```bash
# Generate CPU profile:
go test -bench=. -cpuprofile=cpu.prof
# Memory allocation stats:
go test -bench=. -benchmem
# or similarly to CPU profiling:
go test -bench=. -memprofile=mem.out

# Count (run multiple times)
go test -bench=. -count=5

# Analyze with pprof
go tool pprof cpu.prof
```

## Thoughts and Conclusions

As seen, the `pietroski/go-serializer` library provides extremely, complex, safe, reliable and a powerful way to
convert a wide set of data types into bytes (consequently binary). This is incredibly beneficial to inter-service
communication (communication between services). This library, breaks through the paradigm of performance and reliability
into transforming transferable data and safely challenges the big players. It gives a new perspective on the encoding
world with a new and fresh pair of eyes.
