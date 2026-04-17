package initializecluster

import (
	"fmt"
	"os"

	"Go_K8_Automate/internal/executor/common"
)

// createJoinCommand generates and stores the worker join command.
func (s *Step) createJoinCommand() error {
	fmt.Println("Generating worker join command...")

	joinCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"sudo kubeadm token create --print-join-command | tee join-command.sh > /dev/null",
		},
	}

	if err := s.executor.Run(joinCmd); err != nil {
		return err
	}

	if err := os.Chmod("join-command.sh", 0700); err != nil {
		return fmt.Errorf("failed to set permissions on join-command.sh: %w", err)
	}

	fmt.Println("Worker join command saved to join-command.sh")
	return nil
}
