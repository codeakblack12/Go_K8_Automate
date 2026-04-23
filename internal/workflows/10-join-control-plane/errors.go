package joincontrolplane

import "fmt"

type StaleNodeStateError struct {
	Reason string
}

func (e *StaleNodeStateError) Error() string {
	return fmt.Sprintf(
		"stale Kubernetes state detected on node: %s; rerun with --reset-node to clean the node before joining",
		e.Reason,
	)
}
