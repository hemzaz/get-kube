package kubeconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

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
