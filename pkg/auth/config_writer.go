package auth

import (
	"fmt"

	"github.com/hemzaz/get-kube/pkg/utils"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// UpdateKubeConfig updates the kubeconfig file with the given token and cluster details.
func UpdateKubeConfig(clusterName, masterNode, user, token string) error {
	// Load the current kubeconfig
	kubeconfigPath := utils.GetDefaultKubeConfigPath()
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	// Add or update the cluster
	config.Clusters[clusterName] = &clientcmdapi.Cluster{
		Server: fmt.Sprintf("https://%s:6443", masterNode),
	}

	// Add or update the user
	config.AuthInfos[clusterName] = &clientcmdapi.AuthInfo{
		Token: token,
	}

	// Add or update the context
	config.Contexts[clusterName] = &clientcmdapi.Context{
		Cluster:  clusterName,
		AuthInfo: clusterName,
	}

	// Set the current context
	config.CurrentContext = clusterName

	// Save the updated kubeconfig
	if err := clientcmd.WriteToFile(*config, kubeconfigPath); err != nil {
		return fmt.Errorf("failed to save kubeconfig: %v", err)
	}

	fmt.Printf("Successfully updated kubeconfig for cluster '%s'.\n", clusterName)
	return nil
}
