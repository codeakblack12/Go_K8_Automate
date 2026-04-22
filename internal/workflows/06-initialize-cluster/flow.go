package initializecluster

import (
	"Go_K8_Automate/internal/api/joincode"
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 06: initializing the Kubernetes cluster.
type Step struct {
	config         *config.Config
	executor       *local.Executor
	joinCodeClient *joincode.Client
	joinCommand    string
	joinCode       string
}

// New creates a new initialize-cluster step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:         cfg,
		executor:       local.New(),
		joinCodeClient: joincode.NewClient(cfg.JoinServiceBaseURL),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "06-initialize-cluster"
}

// Run validates prerequisites, initializes the control plane,
// creates a join command, and verifies success.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.initControlPlane(); err != nil {
		return err
	}

	if err := s.createJoinCommand(); err != nil {
		return err
	}

	if err := s.publishJoinCode(); err != nil {
		return err
	}

	if err := s.checkClusterInitialized(); err != nil {
		return err
	}

	return nil
}
