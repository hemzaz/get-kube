package auth

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetEC2Token(masterNode string) (string, error) {
	// Construct the SSH command
	cmd := exec.Command("ssh", "-i", "~/.ssh/id_rsa", fmt.Sprintf("ec2-user@%s", masterNode), "sudo kubeadm token create --print-join-command | awk '{print $5}'")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, %v", err, stderr.String())
	}

	token := stdout.String()

	// Update the kubeconfig
	err = UpdateKubeConfig(masterNode, masterNode, "ec2-user", token)
	if err != nil {
		return "", err
	}

	fmt.Println("Updated EC2 Token for context:", masterNode)

	return token, nil
}
