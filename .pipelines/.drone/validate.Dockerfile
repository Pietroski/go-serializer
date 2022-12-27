FROM alpine:3.17.0

RUN apk update && apk upgrade && apk add --no-cache ca-certificates curl tar jq
RUN curl -L https://github.com/harness/drone-cli/releases/latest/download/drone_linux_amd64.tar.gz | tar zx
RUN install -t /usr/local/bin drone

COPY .pipelines/.drone/.drone.jsonnet .
RUN echo $(drone jsonnet --stream --stdout) > drone-analysis
RUN cut -c 5- drone-analysis > pre-drone-analysis.json
RUN jq . pre-drone-analysis.json > drone-analysis.json
