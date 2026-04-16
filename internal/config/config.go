package config

// Config stores runtime settings for the cluster creator application.
type Config struct {
	OSFamily string
}

// New creates a Config populated with default values.
func New() *Config {
	return &Config{
		OSFamily: defaultOSFamily,
	}
}
