package joincontrolplane

import (
	"fmt"
	"os/exec"
	"strings"
)

func (s *Step) checkPrerequisites() error {
	if s.config == nil {
		return fmt.Errorf("missing configuration")
	}

	if s.executor == nil {
		return fmt.Errorf("missing local executor")
	}

	if _, err := exec.LookPath("kubeadm"); err != nil {
		return fmt.Errorf("kubeadm is not installed or not in PATH")
	}

	if _, err := exec.LookPath("kubelet"); err != nil {
		return fmt.Errorf("kubelet is not installed or not in PATH")
	}

	joinCmd := strings.TrimSpace(s.config.ControlPlaneJoinCommand)
	joinCode := strings.TrimSpace(s.config.ControlPlaneJoinCode)

	if joinCmd == "" && joinCode == "" {
		return fmt.Errorf("control-plane node requires either control-plane join command or control-plane join code")
	}

	if joinCmd == "" && strings.TrimSpace(s.config.JoinServiceBaseURL) == "" {
		return fmt.Errorf("join service base URL is required when using control-plane join code")
	}

	return nil
}
