package auth

import (
	"os/exec"
	"strings"
)

func GetEKSTokens() ([]string, error) {
	cmd := exec.Command("aws", "eks", "list-clusters", "--query", "clusters", "--output", "text")
	clusters, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	clusterList := strings.Split(string(clusters), "\n")
	var tokens []string
	for _, cluster := range clusterList {
		cmd := exec.Command("aws", "eks", "update-kubeconfig", "--name", cluster)
		_, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		cmd = exec.Command("kubectl", "config", "view", "-o", "jsonpath='{.users[0].user.token}'")
		token, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, string(token))
	}
	return tokens, nil
}
