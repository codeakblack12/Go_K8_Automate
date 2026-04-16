package updateos

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 01: updating the operating system packages.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new update OS workflow step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the human-readable name of the workflow step.
func (s *Step) Name() string {
	return "01-update-os"
}

// Run validates prerequisites and executes the OS update workflow.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	switch s.config.OSFamily {
	case "ubuntu":
		return s.runUbuntu()
	default:
		return fmt.Errorf("no update workflow defined for OS family: %s", s.config.OSFamily)
	}
}
