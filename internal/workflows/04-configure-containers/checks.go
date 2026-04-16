package configurecontainers

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

	if _, err := exec.LookPath("containerd"); err != nil {
		return fmt.Errorf("containerd is not installed or not in PATH")
	}

	return nil
}

// checkContainerdConfigured verifies that containerd is configured
// with SystemdCgroup enabled.
func (s *Step) checkContainerdConfigured() error {
	cmd := exec.Command("sh", "-c", "grep 'SystemdCgroup = true' /etc/containerd/config.toml")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify containerd config: %w", err)
	}

	if !strings.Contains(out.String(), "SystemdCgroup = true") {
		return fmt.Errorf("containerd config verification failed: SystemdCgroup is not enabled")
	}

	return nil
}
