package auth

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GetEC2Token fetches a token from an EC2-hosted Kubernetes master node.
func GetEC2Token(masterNode, user, sshKey string) (string, error) {
	// Use default SSH key if none is provided
	if sshKey == "" {
		sshKey = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		if _, err := os.Stat(sshKey); err != nil {
			return "", fmt.Errorf("default SSH key not found at %s: %v", sshKey, err)
		}
	}

	// Construct the SSH command
	sshCommand := []string{
		"ssh",
		"-i", sshKey,
		fmt.Sprintf("%s@%s", user, masterNode),
		"sudo kubeadm token create --print-join-command | awk '{print $5}'",
	}
	cmd := exec.Command(sshCommand[0], sshCommand[1:]...)

	// Prepare buffers to capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute SSH command: %v, stderr: %s", err, stderr.String())
	}

	token := stdout.String()
	if token == "" {
		return "", fmt.Errorf("failed to retrieve token: output was empty")
	}

	// Update the kubeconfig with the new token
	err := UpdateKubeConfig(masterNode, masterNode, user, token)
	if err != nil {
		return "", fmt.Errorf("failed to update kubeconfig: %v", err)
	}

	fmt.Println("Updated EC2 Token for context:", masterNode)

	return token, nil
}
