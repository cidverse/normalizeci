package azuredevops

import (
	"testing"
	"os"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"
)

var testEnvironment = []string{
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI=https://heuer.visualstudio.com/",
	"TAG=8",
	"BUILD_SOURCEBRANCH=refs/heads/master",
	"SYSTEM_TASKDEFINITIONSURI=https://heuer.visualstudio.com/",
	"AGENT_VERSION=2.155.1",
	"SYSTEM_JOBATTEMPT=1",
	"LEIN_HOME=/usr/local/lib/lein",
	"BUILD_QUEUEDBY=GitHub",
	"SYSTEM_HOSTTYPE=build",
	"SYSTEM_COLLECTIONURI=https://heuer.visualstudio.com/",
	"GOROOT_1_11_X64=/usr/local/go1.11",
	"SYSTEM_JOBPARALLELISMTAG=Public",
	"BUILD_REPOSITORY_GIT_SUBMODULECHECKOUT=False",
	"GOROOT_1_9_X64=/usr/local/go1.9",
	"ANDROID_HOME=/usr/local/lib/android/sdk",
	"BUILD_STAGINGDIRECTORY=/home/vsts/work/1/a",
	"RUNNER_TOOLSDIRECTORY=",
	"AGENT_MACHINENAME=fv-az679",
	"SYSTEM_WORKFOLDER=/home/vsts/work",
	"COMMON_TESTRESULTSDIRECTORY=/home/vsts/work/1/TestResults",
	"GRADLE_HOME=/usr/share/gradle",
	"AGENT_JOBNAME=Build",
	"ANDROID_SDK_ROOT=/usr/local/lib/android/sdk",
	"MSDEPLOY_HTTP_USER_AGENT=VSTS_8c7f0be9-cf3c-4627-9df8-fa7d1cb80b37_build_1_0",
	"JAVA_HOME_8_X64=/usr/lib/jvm/zulu-8-azure-amd64",
	"AGENT_OSARCHITECTURE=X64",
	"BUILD_SOURCEVERSIONAUTHOR=Philipp Heuer",
	"BUILD_REQUESTEDFOREMAIL=philipp.heuer@outlook.com",
	"AGENT_ACCEPTTEEEULA=True",
	"SYSTEM_STAGEATTEMPT=1",
	"ANT_HOME=/usr/share/ant",
	"GIT_TERMINAL_PROMPT=0",
	"SYSTEM_DEFINITIONNAME=PhilippHeuer.azure-pipeline-test",
	"SYSTEM_CULTURE=en-US",
	"JAVA_HOME_11_X64=/usr/lib/jvm/zulu-11-azure-amd64",
	"AGENT_TEMPDIRECTORY=/home/vsts/work/_temp",
	"BUILD_REPOSITORY_CLEAN=False",
	"BUILD_SOURCEBRANCHNAME=master",
	"BUILD_REPOSITORY_PROVIDER=GitHub",
	"USER=vsts",
	"SYSTEM_JOBIDENTIFIER=Build.Build.__default",
	"SYSTEM_TEAMFOUNDATIONSERVERURI=https://heuer.visualstudio.com/",
	"TF_BUILD=True",
	"AZURE_HTTP_USER_AGENT=VSTS_8c7f0be9-cf3c-4627-9df8-fa7d1cb80b37_build_1_0",
	"BUILD_QUEUEDBYID=271ebd76-ab97-4f2c-b79e-e90b0bf28f01",
	"SYSTEM_TASKDISPLAYNAME=CmdLine",
	"SYSTEM_STAGENAME=Build",
	"ImageVersion=156.2",
	"AGENT_DISABLELOGPLUGIN_TESTRESULTLOGPLUGIN=false",
	"SYSTEM_TEAMPROJECTID=d1b384a8-f33f-427d-86fd-f021826a54ea",
	"AGENT_ROOTDIRECTORY=/home/vsts/work",
	"VSTS_PROCESS_LOOKUP_ID=vsts_e6211052-b261-4175-b9a8-f7fd395fa8ce",
	"AGENT_HOMEDIRECTORY=/home/vsts/agents/2.155.1",
	"SYSTEM_TEAMPROJECT=azure-test",
	"AGENT_TOOLSDIRECTORY=/opt/hostedtoolcache",
	"BUILD_REPOSITORY_ID=PhilippHeuer/azure-pipeline-test",
	"BUILD_SOURCEVERSIONMESSAGE=feature: initial commit",
	"BUILD_REPOSITORY_LOCALPATH=/home/vsts/work/1/s",
	"SYSTEM_JOBDISPLAYNAME=Build",
	"agent.jobstatus=Succeeded",
	"AGENT_BUILDDIRECTORY=/home/vsts/work/1",
	"SYSTEM=build",
	"BUILD_REASON=IndividualCI",
	"SYSTEM_PIPELINESTARTTIME=2019-08-11 21:29:12+00:00",
	"AGENT_OS=Linux",
	"BUILD_SOURCESDIRECTORY=/home/vsts/work/1/s",
	"DOTNET_SKIP_FIRST_TIME_EXPERIENCE=1",
	"PATH=/usr/share/rust/.cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin",
	"SYSTEM_PHASEATTEMPT=1",
	"LEIN_JAR=/usr/local/lib/lein/self-installs/leiningen-2.9.1-standalone.jar",
	"SYSTEM_ISSCHEDULED=False",
	"VSTS_AGENT_PERFLOG=/home/vsts/agentperf",
	"PWD=/home/vsts/work/1/s",
	"BUILD_BUILDURI=vstfs:///Build/Build/8",
	"SYSTEM_PULLREQUEST_ISFORK=False",
	"CONDA=/usr/share/miniconda",
	"GOROOT_1_10_X64=/usr/local/go1.10",
	"SYSTEM_DEFINITIONID=1",
	"JAVA_HOME=/usr/lib/jvm/zulu-8-azure-amd64",
	"SYSTEM_STAGEID=6884a131-87da-5381-61f3-d7acc3b91d76",
	"AGENT_DISABLELOGPLUGIN_TESTFILEPUBLISHERPLUGIN=true",
	"VCPKG_INSTALLATION_ROOT=/usr/local/share/vcpkg",
	"JAVA_HOME_7_X64=/usr/lib/jvm/zulu-7-azure-amd64",
	"JAVA_HOME_12_X64=/usr/lib/jvm/zulu-12-azure-amd64",
	"SYSTEM_ENABLEACCESSTOKEN=SecretVariable",
	"LANG=en_US.UTF-8",
	"SYSTEM_TASKINSTANCENAME=CmdLine",
	"SYSTEM_PHASEDISPLAYNAME=Build",
	"SYSTEM_SERVERTYPE=Hosted",
	"BUILD_REPOSITORY_NAME=PhilippHeuer/azure-pipeline-test",
	"GOROOT_1_12_X64=/usr/local/go1.12",
	"BUILD_REPOSITORY_URI=https://github.com/PhilippHeuer/azure-pipeline-test",
	"PIPELINE_WORKSPACE=/home/vsts/work/1",
	"BUILD_DEFINITIONNAME=PhilippHeuer.azure-pipeline-test",
	"AGENT_WORKFOLDER=/home/vsts/work",
	"SYSTEM_JOBNAME=Build",
	"BUILD_REQUESTEDFOR=Philipp Heuer",
	"SYSTEM_ARTIFACTSDIRECTORY=/home/vsts/work/1/a",
	"SYSTEM_TIMELINEID=74fa3252-420c-48e3-8a50-1e7827a17aa6",
	"SHLVL=1",
	"AGENT_ID=5",
	"BOOST_ROOT_1_69_0=/usr/local/share/boost/1.69.0",
	"M2_HOME=/usr/share/apache-maven-3.6.1",
	"HOME=/home/vsts",
	"GOROOT=/usr/local/go1.12",
	"AGENT_RETAINDEFAULTENCODING=false",
	"CI=true",
	"SYSTEM_JOBPOSITIONINPHASE=1",
	"BUILD_REQUESTEDFORID=daed2ec0-0b40-4031-b5b4-d5d1a9356542",
	"BUILD_ARTIFACTSTAGINGDIRECTORY=/home/vsts/work/1/a",
	"BUILD_BINARIESDIRECTORY=/home/vsts/work/1/b",
	"BUILD_BUILDID=8",
	"SYSTEM_TASKINSTANCEID=45862c6e-83c5-515f-73e2-e7009eff9f9b",
	"BUILD_SOURCEVERSION=7cb9c91a16950e79d209b710eb33ad56db564e0c",
	"CHROME_BIN=/usr/bin/google-chrome",
	"BOOST_ROOT=/usr/local/share/boost/1.69.0",
	"SYSTEM_DEFAULTWORKINGDIRECTORY=/home/vsts/work/1/s",
	"SYSTEM_JOBID=3dc8fd7e-4368-5a92-293e-d53cefc8c4b3",
	"SYSTEM_TOTALJOBSINPHASE=1",
	"AGENT_NAME=Hosted Agent",
	"SYSTEM_STAGEDISPLAYNAME=Build image",
	"BUILD_DEFINITIONVERSION=1",
	"SYSTEM_PLANID=74fa3252-420c-48e3-8a50-1e7827a17aa6",
	"SYSTEM_PHASEID=a11efe29-9b58-5a6c-3fa4-3e36996dcbd8",
	"SYSTEM_COLLECTIONID=8c7f0be9-cf3c-4627-9df8-fa7d1cb80b37",
	"TASK_DISPLAYNAME=CmdLine",
	"AGENT_JOBSTATUS=Succeeded",
	"ENDPOINT_URL_SYSTEMVSSCONNECTION=https://heuer.visualstudio.com/",
	"BUILD_BUILDNUMBER=20190811.8",
	"SYSTEM_PHASENAME=Build",
	"BUILD_CONTAINERID=3045488",
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
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ID", "5")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_NAME", "fv-az679")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_VERSION", "2.155.1")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ARCH", "linux/amd64")
	// - pipeline
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_TRIGGER", "push")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_NAME", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_SLUG", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_NAME", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_SLUG", "build")
	// - project
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_ID", "d1b384a8-f33f-427d-86fd-f021826a54ea")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_NAME", "azure-test")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_SLUG", "azure-test")
}
