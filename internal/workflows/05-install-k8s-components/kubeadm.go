package installk8scomponents

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// installUbuntuPackages installs kubeadm and related Kubernetes packages
// using the official Kubernetes APT repository.
func (s *Step) installUbuntuPackages() error {
	fmt.Println("STEP 5: Installing Kubernetes components...")

	installPrereqsCmd := common.Command{
		Name: "sudo",
		Args: []string{
			"apt-get", "install", "-y",
			"apt-transport-https",
			"ca-certificates",
			"curl",
			"gpg",
		},
	}
	if err := s.executor.Run(installPrereqsCmd); err != nil {
		return err
	}

	createKeyringDirCmd := common.Command{
		Name: "sudo",
		Args: []string{"mkdir", "-p", "-m", "755", "/etc/apt/keyrings"},
	}
	if err := s.executor.Run(createKeyringDirCmd); err != nil {
		return err
	}

	addRepoKeyCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			fmt.Sprintf(
				"curl -fsSL https://pkgs.k8s.io/core:/stable:/%s/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg",
				s.config.KubernetesRepoVersion,
			),
		},
	}
	if err := s.executor.Run(addRepoKeyCmd); err != nil {
		return err
	}

	addRepoCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			fmt.Sprintf(
				"echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/%s/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list > /dev/null",
				s.config.KubernetesRepoVersion,
			),
		},
	}
	if err := s.executor.Run(addRepoCmd); err != nil {
		return err
	}

	updateCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-get", "update"},
	}
	if err := s.executor.Run(updateCmd); err != nil {
		return err
	}

	installK8sCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-get", "install", "-y", "kubelet", "kubeadm", "kubectl"},
	}
	if err := s.executor.Run(installK8sCmd); err != nil {
		return err
	}

	holdCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-mark", "hold", "kubelet", "kubeadm", "kubectl"},
	}
	if err := s.executor.Run(holdCmd); err != nil {
		return err
	}

	enableKubeletCmd := common.Command{
		Name: "sudo",
		Args: []string{"systemctl", "enable", "--now", "kubelet"},
	}
	if err := s.executor.Run(enableKubeletCmd); err != nil {
		return err
	}

	fmt.Println("STEP 5 complete: kubelet, kubeadm, and kubectl installed")
	return nil
}
