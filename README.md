# normalize.ci

[![Go Reference](https://pkg.go.dev/badge/github.com/cidverse/normalizeci.svg)](https://pkg.go.dev/github.com/cidverse/normalizeci)
[![Go Report Card](https://goreportcard.com/badge/github.com/cidverse/normalizeci)](https://goreportcard.com/report/github.com/cidverse/normalizeci)

> A cli tool (or go library) to provide a foundation for a platform-agnostic CICD process.

## features

- **normalization** - check the env vars and the local repository to provide a [common set of env vars](docs/spec/variables.md) on any ci platform.
- **compatibility** - convert the common env vars into a specific format (ie. gitlab) to run a script made for gitlab on any ci provider.

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

| Id  | Command                                        | Description                                                                       |
|-----|------------------------------------------------|-----------------------------------------------------------------------------------|
| 1   | `normalizeci --format export --output nci.env` | generate nci variables in format export for unix systems, stored as file          |
| 2   | `normalizeci --format powershell`              | generate nci variables in format export for windows powershell, written to stdout |
| 3   | `normalizeci --output nci.env`                 | generate nci variables in the suggested format for the current system             |
| 4   | `normalizeci --hostenv --output nci.env`       | additionally to 3 includes all env vars from the host                             |
| 5   | `normalizeci --format cmd`                     | generate nci variables in format export for windows cmd, written to stdout        |
| 6   | `normalizeci -v`                               | print version information                                                         |

#### file based

Linux/MacOS

```bash
normalizeci --format export --output nci.env
source nci.env
rm nci.env
```

Windows

```powershell
normalizeci --format powershell --output nci.ps1
& .\nci.ps1
rm nci.ps1
```

#### terminal session

The NormalizeCI CLI will return the commands to set the normalized variables in your current terminal session, so you need to run the response of the command.

Linux/MacOS

```bash
eval $(normalizeci)
```

Windows

```powershell
$nenv = normalizeci
Invoke-Expression "$nenv"
```

### library

Install the latest version as library:

```bash
go get -u github.com/cidverse/normalizeci/lib
```

And access the normalized environment, by default it will search for the vcs repo based on the current working directory.

```go
var normalized = normalizeci.RunDefaultNormalization()
```

## normalized variables (!)

- [Specification: Variables](docs/spec/variables.md)

## supported systems

| NAME                                               | SLUG             |
|----------------------------------------------------|------------------|
| [Azure DevOps Pipeline](lib/azuredevops/README.md) | `azure-devops`   |
| [GitLab CI/CD](lib/gitlabci/README.md)             | `gitlab-ci`      |
| [GitHub Actions](lib/githubactions/README.md)      | `github-actions` |
| [Local Git Repository](lib/localgit/README.md)     | `local`          |

*Note:* If none of the above systems is detected, repository information is determined based on the local Git repository.

## supported repository types

- `git`

## planned systems

*Note:*: If you want to contribute, feel free to pick one of the following services and add a package to normalize their variables.

| NAME                                            | SLUG            |
|-------------------------------------------------|-----------------|
| [AppVeyor](pkg_wip/appveyor/README.md)          | `appveyor`      |
| [AWS CodeBuild](pkg_wip/awscodebuild/README.md) | `aws-codebuild` |
| [Bamboo](pkg_wip/bamboo/README.md)              | `bamboo`        |
| [Bitbucket](pkg_wip/bitbucket/README.md)        | `bitbucket`     |
| [Bitrise](pkg_wip/bitrise/README.md)            | `bitrise`       |
| [Buddy](pkg_wip/buddy/README.md)                | `buddy`         |
| [Buildkite](pkg_wip/buildkite/README.md)        | `buildkite`     |
| [CircleCI](pkg_wip/circleci/README.md)          | `circleci`      |
| [Cirrus CI](pkg_wip/cirrusci/README.md)         | `cirrusci`      |
| [Codefresh](pkg_wip/codefresh/README.md)        | `codefresh`     |
| [Codeship](pkg_wip/codeship/README.md)          | `codeship`      |
| [Drone](pkg_wip/drone/README.md)                | `drone`         |
| [Jenkins](pkg_wip/jenkins/README.md)            | `jenkins`       |
| [Sail CI](pkg_wip/sailci/README.md)             | `sailci`        |
| [Semaphore](pkg_wip/semaphore/README.md)        | `semaphore`     |
| [Shippable](pkg_wip/shippable/README.md)        | `shippable`     |
| [TeamCity](pkg_wip/teamcity/README.md)          | `teamcity`      |
| [Travis CI](pkg_wip/travisci/README.md)         | `travis-ci`     |
| [Wercker](pkg_wip/wercker/README.md)            | `wercker`       |

## License

Released under the [MIT license](./LICENSE).
