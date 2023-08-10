package auth

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetDockerDesktopToken(contextName string) (string, error) {
	if contextName == "" {
		contextName = "docker-desktop"
	}

	cmd := exec.Command("kubectl", "config", "view", "-o", fmt.Sprintf("jsonpath='{.users[?(@.name == \"%s\")].user.token}'", contextName))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v, %v", err, stderr.String())
	}

	token := strings.Trim(stdout.String(), "\"")

	// Update the kubeconfig
	err = UpdateKubeConfig(contextName, contextName, "docker-desktop-user", token)
	if err != nil {
		return "", err
	}

	fmt.Println("Updated Docker Desktop Token for context:", contextName)

	return token, nil
}
