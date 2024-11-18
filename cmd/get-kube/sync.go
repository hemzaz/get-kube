package main

import (
	"fmt"
	"os"

	"github.com/hemzaz/get-kube/pkg/kubeconfig"
	"github.com/spf13/cobra"
)

// SyncCmd creates the "sync" subcommand for syncing kubeconfig files.
func SyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync [context]",
		Short: "Sync kubeconfig with remote clusters",
		Long: `The "sync" command updates the local kubeconfig with remote configurations
from Kubernetes clusters. This ensures your local kubeconfig is always in sync
with the latest cluster credentials and configurations.`,
		Args: cobra.MaximumNArgs(1),
		Run:  runSync,
	}

	// Flags for syncing
	cmd.Flags().StringP("host", "H", "", "Cluster host (IP or hostname) for remote sync")
	cmd.Flags().BoolP("all", "a", false, "Sync all contexts in kubeconfig")
	cmd.Flags().StringP("user", "u", "root", "SSH user for remote access")
	cmd.Flags().StringP("password", "p", "", "Password for SSH user (optional)")
	cmd.Flags().StringP("key", "k", "~/.ssh/id_rsa", "Path to SSH private key (optional)")

	return cmd
}

func runSync(cmd *cobra.Command, args []string) {
	kubeconfigPath, _ := cmd.Flags().GetString("config")
	allContexts, _ := cmd.Flags().GetBool("all")
	host, _ := cmd.Flags().GetString("host")
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	key, _ := cmd.Flags().GetString("key")

	localConfig, err := kubeconfig.LoadConfig(kubeconfigPath)
	if err != nil {
		fmt.Printf("Error loading local kubeconfig: %v\n", err)
		os.Exit(1)
	}

	if allContexts {
		fmt.Println("Syncing all contexts from remote clusters...")
		for contextName := range localConfig.Contexts {
			fmt.Printf("Syncing context: %s\n", contextName)
			err := syncContext(contextName, localConfig, host, user, password, key)
			if err != nil {
				fmt.Printf("Error syncing context '%s': %v\n", contextName, err)
			}
		}
	} else {
		if len(args) == 0 {
			fmt.Println("Error: Please specify a context or use the --all flag to sync all contexts.")
			os.Exit(1)
		}
		contextName := args[0]
		err := syncContext(contextName, localConfig, host, user, password, key)
		if err != nil {
			fmt.Printf("Error syncing context '%s': %v\n", contextName, err)
			os.Exit(1)
		}
		fmt.Printf("Context '%s' synced successfully.\n", contextName)
	}
}

func syncContext(contextName string, localConfig *clientcmdapi.Config, host, user, password, key string) error {
	if _, ok := localConfig.Contexts[contextName]; !ok {
		return fmt.Errorf("context '%s' not found in local kubeconfig", contextName)
	}

	remoteConfig, err := kubeconfig.FetchRemoteKubeConfig(host, user, password, key)
	if err != nil {
		return fmt.Errorf("failed to fetch remote kubeconfig: %v", err)
	}

	err = kubeconfig.SyncKubeConfig(localConfig, remoteConfig)
	if err != nil {
		return fmt.Errorf("failed to sync kubeconfig: %v", err)
	}

	return nil
}
