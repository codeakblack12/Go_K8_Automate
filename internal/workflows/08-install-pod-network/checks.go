package installpodnetwork

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

	if _, err := exec.LookPath("kubectl"); err != nil {
		return fmt.Errorf("kubectl is not installed or not in PATH")
	}

	return nil
}

// checkPodNetworkInstalled verifies that network-related pods exist in kube-system.
func (s *Step) checkPodNetworkInstalled() error {
	var command string

	switch s.config.PodNetworkPlugin {
	case "calico":
		command = "kubectl get pods -n kube-system | grep calico"
	case "cilium":
		command = "kubectl get pods -n kube-system | grep cilium"
	default:
		return fmt.Errorf("unsupported pod network plugin: %s", s.config.PodNetworkPlugin)
	}

	cmd := exec.Command("sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify pod network installation: %w", err)
	}

	if strings.TrimSpace(out.String()) == "" {
		return fmt.Errorf("pod network verification failed: no matching plugin pods found")
	}

	return nil
}
