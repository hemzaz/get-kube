package auth

import (
	"fmt"

	"github.com/hemzaz/get-kube/pkg/utils"
)

// GetEC2Token fetches a token from an EC2-hosted Kubernetes master node.
func GetEC2Token(masterNode, user, sshKey string) (string, error) {
	// Validate or get the default SSH key
	var err error
	if sshKey == "" {
		sshKey, err = utils.ValidateSSHKey("")
		if err != nil {
			return "", fmt.Errorf("error validating SSH key: %v", err)
		}
	}

	// Remote command to fetch the Kubernetes token
	remoteCmd := "sudo kubeadm token create --print-join-command | awk '{print $5}'"

	// Execute the SSH command via utils
	token, err := utils.ExecuteSSHCommand(user, masterNode, sshKey, remoteCmd)
	if err != nil {
		return "", fmt.Errorf("failed to fetch EC2 token: %v", err)
	}

	// Ensure the token is not empty
	token = utils.TrimOutput(token)
	if token == "" {
		return "", fmt.Errorf("failed to retrieve token: output was empty")
	}

	// Update the kubeconfig with the fetched token
	if err := UpdateKubeConfig(masterNode, masterNode, user, token); err != nil {
		return "", fmt.Errorf("failed to update kubeconfig: %v", err)
	}

	utils.LogInfo("Updated EC2 Token for context: %s", masterNode)
	return token, nil
}
