package disableswap

import (
	"bytes"
	"fmt"
	"os/exec"
)

// checkPrerequisites validates whether the step has what it needs.
func (s *Step) checkPrerequisites() error {
	if s.config == nil {
		return fmt.Errorf("missing configuration")
	}

	if s.executor == nil {
		return fmt.Errorf("missing local executor")
	}

	return nil
}

// checkSwapDisabled verifies that no active swap remains after execution.
func (s *Step) checkSwapDisabled() error {
	cmd := exec.Command("sh", "-c", "swapon --show | tail -n +2")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to verify swap status: %w", err)
	}

	if out.Len() > 0 {
		return fmt.Errorf("swap is still enabled")
	}

	return nil
}
