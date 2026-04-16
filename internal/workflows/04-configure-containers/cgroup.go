package configurecontainers

import (
	"fmt"

	"Go_K8_Automate/internal/executor/common"
)

// configureContainerd creates the containerd config directory if needed,
// generates the default config file, enables SystemdCgroup,
// then restarts and enables the containerd service.
func (s *Step) configureContainerd() error {
	fmt.Println("STEP 4: Configuring containerd...")

	// Create the config directory if it does not already exist.
	mkdirCmd := common.Command{
		Name: "sudo",
		Args: []string{"mkdir", "-p", "/etc/containerd"},
	}
	if err := s.executor.Run(mkdirCmd); err != nil {
		return err
	}

	// Generate the default containerd configuration file.
	generateConfigCmd := common.Command{
		Name: "sh",
		Args: []string{
			"-c",
			"sudo containerd config default | sudo tee /etc/containerd/config.toml > /dev/null",
		},
	}
	if err := s.executor.Run(generateConfigCmd); err != nil {
		return err
	}

	// Enable the systemd cgroup driver for Kubernetes compatibility.
	setSystemdCgroupCmd := common.Command{
		Name: "sudo",
		Args: []string{
			"sed",
			"-i",
			`s/SystemdCgroup = false/SystemdCgroup = true/`,
			"/etc/containerd/config.toml",
		},
	}
	if err := s.executor.Run(setSystemdCgroupCmd); err != nil {
		return err
	}

	// Restart containerd to apply the new configuration.
	restartCmd := common.Command{
		Name: "sudo",
		Args: []string{"systemctl", "restart", "containerd"},
	}
	if err := s.executor.Run(restartCmd); err != nil {
		return err
	}

	// Enable containerd to start automatically on boot.
	enableCmd := common.Command{
		Name: "sudo",
		Args: []string{"systemctl", "enable", "containerd"},
	}
	if err := s.executor.Run(enableCmd); err != nil {
		return err
	}

	fmt.Println("STEP 4 complete: containerd configured successfully")
	return nil
}
