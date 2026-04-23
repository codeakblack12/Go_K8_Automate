package joinworkernode

import (
	"bytes"
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

	joinCommand := strings.TrimSpace(s.config.JoinCommand)
	joinCode := strings.TrimSpace(s.config.JoinCode)
	joinServiceURL := strings.TrimSpace(s.config.JoinServiceBaseURL)

	if joinCommand == "" && joinCode == "" {
		return fmt.Errorf("worker node requires either join command or join code")
	}

	if joinCommand == "" && joinCode != "" && joinServiceURL == "" {
		return fmt.Errorf("join service base URL is required when using join code")
	}

	return nil
}

// checkWorkerJoined verifies that kubeadm join created kubelet config.
func (s *Step) checkWorkerJoined() error {
	cmd := exec.Command("sh", "-c", "test -f /etc/kubernetes/kubelet.conf && echo ok")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify worker join: %w", err)
	}

	if !strings.Contains(out.String(), "ok") {
		return fmt.Errorf("worker join verification failed: /etc/kubernetes/kubelet.conf not found")
	}

	return nil
}
