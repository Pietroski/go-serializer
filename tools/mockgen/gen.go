package mock_generator

import (
	_ "go.uber.org/mock/mockgen/model"
)

//go:generate mockgen -package mocks -destination ../../mocks/mocked_serializer.go gitlab.com/pietroski-software-company/devex/golang/serializer/models Serializer
//go:generate mockgen -package mocks -destination ../../mocks/mocked_beautifier.go gitlab.com/pietroski-software-company/devex/golang/serializer/models Beautifier
