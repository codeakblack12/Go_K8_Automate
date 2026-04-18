package configurekubectlaccess

import (
	"fmt"
	"os"
	"os/user"

	"Go_K8_Automate/internal/executor/common"
)

// configureKubeconfig copies the admin kubeconfig to the current user's home directory
// so kubectl can access the cluster without requiring an explicit --kubeconfig flag.
func (s *Step) configureKubeconfig() error {
	fmt.Println("STEP 7: Configuring kubectl access...")

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to determine current user: %w", err)
	}

	homeDir := currentUser.HomeDir
	kubeDir := homeDir + "/.kube"
	kubeConfig := kubeDir + "/config"

	if err := os.MkdirAll(kubeDir, 0755); err != nil {
		return fmt.Errorf("failed to create %s: %w", kubeDir, err)
	}

	copyCmd := common.Command{
		Name: "sudo",
		Args: []string{"cp", "-i", "/etc/kubernetes/admin.conf", kubeConfig},
	}
	if err := s.executor.Run(copyCmd); err != nil {
		return err
	}

	chownCmd := common.Command{
		Name: "sudo",
		Args: []string{"chown", currentUser.Uid + ":" + currentUser.Gid, kubeConfig},
	}
	if err := s.executor.Run(chownCmd); err != nil {
		return err
	}

	fmt.Println("STEP 7 complete: kubectl access configured")
	return nil
}
