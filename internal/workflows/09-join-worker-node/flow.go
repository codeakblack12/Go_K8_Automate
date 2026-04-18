package joinworkernode

import (
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 09: joining a worker node to the cluster.
type Step struct {
	config   *config.Config
	executor *local.Executor
}

// New creates a new join-worker-node step.
func New(cfg *config.Config) *Step {
	return &Step{
		config:   cfg,
		executor: local.New(),
	}
}

// Name returns the workflow step name.
func (s *Step) Name() string {
	return "09-join-worker-node"
}

// Run validates prerequisites, joins the worker node,
// and verifies the result.
func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.joinCluster(); err != nil {
		return err
	}

	if err := s.checkWorkerJoined(); err != nil {
		return err
	}

	return nil
}
