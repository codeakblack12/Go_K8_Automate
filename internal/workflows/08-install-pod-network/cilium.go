package installpodnetwork

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// installCilium installs Cilium using its CLI.
// This assumes cilium CLI is already installed if this option is selected.
func (s *Step) installCilium() error {
	fmt.Println("STEP 8: Installing Cilium pod network...")

	installCmd := common.Command{
		Name: "cilium",
		Args: []string{"install"},
	}

	if err := s.executor.Run(installCmd); err != nil {
		return err
	}

	fmt.Println("STEP 8 complete: Cilium installed")
	return nil
}
