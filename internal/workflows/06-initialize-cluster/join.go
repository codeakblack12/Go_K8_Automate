package initializecluster

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func (s *Step) createJoinCommand() error {
	cmd := exec.Command("sudo", "kubeadm", "token", "create", "--print-join-command")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create worker join command: %w: %s", err, stderr.String())
	}

	joinCmd := strings.TrimSpace(out.String())
	if joinCmd == "" {
		return fmt.Errorf("generated worker join command is empty")
	}

	s.joinCommand = joinCmd
	return nil
}
