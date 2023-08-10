package main

import (
	"fmt"
	"os"

	"github.com/hemzaz/get-kube/pkg/auth"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid argument: ec2, eks, docker-desktop, kind")
		return
	}

	switch os.Args[1] {
	case "ec2":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the master node name/IP/FQDN for EC2.")
			return
		}
		_, err := auth.GetEC2Token(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "eks":
		_, err := auth.GetEKSToken()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "docker-desktop":
		contextName := ""
		if len(os.Args) > 2 {
			contextName = os.Args[2]
		}
		_, err := auth.GetDockerDesktopToken(contextName)
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "kind":
		contextName := ""
		if len(os.Args) > 2 {
			contextName = os.Args[2]
		}
		_, err := auth.GetKindToken(contextName)
		if err != nil {
			fmt.Println("Error:", err)
		}

	default:
		fmt.Println("Invalid argument. Please provide one of the following: ec2, eks, docker-desktop, kind")
	}
}
