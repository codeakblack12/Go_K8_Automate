package installpodnetwork

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 08: installing the pod network plugin.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new install-pod-network step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "08-install-pod-network"
}

// Run validates prerequisites, installs the selected network plugin,
// and verifies that the plugin pods are present.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	switch s.config.PodNetworkPlugin {
	case "calico":
		if err := s.installCalico(); err != nil {
			return err
		}
	case "cilium":
		if err := s.installCilium(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported pod network plugin: %s", s.config.PodNetworkPlugin)
	}

	if err := s.checkPodNetworkInstalled(); err != nil {
		return err
	}

	return nil
}
