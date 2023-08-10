package auth

import (
	// "fmt"
	"os/exec"
)

func GetKindToken(clusterName string) (string, error) {
	if clusterName == "" {
		clusterName = "kind-kind" // default kind cluster name
	}
	cmd := exec.Command("kubectl", "--context", clusterName, "config", "view", "-o", "jsonpath='{.users[0].user.token}'")
	token, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(token), nil
}
