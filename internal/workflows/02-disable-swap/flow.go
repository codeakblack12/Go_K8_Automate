package disableswap

import (
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 02: disabling swap.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new disable-swap workflow step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "02-disable-swap"
}

// Run validates prerequisites, disables swap, and verifies the result.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.disableSwap(); err != nil {
		return err
	}

	if err := s.checkSwapDisabled(); err != nil {
		return err
	}

	return nil
}
