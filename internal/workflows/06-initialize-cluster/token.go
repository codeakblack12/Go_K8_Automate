package initializecluster

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// createJoinCommand generates the kubeadm join command for worker nodes.
func (s *Step) createJoinCommand() error {
	cmd := exec.Command("sudo", "kubeadm", "token", "create", "--print-join-command")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate join command: %w", err)
	}

	s.joinCommand = strings.TrimSpace(out.String())
	if s.joinCommand == "" {
		return fmt.Errorf("generated join command is empty")
	}

	return nil
}
