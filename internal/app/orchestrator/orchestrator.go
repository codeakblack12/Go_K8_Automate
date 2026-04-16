package orchestrator

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/models"
	updateos "Go_K8_Automate/internal/workflows/01-update-os"
)

// Orchestrator coordinates execution of workflow workflows.
type Orchestrator struct {
	workflows []models.Workflow
}

// New creates a new Orchestrator with the configured workflow workflows.
func New(cfg *config.Config) *Orchestrator {
	workflows := []models.Workflow{
		updateos.New(cfg),
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
