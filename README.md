# Environment Variable Mapper
[![GoDoc](https://godoc.org/github.com/asecurityteam/env-mapper?status.svg)](https://godoc.org/github.com/asecurityteam/env-mapper)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=bugs)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=code_smells)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=coverage)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=ncloc)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=alert_status)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=security_rating)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=sqale_index)](https://sonarcloud.io/dashboard?id=env-mapper)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=env-mapper&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=env-mapper)

<https://github.com/asecurityteam/env-mapper>

- [Overview](#overview)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
    - [Logging](#logging)
    - [Stats](#stats)
- [Supported Resources](#supported-resources)
- [Status](#status)
- [Contributing](#contributing)
    - [Building And Testing](#building-and-testing)
    - [Quality Gates](#quality-gates)
    - [License](#license)
    - [Contributing Agreement](#contributing-agreement)

<a id="markdown-overview" name="overview"></a>
## Overview
![When and why do I need Env-Mapper?](doc/do-i-need-env-mapper.png)

## Quick Start
[Usage instructions](usage.md)

## Status

This project is in incubation which means we are not yet operating this tool in production
and the interfaces are subject to change.

<a id="markdown-contributing" name="contributing"></a>
## Contributing

If you are interested in contributing to the project, feel free to open an issue or PR.

<a id="markdown-building-and-testing" name="building-and-testing"></a>
### Building And Testing

We publish a docker image called [SDCLI](https://github.com/asecurityteam/sdcli) that
bundles all of our build dependencies. It is used by the included Makefile to help make
building and testing a bit easier. The following actions are available through the Makefile:

-   make dep

    Install the project dependencies into a vendor directory

-   make lint

    Run our static analysis suite

-   make test

    Run unit tests and generate a coverage artifact

-   make integration

    Run integration tests and generate a coverage artifact

-   make coverage

    Report the combined coverage for unit and integration tests

-   make build

    Generate a local build of the project (if applicable)

-   make run

    Run a local instance of the project (if applicable)

-   make doc

    Generate the project code documentation and make it viewable
    locally.

<a id="markdown-quality-gates" name="quality-gates"></a>
### Quality Gates

Our build process will run the following checks before going green:

-   make lint
-   make test
-   make integration
-   make coverage (combined result must be 85% or above for the project)

Running these locally, will give early indicators of pass/fail.

<a id="markdown-license" name="license"></a>
### License

This project is licensed under Apache 2.0. See LICENSE.txt for details.

<a id="markdown-contributing-agreement" name="contributing-agreement"></a>
### Contributing Agreement

Atlassian requires signing a contributor's agreement before we can accept a
patch. If you are an individual you can fill out the
[individual CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d).
If you are contributing on behalf of your company then please fill out the
[corporate CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b).
