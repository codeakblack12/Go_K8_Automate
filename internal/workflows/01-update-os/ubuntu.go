package updateos

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// runUbuntu performs the Ubuntu/Debian package update workflow.
// It refreshes package metadata first, then upgrades installed packages.
func (s *Step) runUbuntu() error {
	fmt.Println("STEP 1: Updating OS packages for Ubuntu...")

	updateCmd := common.Command{
		Name: "sudo",
		Args: []string{"apt-get", "update"},
	}

	if err := s.executor.Run(updateCmd); err != nil {
		return err
	}

	upgradeCmd := common.Command{
		Name: "sudo",
		Args: []string{
			"env",
			"DEBIAN_FRONTEND=noninteractive",
			"apt-get",
			"upgrade",
			"-y",
		},
	}

	if err := s.executor.Run(upgradeCmd); err != nil {
		return err
	}

	fmt.Println("STEP 1 complete: OS packages updated successfully")
	return nil
}
