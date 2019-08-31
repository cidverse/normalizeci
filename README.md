# normalize.ci

> A tool to turn the continious integration / deployment variables into a common format for generally usable scripts without any dependencies.

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
