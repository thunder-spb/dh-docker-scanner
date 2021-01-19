# Docker Container Scanner

This container include Docker Image Scanner tools like [Trivy](https://github.com/aquasecurity/trivy/) and [Dockle](https://github.com/goodwithtech/dockle/).

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

# Versions

- Alpine version: _3.12_
- Trivy (**trivy**) version: _0.15.0_, Home: https://github.com/aquasecurity/trivy/
- Dockle (**dockle**) version: _0.3.1_, Home: https://github.com/goodwithtech/dockle/

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
trivy image --format template --template \"@\${TRIVY_TPL_JUNIT}\" -o junit-report.xml <your-docker-image>
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

# Github home

https://github.com/thunder-spb/dh-docker-scanner

# Docker Hub home

Here is the link on Docker Hub: https://hub.docker.com/r/thunderspb/docker-scanner
