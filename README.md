# normalize.ci

[![Go Reference](https://pkg.go.dev/badge/github.com/cidverse/normalizeci.svg)](https://pkg.go.dev/github.com/cidverse/normalizeci)
[![Go Report Card](https://goreportcard.com/badge/github.com/cidverse/normalizeci)](https://goreportcard.com/report/github.com/cidverse/normalizeci)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/cidverse/normalizeci/badge)](https://securityscorecards.dev/viewer/?uri=github.com/cidverse/normalizeci)

> A cli tool (or go library) to provide a foundation for a platform-agnostic CICD process.

**Documentation**: https://cidverse.github.io/normalizeci/

**Quick Links**:

- [standardized variables](https://cidverse.github.io/normalizeci/spec/)

## features

- **normalization** - check the env vars and the local repository to provide a [common set of env vars](docs/spec/variables.md) on any ci platform.
- **compatibility** - convert the common env vars into a specific format (i.e. gitlab) to run a script made for gitlab on any ci provider.

## installation

You can download the binaries from the project releases: https://github.com/cidverse/normalizeci/releases

Linux:
```bash
sudo curl -L -s -o /usr/local/bin/normalizeci https://github.com/cidverse/normalizeci/releases/download/v1.0.0/linux_amd64
sudo chmod +x /usr/local/bin/normalizeci
```

(available builds: linux_386, linux_amd64, windows_386, windows_amd64, darwin_386, darwin_amd64)

## usage

### cli

Examples:

| Id  | Command                                                  | Description                                                                       |
|-----|----------------------------------------------------------|-----------------------------------------------------------------------------------|
| 1   | `normalizeci normalize --format export --output nci.env` | generate nci variables in format export for unix systems, stored as file          |
| 2   | `normalizeci normalize --format powershell`              | generate nci variables in format export for windows powershell, written to stdout |
| 3   | `normalizeci normalize --output nci.env`                 | generate nci variables in the suggested format for the current system             |
| 4   | `normalizeci normalize --hostenv --output nci.env`       | additionally to 3 includes all env vars from the host                             |
| 5   | `normalizeci normalize --format cmd`                     | generate nci variables in format export for windows cmd, written to stdout        |
| 6   | `normalizeci denormalize --target gitlab`                | generate a gitlab ci like environment based on the normalized environment         |
| 7   | `normalizeci version`                                    | print version information                                                         |

#### file based

Linux/MacOS

```bash
normalizeci normalize --format export --output nci.env
source nci.env
rm nci.env
```

Windows

```powershell
normalizeci normalize --format powershell --output nci.ps1
& .\nci.ps1
rm nci.ps1
```

#### terminal session

The NormalizeCI CLI will return the commands to set the normalized variables in your current terminal session, so you need to run the response of the command.

Linux/MacOS

```bash
eval $(normalizeci normalize)
```

Windows

```powershell
$nenv = normalizeci normalize
Invoke-Expression "$nenv"
```

### library

Install the latest version as library:

```bash
go get -u github.com/cidverse/normalizeci/pkg
```

And access the normalized environment, by default it will search for the vcs repo based on the current working directory.

```go
var normalized = normalizeci.RunDefaultNormalization()
```

## normalized variables (!)

- [Specification: Variables](docs/spec/variables.md)

## supported systems

| NAME                  | SLUG             |
|-----------------------|------------------|
| Azure DevOps Pipeline | `azure-devops`   |
| GitLab CI/CD          | `gitlab-ci`      |
| GitHub Actions        | `github-actions` |
| Local Git Repository  | `local`          |

*Note:* If none of the above systems is detected, repository information is determined based on the local Git repository.

## supported repository types

- `git`

## planned systems

*Note:*: If you want to contribute, feel free to pick one of the following services and add a package to normalize their variables.

| NAME          | SLUG            |
|---------------|-----------------|
| AppVeyor      | `appveyor`      |
| AWS CodeBuild | `aws-codebuild` |
| Bamboo        | `bamboo`        |
| Bitbucket     | `bitbucket`     |
| Bitrise       | `bitrise`       |
| Buddy         | `buddy`         |
| Buildkite     | `buildkite`     |
| CircleCI      | `circleci`      |
| Cirrus CI     | `cirrusci`      |
| Codefresh     | `codefresh`     |
| Codeship      | `codeship`      |
| Drone         | `drone`         |
| Jenkins       | `jenkins`       |
| Sail CI       | `sailci`        |
| Semaphore     | `semaphore`     |
| Shippable     | `shippable`     |
| TeamCity      | `teamcity`      |
| Travis CI     | `travis-ci`     |
| Wercker       | `wercker`       |

If a system is missing in  this list, please open an issue.

## License

Released under the [MIT license](./LICENSE).
