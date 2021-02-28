# normalize.ci

> A cli tool to convert continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.

This project can also be used as library for other tools.

## installation

```bash
sudo curl -L -s -o /usr/local/bin/normalizeci https://www.philippheuer.me/linux_amd64
sudo chmod +x /usr/local/bin/normalizeci
```

## usage

```bash
eval $(normalizeci)
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
