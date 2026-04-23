package initializecluster

import "fmt"

func (s *Step) publishJoinCode() error {
	resp, err := s.joinCodeClient.Create(s.joinCommand, "worker")
	if err != nil {
		return err
	}

	s.joinCode = resp.JoinCode
	fmt.Printf("Worker join code created successfully: %s\n", s.joinCode)
	return nil
}

func (s *Step) publishControlPlaneJoinCode() error {
	resp, err := s.joinCodeClient.Create(s.config.ControlPlaneJoinCommand, "control-plane")
	if err != nil {
		return err
	}

	s.config.ControlPlaneJoinCode = resp.JoinCode
	fmt.Printf("Control-plane join code created successfully: %s\n", s.config.ControlPlaneJoinCode)
	return nil
}
