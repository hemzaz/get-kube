package main

import (
	"fmt"
	"os"

	"github.com/hemzaz/get-kube/pkg/auth"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: get-kube <eks|ec2|kind> [additional-args]")
		return
	}

	switch args[0] {
	case "eks":
		tokens, err := auth.GetEKSTokens()
		if err != nil {
			fmt.Println("Error fetching EKS tokens:", err)
			return
		}
		for _, token := range tokens {
			fmt.Println("EKS Token:", token)
		}

	case "kind":
		if len(args) < 2 {
			fmt.Println("Usage: get-kube kind <cluster-name>")
			return
		}
		token, err := auth.GetKindToken(args[1])
		if err != nil {
			fmt.Println("Error fetching Kind token:", err)
			return
		}
		fmt.Println("Kind Token:", token)

	case "ec2":
		if len(args) < 2 {
			fmt.Println("Usage: get-kube ec2 <master-node-name/master-node-ip/master-node-fqdn>")
			return
		}
		// Assuming you have a function in auth for EC2
		token, err := auth.GetEC2Token(args[1])
		if err != nil {
			fmt.Println("Error fetching EC2 token:", err)
			return
		}
		fmt.Println("EC2 Token:", token)

	default:
		fmt.Println("Invalid argument. Please provide one of the following: eks, ec2, kind, docker-desktop")
	}
}
