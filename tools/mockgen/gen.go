package mock_generator

import (
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -package mocked_serializer -destination ../../pkg/tools/serializer/mocks/serializer/mocked_serializer.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer Serializer
//go:generate mockgen -package mocked_beautifier -destination ../../pkg/tools/serializer/mocks/beautifier/mocked_beautifier.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer Beautifier

// mocks
// // proto
//go:generate mockgen -package go_mocked_proto -destination ../../internal/mocks/proto/mocks/mocked_proto_message.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/internal/mocks/proto ProtoMessage
// //go:generate mockgen -package go_mocked_proto -destination ../../internal/mocks/proto/mocks/mocked_proto_reflect_message.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/internal/mocks/proto ProtoReflectMessage
//go:generate mockgen -package go_mocked_proto -destination ../../internal/mocks/proto/mocks/mocked_proto_reflect_proto_message.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/internal/mocks/proto ProtoReflectProtoMessage
