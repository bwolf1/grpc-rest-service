# A containerized REST and GRPC service

This should serve as a good template for creating a containerized GRPC service
in Golang with a REST gateway, metrics, logging and tracing.

A small GRPC client is also included, for testing.

Built with:

- GO 1.17.7
- PostgreSQL
- Cobra
- Viper
- GRPC 1.48.0
- Zap
- Jaeger
- Prometheus
- Grafana
- Docker
- Kubernetes

## Generating the RPC code from the proto file

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto
```

## Running the server

```shell
go run server/main.go
```

## Running the client

```shell
go run client/main.go --word <word-to-echo>
```

## Example

```shell
✗ go run client/main.go --word hello
2022/07/31 12:20:55 "hello" echoed back as "hello"
```
