package joinworkernode

import (
	"fmt"
	"strings"

	"Go_K8_Automate/internal/executor/common"
)

// joinCluster runs the kubeadm join command received from the control-plane node.
func (s *Step) joinNode() error {
	joinCmd := strings.TrimSpace(s.config.JoinCommand)
	if joinCmd == "" {
		return fmt.Errorf("worker join command is empty")
	}

	if strings.Contains(joinCmd, "--control-plane") {
		return fmt.Errorf("resolved worker join command unexpectedly contains --control-plane")
	}

	fmt.Println("Joining node as worker...")

	command := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"sudo " + joinCmd,
		},
	}

	if err := s.executor.Run(command); err != nil {
		return err
	}

	fmt.Println("Worker node joined successfully")
	return nil
}
