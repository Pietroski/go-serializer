# Makefile command list

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

TAG := $(shell cat VERSION)
tag:
	git tag $(TAG)
