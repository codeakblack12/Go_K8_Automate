package installpodnetwork

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// installCalico installs the Calico pod network using the official manifest.
func (s *Step) installCalico() error {
	fmt.Println("STEP 8: Installing Calico pod network...")

	applyCmd := common.Command{
		Name: "kubectl",
		Args: []string{
			"apply",
			"-f",
			"https://raw.githubusercontent.com/projectcalico/calico/v3.31.0/manifests/calico.yaml",
		},
	}

	if err := s.executor.Run(applyCmd); err != nil {
		return err
	}

	fmt.Println("STEP 8 complete: Calico installed")
	return nil
}
