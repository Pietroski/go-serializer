# Schema scripts

## creates a swagger folder if does not exists
make-swagger-dir:
	@mkdir -p ./generated/swagger

## install only gRPC go plug-ins - it does nothing else
go-plugin-install-only:
	@./scripts/schemas/plugins/go-plugin-install-only.sh

## install all gRPC go plug-ins
go-plugins:
	@./scripts/schemas/plugins/go-plugins.sh

## generates protobuf
proto-gen: make-swagger-dir
	@./scripts/schemas/proto-gen/gen.sh
