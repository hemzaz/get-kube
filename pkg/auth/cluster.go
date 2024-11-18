package auth

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GetClusterToken retrieves the authentication token for a vanilla Kubernetes cluster.
func GetClusterToken(host, user, password, key string) (string, error) {
	if host == "" {
		return "", fmt.Errorf("host is required for retrieving the cluster token")
	}

	// Use the default SSH key if none is explicitly provided
	if key == "" {
		key = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		if _, err := os.Stat(key); err != nil {
			return "", fmt.Errorf("default SSH key not found at %s: %v", key, err)
		}
	}

	// Construct the SSH command to retrieve the token
	sshCommand := []string{"ssh", "-i", key, fmt.Sprintf("%s@%s", user, host), "kubectl config view -o jsonpath='{.users[0].user.token}'"}
	cmd := exec.Command(sshCommand[0], sshCommand[1:]...)

	// Prepare buffers to capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the SSH command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute SSH command: %v, stderr: %s", err, stderr.String())
	}

	// Capture the token from the command output
	token := stdout.String()
	if token == "" {
		return "", fmt.Errorf("failed to retrieve token from cluster at %s", host)
	}

	return token, nil
}
