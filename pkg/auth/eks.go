package auth

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetEKSToken() ([]string, error) {
	cmd := exec.Command("aws", "eks", "list-clusters", "--query", "clusters", "--output", "text")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %v, %v", err, stderr.String())
	}

	clusters := strings.Split(stdout.String(), "\n")
	var tokens []string

	for _, cluster := range clusters {
		if cluster == "" {
			continue
		}

		tokenCmd := exec.Command("aws", "eks", "get-token", "--cluster-name", cluster)
		var tokenOut, tokenErr bytes.Buffer
		tokenCmd.Stdout = &tokenOut
		tokenCmd.Stderr = &tokenErr

		err := tokenCmd.Run()
		if err != nil {
			return nil, fmt.Errorf("failed to get token for cluster %s: %v, %v", cluster, err, tokenErr.String())
		}

		token := strings.Trim(tokenOut.String(), "\"")
		tokens = append(tokens, token)

		// Update the kubeconfig
		err = UpdateKubeConfig(cluster, cluster, "eks-user", token)
		if err != nil {
			return nil, err
		}

		fmt.Println("Updated EKS Token for context:", cluster)
	}

	return tokens, nil
}
