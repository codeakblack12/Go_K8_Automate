package initializecluster

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"Go_K8_Automate/internal/executor/common"
)

func (s *Step) initControlPlane() error {
	fmt.Println("STEP 6: Initializing Kubernetes cluster...")

	configPath, err := s.writeKubeadmConfig()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configPath)
	if err == nil {
		fmt.Printf("kubeadm config:\n%s\n", string(data))
	}

	initCmd := common.Command{
		Name: "sudo",
		Args: []string{
			"kubeadm",
			"init",
			"--config", configPath,
			"--upload-certs",
		},
	}

	if err := s.executor.Run(initCmd); err != nil {
		return fmt.Errorf("failed to initialize control plane: %w", err)
	}

	fmt.Println("STEP 6 complete: Kubernetes control plane initialized")
	return nil
}

func (s *Step) writeKubeadmConfig() (string, error) {
	if strings.TrimSpace(s.config.APIServerAddress) == "" {
		return "", fmt.Errorf("API server address is required")
	}

	if strings.TrimSpace(s.config.ControlPlaneEndpoint) == "" {
		return "", fmt.Errorf("control plane endpoint is required for HA-ready initialization")
	}

	k8sVersion := strings.TrimSpace(s.config.KubernetesVersion)
	if k8sVersion == "" {
		k8sVersion = "stable"
	}

	content := fmt.Sprintf(
		"apiVersion: kubeadm.k8s.io/v1beta4\n"+
			"kind: InitConfiguration\n"+
			"localAPIEndpoint:\n"+
			"  advertiseAddress: %s\n"+
			"  bindPort: 6443\n"+
			"nodeRegistration:\n"+
			"  criSocket: unix:///var/run/containerd/containerd.sock\n"+
			"---\n"+
			"apiVersion: kubeadm.k8s.io/v1beta4\n"+
			"kind: ClusterConfiguration\n"+
			"kubernetesVersion: %s\n"+
			"controlPlaneEndpoint: \"%s\"\n"+
			"networking:\n"+
			"  podSubnet: \"%s\"\n",
		s.config.APIServerAddress,
		k8sVersion,
		s.config.ControlPlaneEndpoint,
		s.config.PodNetworkCIDR,
	)

	configPath := filepath.Join(os.TempDir(), "kubeadm-init.yaml")
	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		return "", fmt.Errorf("failed to write kubeadm config file: %w", err)
	}

	fmt.Printf("Wrote kubeadm config to %s\n", configPath)
	return configPath, nil
}
