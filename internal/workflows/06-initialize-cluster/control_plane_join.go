package initializecluster

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func (s *Step) uploadControlPlaneCerts() error {
	cmd := exec.Command("sudo", "kubeadm", "init", "phase", "upload-certs", "--upload-certs")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to upload control-plane certs: %w: %s", err, stderr.String())
	}

	re := regexp.MustCompile(`(?m)^[a-f0-9]{64}$`)
	match := re.FindString(out.String())
	if match == "" {
		return fmt.Errorf("failed to extract certificate key from kubeadm output")
	}

	s.config.CertificateKey = strings.TrimSpace(match)
	return nil
}

func (s *Step) createControlPlaneJoinCommand() error {
	if s.config.CertificateKey == "" {
		return fmt.Errorf("certificate key is empty")
	}

	cmd := exec.Command(
		"sudo",
		"kubeadm",
		"token",
		"create",
		"--print-join-command",
		"--certificate-key",
		s.config.CertificateKey,
	)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate control-plane join command: %w: %s", err, stderr.String())
	}

	joinCmd := strings.TrimSpace(out.String())
	if joinCmd == "" {
		return fmt.Errorf("generated control-plane join command is empty")
	}

	if !strings.Contains(joinCmd, "--control-plane") {
		return fmt.Errorf("generated control-plane join command missing --control-plane")
	}

	if !strings.Contains(joinCmd, "--certificate-key") {
		return fmt.Errorf("generated control-plane join command missing --certificate-key")
	}

	s.config.ControlPlaneJoinCommand = joinCmd
	return nil
}
