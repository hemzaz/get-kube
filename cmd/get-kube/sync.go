package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hemzaz/get-kube/pkg/kubeconfig"
	"github.com/spf13/cobra"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func SyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync [context]",
		Short: "Sync kubeconfig with remote clusters",
		Long: `The "sync" command updates the local kubeconfig with remote configurations 
from Kubernetes clusters. It ensures your local kubeconfig is always in sync 
with the latest cluster credentials and configurations.`,
		Args: cobra.MaximumNArgs(1),
		RunE: runSync,
	}

	// Add flags for syncing
	cmd.Flags().StringP("host", "H", "", "Cluster host (IP or hostname) for remote sync")
	cmd.Flags().BoolP("all", "a", false, "Sync all contexts in kubeconfig")
	cmd.Flags().StringP("user", "u", "root", "SSH user for remote access")
	cmd.Flags().StringP("password", "p", "", "Password for SSH user (optional)")
	cmd.Flags().StringP("key", "k", "~/.ssh/id_rsa", "Path to SSH private key (optional)")
	cmd.Flags().StringP("config", "c", "", "Path to the kubeconfig file (optional)")

	return cmd
}

func runSync(cmd *cobra.Command, args []string) error {
	kubeconfigPath, _ := cmd.Flags().GetString("config")
	allContexts, _ := cmd.Flags().GetBool("all")
	host, _ := cmd.Flags().GetString("host")
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	key, _ := cmd.Flags().GetString("key")

	// Validate the required arguments
	if !allContexts && len(args) == 0 {
		return errors.New("please specify a context or use the --all flag to sync all contexts")
	}
	if host == "" {
		return errors.New("the --host flag is required for syncing remote kubeconfig")
	}

	// Load the local kubeconfig
	localConfig, err := kubeconfig.LoadConfig(kubeconfigPath)
	if err != nil {
		return fmt.Errorf("error loading local kubeconfig: %v", err)
	}

	if allContexts {
		fmt.Println("Syncing all contexts from remote clusters...")
		for contextName := range localConfig.Contexts {
			fmt.Printf("Syncing context: %s\n", contextName)
			if err := syncContext(contextName, localConfig, host, user, password, key); err != nil {
				fmt.Printf("Error syncing context '%s': %v\n", contextName, err)
			}
		}
	} else {
		// Sync a specific context
		contextName := args[0]
		if err := syncContext(contextName, localConfig, host, user, password, key); err != nil {
			return fmt.Errorf("error syncing context '%s': %v", contextName, err)
		}
		fmt.Printf("Context '%s' synced successfully.\n", contextName)
	}

	return nil
}

func syncContext(contextName string, localConfig *clientcmdapi.Config, host, user, password, key string) error {
	// Check if the context exists locally
	if _, ok := localConfig.Contexts[contextName]; !ok {
		return fmt.Errorf("context '%s' not found in local kubeconfig", contextName)
	}

	// Fetch the remote kubeconfig for the cluster
	fmt.Printf("Fetching remote kubeconfig for context '%s'...\n", contextName)
	remoteConfig, err := kubeconfig.FetchRemoteKubeConfig(host, user, password, key)
	if err != nil {
		return fmt.Errorf("failed to fetch remote kubeconfig: %v", err)
	}

	// Sync the remote kubeconfig with the local one
	fmt.Printf("Merging remote kubeconfig for context '%s'...\n", contextName)
	if err := kubeconfig.SyncKubeConfig(localConfig, remoteConfig); err != nil {
		return fmt.Errorf("failed to sync kubeconfig: %v", err)
	}

	fmt.Printf("Successfully synced context '%s'.\n", contextName)
	return nil
}
