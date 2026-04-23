package joinworkernode

import (
	"fmt"
	"strings"
)

// resolveJoinCode exchanges a short join-code for the full kubeadm join command.
// If JoinCommand is already provided directly, this step does nothing.
func (s *Step) resolveJoinCode() error {
	if strings.TrimSpace(s.config.JoinCommand) != "" {
		return nil
	}

	if strings.TrimSpace(s.config.JoinCode) == "" {
		return fmt.Errorf("join code is empty")
	}

	resp, err := s.joinCodeClient.Resolve(s.config.JoinCode, "worker")
	if err != nil {
		return fmt.Errorf("failed to resolve worker join code: %w", err)
	}

	if strings.TrimSpace(resp.JoinCommand) == "" {
		return fmt.Errorf("resolved worker join command is empty")
	}

	if resp.NodeRole != "worker" {
		return fmt.Errorf("resolved join role mismatch: expected worker, got %s", resp.NodeRole)
	}

	s.config.JoinCommand = strings.TrimSpace(resp.JoinCommand)
	return nil
}
