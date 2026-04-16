package installk8scomponents

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 05: installing kubeadm, kubelet, and kubectl.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new install-k8s-components step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the step name.
func (s *Step) Name() string {
	return "05-install-k8s-components"
}

// Run validates prerequisites, installs Kubernetes components,
// and verifies that they are available.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	switch s.config.OSFamily {
	case "ubuntu":
		if err := s.installUbuntuPackages(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("no Kubernetes component install workflow defined for OS family: %s", s.config.OSFamily)
	}

	if err := s.checkInstalled(); err != nil {
		return err
	}

	return nil
}
