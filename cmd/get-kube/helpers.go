package getkube

import (
	"fmt"
)

// PrintError prints a formatted error message to the console.
func PrintError(msg string, err error) {
	if err != nil {
		fmt.Printf("Error: %s: %v\n", msg, err)
	} else {
		fmt.Printf("Error: %s\n", msg)
	}
}

// ValidateFlag ensures that a required flag is provided.
func ValidateFlag(value, flagName string) error {
	if value == "" {
		return fmt.Errorf("the required flag --%s is missing", flagName)
	}
	return nil
}
