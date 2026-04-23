package initializecluster

import (
	"Go_K8_Automate/internal/api/joincode"
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
	"fmt"
)

// Step handles workflow step 06: initializing the Kubernetes cluster.
type Step struct {
	config         *config.Config
	executor       *local.Executor
	joinCodeClient *joincode.Client
	joinCommand    string
	joinCode       string
}

func New(cfg *config.Config) *Step {
	return &Step{
		config:         cfg,
		executor:       local.New(),
		joinCodeClient: joincode.NewClient(cfg.JoinServiceBaseURL),
	}
}

func (s *Step) Name() string {
	return "06-initialize-cluster"
}

func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	exists, err := s.checkExistingControlPlane()
	if err != nil {
		return err
	}

	if exists {
		if !s.config.ResetNode {
			return fmt.Errorf("existing control-plane state detected on this node; rerun with --reset-node if you want to reset it")
		}

		if err := s.resetNode(); err != nil {
			return err
		}
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

	if err := s.uploadControlPlaneCerts(); err != nil {
		return err
	}

	if err := s.createControlPlaneJoinCommand(); err != nil {
		return err
	}

	if err := s.publishControlPlaneJoinCode(); err != nil {
		return err
	}

	if err := s.checkClusterInitialized(); err != nil {
		return err
	}

	return nil
}
