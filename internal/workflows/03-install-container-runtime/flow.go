package installcontainerruntime

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 03: installing the container runtime.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new install-container-runtime workflow step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "03-install-container-runtime"
}

// Run validates prerequisites, installs the runtime, and verifies it.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	switch s.config.OSFamily {
	case "ubuntu":
		if err := s.installContainerdUbuntu(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("no container runtime install workflow defined for OS family: %s", s.config.OSFamily)
	}

	if err := s.checkContainerdInstalled(); err != nil {
		return err
	}

	return nil
}
