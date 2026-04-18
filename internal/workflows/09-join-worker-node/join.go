package joinworkernode

import (
	"fmt"
	"strings"

	"Go_K8_Automate/internal/executor/common"
)

// joinCluster runs the kubeadm join command received from the control-plane node.
func (s *Step) joinCluster() error {
	fmt.Println("Joining worker node to the Kubernetes cluster...")

	joinCmd := strings.TrimSpace(s.config.JoinCommand)

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

	fmt.Println("Worker node joined the cluster successfully")
	return nil
}
