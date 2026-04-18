package orchestrator

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/models"
	updateos "Go_K8_Automate/internal/workflows/01-update-os"
	disableswap "Go_K8_Automate/internal/workflows/02-disable-swap"
	installcontainerruntime "Go_K8_Automate/internal/workflows/03-install-container-runtime"
	configurecontainers "Go_K8_Automate/internal/workflows/04-configure-containers"
	installk8scomponents "Go_K8_Automate/internal/workflows/05-install-k8s-components"
	initializecluster "Go_K8_Automate/internal/workflows/06-initialize-cluster"
	configurekubectlaccess "Go_K8_Automate/internal/workflows/07-configure-kubectl-access"
	installpodnetwork "Go_K8_Automate/internal/workflows/08-install-pod-network"
	joinworkernode "Go_K8_Automate/internal/workflows/09-join-worker-node"
)

// Orchestrator coordinates execution of workflow workflows.
type Orchestrator struct {
	workflows []models.Workflow
}

// New creates a new Orchestrator with the configured workflow workflows.
func New(cfg *config.Config) *Orchestrator {
	// workflows := []models.Workflow{
	// 	updateos.New(cfg),
	// 	disableswap.New(cfg),
	// 	installcontainerruntime.New(cfg),
	// 	configurecontainers.New(cfg),
	// 	installk8scomponents.New(cfg),
	// 	initializecluster.New(cfg),
	// 	configurekubectlaccess.New(cfg),

	// }
	workflows := []models.Workflow{
		updateos.New(cfg),
		disableswap.New(cfg),
		installcontainerruntime.New(cfg),
		configurecontainers.New(cfg),
		installk8scomponents.New(cfg),
	}

	switch cfg.NodeRole {
	case "master":
		workflows = append(workflows,
			initializecluster.New(cfg),
			configurekubectlaccess.New(cfg),
			installpodnetwork.New(cfg),
		)
	case "worker":
		workflows = append(workflows,
			joinworkernode.New(cfg),
		)
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
