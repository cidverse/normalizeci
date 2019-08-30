# normalize.ci

> A tool to turn the continious integration / deployment variables into a common format for generally usable scripts without any dependencies.

## usage

dowonload the provided library and run it ...

```bash
// TODO
```

## normalized variables

- [Specification: Variables](docs/spec/variables.md)

## supported systems

NAME | SLUG
--- | --- |
[Azure DevOps Pipeline](docs/system/azure-devops-pipeline.md) | `azure-devops`
[GitLab CI/CD](docs/system/gitlab-ci.md) | `gitlab-ci`
[GitHub Actions](docs/system/github-actions.md) | `github-actions`

*Note:* If none of the above systems is detected, repository information is determined based on the local Git repository.
