package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetDefaultKubeConfigPath returns the default kubeconfig path for the current user.
func GetDefaultKubeConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".kube", "config")
}

// ValidateSSHKey returns the validated SSH key path or an error if it doesn't exist.
func ValidateSSHKey(customKey string) (string, error) {
	if customKey != "" {
		return customKey, nil
	}
	defaultKey := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	if _, err := os.Stat(defaultKey); os.IsNotExist(err) {
		return "", fmt.Errorf("default SSH key not found at %s", defaultKey)
	}
	return defaultKey, nil
}

// LogInfo prints an info-level message.
func LogInfo(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// LogError prints an error-level message.
func LogError(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// TrimOutput removes trailing and leading whitespace from a string.
func TrimOutput(output string) string {
	return strings.TrimSpace(output)
}

// ValidateKubeConfigPath checks if the kubeconfig path is valid, defaulting to ~/.kube/config.
func ValidateKubeConfigPath(path string) (string, error) {
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("kubeconfig file not found at %s", path)
	}
	return path, nil
}

// ExecuteSSHCommand runs a command on a remote host via SSH.
// It requires the user, host, SSH key path, and the command to run.
func ExecuteSSHCommand(user, host, sshKey, remoteCmd string) (string, error) {
	// Construct the SSH command
	sshCommand := []string{
		"ssh",
		"-i", sshKey,
		fmt.Sprintf("%s@%s", user, host),
		remoteCmd,
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

	// Return the command output
	return stdout.String(), nil
}
