package joinworkernode

import "fmt"

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
		return err
	}

	s.config.JoinCommand = resp.JoinCommand
	return nil
}
