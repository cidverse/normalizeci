package gitlabci

import (
	"testing"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"
)

var testEnvironment = []string{
	"FF_CMD_DISABLE_DELAYED_ERROR_LEVEL_EXPANSION=false",
	"FF_USE_LEGACY_BUILDS_DIR_FOR_DOCKER=false",
	"FF_USE_LEGACY_VOLUMES_MOUNTING_ORDER=false",
	"CI_BUILDS_DIR=/builds",
	"CI_PROJECT_DIR=/builds/PhilippHeuer/citest",
	"CI_CONCURRENT_ID=18",
	"CI_CONCURRENT_PROJECT_ID=0",
	"CI_SERVER=yes",
	"CI_PIPELINE_ID=77135359",
	"CI_PIPELINE_URL=https://gitlab.com/PhilippHeuer/citest/pipelines/77135359",
	"CI_JOB_ID=275323310",
	"CI_JOB_URL=https://gitlab.com/PhilippHeuer/citest/-/jobs/275323310",
	"CI_JOB_TOKEN=secret",
	"CI_BUILD_ID=275323310",
	"CI_BUILD_TOKEN=secret",
	"CI_REGISTRY_USER=gitlab-ci-token",
	"CI_REGISTRY_PASSWORD=secret",
	"CI_REPOSITORY_URL=https://gitlab-ci-token:[MASKED]@gitlab.com/PhilippHeuer/citest.git",
	"CI=true",
	"GITLAB_CI=true",
	"GITLAB_FEATURES=audit_events,burndown_charts,code_owners,contribution_analytics,elastic_search,export_issues,group_bulk_edit,group_burndown_charts,group_webhooks,issuable_default_templates,issue_board_focus_mode,issue_weights,jenkins_integration,ldap_group_sync,member_lock,merge_request_approvers,multiple_ldap_servers,multiple_issue_assignees,multiple_merge_request_assignees,push_rules,protected_refs_for_users,related_issues,repository_mirrors,repository_size_limit,scoped_issue_board,usage_quotas,visual_review_app,admin_audit_log,auditor_user,blocking_merge_requests,board_assignee_lists,board_milestone_lists,cross_project_pipelines,custom_file_templates,custom_file_templates_for_namespace,email_additional_text,db_load_balancing,deploy_board,extended_audit_events,file_locks,geo,github_project_service_integration,jira_dev_panel_integration,scoped_labels,ldap_group_sync_filter,multiple_clusters,multiple_group_issue_boards,multiple_approval_rules,merge_request_performance_metrics,object_storage,group_saml,service_desk,smartcard_auth,unprotection_restrictions,reject_unsigned_commits,commit_committer_check,external_authorization_service_api_management,ci_cd_projects,default_project_deletion_protection,protected_environments,custom_project_templates,packages,code_owner_approval_required,feature_flags,batch_comments,issues_analytics,merge_pipelines,merge_trains,design_management,operations_dashboard,dependency_proxy,metrics_reports,custom_prometheus_metrics,required_ci_templates,project_aliases,cycle_analytics_for_groups,security_dashboard,dependency_scanning,dependency_list,license_management,sast,container_scanning,cluster_health,dast,epics,pod_logs,pseudonymizer,prometheus_alerts,tracing,insights,web_ide_terminal,incident_management,report_approver_rules,group_ip_restriction",
	"CI_SERVER_HOST=gitlab.com",
	"CI_SERVER_NAME=GitLab",
	"CI_SERVER_VERSION=12.2.0-pre",
	"CI_SERVER_VERSION_MAJOR=12",
	"CI_SERVER_VERSION_MINOR=2",
	"CI_SERVER_VERSION_PATCH=0",
	"CI_SERVER_REVISION=36c4e152270",
	"CI_JOB_NAME=build",
	"CI_JOB_STAGE=build",
	"CI_COMMIT_SHA=c2cc4ef1015fce6cd4ba80ca9f43077e8bbc843c",
	"CI_COMMIT_SHORT_SHA=c2cc4ef1",
	"CI_COMMIT_BEFORE_SHA=a9a3893b44a4e5d2292926e15ff1dc39200beae8",
	"CI_COMMIT_REF_NAME=master",
	"CI_COMMIT_REF_SLUG=master",
	"CI_NODE_TOTAL=1",
	"CI_BUILD_REF=c2cc4ef1015fce6cd4ba80ca9f43077e8bbc843c",
	"CI_BUILD_BEFORE_SHA=a9a3893b44a4e5d2292926e15ff1dc39200beae8",
	"CI_BUILD_REF_NAME=master",
	"CI_BUILD_REF_SLUG=master",
	"CI_BUILD_NAME=build",
	"CI_BUILD_STAGE=build",
	"CI_PROJECT_ID=13882743",
	"CI_PROJECT_NAME=citest",
	"CI_PROJECT_PATH=PhilippHeuer/citest",
	"CI_PROJECT_PATH_SLUG=philippheuer-citest",
	"CI_PROJECT_NAMESPACE=PhilippHeuer",
	"CI_PROJECT_URL=https://gitlab.com/PhilippHeuer/citest",
	"CI_PROJECT_VISIBILITY=public",
	"CI_PAGES_DOMAIN=gitlab.io",
	"CI_PAGES_URL=https://philippheuer.gitlab.io/citest",
	"CI_REGISTRY=registry.gitlab.com",
	"CI_REGISTRY_IMAGE=registry.gitlab.com/philippheuer/citest",
	"CI_API_V4_URL=https://gitlab.com/api/v4",
	"CI_PIPELINE_IID=5",
	"CI_CONFIG_PATH=.gitlab-ci.yml",
	"CI_PIPELINE_SOURCE=push",
	"CI_COMMIT_MESSAGE=Update .gitlab-ci.yml",
	"CI_COMMIT_TITLE=Update .gitlab-ci.yml",
	"CI_COMMIT_DESCRIPTION=",
	"CI_COMMIT_REF_PROTECTED=true",
	"CI_RUNNER_ID=380987",
	"CI_RUNNER_DESCRIPTION=shared-runners-manager-6.gitlab.com",
	"CI_RUNNER_TAGS=docker, gce",
	"GITLAB_USER_ID=1193975",
	"GITLAB_USER_EMAIL=git@philippheuer.me",
	"GITLAB_USER_LOGIN=PhilippHeuer",
	"GITLAB_USER_NAME=Philipp Heuer",
	"CI_DISPOSABLE_ENVIRONMENT=true",
	"CI_RUNNER_VERSION=12.1.0",
	"CI_RUNNER_REVISION=de7731dd",
	"CI_RUNNER_EXECUTABLE_ARCH=linux/amd64",
	"GIT_LFS_SKIP_SMUDGE=1",
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(testEnvironment) != true {
		t.Errorf("Check should succeed with the provided gitlab ci sample data!")
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
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_NAME", "GitLab")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_HOST", "gitlab.com")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_VERSION", "12.2.0-pre")
	// - worker
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ID", "380987")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_NAME", "shared-runners-manager-6.gitlab.com")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_VERSION", "12.1.0")
	common.AssertThatEnvEquals(t, normalized, "NCI_WORKER_ARCH", "linux/amd64")
	// - pipeline
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_TRIGGER", "push")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_NAME", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_STAGE_SLUG", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_NAME", "build")
	common.AssertThatEnvEquals(t, normalized, "NCI_PIPELINE_JOB_SLUG", "build")
	// - container registry
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_HOST", "registry.gitlab.com")
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_REPOSITORY", "registry.gitlab.com/philippheuer/citest")
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_USERNAME", "gitlab-ci-token")
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_PASSWORD", "secret")
	// - project
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_ID", "13882743")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_NAME", "citest")
	common.AssertThatEnvEquals(t, normalized, "NCI_PROJECT_SLUG", "philippheuer-citest")
}
