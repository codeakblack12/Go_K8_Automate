package joincontrolplane

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func (s *Step) hasStaleNodeState() (bool, string, error) {
	var reasons []string

	if _, err := os.Stat("/etc/kubernetes/kubelet.conf"); err == nil {
		reasons = append(reasons, "/etc/kubernetes/kubelet.conf already exists")
	} else if !os.IsNotExist(err) {
		return false, "", fmt.Errorf("failed to check /etc/kubernetes/kubelet.conf: %w", err)
	}

	if _, err := os.Stat("/etc/kubernetes/pki/ca.crt"); err == nil {
		reasons = append(reasons, "/etc/kubernetes/pki/ca.crt already exists")
	} else if !os.IsNotExist(err) {
		return false, "", fmt.Errorf("failed to check /etc/kubernetes/pki/ca.crt: %w", err)
	}

	if _, err := os.Stat("/etc/kubernetes/manifests"); err == nil {
		reasons = append(reasons, "/etc/kubernetes/manifests already exists")
	} else if !os.IsNotExist(err) {
		return false, "", fmt.Errorf("failed to check /etc/kubernetes/manifests: %w", err)
	}

	ln, err := net.Listen("tcp", ":10250")
	if err != nil {
		reasons = append(reasons, "port 10250 is already in use")
	} else {
		_ = ln.Close()
	}

	if len(reasons) == 0 {
		return false, "", nil
	}

	return true, strings.Join(reasons, "; "), nil
}
