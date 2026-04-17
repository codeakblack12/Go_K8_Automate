package initializecluster

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// initControlPlane runs kubeadm init to bootstrap the control-plane node.
func (s *Step) initControlPlane() error {
	fmt.Println("STEP 6: Initializing Kubernetes cluster...")

	args := []string{
		"kubeadm",
		"init",
		"--apiserver-advertise-address", s.config.APIServerAddress,
		"--pod-network-cidr", s.config.PodNetworkCIDR,
	}

	if s.config.KubernetesVersion != "" {
		args = append(args, "--kubernetes-version", s.config.KubernetesVersion)
	}

	initCmd := common.Command{
		Name: "sudo",
		Args: args,
	}

	if err := s.executor.Run(initCmd); err != nil {
		return err
	}

	fmt.Println("STEP 6 complete: Kubernetes control plane initialized")
	return nil
}
