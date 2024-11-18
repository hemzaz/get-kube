package getkube

import (
	"fmt"

	"github.com/hemzaz/get-kube/pkg/auth"
	"github.com/spf13/cobra"
)

// GetCmd creates the "get" subcommand for retrieving tokens or kubeconfig.
func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [eks|ec2|cluster]",
		Short: "Retrieve kubeconfig or tokens from EKS, EC2, or a generic cluster",
		Long: `The "get" command retrieves Kubernetes authentication tokens or kubeconfig 
from the specified source: 
- "eks" for Amazon EKS clusters.
- "ec2" for EC2-hosted Kubernetes clusters.
- "cluster" for generic Kubernetes clusters using SSH.`,
		Args: cobra.ExactArgs(1),
		Run:  runGet,
	}

	// Flags for the "cluster" option
	cmd.Flags().StringP("host", "H", "", "Cluster host (IP or hostname) [required for cluster]")
	cmd.Flags().StringP("user", "u", "root", "SSH user for remote access")
	cmd.Flags().StringP("password", "p", "", "Password for SSH user (optional)")
	cmd.Flags().StringP("key", "k", "~/.ssh/id_rsa", "Path to SSH private key (optional)")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	source := args[0]

	switch source {
	case "eks":
		getEKS()
	case "ec2":
		getEC2(cmd)
	case "cluster":
		getCluster(cmd)
	default:
		fmt.Println("Invalid source. Use one of: eks, ec2, cluster.")
	}
}

func getEKS() {
	tokens, err := auth.GetEKSTokens()
	if err != nil {
		fmt.Printf("Error fetching EKS tokens: %v\n", err)
		return
	}
	for _, token := range tokens {
		fmt.Println("EKS Token:", token)
	}
}

func getEC2(cmd *cobra.Command) {
	host, _ := cmd.Flags().GetString("host")
	if host == "" {
		fmt.Println("Error: --host is required for EC2.")
		return
	}

	token, err := auth.GetEC2Token(host)
	if err != nil {
		fmt.Printf("Error fetching EC2 token: %v\n", err)
		return
	}
	fmt.Println("EC2 Token:", token)
}

func getCluster(cmd *cobra.Command) {
	host, _ := cmd.Flags().GetString("host")
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	key, _ := cmd.Flags().GetString("key")

	if host == "" {
		fmt.Println("Error: --host is required for cluster.")
		return
	}

	token, err := auth.GetClusterToken(host, user, password, key)
	if err != nil {
		fmt.Printf("Error fetching cluster token: %v\n", err)
		return
	}
	fmt.Println("Cluster Token:", token)
}
