package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	usageMessage  = "Usage: eks login {hangar} {env} {region}"
	commandFormat = "aws eks update-kubeconfig --name %s-%s-%s --region %s"
)

func main() {
	// Get command line arguments for hangar, env, and region
	args := os.Args[1:]

	// Check if the required number of arguments is provided
	if len(args) != 3 {
		fmt.Println(usageMessage)
		return
	}

	hangar := args[0]
	env := args[1]
	region := args[2]

	// Validate input values
	if hangar == "" || env == "" || region == "" {
		fmt.Println("Error: Hangar, env, and region cannot be empty.")
		fmt.Println(usageMessage)
		return
	}

	// Generate the command
	command := fmt.Sprintf(commandFormat, hangar, env, region, region)

	// Execute the command
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Println("Command executed successfully.")
}
