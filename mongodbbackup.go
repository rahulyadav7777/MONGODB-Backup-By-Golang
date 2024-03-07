package main

import (
	"fmt"
	"os/exec"
)

func main() {

	host := "localhost"
	port := "27017"
	username := "yadav"
	password := "yadav"
	authenticationDatabase := "aeroplane" // or the database where the user is defined

	// Command to execute mongodump
	cmd := exec.Command("mongodump",
		"--host", host+":"+port,
		"--username", username,
		"--password", password,
		"--authenticationDatabase", authenticationDatabase,
	)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing mongodump:", err)
		return
	}

	// Print the output
	fmt.Println(string(output))
}
