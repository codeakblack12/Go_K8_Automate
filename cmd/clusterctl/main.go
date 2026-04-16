package main

import (
	"fmt"
	"os"

	"Go_K8_Automate/internal/app/orchestrator"
	"Go_K8_Automate/internal/config"
)

// main is the CLI entry point.
// It loads configuration, builds the orchestrator, runs the workflow,
// and exits with a non-zero code if execution fails.
func main() {
	// Load application configuration.
	cfg := config.New()

	// Validate configuration before running any workflow.
	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	// Build the orchestrator that will execute workflow steps.
	orch := orchestrator.New(cfg)

	// Run the workflow.
	if err := orch.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cluster setup failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("cluster setup completed successfully")
}
