
## 为什么要在容器化环境下测试

1. 此基准测试走网络,需要有方法排除网络因素的影响

2. 使用多协程的方式运行,引入并发

3. 使用k8s时间片limit功能限制程序对cpu的使用,压满cpu对应的处理能力以达到基准测试效果(即在某些cpu限制条件下,处理能力表现为线性增长)

---

### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.11

### Run tests

Run benchmarks:
```
GO111MODULE=on go test -bench=. -benchmem
```

### Results

```
goos: darwin
goarch: amd64
BenchmarkGRPCProtobuf-8            10000            117649 ns/op            7686 B/op        154 allocs/op
BenchmarkHTTPJSON-8                10000            105837 ns/op            8932 B/op        116 allocs/op
PASS
ok      github.com/plutov/benchmark-grpc-protobuf-vs-http-json  4.340s
```

They are almost the same, HTTP+JSON is a bit faster and has less allocs/op.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```
GO111MODULE=on go test -bench=BenchmarkGRPCProtobuf -cpuprofile=grpcprotobuf.cpu
GO111MODULE=on go test -bench=BenchmarkHTTPJSON -cpuprofile=httpjson.cpu
```

Check CPU usage per approach using:

```
go tool pprof grpcprotobuf.cpu
go tool pprof httpjson.cpu
```

My results show that Protobuf consumes less ressources, around **30% less**.

### gRPC definition

 - Install [Go](https://golang.org/dl/)
 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

```
protoc --go_out=plugins=grpc:. grpc-protobuf/proto/api.proto
```