package configurecontainers

import "Go_K8_Automate/internal/executor/common"

// configureKernelModules loads common kernel modules used by Kubernetes networking
// and persists them across reboots.
func (s *Step) configureKernelModules() error {
	writeModulesCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"printf 'overlay\nbr_netfilter\n' | sudo tee /etc/modules-load.d/k8s.conf > /dev/null",
		},
	}
	if err := s.executor.Run(writeModulesCmd); err != nil {
		return err
	}

	loadOverlayCmd := common.Command{
		Name: "sudo",
		Args: []string{"modprobe", "overlay"},
	}
	if err := s.executor.Run(loadOverlayCmd); err != nil {
		return err
	}

	loadBrNetfilterCmd := common.Command{
		Name: "sudo",
		Args: []string{"modprobe", "br_netfilter"},
	}
	if err := s.executor.Run(loadBrNetfilterCmd); err != nil {
		return err
	}

	return nil
}
