package joinworkernode

import (
	"Go_K8_Automate/internal/api/joincode"
	"Go_K8_Automate/internal/config"
	"Go_K8_Automate/internal/executor/local"
)

// Step handles workflow step 09: joining a worker node to the cluster.
type Step struct {
	config         *config.Config
	executor       *local.Executor
	joinCodeClient *joincode.Client
}

func New(cfg *config.Config) *Step {
	return &Step{
		config:         cfg,
		executor:       local.New(),
		joinCodeClient: joincode.NewClient(cfg.JoinServiceBaseURL),
	}
}

func (s *Step) Name() string {
	return "09-join-worker-node"
}

func (s *Step) Run() error {
	if err := s.checkPrerequisites(); err != nil {
		return err
	}

	if err := s.resolveJoinCode(); err != nil {
		return err
	}

	if err := s.joinNode(); err != nil {
		return err
	}

	return nil
}
