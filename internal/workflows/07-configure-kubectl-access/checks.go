package configurekubectlaccess

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
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

	if _, err := exec.LookPath("kubectl"); err != nil {
		return fmt.Errorf("kubectl is not installed or not in PATH")
	}

	return nil
}

// checkKubectlAccess verifies that the kubeconfig file exists
// and kubectl can read the current context.
func (s *Step) checkKubectlAccess() error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to determine current user: %w", err)
	}

	kubeConfig := currentUser.HomeDir + "/.kube/config"

	if _, err := os.Stat(kubeConfig); err != nil {
		return fmt.Errorf("kubeconfig verification failed: %w", err)
	}

	cmd := exec.Command("kubectl", "--kubeconfig", kubeConfig, "config", "current-context")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify kubectl access: %w", err)
	}

	if strings.TrimSpace(out.String()) == "" {
		return fmt.Errorf("kubectl access verification failed: current context is empty")
	}

	return nil
}
