package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "get-kube",
		Aliases: []string{"gkb"},
		Short:   "A CLI tool for managing Kubernetes kubeconfig and tokens",
		Long: `get-kube (gkb) is a unified CLI tool for managing Kubernetes authentication and
kubeconfig files. It supports retrieving tokens or kubeconfig from EKS, EC2, and generic clusters, 
as well as syncing local kubeconfig with remote sources.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get-kube: use 'get-kube --help' for available commands.")
		},
	}

	// Add subcommands
	rootCmd.AddCommand(GetCmd())
	rootCmd.AddCommand(SyncCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
