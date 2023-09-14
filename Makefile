# Makefile command list

# Makefile Documentation parser
-include ./scripts/docs/Makefile

# Drone's pipeline validation
-include .pipelines/.drone/Makefile

# Schemas's Makefile
-include ./scripts/schemas/Makefile

########################################################################################################################

report-dir:
	mkdir -p ./docs/reports/tests/unit/

test-unit:
	go test -race -v `go list ./... | grep -v ./pkg/mocks`

test-unit-cover: report-dir
	go test -race -v -coverprofile ./docs/reports/tests/unit/cover.out `go list ./... | grep -v ./pkg/mocks`

test-unit-cover-silent: report-dir
	go test -race -coverprofile ./docs/reports/tests/unit/cover.out `go list ./... | grep -v ./pkg/mocks`

test-unit-cover-all: report-dir
	go test -race -v -coverprofile ./docs/reports/tests/unit/cover-all.out ./...

test-unit-cover-all-silent: report-dir
	go test -race -coverprofile ./docs/reports/tests/unit/cover-all.out ./...

test-unit-cover-report: report-dir
	go tool cover -html=docs/reports/tests/unit/cover.out

test-unit-cover-all-report: report-dir
	go tool cover -html=docs/reports/tests/unit/cover-all.out

########################################################################################################################

get-mocker:
	go get -d github.com/golang/mock@v1.6.0
	go install github.com/golang/mock/mockgen@v1.6.0

## generates mocks
mock-generate:
	go get -d github.com/golang/mock/mockgen
	go mod download
	go generate ./...
	go mod tidy
	go mod download

########################################################################################################################

count-written-lines:
	./scripts/metrics/line-counter

########################################################################################################################

TAG := $(shell cat VERSION)
tag:
	git tag $(TAG)
