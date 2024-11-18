package kubeconfig

import (
	"bytes"
	"fmt"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// SyncKubeConfig synchronizes the local kubeconfig with data from a remote kubeconfig.
func SyncKubeConfig(localConfig *clientcmdapi.Config, remoteConfig *clientcmdapi.Config) error {
	if localConfig == nil {
		return fmt.Errorf("local kubeconfig is nil")
	}
	if remoteConfig == nil {
		return fmt.Errorf("remote kubeconfig is nil")
	}

	updated := false

	// Synchronize clusters
	for name, remoteCluster := range remoteConfig.Clusters {
		if localCluster, exists := localConfig.Clusters[name]; !exists || !bytes.Equal(localCluster.CertificateAuthorityData, remoteCluster.CertificateAuthorityData) {
			localConfig.Clusters[name] = remoteCluster
			updated = true
			fmt.Printf("Updated cluster: %s\n", name)
		}
	}

	// Synchronize users
	for name, remoteAuth := range remoteConfig.AuthInfos {
		if localAuth, exists := localConfig.AuthInfos[name]; !exists || !bytes.Equal(localAuth.ClientCertificateData, remoteAuth.ClientCertificateData) || !bytes.Equal(localAuth.ClientKeyData, remoteAuth.ClientKeyData) {
			localConfig.AuthInfos[name] = remoteAuth
			updated = true
			fmt.Printf("Updated user: %s\n", name)
		}
	}

	// Synchronize contexts
	for name, remoteContext := range remoteConfig.Contexts {
		if localContext, exists := localConfig.Contexts[name]; !exists || localContext.Cluster != remoteContext.Cluster || localContext.AuthInfo != remoteContext.AuthInfo {
			localConfig.Contexts[name] = remoteContext
			updated = true
			fmt.Printf("Updated context: %s\n", name)
		}
	}

	// Update the current context if needed
	if localConfig.CurrentContext != remoteConfig.CurrentContext {
		localConfig.CurrentContext = remoteConfig.CurrentContext
		updated = true
		fmt.Printf("Updated current context to: %s\n", remoteConfig.CurrentContext)
	}

	if !updated {
		fmt.Println("No updates needed; kubeconfig is already up-to-date.")
		return nil
	}

	fmt.Println("Local kubeconfig updated with remote data.")
	return nil
}
