package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	if _, err := os.Stat("op-geth"); os.IsNotExist(err) {
		log.Println("Cloning the op-geth repo... ")
		cloneCmd := exec.Command("git", "clone", "https://github.com/ethereum-optimism/op-geth.git")
		err := runCommandWithOutput(cloneCmd)
		if err != nil {
			log.Fatal("Error cloning the op-geth repo: ", err)
		}
	} else {
		log.Println("op-geth repo is already cloned")
	}

	err := os.Chdir("op-geth")
	if err != nil {
		log.Fatal("Error changing working directory: ", err)
	}

	log.Println("Building the op-geth...")
	makeCmd := exec.Command("make", "geth")
	err = runCommandWithOutput(makeCmd)
	if err != nil {
		log.Fatal("Error building the op-geth: ", err)
	}
	fmt.Println("op-geth repo and packages built successfully!")
}
