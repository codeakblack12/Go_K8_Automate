package configurekubectlaccess

import (
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 07: configuring kubectl access for the current user.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new configure-kubectl-access step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "07-configure-kubectl-access"
}

// Run validates prerequisites, configures kubeconfig access,
// and verifies that kubectl config is available.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.configureKubeconfig(); err != nil {
		return err
	}

	if err := s.checkKubectlAccess(); err != nil {
		return err
	}

	return nil
}
