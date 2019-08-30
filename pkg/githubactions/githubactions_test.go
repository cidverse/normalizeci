package githubactions

import (
	"os"
	"testing"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"
)

var testEnvironment = []string{
	"LEIN_HOME=/usr/local/lib/lein",
	"M2_HOME=/usr/share/apache-maven-3.6.1",
	"BOOST_ROOT=/usr/local/share/boost/1.69.0",
	"GOROOT_1_11_X64=/usr/local/go1.11",
	"ANDROID_HOME=/usr/local/lib/android/sdk",
	"JAVA_HOME_11_X64=/usr/lib/jvm/zulu-11-azure-amd64",
	"ImageVersion=157.1",
	"LANG=C.UTF-8",
	"INVOCATION_ID=f571ba9eda014d2b85ac026677077d76",
	"JAVA_HOME_12_X64=/usr/lib/jvm/zulu-12-azure-amd64",
	"ANDROID_SDK_ROOT=/usr/local/lib/android/sdk",
	"RUNNER_TOOL_CACHE=/opt/hostedtoolcache",
	"JAVA_HOME=/usr/lib/jvm/zulu-8-azure-amd64",
	"RUNNER_TRACKING_ID=github_1f3d9475-6c94-40ee-a160-8b3fd282c3a1",
	"GITHUB_ACTIONS=true",
	"DOTNET_SKIP_FIRST_TIME_EXPERIENCE=1",
	"USER=runner",
	"GITHUB_HEAD_REF=",
	"GITHUB_ACTOR=PhilippHeuer",
	"GITHUB_ACTION=run",
	"GRADLE_HOME=/usr/share/gradle",
	"PWD=/home/runner/work/normalize-ci/normalize-ci",
	"HOME=/home/runner",
	"GOROOT=/usr/local/go1.12",
	"JOURNAL_STREAM=9:31931",
	"JAVA_HOME_8_X64=/usr/lib/jvm/zulu-8-azure-amd64",
	"RUNNER_TEMP=/home/runner/work/_temp",
	"CONDA=/usr/share/miniconda",
	"BOOST_ROOT_1_69_0=/usr/local/share/boost/1.69.0",
	"RUNNER_WORKSPACE=/home/runner/work/normalize-ci",
	"GITHUB_REF=refs/heads/master",
	"GITHUB_SHA=abe7b23a948559a871f1158ec2db3caaef726854",
	"GOROOT_1_12_X64=/usr/local/go1.12",
	"DEPLOYMENT_BASEPATH=/opt/runner",
	"GITHUB_EVENT_PATH=/home/runner/work/_temp/_github_workflow/event.json",
	"RUNNER_OS=Linux",
	"GITHUB_BASE_REF=",
	"VCPKG_INSTALLATION_ROOT=/usr/local/share/vcpkg",
	"PERFLOG_LOCATION_SETTING=RUNNER_PERFLOG",
	"JAVA_HOME_7_X64=/usr/lib/jvm/zulu-7-azure-amd64",
	"RUNNER_USER=runner",
	"SHLVL=1",
	"GITHUB_REPOSITORY=PhilippHeuer/normalize-ci",
	"GITHUB_EVENT_NAME=push",
	"LEIN_JAR=/usr/local/lib/lein/self-installs/leiningen-2.9.1-standalone.jar",
	"RUNNER_PERFLOG=/home/runner/perflog",
	"GITHUB_WORKFLOW=CI",
	"ANT_HOME=/usr/share/ant",
	"PATH=/usr/share/rust/.cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin",
	"GITHUB_WORKSPACE=/home/runner/work/normalize-ci/normalize-ci",
	"CHROME_BIN=/usr/bin/google-chrome",
}

func TestMain(m *testing.M) {
	common.SetupTestLogger()
	code := m.Run()
	os.Exit(code)
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(testEnvironment) != true {
		t.Errorf("Check should succeed with the provided sample data!")
	}
}

func TestEnvironmentNormalizer(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(testEnvironment)

	// log all normalized values
	for _, envvar := range normalized {
		t.Log(envvar)
	}

	// validate fields
	// - common
	common.AssertThatEnvEquals(t, normalized, "NCI", "true")
	common.AssertThatEnvEquals(t, normalized, "NCI_VERSION", normalizer.version)
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVICE_NAME", normalizer.name)
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVICE_SLUG", normalizer.slug)
	// - server
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_NAME", "GitHub")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_HOST", "github.com")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_VERSION", "")
	// - worker
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ID", "github_1f3d9475-6c94-40ee-a160-8b3fd282c3a1")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_NAME", "github_1f3d9475-6c94-40ee-a160-8b3fd282c3a1")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_VERSION", "157.1")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ARCH", "linux/amd64")
	// - pipeline
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_TRIGGER", "push")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_NAME", "CI")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_SLUG", "ci")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_NAME", "run")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_SLUG", "run")
	// - project
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_ID", "philippheuer-normalize-ci")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_NAME", "PhilippHeuer/normalize-ci")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_SLUG", "philippheuer-normalize-ci")
}
