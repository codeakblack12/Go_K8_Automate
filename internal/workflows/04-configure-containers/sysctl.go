package configurecontainers

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// configureSysctl persists and applies kernel network settings
// required by Kubernetes.
func (s *Step) configureSysctl() error {
	fmt.Println("Configuring sysctl settings for Kubernetes...")

	writeSysctlCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			`cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf > /dev/null
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF`,
		},
	}
	if err := s.executor.Run(writeSysctlCmd); err != nil {
		return err
	}

	applySysctlCmd := common.Command{
		Name: "sudo",
		Args: []string{"sysctl", "--system"},
	}
	if err := s.executor.Run(applySysctlCmd); err != nil {
		return err
	}

	return nil
}
