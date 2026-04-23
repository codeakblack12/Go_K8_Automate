package joincontrolplane

import (
	"fmt"
	"strings"

	"Go_K8_Automate/internal/executor/common"
)

func (s *Step) joinControlPlane() error {
	joinCmd := strings.TrimSpace(s.config.ControlPlaneJoinCommand)
	if joinCmd == "" {
		return fmt.Errorf("control-plane join command is empty")
	}

	if !strings.Contains(joinCmd, "--control-plane") {
		return fmt.Errorf("resolved command is not a control-plane join command")
	}

	fmt.Println("Joining node as an additional control-plane...")

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

	fmt.Println("Control-plane node joined successfully")
	return nil
}
