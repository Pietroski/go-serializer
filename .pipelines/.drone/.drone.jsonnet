local Pipeline(name, image) = {
  kind: "pipeline",
  name: name,
  steps: [
    {
      name: "test",
      image: image,
      commands: [
        "apk update && apk upgrade && apk add git bash make build-base",
        'echo -e "machine gitlab.com\nlogin Pietroski\npassword glpat-ipcHZa8HtSHXdPyoAYFK" > ~/.netrc',
        "chmod 640 ~/.netrc",
        "go env -w GONOSUMDB=gitlab.com/pietroski-software-company",
        "go env -w GONOPROXY=gitlab.com/pietroski-software-company",
        "go env -w GOPRIVATE=gitlab.com/pietroski-software-company",
        "go env -w GO111MODULE=on",
        "go mod download",
        "go test $(go list ./... | grep -v /tests/ | grep -v /mocks/ | grep -v /schemas/)",
        "git checkout -b release/merging-branch",
        "git remote add gitlab https://gitlab.com/pietroski-software-company/tools/tracer/go-tracer.git",
        "git push gitlab release/merging-branch",
      ],
    },
  ]
};

[
  Pipeline("go-serializer-pipeline", "golang:1.19.4-alpine3.16"),
]
