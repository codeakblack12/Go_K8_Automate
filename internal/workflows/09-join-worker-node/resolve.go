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

	fmt.Printf("DEBUG resolving join code %q against %s\n",
		s.config.JoinCode,
		s.config.JoinServiceBaseURL,
	)

	resp, err := s.joinCodeClient.Resolve(s.config.JoinCode)
	if err != nil {
		return fmt.Errorf("failed to resolve join code: %w", err)
	}

	if strings.TrimSpace(resp.JoinCommand) == "" {
		return fmt.Errorf("resolved join command is empty")
	}

	s.config.JoinCommand = strings.TrimSpace(resp.JoinCommand)
	return nil
}
