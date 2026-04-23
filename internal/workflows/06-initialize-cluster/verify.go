package initializecluster

import "os"

func (s *Step) checkClusterInitialized() error {
	_, err := os.Stat("/etc/kubernetes/admin.conf")
	return err
}
