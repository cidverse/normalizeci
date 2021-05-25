# normalize.ci

> A cli tool to convert continuous integration variables into a common format for generally usable scripts.

This project can also be used as library for other tools.

## installation

You can download the binaries from the gh release page: https://github.com/qubid/normalizeci/releases

- linux_386
- linux_amd64
- windows_386
- windows_amd64
- darwin_386
- darwin_amd64

Linux:
```bash
sudo curl -L -s -o /usr/local/bin/normalizeci https://github.com/qubid/normalizeci/releases/download/v1.0.0/linux_amd64
sudo chmod +x /usr/local/bin/normalizeci
```

## usage

### cli

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
go get -u github.com/qubid/normalizeci
```

## normalized variables

- [Specification: Variables](docs/spec/variables.md)

## supported systems

NAME | SLUG
--- | --- |
[Azure DevOps Pipeline](pkg/azuredevops/README.md) | `azure-devops`
[GitLab CI/CD](pkg/gitlabci/README.md) | `gitlab-ci`
[GitHub Actions](pkg/githubactions/README.md) | `github-actions`
[Local Git Repository](pkg/localgit/README.md) | `local`

*Note:* If none of the above systems is detected, repository information is determined based on the local Git repository.

## supported repository types

NAME |
--- |
`git` |

*Note:* Some information will be extracted directly from the vcs repository, because not all services provide the same level of information.

## planned systems

*Note:*: If you want to contribute, feel free to pick one of the following services and add a package to normalize their variables.

NAME | SLUG
--- | --- |
[AppVeyor](pkg_wip/appveyor/README.md) | `appveyor`
[AWS CodeBuild](pkg_wip/awscodebuild/README.md) | `aws-codebuild`
[Bamboo](pkg_wip/bamboo/README.md) | `bamboo`
[Bitbucket](pkg_wip/bitbucket/README.md) | `bitbucket`
[Bitrise](pkg_wip/bitrise/README.md) | `bitrise`
[Buddy](pkg_wip/buddy/README.md) | `buddy`
[Buildkite](pkg_wip/buildkite/README.md) | `buildkite`
[CircleCI](pkg_wip/circleci/README.md) | `circleci`
[Cirrus CI](pkg_wip/cirrusci/README.md) | `cirrusci`
[Codefresh](pkg_wip/codefresh/README.md) | `codefresh`
[Codeship](pkg_wip/codeship/README.md) | `codeship`
[Drone](pkg_wip/drone/README.md) | `drone`
[Jenkins](pkg_wip/jenkins/README.md) | `jenkins`
[Sail CI](pkg_wip/sailci/README.md) | `sailci`
[Semaphore](pkg_wip/semaphore/README.md) | `semaphore`
[Shippable](pkg_wip/shippable/README.md) | `shippable`
[TeamCity](pkg_wip/teamcity/README.md) | `teamcity`
[Travis CI](pkg_wip/travisci/README.md) | `travis-ci`
[Wercker](pkg_wip/wercker/README.md) | `wercker`
