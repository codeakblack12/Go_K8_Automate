package installk8scomponents

import (
	"fmt"
	"os/exec"
)

// checkPrerequisites validates whether this step can run.
func (s *Step) checkPrerequisites() error {
	if s.config == nil {
		return fmt.Errorf("missing configuration")
	}

	if s.executor == nil {
		return fmt.Errorf("missing local executor")
	}

	return nil
}

// checkInstalled verifies that kubeadm, kubelet, and kubectl are installed.
func (s *Step) checkInstalled() error {
	if _, err := exec.LookPath("kubeadm"); err != nil {
		return fmt.Errorf("kubeadm installation verification failed: %w", err)
	}

	if _, err := exec.LookPath("kubelet"); err != nil {
		return fmt.Errorf("kubelet installation verification failed: %w", err)
	}

	if _, err := exec.LookPath("kubectl"); err != nil {
		return fmt.Errorf("kubectl installation verification failed: %w", err)
	}

	return nil
}
