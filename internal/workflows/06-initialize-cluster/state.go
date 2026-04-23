package initializecluster

import (
	"os"
)

func (s *Step) checkExistingControlPlane() (bool, error) {
	if _, err := os.Stat("/etc/kubernetes/admin.conf"); err == nil {
		return true, nil
	} else if !os.IsNotExist(err) {
		return false, err
	}

	return false, nil
}
