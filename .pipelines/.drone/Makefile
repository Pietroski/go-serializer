# Drone docker commands list

DOCKER_TAG := $(shell cat .pipelines/.drone/DOCKER_VERSION)

drone-clean-validate:
	docker image build --no-cache --progress=plain -t pietroski/drone-jsonnet-validate:"${DOCKER_TAG}" -f .pipelines/.drone/validate.Dockerfile .

build-drone-validate:
	docker image build --no-cache -t pietroski/drone-jsonnet-validate:"${DOCKER_TAG}" -f .pipelines/.drone/validate.Dockerfile . # --progress=plain (after build)

run-drone-validate:
	docker run --name drone-validate pietroski/drone-jsonnet-validate:"${DOCKER_TAG}" # --rm
	docker cp drone-validate:drone-analysis .pipelines/.drone/reports/drone-analysis
	docker cp drone-validate:drone-analysis.json .pipelines/.drone/reports/drone-analysis.json
	docker stop drone-validate
	docker container rm drone-validate

rm-drone-validate:
	docker container stop drone-validate
	docker container rm drone-validate

drone-validate: build-drone-validate run-drone-validate
	@echo "done validating drone jsonnet file"
