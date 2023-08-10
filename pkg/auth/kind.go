package auth

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetKindToken(contextName string) (string, error) {
	if contextName == "" {
		contextName = "kind-kind"
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
	err = UpdateKubeConfig(contextName, contextName, "kind-user", token)
	if err != nil {
		return "", err
	}

	fmt.Println("Updated Kind Token for context:", contextName)

	return token, nil
}
