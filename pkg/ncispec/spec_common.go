package ncispec

// NCISpec is a common interface for all versions of the spec
type NCISpec interface {
	Validate() error
}

const (
	PipelineTriggerCLI         = "cli"
	PipelineTriggerManual      = "manual"
	PipelineTriggerPush        = "push"
	PipelineTriggerPullRequest = "pull_request"
	PipelineTriggerSchedule    = "schedule"
	PipelineTriggerBuild       = "build" // triggered by the completion of a different build
	PipelineTriggerUnknown     = "unknown"
)

const (
	PipelineStageDefault = "default"
	PipelineJobDefault   = "default"
)
