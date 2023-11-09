package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get command line arguments for hangar, env, and region
	args := os.Args[1:]

	// Check if the required number of arguments is provided
	if len(args) != 3 {
		fmt.Println("Usage: eks login {hangar} {env} {region}")
		return
	}

	hangar := args[0]
	env := args[1]
	region := args[2]

	// Generate the command
	command := fmt.Sprintf("aws eks update-kubeconfig --name %s-%s-%s --region %s", hangar, env, region, region)

	// Execute the command
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
