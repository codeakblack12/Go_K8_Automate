package config

import "fmt"

// Validate checks whether the configuration contains supported values.
func (c *Config) Validate() error {
	switch c.OSFamily {
	case "ubuntu":
		return nil
	default:
		return fmt.Errorf("unsupported OS family: %s", c.OSFamily)
	}
}
