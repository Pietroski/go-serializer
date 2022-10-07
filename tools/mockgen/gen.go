package mock_generator

import (
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -package mocked_serializer -destination ../../pkg/tools/serializer/mocks/serializer/mocked_serializer.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer Serializer
//go:generate mockgen -package mocked_beautifier -destination ../../pkg/tools/serializer/mocks/beautifier/mocked_beautifier.go gitlab.com/pietroski-software-company/tools/serializer/go-serializer/pkg/tools/serializer Beautifier
