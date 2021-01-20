FROM alpine:3.12

ARG TOOL_TRIVY_VERSION=0.15.0
ARG TOOL_DOCKLE_VERSION=0.3.1
ARG TOOL_HADOLINT_VERSION=1.19.0

RUN apk add --no-cache bash

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
