package ncispec

// NCISpec is a common interface for all versions of the spec
type NCISpec interface {
	Validate() error
}

type PipelineTrigger string

const (
	PipelineTriggerCLI         PipelineTrigger = "cli"
	PipelineTriggerManual      PipelineTrigger = "manual"
	PipelineTriggerPush        PipelineTrigger = "push"
	PipelineTriggerPullRequest PipelineTrigger = "pull_request"
	PipelineTriggerSchedule    PipelineTrigger = "schedule"
	PipelineTriggerBuild       PipelineTrigger = "build" // triggered by the completion of a different build
	PipelineTriggerUnknown     PipelineTrigger = "unknown"
)

const (
	PipelineStageDefault = "default"
	PipelineJobDefault   = "default"
)
