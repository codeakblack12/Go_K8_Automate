package local

import (
	"fmt"
	"os"
	"os/exec"

	"Go_K8_Automate/internal/executor/common"
)

// Executor runs commands on the local machine.
type Executor struct{}

// New creates a new local executor.
func New() *Executor {
	return &Executor{}
}

// Run executes a command and streams output to the current terminal.
func (e *Executor) Run(cmd common.Command) error {
	command := exec.Command(cmd.Name, cmd.Args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to run command %s %v: %w", cmd.Name, cmd.Args, err)
	}

	return nil
}
