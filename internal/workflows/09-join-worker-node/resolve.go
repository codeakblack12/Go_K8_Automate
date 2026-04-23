package joinworkernode

import (
	"fmt"
	"strings"
)

// resolveJoinCode exchanges a short join-code for the full kubeadm join command.
// If JoinCommand is already provided directly, this step does nothing.
func (s *Step) resolveJoinCode() error {
	if s.config.JoinCommand != "" {
		return nil
	}

	if s.config.JoinCode == "" {
		return fmt.Errorf("join code is empty")
	}

	resp, err := s.joinCodeClient.Resolve(s.config.JoinCode)
	if err != nil {
		return fmt.Errorf("failed to resolve join code %s: %w", s.config.JoinCode, err)
	}

	if strings.TrimSpace(resp.JoinCommand) == "" {
		return fmt.Errorf("resolved join command is empty for join code %s", s.config.JoinCode)
	}

	s.config.JoinCommand = resp.JoinCommand
	return nil
}
