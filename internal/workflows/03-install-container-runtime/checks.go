package installcontainerruntime

import (
	"fmt"
	"os/exec"
)

// checkPrerequisites validates whether the step has what it needs.
func (s *Step) checkPrerequisites() error {
	if s.config == nil {
		return fmt.Errorf("missing configuration")
	}

	if s.executor == nil {
		return fmt.Errorf("missing local executor")
	}

	return nil
}

// checkContainerdInstalled verifies that containerd is available in PATH.
func (s *Step) checkContainerdInstalled() error {
	if _, err := exec.LookPath("containerd"); err != nil {
		return fmt.Errorf("containerd installation verification failed: %w", err)
	}

	return nil
}
