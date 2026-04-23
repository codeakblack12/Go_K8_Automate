package joincontrolplane

import (
	"Go_K8_Automate/internal/executor/common"
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
	joinCode := strings.TrimSpace(s.config.JoinCode)

	if joinCmd == "" && joinCode == "" {
		return fmt.Errorf("control-plane node requires either control-plane join command or shared join code")
	}

	if joinCmd == "" && strings.TrimSpace(s.config.JoinServiceBaseURL) == "" {
		return fmt.Errorf("join service base URL is required when using shared join code")
	}

	return nil
}

func (s *Step) ensureContainerRuntimeReady() error {
	checkCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"sudo systemctl is-active --quiet containerd",
		},
	}

	if err := s.executor.Run(checkCmd); err != nil {
		restartCmd := common.Command{
			Name: "sh",
			Args: []string{
				"-c",
				"sudo systemctl restart containerd && sudo systemctl enable containerd",
			},
		}

		if err := s.executor.Run(restartCmd); err != nil {
			return fmt.Errorf("failed to restart containerd: %w", err)
		}
	}

	verifyCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"test -S /var/run/containerd/containerd.sock",
		},
	}

	if err := s.executor.Run(verifyCmd); err != nil {
		return fmt.Errorf("containerd socket is unavailable: %w", err)
	}

	return nil
}
