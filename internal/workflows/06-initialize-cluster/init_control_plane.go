package initializecluster

import (
	"fmt"
	"strings"

	"Go_K8_Automate/internal/executor/common"
)

// initControlPlane runs kubeadm init to bootstrap the control-plane node.
func (s *Step) initControlPlane() error {
	fmt.Println("STEP 6: Initializing Kubernetes cluster...")

	if strings.TrimSpace(s.kubeadmConfigPath) == "" {
		return fmt.Errorf("kubeadm config path is empty")
	}

	command := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			fmt.Sprintf("sudo kubeadm init --config %s --upload-certs", s.kubeadmConfigPath),
		},
	}

	if err := s.executor.Run(command); err != nil {
		return fmt.Errorf("failed to initialize control plane: %w", err)
	}

	fmt.Println("STEP 6 complete: Kubernetes control plane initialized")
	return nil
}
