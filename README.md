# A containerized REST and gRPC service

A working template for creating a containerized Golang gRPC
service with a REST gateway, residing in a kubernetes cluster.

A small Golang gRPC client is also included, for testing.

## Generating the protobuf code

```shell
protoc \
-I. \
-I  $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
--grpc-gateway_out=logtostderr=true:. --grpc-gateway_opt=paths=source_relative \
proto/service.proto
```

## Running the server natively

```shell
go run main.go grpcServer
```

## Running the gRPC client natively

```shell
go run main.go grpcClient --word abcdef
```

### Example gRPC request and response

```shell
➜ go run main.go grpcClient --word abcdef
2022/10/19 15:49:43 echo:"abcdef" timestamp:"2022-10-19 20:49:43.293029 +0000 UTC"
```

## Running the REST proxy natively

```shell
go run main.go restProxy
```

### Example REST request and response

```shell
➜ curl http://localhost:8080/echo/abcdef
{"echo":"abcdef","timestamp":"2022-10-20 22:26:46.416726 +0000 UTC"}%
```

#### Using [JQ](https://github.com/stedolan/jq) to make the response more readable

```shell
➜ curl http://localhost:8080/echo/abcdefg | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    69  100    69    0     0   5978      0 --:--:-- --:--:-- --:--:-- 69000
{
  "echo": "abcdefg",
  "timestamp": "2022-10-20 22:29:21.706632 +0000 UTC"
}
```

## Running the gRPC server and REST proxy in Docker with docker-compose

```shell
docker-compose up --force-recreate --remove-orphans --build
```

### Example REST request and response using Docker

```shell
✗ curl http://localhost:8080/echo/abcdefg
{"echo":"abcdefg","timestamp":"2023-06-06 20:29:06.774947963 +0000 UTC"}%
```

#### Using [JQ](https://github.com/stedolan/jq) to make the response more readable with Docker

```shell
✗ curl http://localhost:8080/echo/abcdefg | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    72  100    72    0     0   4489      0 --:--:-- --:--:-- --:--:--  7200
{
  "echo": "abcdefg",
  "timestamp": "2023-06-06 20:25:26.109105521 +0000 UTC"
}
```
