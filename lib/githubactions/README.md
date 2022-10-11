# GitHub Actions

## sources

- [Default Variables](https://help.github.com/en/articles/virtual-environments-for-github-actions#default-environment-variables)
- [Predefined variables](https://developer.github.com/actions/creating-github-actions/accessing-the-runtime-environment/#environment-variables)

## detection

GitHub Actions are detected by the presence of `GITHUB_ACTIONS`.

## resources

- [Run locally](https://github.com/nektos/act)

## example variables

```bash
LEIN_HOME=/usr/local/lib/lein
M2_HOME=/usr/share/apache-maven-3.6.1
BOOST_ROOT=/usr/local/share/boost/1.69.0
GOROOT_1_11_X64=/usr/local/go1.11
ANDROID_HOME=/usr/local/lib/android/sdk
JAVA_HOME_11_X64=/usr/lib/jvm/zulu-11-azure-amd64
LANG=C.UTF-8
INVOCATION_ID=c5ef6287439c412aa6edcf0dc9dcd685
JAVA_HOME_12_X64=/usr/lib/jvm/zulu-12-azure-amd64
ANDROID_SDK_ROOT=/usr/local/lib/android/sdk
RUNNER_TOOL_CACHE=/opt/hostedtoolcache
JAVA_HOME=/usr/lib/jvm/zulu-8-azure-amd64
RUNNER_TRACKING_ID=github_a9744cbc-8514-42e9-ac4a-4fdcad31014c
GITHUB_ACTIONS=true
DOTNET_SKIP_FIRST_TIME_EXPERIENCE=1
USER=runner
GITHUB_HEAD_REF=
GITHUB_ACTOR=PhilippHeuer
GITHUB_ACTION=run
GRADLE_HOME=/usr/share/gradle
PWD=/home/runner/work/normalize.ci/normalize.ci
HOME=/home/runner
GOROOT=/usr/local/go1.12
JOURNAL_STREAM=9:27802
JAVA_HOME_8_X64=/usr/lib/jvm/zulu-8-azure-amd64
RUNNER_TEMP=/home/runner/work/_temp
CONDA=/usr/share/miniconda
BOOST_ROOT_1_69_0=/usr/local/share/boost/1.69.0
RUNNER_WORKSPACE=/home/runner/work/normalize.ci
GITHUB_REF=refs/heads/master
GITHUB_SHA=dfe4a05873251d35f5fc4771106074a932751be1
GOROOT_1_12_X64=/usr/local/go1.12
DEPLOYMENT_BASEPATH=/opt/runner
GITHUB_EVENT_PATH=/home/runner/work/_temp/_github_workflow/event.json
RUNNER_OS=Linux
GITHUB_BASE_REF=
VCPKG_INSTALLATION_ROOT=/usr/local/share/vcpkg
PERFLOG_LOCATION_SETTING=RUNNER_PERFLOG
JAVA_HOME_7_X64=/usr/lib/jvm/zulu-7-azure-amd64
RUNNER_USER=runner
SHLVL=1
GITHUB_REPOSITORY=PhilippHeuer/normalize.ci
GITHUB_EVENT_NAME=push
LEIN_JAR=/usr/local/lib/lein/self-installs/leiningen-2.9.1-standalone.jar
RUNNER_PERFLOG=/home/runner/perflog
GITHUB_WORKFLOW=CI
ANT_HOME=/usr/share/ant
PATH=/usr/share/rust/.cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin
GITHUB_WORKSPACE=/home/runner/work/normalize.ci/normalize.ci
CHROME_BIN=/usr/bin/google-chrome
```
