package config

import "fmt"

// Validate checks whether the configuration contains supported values.
func (c *Config) Validate() error {
	switch c.OSFamily {
	case "ubuntu":
	default:
		return fmt.Errorf("unsupported OS family: %s", c.OSFamily)
	}

	if c.KubernetesRepoVersion == "" {
		return fmt.Errorf("kubernetes repo version cannot be empty")
	}

	if c.APIServerAddress == "" {
		return fmt.Errorf("API server address cannot be empty")
	}

	if c.PodNetworkCIDR == "" {
		return fmt.Errorf("pod network CIDR cannot be empty")
	}

	return nil
}
