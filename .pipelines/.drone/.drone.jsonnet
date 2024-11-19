local FromSecret(secret) = {
  from_secret: secret,
};

local envs() = {
  GITLAB_ACCESS_TOKEN: FromSecret('GITLAB_ACCESS_TOKEN'),
  GITLAB_USERNAME: FromSecret('GITLAB_USERNAME'),
  GITHUB_ACCESS_TOKEN: FromSecret('GITHUB_ACCESS_TOKEN'),
  GITHUB_USERNAME: FromSecret('GITHUB_USERNAME'),
};

local set_netrc = [
  'apk update && apk upgrade && apk add git bash make build-base',
  'echo -e "machine gitlab.com\nlogin $${GITLAB_USERNAME}\npassword $${GITLAB_ACCESS_TOKEN}" > ~/.netrc',
  'echo -e "\n" >> ~/.netrc',
  'echo -e "machine gitlab.com\nlogin $${GITHUB_USERNAME}\npassword $${GITHUB_ACCESS_TOKEN}" >> ~/.netrc',
  'chmod 640 ~/.netrc',
];

local set_ptk_envs = [
  'go env -w GONOSUMDB=gitlab.com/pietroski-software-company',
  'go env -w GONOPROXY=gitlab.com/pietroski-software-company',
  'go env -w GOPRIVATE=gitlab.com/pietroski-software-company',
];

local run_tests = [
  'go test -race --tags=unit -cover ./...',
];

local remote_gitlab_repo_address = 'https://gitlab.com/pietroski-software-company/devex/golang/serializer';
local remote_github_repo_address = 'https://github.com/Pietroski/go-serializer';

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

local remote_push = [
  'git checkout -b release/merging-branch',
  'git remote add gitlab '+remote_gitlab_repo_address,
  'git push gitlab release/merging-branch -f',
//  'git remote add github '+remote_github_repo_address,
//  'git push github release/merging-branch -f',
];

local remotePushStep(image, envs) = {
  name: 'remote-push',
  image: image,
  environment: envs,
  commands: std.flattenArrays([set_netrc, remote_push]),
};

local remote_tag = [
  'make tag',
  'git remote add gitlab '+remote_gitlab_repo_address,
  'git push gitlab --tags',
//  'git remote add github '+remote_github_repo_address,
//  'git push github --tags',
];

local remoteTagStep(image, envs) = {
  name: 'remote-tag',
  image: image,
  environment: envs,
  commands: std.flattenArrays([set_netrc, remote_tag]),
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
  tests('test-non-main', image, envs),
]);

local whenCommitToMaster(step) = step {
  when: {
    event: ['push'],
    branch: ['main'],
  },
};

local commitToMasterSteps(image, envs) = std.map(whenCommitToMaster, [
  tests('test-main', image, envs),
  remotePushStep(image, envs),
]);

local whenTag(step) = step {
  when: {
    event: ['tag'],
    branch: ['*'],
  },
};

local tagSteps(image, envs) = std.map(whenTag, [
  tests('test-tag', image, envs),
  remoteTagStep(image, envs),
]);

local Pipeline(name, image) = {
  kind: 'pipeline',
  name: name,
  steps: std.flattenArrays([
    commitToNonMasterSteps(image, envs()),
    commitToMasterSteps(image, envs()),
    tagSteps(image, envs()),
  ]),
};

[
  Pipeline('go-serializer-pipeline', 'golang:1.23.3-alpine3.20'),
]
