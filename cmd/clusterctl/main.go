package main

import (
	"flag"
	"fmt"
	"os"

	"Go_K8_Automate/internal/app/orchestrator"
	"Go_K8_Automate/internal/config"
)

// main is the CLI entry point.
// It parses flags, loads config values, validates them,
// and runs the correct workflow for the selected node role.
func main() {
	// Start with default application configuration.
	cfg := config.New()

	// Define CLI flags.
	role := flag.String("role", cfg.NodeRole, "Node role: master or worker")
	joinCommand := flag.String("join-command", cfg.JoinCommand, "Full kubeadm join command for worker nodes")
	apiServerAddress := flag.String("apiserver-address", cfg.APIServerAddress, "API server advertise address for master node")
	podNetworkCIDR := flag.String("pod-network-cidr", cfg.PodNetworkCIDR, "Pod network CIDR for cluster initialization")
	podNetworkPlugin := flag.String("pod-network-plugin", cfg.PodNetworkPlugin, "Pod network plugin: calico or cilium")
	kubernetesVersion := flag.String("kubernetes-version", cfg.KubernetesVersion, "Optional Kubernetes version for kubeadm init")
	repoVersion := flag.String("repo-version", cfg.KubernetesRepoVersion, "Kubernetes apt repository version, e.g. v1.35")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of clusterctl:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), `
	Examples:
	Master node:
		go run ./cmd/clusterctl --role master --apiserver-address 192.168.1.20

	Worker node:
		go run ./cmd/clusterctl --role worker --join-command "kubeadm join 192.168.1.20:6443 --token <token> --discovery-token-ca-cert-hash sha256:<hash>"
	`)
	}
	// Parse CLI flags.
	flag.Parse()

	// Apply CLI flag values to config.
	cfg.NodeRole = *role
	cfg.JoinCommand = *joinCommand
	cfg.APIServerAddress = *apiServerAddress
	cfg.PodNetworkCIDR = *podNetworkCIDR
	cfg.PodNetworkPlugin = *podNetworkPlugin
	cfg.KubernetesVersion = *kubernetesVersion
	cfg.KubernetesRepoVersion = *repoVersion

	// Validate configuration before running any workflow.
	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	// Create the orchestrator and run the workflow.
	orch := orchestrator.New(cfg)

	if err := orch.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cluster setup failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("cluster setup completed successfully")
}
