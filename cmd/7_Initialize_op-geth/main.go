package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	gethExecutablePath = "build/bin/geth"
)

func runCommandWithOutput(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
  
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution error: %v\n", err)
	}
	return nil
}

func main() {
	log.Println("Entering the op-geth directory...")
	err := os.Chdir("op-geth")
	if err != nil {
		log.Fatal("Failed to change the working directory:", err)
	}

	if _, err := os.Stat("datadir"); os.IsNotExist(err) {
		log.Println("Creating the data directory folder...")
		mkdirCmd := exec.Command("mkdir", "datadir")
		err = runCommandWithOutput(mkdirCmd)
		if err != nil {
			log.Fatal("Error creating the data directory: ", err)
		}
	} else {
		log.Println("the data dir already exists")
	}
	
	
	log.Println("Initializing op-geth with genesis...")
	initGethCmd := exec.Command(gethExecutablePath, "init", "--datadir=datadir", "genesis.json")
	err = runCommandWithOutput(initGethCmd)
	if err != nil {
		log.Fatal("Error initializing op-geth: ", err)
	}

	fmt.Println("op-geth initialized successfully!")
}
