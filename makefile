# TODO (bwolf1): Make running thie makefile part of the build when built via docker
.PHONY: build-service
## build-service: builds the service from the protobuf spec
build-service:
	@ protoc \
	-I. \
	-I  $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=logtostderr=true:. --grpc-gateway_opt=paths=source_relative \
	proto/service.proto

.PHONY: build-docker-image-multistage
## build-docker-image-multistage: builds a smaller docker image
build-docker-image-multistage:
	@ docker build . -t grpc-rest-kubernetes
	@ docker tag grpc-rest-kubernetes:latest grpc-rest-kubernetes:v0.0.1


.PHONY: run-docker-multistage
## run-docker-multistage: runs the server as a Docker container, using the smaller image
run-docker-multistage: build-docker-image-multistage
	@ docker run -p 50051:50051 grpc-rest-kubernetes
