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
	git commit --amend

rebase-continue:
	git rebase --continue

trigger-pipeline:
	git commit --amend
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
