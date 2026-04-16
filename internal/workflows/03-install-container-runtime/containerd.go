package installcontainerruntime

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// installContainerdUbuntu installs containerd on Ubuntu/Debian systems.
func (s *Step) installContainerdUbuntu() error {
	fmt.Println("STEP 3: Installing container runtime (containerd)...")

	updateCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-get", "update"},
	}
	if err := s.executor.Run(updateCmd); err != nil {
		return err
	}

	installCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-get", "install", "-y", "containerd"},
	}
	if err := s.executor.Run(installCmd); err != nil {
		return err
	}

	fmt.Println("STEP 3 complete: containerd installed")
	return nil
}
