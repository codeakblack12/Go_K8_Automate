package joincontrolplane

import (
	"fmt"
	"strings"
)

func (s *Step) resolveJoinCode() error {
	if strings.TrimSpace(s.config.ControlPlaneJoinCommand) != "" {
		return nil
	}

	if strings.TrimSpace(s.config.JoinCode) == "" {
		return fmt.Errorf("shared join code is empty")
	}

	resp, err := s.joinCodeClient.Resolve(s.config.JoinCode, "control-plane")
	if err != nil {
		return fmt.Errorf("failed to resolve control-plane join code: %w", err)
	}

	if strings.TrimSpace(resp.JoinCommand) == "" {
		return fmt.Errorf("resolved control-plane join command is empty")
	}

	if resp.NodeRole != "control-plane" {
		return fmt.Errorf("resolved join role mismatch: expected control-plane, got %s", resp.NodeRole)
	}

	s.config.ControlPlaneJoinCommand = strings.TrimSpace(resp.JoinCommand)
	return nil
}
