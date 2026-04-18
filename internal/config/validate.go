package config

import "fmt"

// Validate checks whether the configuration contains supported values.
func (c *Config) Validate() error {
	switch c.OSFamily {
	case "ubuntu":
	default:
		return fmt.Errorf("unsupported OS family: %s", c.OSFamily)
	}

	switch c.NodeRole {
	case "master", "worker":
	default:
		return fmt.Errorf("unsupported node role: %s", c.NodeRole)
	}

	switch c.PodNetworkPlugin {
	case "calico", "cilium":
	default:
		return fmt.Errorf("unsupported pod network plugin: %s", c.PodNetworkPlugin)
	}

	if c.KubernetesRepoVersion == "" {
		return fmt.Errorf("kubernetes repo version cannot be empty")
	}

	if c.NodeRole == "master" {
		if c.APIServerAddress == "" {
			return fmt.Errorf("API server address cannot be empty for master nodes")
		}
		if c.PodNetworkCIDR == "" {
			return fmt.Errorf("pod network CIDR cannot be empty for master nodes")
		}
	}

	if c.NodeRole == "worker" && c.JoinCommand == "" {
		return fmt.Errorf("join command cannot be empty for worker nodes")
	}

	return nil
}
