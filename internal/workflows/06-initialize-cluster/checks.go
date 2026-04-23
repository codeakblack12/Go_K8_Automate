package initializecluster

import (
	"fmt"
	"os/exec"
	"strings"
)

var controlPlanePorts = []int{
	6443,  // kube-apiserver
	2379,  // etcd client
	2380,  // etcd peer
	10250, // kubelet
	10257, // kube-controller-manager
	10259, // kube-scheduler
}

// checkExistingControlPlane looks for signs that this node already has
// control-plane state from a previous kubeadm run.
// func (s *Step) checkExistingControlPlane() (bool, error) {
// 	if _, err := os.Stat("/etc/kubernetes/manifests/kube-apiserver.yaml"); err == nil {
// 		return true, nil
// 	}

// 	if _, err := os.Stat("/etc/kubernetes/manifests/etcd.yaml"); err == nil {
// 		return true, nil
// 	}

// 	// Check control-plane ports
// 	cmd := exec.Command("sh", "-c", "ss -ltnp | awk 'NR>1 {print $4}'")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	if err := cmd.Run(); err != nil {
// 		// Non-fatal; just log via error and assume no ports
// 		return false, fmt.Errorf("failed to inspect ports: %w", err)
// 	}

// 	inUse := out.String()

// 	for _, port := range controlPlanePorts {
// 		token := fmt.Sprintf(":%d", port)
// 		if strings.Contains(inUse, token) {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }

// resetNode runs kubeadm reset and basic cleanup.
// func (s *Step) resetNode() error {
// 	fmt.Println("Detected existing control-plane state. Resetting node with 'kubeadm reset -f'...")

// 	resetCmd := exec.Command("sudo", "kubeadm", "reset", "-f")
// 	if err := resetCmd.Run(); err != nil {
// 		return fmt.Errorf("kubeadm reset failed: %w", err)
// 	}

// 	// Best-effort cleanup; ignore individual errors.
// 	_ = os.RemoveAll("/etc/cni/net.d")
// 	_ = os.RemoveAll("/var/lib/etcd")
// 	_ = os.RemoveAll("/etc/kubernetes/manifests")
// 	_ = os.RemoveAll(os.ExpandEnv("$HOME/.kube/config"))

// 	_ = exec.Command("sudo", "systemctl", "restart", "containerd").Run()
// 	_ = exec.Command("sudo", "systemctl", "restart", "kubelet").Run()

// 	fmt.Println("Node reset completed")
// 	return nil
// }

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

	if _, err := exec.LookPath("kubelet"); err != nil {
		return fmt.Errorf("kubelet is not installed or not in PATH")
	}

	if strings.TrimSpace(s.config.APIServerAddress) == "" {
		return fmt.Errorf("API server address is required")
	}

	if strings.TrimSpace(s.config.ControlPlaneEndpoint) == "" {
		return fmt.Errorf("control plane endpoint is required")
	}

	return nil
}

// checkClusterInitialized verifies that admin.conf exists after kubeadm init.
// func (s *Step) checkClusterInitialized() error {
// 	cmd := exec.Command("sh", "-c", "test -f /etc/kubernetes/admin.conf && echo ok")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out

// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("failed to verify cluster initialization: %w", err)
// 	}

// 	if !strings.Contains(out.String(), "ok") {
// 		return fmt.Errorf("cluster initialization verification failed: /etc/kubernetes/admin.conf not found")
// 	}

// 	return nil
// }
