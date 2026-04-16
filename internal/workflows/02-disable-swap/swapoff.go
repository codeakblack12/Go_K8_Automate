package disableswap

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// disableSwap turns off active swap immediately and comments out swap entries
// in /etc/fstab so the change survives reboot.
func (s *Step) disableSwap() error {
	fmt.Println("STEP 2: Disabling swap...")

	swapOffCmd := common.Command{
		Name: "sudo",
		Args: []string{"swapoff", "-a"},
	}

	if err := s.executor.Run(swapOffCmd); err != nil {
		return err
	}

	commentFstabCmd := common.Command{
		Name: "sudo",
		Args: []string{
			"sed",
			"-i",
			`/ swap / s/^/#/`,
			"/etc/fstab",
		},
	}

	if err := s.executor.Run(commentFstabCmd); err != nil {
		return err
	}

	fmt.Println("STEP 2 complete: Swap disabled")
	return nil
}
