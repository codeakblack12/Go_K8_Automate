package initializecluster

import "fmt"

func (s *Step) publishSharedJoinCode() error {
	resp, err := s.joinCodeClient.Create(s.joinCommand, s.config.ControlPlaneJoinCommand)
	if err != nil {
		return err
	}

	s.joinCode = resp.JoinCode
	s.config.JoinCode = resp.JoinCode

	fmt.Printf("Shared join code created successfully: %s\n", s.joinCode)
	return nil
}
