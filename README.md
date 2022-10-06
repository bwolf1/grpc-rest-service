# A containerized REST and gRPC service

This is built to be used as a template for creating a containerized Golang gRPC
service with a REST gateway, metrics and logging.

A small Golang gRPC client is also included, for testing.

Built with:

- GO 1.17.7
- PostgreSQL
- Cobra
- Viper
- gRPC 1.48.0
- gRPC Gateway 1.16.0
- Zap
- Prometheus
- Grafana
- Docker
- Kubernetes

## Generating the gRPC and REST gateway code from the protobuf file

Note: I'm specifying version `1.16.0` here to keep my development environment
compatible with multiple projects (some of which require this specific version).
This is not required. I left it in only to demonstrate that it can be done, and
for my own convenience.

```shell
protoc \
-I. \
-I  $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
--grpc-gateway_out=logtostderr=true:. --grpc-gateway_opt=paths=source_relative \
proto/service.proto
```

## Running the server

```shell
go run main.go grpcServer
```

## Running the gRPC client

```shell
go run main.go grpcClient --word abcdef
```

### Example gRPC request and response

```shell
➜ go run main.go grpcClient --word abcdef
2022/10/19 15:49:43 echo:"abcdef" timestamp:"2022-10-19 20:49:43.293029 +0000 UTC"
```

## Running the REST proxy

```shell
go run main.go restProxy
```

### Example REST request and response

```shell
➜ curl http://localhost:50052/echo/abcdef
{"echo":"abcdef","timestamp":"2022-10-20 22:26:46.416726 +0000 UTC"}%
```

#### Using [JQ](https://github.com/stedolan/jq) to make the response more readable

```shell
➜ curl http://localhost:50052/echo/abcdefg | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    69    0     0   5978      0 --:--:-- --:--:-- --:--:-- 69000
{
  "echo": "abcdefg",
  "timestamp": "2022-10-20 22:29:21.706632 +0000 UTC"
}
```
