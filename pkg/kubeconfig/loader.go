package kubeconfig

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// FetchRemoteKubeConfig fetches the kubeconfig file from a remote server over SSH.
func FetchRemoteKubeConfig(host, user, password, key string) (*clientcmdapi.Config, error) {
	if host == "" {
		return nil, fmt.Errorf("host is required for fetching remote kubeconfig")
	}

	// Use default SSH key if none is explicitly provided
	if key == "" {
		key = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
		if _, err := os.Stat(key); err != nil {
			return nil, fmt.Errorf("default SSH key not found at %s: %v", key, err)
		}
	}

	// Construct the SSH command to retrieve the kubeconfig
	sshCommand := []string{"ssh", "-i", key, fmt.Sprintf("%s@%s", user, host), "cat ~/.kube/config"}
	cmd := exec.Command(sshCommand[0], sshCommand[1:]...)

	// Prepare buffers to capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the SSH command
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to execute SSH command: %v, stderr: %s", err, stderr.String())
	}

	// Parse the kubeconfig file content into a clientcmdapi.Config object
	kubeconfigData := stdout.Bytes()
	remoteConfig, err := clientcmd.Load(kubeconfigData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse remote kubeconfig: %v", err)
	}

	return remoteConfig, nil
}

// LoadConfig loads the Kubernetes kubeconfig file from the specified path or the default location.
func LoadConfig(path string) (*clientcmdapi.Config, error) {
	// Use the default kubeconfig path if none is provided
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	// Check if the kubeconfig file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("kubeconfig file not found at %s", path)
	}

	// Load the kubeconfig file into a clientcmdapi.Config object
	config, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig from %s: %v", path, err)
	}

	return config, nil
}

// SaveConfig writes the given Kubernetes configuration back to the specified file.
func SaveConfig(config *clientcmdapi.Config, path string) error {
	// Use the default kubeconfig path if none is provided
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}

	// Write the configuration to the kubeconfig file
	err := clientcmd.WriteToFile(*config, path)
	if err != nil {
		return fmt.Errorf("failed to save kubeconfig to %s: %v", path, err)
	}

	return nil
}

// GetDefaultKubeConfigPath returns the default kubeconfig path for the current user.
func GetDefaultKubeConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".kube", "config")
}
