#!/usr/bin/env bash

MOD_NAME=gitlab.com/pietroski-software-company/tools/serializer/go-serializer/protos/schemas/generated
PROTO_PATH=protos
OUTPUT_PATH=generated
SWAGGER_PATH="$OUTPUT_PATH/swagger"

declare -a proto_list=(
	"item"
	"item-error"

	"data/graph"
)

for domain in "${proto_list[@]}"; do
	protoc --proto_path="$PROTO_PATH" "$PROTO_PATH"/"${domain}"/*.proto \
		--go_out=:"$OUTPUT_PATH" \
		--go_opt=module="$MOD_NAME" \
		--go-grpc_out=:"$OUTPUT_PATH" \
		--go-grpc_opt=module="$MOD_NAME" \
		--grpc-gateway_out=:"$OUTPUT_PATH" \
		--grpc-gateway_opt=module="$MOD_NAME" \
		--openapiv2_out=:"$SWAGGER_PATH"
done

#go mod tidy
#go mod vendor
