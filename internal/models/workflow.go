package models

// Step defines the behavior required for any workflow step.
type Workflow interface {
	Name() string
	Run() error
}
