package initializecluster

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// checkPrerequisites validates whether this step can run.
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

	return nil
}

// checkClusterInitialized verifies that admin.conf exists after kubeadm init.
func (s *Step) checkClusterInitialized() error {
	cmd := exec.Command("sh", "-c", "test -f /etc/kubernetes/admin.conf && echo ok")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify cluster initialization: %w", err)
	}

	if !strings.Contains(out.String(), "ok") {
		return fmt.Errorf("cluster initialization verification failed: /etc/kubernetes/admin.conf not found")
	}

	return nil
}
