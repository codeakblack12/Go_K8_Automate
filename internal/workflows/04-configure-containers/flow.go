package configurecontainers

import (
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 04: configuring the container runtime
// and required host networking settings for Kubernetes.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new configure-containers step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the step name.
func (s *Step) Name() string {
	return "04-configure-containers"
}

// Run validates prerequisites, configures kernel modules,
// sysctl settings, containerd, and verifies the result.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.configureKernelModules(); err != nil {
		return err
	}

	if err := s.configureSysctl(); err != nil {
		return err
	}

	if err := s.configureContainerd(); err != nil {
		return err
	}

	if err := s.checkContainerdConfigured(); err != nil {
		return err
	}

	if err := s.checkIPForwardingEnabled(); err != nil {
		return err
	}

	return nil
}
