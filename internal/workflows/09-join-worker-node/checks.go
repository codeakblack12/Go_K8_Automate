package joinworkernode

import (
	// "bytes"
	"fmt"
	"os/exec"
	"strings"
)

// checkPrerequisites validates whether the worker node can proceed
// with resolving the join code and joining the cluster.
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

	joinCmd := strings.TrimSpace(s.config.JoinCommand)
	joinCode := strings.TrimSpace(s.config.JoinCode)

	if joinCmd == "" && joinCode == "" {
		return fmt.Errorf("worker node requires either join command or join code")
	}

	if joinCmd == "" && strings.TrimSpace(s.config.JoinServiceBaseURL) == "" {
		return fmt.Errorf("join service base URL is required when using join code")
	}

	return nil
}
