local FromSecret(secret) = {
	"from_secret": secret,
};

local envs() = {
	GITLAB_ACCESS_TOKEN: FromSecret("GITLAB_ACCESS_TOKEN"),
	GITLAB_USERNAME: FromSecret("GITLAB_USERNAME"),
};

local set_netrc = [
	"apk update && apk upgrade && apk add git bash make build-base",
	'echo -e "machine gitlab.com\nlogin $${GITLAB_USERNAME}\npassword $${GITLAB_ACCESS_TOKEN}" > ~/.netrc',
	"chmod 640 ~/.netrc",
];

local set_ptk_envs = [
	"go env -w GONOSUMDB=gitlab.com/pietroski-software-company",
	"go env -w GONOPROXY=gitlab.com/pietroski-software-company",
	"go env -w GOPRIVATE=gitlab.com/pietroski-software-company",
];

local run_tests = [
	"go test -race $(go list ./... | grep -v /tests/ | grep -v /mocks/ | grep -v /schemas/)",
];

local remote_git_repo_address = 'https://gitlab.com/pietroski-software-company/tools/serializer/go-serializer.git';

local tests_cmd = std.flattenArrays([
	set_netrc,
	set_ptk_envs,
	run_tests,
]);

local tests(name, image, envs) = {
	name: name,
  image: image,
  environment: envs,
  commands: tests_cmd,
};

local gitlab_push = [
	"git checkout -b release/merging-branch",
  "git remote add gitlab "+remote_git_repo_address,
  "git push gitlab release/merging-branch -f",
];

local gitlabPushStep(image, envs) = {
	name: "gitlab-push",
  image: image,
  environment: envs,
  commands: std.flattenArrays([set_netrc, gitlab_push]),
};

local gitlab_tag = [
  "git remote add gitlab "+remote_git_repo_address,
  "make tag",
  "git push gitlab --tags",
];

local gitlabTagStep(image, envs) = {
	name: "gitlab-tag",
  image: image,
  environment: envs,
  commands: std.flattenArrays([set_netrc, gitlab_tag]),
};

local whenCommitToNonMaster(step) = step {
  when: {
    event: ['push'],
    branch: {
      exclude: ['main'],
    },
  },
};

local commitToNonMasterSteps(image, envs) = std.map(whenCommitToNonMaster, [
  tests("test-non-main", image, envs),
]);

local whenCommitToMaster(step) = step {
	when: {
		event: ['push'],
		branch: ['main'],
	},
};

local commitToMasterSteps(image, envs) = std.map(whenCommitToMaster, [
	tests("test-main", image, envs),
	gitlabPushStep(image, envs),
]);

local whenTag(step) = step {
	when: {
		event: ['tag'],
		branch: ['*'],
	},
};

local tagSteps(image, envs) = std.map(whenTag, [
	tests("test-tag", image, envs),
	gitlabTagStep(image, envs),
]);

local Pipeline(name, image) = {
  kind: "pipeline",
  name: name,
  steps: std.flattenArrays([
    commitToNonMasterSteps(image, envs()),
    commitToMasterSteps(image, envs()),
    tagSteps(image, envs()),
  ]),
};

[
  Pipeline("go-serializer-pipeline", "golang:1.21.3-alpine3.18"),
]
