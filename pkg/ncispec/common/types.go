package common

const (
	PipelineTriggerCLI          = "cli"
	PipelineTriggerManual       = "manual"
	PipelineTriggerPush         = "push"
	PipelineTriggerMergeRequest = "merge_request"
	PipelineTriggerSchedule     = "schedule"
	PipelineTriggerBuild        = "build" // triggered by the completion of a different build
	PipelineTriggerUnknown      = "unknown"
)

const (
	PipelineStageDefault = "default"
	PipelineJobDefault   = "default"
)
