package installpodnetwork

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

	if _, err := exec.LookPath("kubectl"); err != nil {
		return fmt.Errorf("kubectl is not installed or not in PATH")
	}

	return nil
}

// checkPodNetworkInstalled waits for the selected CNI resources to roll out.
func (s *Step) checkPodNetworkInstalled() error {
	switch s.config.PodNetworkPlugin {
	case "calico":
		if err := exec.Command(
			"kubectl", "rollout", "status",
			"daemonset/calico-node",
			"-n", "kube-system",
			"--timeout=180s",
		).Run(); err != nil {
			return fmt.Errorf("failed waiting for calico-node daemonset rollout: %w", err)
		}

		if err := exec.Command(
			"kubectl", "rollout", "status",
			"deployment/calico-kube-controllers",
			"-n", "kube-system",
			"--timeout=180s",
		).Run(); err != nil {
			return fmt.Errorf("failed waiting for calico-kube-controllers rollout: %w", err)
		}

		return nil

	case "cilium":
		if err := exec.Command(
			"kubectl", "rollout", "status",
			"daemonset/cilium",
			"-n", "kube-system",
			"--timeout=180s",
		).Run(); err != nil {
			return fmt.Errorf("failed waiting for cilium daemonset rollout: %w", err)
		}

		if err := exec.Command(
			"kubectl", "rollout", "status",
			"deployment/cilium-operator",
			"-n", "kube-system",
			"--timeout=180s",
		).Run(); err != nil {
			return fmt.Errorf("failed waiting for cilium-operator rollout: %w", err)
		}

		return nil

	default:
		return fmt.Errorf("unsupported pod network plugin: %s", s.config.PodNetworkPlugin)
	}
}
