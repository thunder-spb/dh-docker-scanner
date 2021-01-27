# Docker Container Scanner

This container include Docker Image Scanner tools like [Trivy](https://github.com/aquasecurity/trivy/) and [Dockle](https://github.com/goodwithtech/dockle/). For Dockerfile linting this image include [hadolint](https://github.com/hadolint/hadolint/)

As a bonus, this image also has 2 binaries which help to convert Dockle and Hadolint JSON output into JUnit report -- `dockle2junit` and `hadolint2junit`. Both located in `/usr/bin`. See usage examples below. Source code you can find under `scripts/convert2junit` directory.

# What is Trivy

> A Simple and Comprehensive Vulnerability Scanner for Containers and other Artifacts, Suitable for CI.

Project homepage and documentation is here: https://github.com/aquasecurity/trivy/

In addition, this image include JUnit report template which could be consumed by CI system, like Jenkins.

# What is Dockle

> Dockle - Container Image Linter for Security, Helping build the Best-Practice Docker Image, Easy to start

This tool do two types of checks:

- Following official [Best Practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)

Project homepage and documentation is here: https://github.com/goodwithtech/dockle/

# What is Hadolint

> A smarter Dockerfile linter that helps you build [best practice](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices) Docker images. The linter is parsing the Dockerfile into an AST and performs rules on top of the AST. It is standing on the shoulders of [ShellCheck](https://github.com/koalaman/shellcheck) to lint the Bash code inside `RUN` instructions.

Project homepage and documentation is here: https://github.com/hadolint/hadolint/

# Versions

- Alpine version: _3.12_
- Trivy (`trivy`) version: _0.15.0_, Home: https://github.com/aquasecurity/trivy/
- Dockle (`dockle`) version: _0.3.1_, Home: https://github.com/goodwithtech/dockle/
- Hadolint (`hadolint`) version: _1.19.0_, Home: https://github.com/hadolint/hadolint/
- Dockle and Hadolint converters (`dockle2junit`, `hadolint2junit`) -- not versioned, Home: https://github.com/thunder-spb/dh-docker-scanner/tree/master/scripts/convert2junit

# Usage

Here is some examples on how to run scans within this container.

## Trivy scan

Pull this image by invoking this command:

```
docker pull thunderspb/docker-scanner
```

Attach to pulled container with mounting docker socket:

```
docker run -ti --rm --name docker-scanner \
  -v ${PWD}:/work -w /work \
  -v /var/run/docker.sock:/var/run/docker.sock \
  thunderspb/docker-scanner /bin/bash
```

and then, inside container, run:

```
trivy image --format template --template "@${TRIVY_TPL_JUNIT}" -o junit-report.xml <your-docker-image>
```

## Dockle scan

Pull this image by invoking this command:

```
docker pull thunderspb/docker-scanner
```

Attach to pulled container with mounting docker socket:

```
docker run -ti --rm --name docker-scanner \
  -v ${PWD}:/work -w /work \
  -v /var/run/docker.sock:/var/run/docker.sock \
  thunderspb/docker-scanner /bin/bash
```

and then, inside container, just run:

```
dockle thunderspb/docker-scanner
```

### `dockle2junit` converter

In order to use this tool, you must run `dockle` with `--format json --output filename.json` arguments.

After you got Dockle report in JSON, you can invoke `dockle2junit`:

```
dockle2junit -in-file filename.json -out-file junit-report.xml -image-name my/imagename:tag
```

`-image-name` -- to add scanned image name into report for convenience.

Here is an example of Dockle scan results in JSON format for this (`thunderspb/docker-scanner`) image:

```json
{
  "summary": {
    "fatal": 0,
    "warn": 2,
    "info": 2,
    "skip": 0,
    "pass": 12
  },
  "details": [
    {
      "code": "CIS-DI-0001",
      "title": "Create a user for the container",
      "level": "WARN",
      "alerts": ["Last user should not be root"]
    },
    {
      "code": "DKL-DI-0006",
      "title": "Avoid latest tag",
      "level": "WARN",
      "alerts": ["Avoid 'latest' tag"]
    },
    {
      "code": "CIS-DI-0005",
      "title": "Enable Content trust for Docker",
      "level": "INFO",
      "alerts": ["export DOCKER_CONTENT_TRUST=1 before docker pull/build"]
    },
    {
      "code": "CIS-DI-0006",
      "title": "Add HEALTHCHECK instruction to the container image",
      "level": "INFO",
      "alerts": ["not found HEALTHCHECK statement"]
    }
  ]
}
```

And an example of generated JUnit report by `dockle2junit` binary:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<testsuites>
  <testsuite tests="4" failures="0" time="" name="thunderspb/docker-scanner:latest">
    <properties></properties>
    <testcase classname="thunderspb/docker-scanner:latest" name="[WARN] CIS-DI-0001: Create a user for the container" time="">
      <system-err>Last user should not be root</system-err>
    </testcase>
    <testcase classname="thunderspb/docker-scanner:latest" name="[WARN] DKL-DI-0006: Avoid latest tag" time="">
      <system-err>Avoid &#39;latest&#39; tag</system-err>
    </testcase>
    <testcase classname="thunderspb/docker-scanner:latest" name="[INFO] CIS-DI-0005: Enable Content trust for Docker" time="">
      <system-out>export DOCKER_CONTENT_TRUST=1 before docker pull/build</system-out>
    </testcase>
    <testcase classname="thunderspb/docker-scanner:latest" name="[INFO] CIS-DI-0006: Add HEALTHCHECK instruction to the container image" time="">
      <system-out>not found HEALTHCHECK statement</system-out>
    </testcase>
  </testsuite>
</testsuites>
```

Now, you can order your CI system consume this file :)

## Hadolint scan

This tool scans Dockerfile, not the image, so working directory should contain your Dockerfile.

Pull this image by invoking this command:

```
docker pull thunderspb/docker-scanner
```

Attach to pulled container with mounting docker socket:

```
docker run -ti --rm --name docker-scanner \
  -v ${PWD}:/work -w /work \
  -v /var/run/docker.sock:/var/run/docker.sock \
  thunderspb/docker-scanner /bin/bash
```

and then, inside container, just run:

```
hadolint Dockerfile
```

### `hadolint2junit` converter

In order to use this tool, you must run `hadolint` with `--format json` arguments and redirect its output to some file, like this:

```
hadolint --format json Dockerfile > hadolint-docker-scanner.json
```

After you got Hadolint report in JSON, you can invoke `hadolint2junit`:

```
hadolint2junit -in-file hadolint-docker-scanner.json -out-file junit-hadolint-docker-scanner.xml
```

Here is an example of Hadolint scan results in JSON format for `Dockerfile` of this (`thunderspb/docker-scanner`) image:

```json
[
  {
    "line": 16,
    "code": "DL3018",
    "message": "Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  },
  {
    "line": 20,
    "code": "DL3003",
    "message": "Use WORKDIR to switch to a directory",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  },
  {
    "line": 20,
    "code": "DL4006",
    "message": "Set the SHELL option -o pipefail before RUN with a pipe in it. If you are using /bin/sh in an alpine image or if your shell is symlinked to busybox then consider explicitly setting your SHELL to /bin/ash, or disable this check",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  },
  {
    "line": 26,
    "code": "DL3003",
    "message": "Use WORKDIR to switch to a directory",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  },
  {
    "line": 26,
    "code": "DL4006",
    "message": "Set the SHELL option -o pipefail before RUN with a pipe in it. If you are using /bin/sh in an alpine image or if your shell is symlinked to busybox then consider explicitly setting your SHELL to /bin/ash, or disable this check",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  },
  {
    "line": 34,
    "code": "DL3003",
    "message": "Use WORKDIR to switch to a directory",
    "column": 1,
    "file": "Dockerfile",
    "level": "warning"
  }
]
```

And an example of generated JUnit report by `hadolint2junit` binary:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<testsuites>
  <testsuite tests="6" failures="0" time="" name="Dockerfile">
    <properties></properties>
    <testcase classname="Dockerfile:16:1" name="[WARN] DL3018" time="">
      <system-err>Pin versions in apk add. Instead of `apk add &lt;package&gt;` use `apk add &lt;package&gt;=&lt;version&gt;`</system-err>
    </testcase>
    <testcase classname="Dockerfile:20:1" name="[WARN] DL3003" time="">
      <system-err>Use WORKDIR to switch to a directory</system-err>
    </testcase>
    <testcase classname="Dockerfile:20:1" name="[WARN] DL4006" time="">
      <system-err>Set the SHELL option -o pipefail before RUN with a pipe in it. If you are using /bin/sh in an alpine image or if your shell is symlinked to busybox then consider explicitly setting your SHELL to /bin/ash, or disable this check</system-err>
    </testcase>
    <testcase classname="Dockerfile:26:1" name="[WARN] DL3003" time="">
      <system-err>Use WORKDIR to switch to a directory</system-err>
    </testcase>
    <testcase classname="Dockerfile:26:1" name="[WARN] DL4006" time="">
      <system-err>Set the SHELL option -o pipefail before RUN with a pipe in it. If you are using /bin/sh in an alpine image or if your shell is symlinked to busybox then consider explicitly setting your SHELL to /bin/ash, or disable this check</system-err>
    </testcase>
    <testcase classname="Dockerfile:34:1" name="[WARN] DL3003" time="">
      <system-err>Use WORKDIR to switch to a directory</system-err>
    </testcase>
  </testsuite>
</testsuites>
```

Now, you can order your CI system consume this file :)

# Github home

Sources: https://github.com/thunder-spb/dh-docker-scanner

# Docker Hub home

Here is the link on Docker Hub: https://hub.docker.com/r/thunderspb/docker-scanner
