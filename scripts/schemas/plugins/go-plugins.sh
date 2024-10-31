#!/usr/bin/env bash

#go mod tidy
go get -u google.golang.org/grpc \
	google.golang.org/protobuf/proto \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	google.golang.org/protobuf/reflect/protoreflect \
	google.golang.org/protobuf/runtime/protoimpl

#go install google.golang.org/grpc \
#	google.golang.org/protobuf/proto \
#	google.golang.org/protobuf/cmd/protoc-gen-go \
#	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
#	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
#	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
#	google.golang.org/protobuf/reflect/protoreflect \
#    google.golang.org/protobuf/runtime/protoimpl

#go mod tidy
go mod vendor
