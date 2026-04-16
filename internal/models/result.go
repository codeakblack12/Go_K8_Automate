package models

// Result captures the outcome of a workflow step.
type Result struct {
	WorkflowName string
	Success      bool
	Message      string
}
