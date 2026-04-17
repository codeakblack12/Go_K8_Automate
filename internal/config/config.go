package config

// Config stores runtime settings for the cluster creator application.
type Config struct {
	OSFamily              string
	KubernetesRepoVersion string
	APIServerAddress      string
	PodNetworkCIDR        string
	KubernetesVersion     string
}

// New creates a Config populated with default values.
func New() *Config {
	return &Config{
		OSFamily:              defaultOSFamily,
		KubernetesRepoVersion: defaultKubernetesRepoVersion,
		APIServerAddress:      defaultAPIServerAddress,
		PodNetworkCIDR:        defaultPodNetworkCIDR,
		KubernetesVersion:     defaultKubernetesVersion,
	}
}
