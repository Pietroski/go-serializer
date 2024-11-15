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

bench:
	go test -bench BenchmarkAll -benchmem ./...

bench-to-file:
	go test -bench BenchmarkAll -benchmem ./... &> tests/benchmarks/serializer/results/BenchmarkAll.log

generate-bench-report:
	go test -run TestParseBenchResults -v ./...

generate-full-bench-report:
	echo 'implement me!'

########################################################################################################################

clean-all-caches:
	go clean -cache
	go clean -modcache
	go clean -testcache

test: clean-all-caches
	go test -race --tags=unit,test_validation -cover ./...

quick-test:
	go test -race --tags=unit,test_validation -cover ./...

########################################################################################################################

## generates mocks
mock-generate:
	go get go.uber.org/mock/mockgen
	go mod vendor
	go generate ./...
	go mod tidy
	go mod vendor

########################################################################################################################

tag-delete-all:
	git tag | xargs git tag -d

pull-latest:
	git pull gitea main

add-all:
	git add .

commit-with:
	git commit -m "$${m}"

chore-version-bump:
	git add .
	git commit -m "chore: version bump"

TAG := $(shell cat VERSION)
tag:
	git tag $(TAG)

changelog:
	@./scripts/docs/changelog.sh

commit-changelog:
	git add .
	git commit -m "chore: changelog"

gitea-push-main:
	git push gitea main

gitlab-push-main:
	git push gitlab main

push-main-all: gitea-push-main gitlab-push-main

amend:
	git commit --amend --no-edit

rebase-continue:
	git rebase --continue

trigger-pipeline: amend
	git push gitea main --force-with-lease

gitea-push-tags:
	git push gitea --tags

gitlab-push-tags:
	git push gitlab --tags

push-tags: gitea-push-tags gitlab-push-tags

publish:
	make chore-version-bump
	make tag
	make changelog
	make commit-changelog
	make changelog
	make commit-changelog

########################################################################################################################

count-written-lines:
	./scripts/metrics/line-counter

########################################################################################################################
