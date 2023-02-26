# Variables

If a system already provides `NCI` compliant variables, then it can set `NCI` to true to prevent this script from being executed.
In that case it's required to set `NCI_VERSION` to the implemented specification, to allow scripts to run actions if a newer spec is released.

Spec-Version: `0.3.0`

## common

| Variable           | Description                                                                                |
|--------------------|--------------------------------------------------------------------------------------------|
| `NCI`              | Will be set the true, if the variables have been normalized. (this script)                 |
| `NCI_VERSION`      | The revision of nci that was used to generate the normalized variables.                    |
| `NCI_SERVICE_NAME` | The commercial name of the used ci service. (e.g. GitLab CI, Travis CI, CircleCI, Jenkins) |
| `NCI_SERVICE_SLUG` | The commercial name normalized as slug for use in scripts, will not be changed.            |

## worker

| Variable             | Description                                           |
|----------------------|-------------------------------------------------------|
| `NCI_WORKER_ID`      | A unique id of the ci worker.                         |
| `NCI_WORKER_NAME`    | The human readable name of the ci worker.             |
| `NCI_WORKER_TYPE`    | Runner Type (github_hosted_vm, gitlab_hosted_vm, ...) |
| `NCI_WORKER_OS`      | OS / OS Image of the ci worker.                       |
| `NCI_WORKER_VERSION` | The version of the ci worker.                         |
| `NCI_WORKER_ARCH`    | The arch of the ci worker. (ie. linux/amd64)          |

## pipeline

| Variable                       | Description                                                                               |
|--------------------------------|-------------------------------------------------------------------------------------------|
| `NCI_PIPELINE_ID`              | Pipeline ID                                                                               |
| `NCI_PIPELINE_TRIGGER`         | What triggered the pipeline. (ie. manual/push/trigger/api/schedule/pull_request/build)    |
| `NCI_PIPELINE_STAGE_ID`        | Pipeline Stage ID                                                                         |
| `NCI_PIPELINE_STAGE_NAME`      | Human readable name of the current stage.                                                 |
| `NCI_PIPELINE_STAGE_SLUG`      | Slug of the current stage.                                                                |
| `NCI_PIPELINE_JOB_ID`          | Pipeline Job ID                                                                           |
| `NCI_PIPELINE_JOB_NAME`        | Human readable name of the current job.                                                   |
| `NCI_PIPELINE_JOB_SLUG`        | Slug of the current job.                                                                  |
| `NCI_PIPELINE_JOB_STARTED_AT`  | Pipeline Started At                                                                       |
| `NCI_PIPELINE_ATTEMPT`         | Run Attempt                                                                               |
| `NCI_PIPELINE_URL`             | Pipeline URL                                                                              |
| `NCI_PIPELINE_PULL_REQUEST_ID` | The number of the pull request, is only present if `NCI_PIPELINE_TRIGGER` = pull_request. |

## repository

| Variable                             | Description                                                                                                 |
|--------------------------------------|-------------------------------------------------------------------------------------------------------------|
| `NCI_REPOSITORY_KIND`                | The used version control system. (git)                                                                      |
| `NCI_REPOSITORY_REMOTE`              | The remote url pointing at the repository. (git remote url or `local` if no remote was found)               |

## project

| Variable                     | Description                                                    |
|------------------------------|----------------------------------------------------------------|
| `NCI_PROJECT_ID`             | Unique project id, can be used in deployments.                 |
| `NCI_PROJECT_NAME`           | Unique project id, can be used in deployments.                 |
| `NCI_PROJECT_PATH`           | Path of the Namespace and the project.                         |
| `NCI_PROJECT_SLUG`           | Project slug, that can be used in deployments.                 |
| `NCI_PROJECT_DESCRIPTION`    | The project description.                                       |
| `NCI_PROJECT_TOPICS`         | The topics / tags of the project.                              |
| `NCI_PROJECT_ISSUE_URL`      | A template for links to issues, contains a `{ID}` placeholder. |
| `NCI_PROJECT_STARGAZERS`     | The number of people who `follow` / `bookmarked` the project.  |
| `NCI_PROJECT_FORKS`          | The number of forks of the project.                            |
| `NCI_PROJECT_DEFAULT_BRANCH` | The default branch                                             |
| `NCI_PROJECT_DIR`            | Project directory on the local filesystem.                     |

## container registry

| Variable                           | Description                                                  |
|------------------------------------|--------------------------------------------------------------|
| `NCI_CONTAINERREGISTRY_HOST`       | The hostname of the container registry.                      |
| `NCI_CONTAINERREGISTRY_USERNAME`   | The username used for container registry authentication.     |
| `NCI_CONTAINERREGISTRY_PASSWORD`   | The password used for container registry authentication.     |
| `NCI_CONTAINERREGISTRY_REPOSITORY` | The repository, that should be used for the current project. |
| `NCI_CONTAINERREGISTRY_TAG`        | The tag that should be build.                                |

## commit

| Variable                             | Description                                                                                                 |
|--------------------------------------|-------------------------------------------------------------------------------------------------------------|
| `NCI_COMMIT_REF_TYPE`                | The reference type. (branch / tag)                                                                          |
| `NCI_COMMIT_REF_NAME`                | Human readable name of the current repository reference.                                                    |
| `NCI_COMMIT_REF_PATH`                | Combination of the ref type and ref name. (tag/v1.0.0 or branch/main)                                       |
| `NCI_COMMIT_REF_SLUG`                | Slug of the current repository reference.                                                                   |
| `NCI_COMMIT_REF_VCS`                 | Holds the vcs specific absolute reference name. (ex: `refs/heads/main`)                                     |
| `NCI_COMMIT_REF_RELEASE`             | Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`. |
| `NCI_COMMIT_SHA`                     | A unique hash, that each commit gets.                                                                       |
| `NCI_COMMIT_SHA_SHORT`               | A short form of the unique commit hash. (8 chars)                                                           |
| `NCI_COMMIT_AUTHOR_NAME`             | author name                                                                                                 |
| `NCI_COMMIT_AUTHOR_EMAIL`            | author email                                                                                                |
| `NCI_COMMIT_COMMITTER_NAME`          | committer name                                                                                              |
| `NCI_COMMIT_COMMITTER_EMAIL`         | committer email                                                                                             |
| `NCI_COMMIT_TITLE`                   | The title of the latest commit on the current reference.                                                    |
| `NCI_COMMIT_DESCRIPTION`             | The description of the latest commit on the current reference.                                              |
| `NCI_COMMIT_COUNT`                   | The total amount of commits inside of the current reference, can be used as build number.                   |
| `NCI_LASTRELEASE_REF_NAME`           | Human readable name of the last stable release.                                                             |
| `NCI_LASTRELEASE_REF_SLUG`           | Slug of the last stable release.                                                                            |
| `NCI_LASTRELEASE_REF_VCS`            | Holds the vcs specific absolute reference name of the last stable release. (ex: `refs/heads/main`)          |
| `NCI_LASTRELEASE_COMMIT_AFTER_COUNT` | Holds the count of commits since the last stable release.                                                   |

## control

| Variable            | Description                                            |
|---------------------|--------------------------------------------------------|
| `NCI_DEPLOY_FREEZE` | Currently in a deploy freeze window? (`true`, `false`) |
