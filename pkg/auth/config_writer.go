package auth

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func UpdateKubeConfig(contextName, clusterName, userName, token string) error {
	// Load the kubeconfig file
	kubeconfig := clientcmd.NewDefaultPathOptions()

	// Load the existing config
	config, err := kubeconfig.GetStartingConfig()
	if err != nil {
		return err
	}

	// Check if the context already exists
	if _, exists := config.Contexts[contextName]; !exists {
		config.Contexts[contextName] = api.NewContext()
	}
	config.Contexts[contextName].Cluster = clusterName
	config.Contexts[contextName].AuthInfo = userName

	// Check if the user already exists
	if _, exists := config.AuthInfos[userName]; !exists {
		config.AuthInfos[userName] = api.NewAuthInfo()
	}
	config.AuthInfos[userName].Token = token

	// Save the updated config
	return clientcmd.ModifyConfig(kubeconfig, *config, false)
}
