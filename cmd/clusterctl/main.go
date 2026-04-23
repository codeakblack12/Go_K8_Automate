package main

import (
	"flag"
	"fmt"
	"os"

	"Go_K8_Automate/internal/app/orchestrator"
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/utils/network"
)

// main is the CLI entry point.
// It parses flags, loads config values, validates them,
// and runs the correct workflow for the selected node role.
func main() {
	cfg := config.New()

	if cfg.NodeRole == "master" && cfg.APIServerAddress == "" {
		detectedIP, err := network.DetectPrimaryIPv4()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to detect API server address automatically: %v\n", err)
			os.Exit(1)
		}

		cfg.APIServerAddress = detectedIP
		fmt.Printf("Detected API server advertise address: %s\n", cfg.APIServerAddress)
	}

	role := flag.String("role", cfg.NodeRole, "Node role: master or worker")
	joinCommand := flag.String("join-command", cfg.JoinCommand, "Full kubeadm join command for worker nodes")
	joinCode := flag.String("join-code", cfg.JoinCode, "Short join code for worker nodes")
	joinServiceURL := flag.String("join-service-url", cfg.JoinServiceBaseURL, "Base URL for the join-code service")
	apiServerAddress := flag.String("apiserver-address", cfg.APIServerAddress, "API server advertise address for master node")
	podNetworkCIDR := flag.String("pod-network-cidr", cfg.PodNetworkCIDR, "Pod network CIDR for cluster initialization")
	podNetworkPlugin := flag.String("pod-network-plugin", cfg.PodNetworkPlugin, "Pod network plugin: calico or cilium")
	kubernetesVersion := flag.String("kubernetes-version", cfg.KubernetesVersion, "Optional Kubernetes version for kubeadm init")
	repoVersion := flag.String("repo-version", cfg.KubernetesRepoVersion, "Kubernetes apt repository version, e.g. v1.35")
	resetNode := flag.Bool("reset-node", false, "If true, reset the node with kubeadm reset before initializing")

	flag.Parse()

	cfg.NodeRole = *role
	cfg.JoinCommand = *joinCommand
	cfg.JoinCode = *joinCode
	cfg.JoinServiceBaseURL = *joinServiceURL
	cfg.APIServerAddress = *apiServerAddress
	cfg.PodNetworkCIDR = *podNetworkCIDR
	cfg.PodNetworkPlugin = *podNetworkPlugin
	cfg.KubernetesVersion = *kubernetesVersion
	cfg.KubernetesRepoVersion = *repoVersion
	cfg.ResetNode = *resetNode

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	orch := orchestrator.New(cfg)

	if err := orch.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cluster setup failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("cluster setup completed successfully")
}
