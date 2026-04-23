package joincontrolplane

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

func (s *Step) resetNode() error {
	fmt.Println("Stale Kubernetes state detected; resetting node before control-plane join...")

	commands := []common.Command{
		{
			Name: "sh",
			Args: []string{"-c", "sudo kubeadm reset -f"},
		},
		{
			Name: "sh",
			Args: []string{"-c", "sudo systemctl stop kubelet || true"},
		},
		{
			Name: "sh",
			Args: []string{"-c", "sudo rm -rf /etc/kubernetes /var/lib/etcd /etc/cni/net.d"},
		},
		{
			Name: "sh",
			Args: []string{"-c", "sudo systemctl restart containerd || sudo systemctl start containerd"},
		},
	}

	for _, cmd := range commands {
		if err := s.executor.Run(cmd); err != nil {
			return fmt.Errorf("failed to reset node: %w", err)
		}
	}

	fmt.Println("Node reset completed successfully")
	return nil
}
