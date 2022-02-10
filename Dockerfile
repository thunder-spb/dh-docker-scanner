FROM golang:alpine AS builder
# FROM golang:1.14.2
COPY scripts/convert2junit .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -a -installsuffix cgo -o dockle2junit dockle2junit.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -a -installsuffix cgo -o hadolint2junit hadolint2junit.go

FROM alpine:3.15

ARG TOOL_TRIVY_VERSION=0.23.0
ARG TOOL_DOCKLE_VERSION=0.4.3
ARG TOOL_HADOLINT_VERSION=2.8.0

COPY --from=builder /go/dockle2junit /usr/bin/dockle2junit
COPY --from=builder /go/hadolint2junit /usr/bin/hadolint2junit

RUN apk add --no-cache bash wget

ENV DOCKER_TOOLS_HOME=/srv/tools

RUN mkdir -p ${DOCKER_TOOLS_HOME}/dockle \
  && cd ${DOCKER_TOOLS_HOME}/dockle \
  && wget -nv -O - https://github.com/goodwithtech/dockle/releases/download/v${TOOL_DOCKLE_VERSION}/dockle_${TOOL_DOCKLE_VERSION}_Linux-64bit.tar.gz | tar -xz \
  && chmod 755 ${DOCKER_TOOLS_HOME}/dockle/dockle \
  && ln -sf ${DOCKER_TOOLS_HOME}/dockle/dockle /usr/bin/dockle

RUN mkdir -p ${DOCKER_TOOLS_HOME}/trivy \
  && cd ${DOCKER_TOOLS_HOME}/trivy \
  && export TRIVY_TPL_JUNIT=${DOCKER_TOOLS_HOME}/trivy/contrib/junit.tpl \
  && wget -nv -O- https://github.com/aquasecurity/trivy/releases/download/v${TOOL_TRIVY_VERSION}/trivy_${TOOL_TRIVY_VERSION}_Linux-64bit.tar.gz | tar -xz \
  && wget -nv -O ${TRIVY_TPL_JUNIT} https://raw.githubusercontent.com/aquasecurity/trivy/v${TOOL_TRIVY_VERSION}/contrib/junit.tpl \
  && chmod 755 ${DOCKER_TOOLS_HOME}/trivy/trivy \
  && ln -sf ${DOCKER_TOOLS_HOME}/trivy/trivy /usr/bin/trivy

RUN mkdir -p ${DOCKER_TOOLS_HOME}/hadolint \
  && cd ${DOCKER_TOOLS_HOME}/hadolint \
  && wget -nv -O ${DOCKER_TOOLS_HOME}/hadolint/hadolint https://github.com/hadolint/hadolint/releases/download/v${TOOL_HADOLINT_VERSION}/hadolint-Linux-x86_64 \
  && chmod 755 ${DOCKER_TOOLS_HOME}/hadolint/hadolint \
  && ln -sf ${DOCKER_TOOLS_HOME}/hadolint/hadolint /usr/bin/hadolint

ENV TRIVY_TPL_JUNIT=${DOCKER_TOOLS_HOME}/trivy/contrib/junit.tpl

LABEL name="Docker Image security scanner with Trivy, Dockle and Hadolint based on Alpine"
LABEL maintainer="Alexander thunder Shevchenko <iam@thunder.spb.ru>"
LABEL tools.dockle.verison="${TOOL_DOCKLE_VERSION}"
LABEL tools.dockle.homepage="https://github.com/goodwithtech/dockle/"
LABEL tools.trivy.verison="${TOOL_TRIVY_VERSION}"
LABEL tools.trivy.homepage="https://github.com/aquasecurity/trivy/"
LABEL tools.hadolint.verison="${TOOL_HADOLINT_VERSION}"
LABEL tools.hadolint.homepage="https://github.com/hadolint/hadolint/"
