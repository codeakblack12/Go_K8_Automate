package orchestrator

import (
	"fmt"

	"Go_K8_Automate/internal/config"
	// "Go_K8_Automate/internal/models"
	updateos "Go_K8_Automate/internal/workflows/01-update-os"
	disableswap "Go_K8_Automate/internal/workflows/02-disable-swap"
	installcontainerruntime "Go_K8_Automate/internal/workflows/03-install-container-runtime"
	configurecontainers "Go_K8_Automate/internal/workflows/04-configure-containers"
	installk8scomponents "Go_K8_Automate/internal/workflows/05-install-k8s-components"
	initializecluster "Go_K8_Automate/internal/workflows/06-initialize-cluster"
	configurekubectlaccess "Go_K8_Automate/internal/workflows/07-configure-kubectl-access"
	installpodnetwork "Go_K8_Automate/internal/workflows/08-install-pod-network"
	joinworkernode "Go_K8_Automate/internal/workflows/09-join-worker-node"
	joincontrolplane "Go_K8_Automate/internal/workflows/10-join-control-plane"
)

// Orchestrator coordinates execution of workflow workflows.
type Runnable interface {
	Name() string
	Run() error
}

type Orchestrator struct {
	config *config.Config
	steps  []Runnable
}

func New(cfg *config.Config) *Orchestrator {
	o := &Orchestrator{
		config: cfg,
	}

	o.steps = o.buildSteps()
	return o
}

func (o *Orchestrator) buildSteps() []Runnable {
	common := []Runnable{
		updateos.New(o.config),
		disableswap.New(o.config),
		installcontainerruntime.New(o.config),
		configurecontainers.New(o.config),
		installk8scomponents.New(o.config),
	}

	switch o.config.NodeRole {
	case "master":
		return append(common,
			initializecluster.New(o.config),
			configurekubectlaccess.New(o.config),
			installpodnetwork.New(o.config),
		)
	case "worker":
		return append(common,
			joinworkernode.New(o.config),
		)
	case "control-plane":
		return append(common,
			joincontrolplane.New(o.config),
		)
	default:
		return common
	}
}

func (o *Orchestrator) Run() error {
	for _, step := range o.steps {
		fmt.Printf("Running %s...\n", step.Name())
		if err := step.Run(); err != nil {
			return fmt.Errorf("%s failed: %w", step.Name(), err)
		}
	}
	return nil
}
