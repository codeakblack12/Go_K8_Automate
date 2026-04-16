package orchestrator

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/models"
	updateos "Go_K8_Automate/internal/workflows/01-update-os"
	disableswap "Go_K8_Automate/internal/workflows/02-disable-swap"
	installcontainerruntime "Go_K8_Automate/internal/workflows/03-install-container-runtime"
	configurecontainers "Go_K8_Automate/internal/workflows/04-configure-containers"
)

// Orchestrator coordinates execution of workflow workflows.
type Orchestrator struct {
	workflows []models.Workflow
}

// New creates a new Orchestrator with the configured workflow workflows.
func New(cfg *config.Config) *Orchestrator {
	workflows := []models.Workflow{
		updateos.New(cfg),
		disableswap.New(cfg),
		installcontainerruntime.New(cfg),
		configurecontainers.New(cfg),
	}

	return &Orchestrator{
		workflows: workflows,
	}
}

// Run executes the configured workflows in order.
func (o *Orchestrator) Run() error {
	for _, workflow := range o.workflows {
		fmt.Printf("Running %s...\n", workflow.Name())

		if err := workflow.Run(); err != nil {
			return fmt.Errorf("%s failed: %w", workflow.Name(), err)
		}
	}

	return nil
}
