package initializecluster

import "fmt"

// publishJoinCode sends the generated kubeadm join command
// to the join-code service and receives a short join-code back.
func (s *Step) publishJoinCode() error {
	resp, err := s.joinCodeClient.Create(s.joinCommand)
	if err != nil {
		return err
	}

	s.joinCode = resp.JoinCode
	fmt.Printf("Join code created successfully: %s\n", s.joinCode)
	return nil
}
