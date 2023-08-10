package main

import (
	"fmt"
	"os"
	"github.com/hemzaz/get-kube/pkg/auth"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: get-kube <eks|ec2|kind|docker-desktop>")
		return
	}

	switch os.Args[1] {
	case "eks":
		token, err := auth.GetEKSToken()
		// Handle token and error
	case "ec2":
		token, err := auth.GetEC2Token()
		// Handle token and error
	case "kind":
		token, err := auth.GetKindToken()
		// Handle token and error
	case "docker-desktop":
		token, err := auth.GetDockerDesktopToken()
		// Handle token and error
	default:
		fmt.Println("Invalid argument. Use one of: eks|ec2|kind|docker-desktop")
	}
}
