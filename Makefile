# Makefile command list

# Makefile Documentation parser
-include ./scripts/docs/Makefile

# Drone's pipeline validation
-include .pipelines/.drone/Makefile

# Schemas's Makefile
-include ./scripts/schemas/Makefile

########################################################################################################################

REPORT_DIR := tests/reports

report-dir:
	mkdir -p $(REPORT_DIR)

test-unit: clean-test-cache
	go test -race --tags=unit ./...

test-bench: clean-test-cache
	go test -race --tags=benchmark ./...

test-unit-cover: report-dir clean-test-cache
	go test -race --tags=unit -coverprofile $(REPORT_DIR)/unit_cover.out ./...

test-unit-cover-verbose: report-dir clean-test-cache
	go test -race --tags=unit -v -coverprofile $(REPORT_DIR)/unit_cover.out ./...

test-unit-cover-report-view: report-dir clean-test-cache
	go tool cover -html=$(REPORT_DIR)/unit_cover.out

quick-test-cover:
	go test -race --tags=unit -cover ./...

########################################################################################################################

BENCH_DIR := tests/benchmarks/serializer/results/BenchmarkAll.log

bench:
	go test --tags=benchmark -bench BenchmarkAll -benchmem ./...

bench-to-file:
	go test --tags=benchmark -bench BenchmarkAll -benchmem ./... &> $(BENCH_DIR)

generate-bench-report:
	go test --tags=benchmark -run TestParseBenchResults -v ./...

generate-full-bench-report:
	echo 'implement me!'

########################################################################################################################

clean-test-cache:
	go clean -testcache

clean-all-caches:
	go clean -cache
	go clean -modcache
	go clean -testcache

########################################################################################################################

## generates mocks
mock-generate:
	go get go.uber.org/mock/mockgen
	go get github.com/maxbrunsfeld/counterfeiter/v6
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

github-push-main:
	git push github main

push-main-all: gitea-push-main gitlab-push-main github-push-main

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

github-push-tags:
	git push github --tags

push-tags: gitea-push-tags gitlab-push-tags github-push-tags

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
