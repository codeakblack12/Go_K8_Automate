package initializecluster

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func (s *Step) writeKubeadmConfig() error {
// 	k8sVersion := strings.TrimSpace(s.config.KubernetesVersion)
// 	if k8sVersion == "" {
// 		k8sVersion = "stable"
// 	}

// 	content := fmt.Sprintf(`apiVersion: kubeadm.k8s.io/v1beta4
// 		kind: InitConfiguration
// 		localAPIEndpoint:
// 		advertiseAddress: %s
// 		bindPort: 6443
// 		nodeRegistration:
// 		criSocket: unix:///var/run/containerd/containerd.sock
// 		---
// 		apiVersion: kubeadm.k8s.io/v1beta4
// 		kind: ClusterConfiguration
// 		kubernetesVersion: %s
// 		controlPlaneEndpoint: "%s"
// 		networking:
// 		podSubnet: "%s"
// 	`,
// 		s.config.APIServerAddress,
// 		k8sVersion,
// 		s.config.ControlPlaneEndpoint,
// 		s.config.PodNetworkCIDR,
// 	)

// 	path := filepath.Join(os.TempDir(), "kubeadm-init.yaml")
// 	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
// 		return fmt.Errorf("failed to write kubeadm config file: %w", err)
// 	}

// 	s.kubeadmConfigPath = path
// 	return nil
// }
