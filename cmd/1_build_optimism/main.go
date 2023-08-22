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
	if _, err := os.Stat("optimism"); os.IsNotExist(err) {
		log.Println("Cloning the Optimism Monorepo...")
		cloneCmd := exec.Command("git", "clone", "https://github.com/ethereum-optimism/optimism.git")
		err := runCommandWithOutput(cloneCmd)
		if err != nil {
			log.Fatal("Error cloning the Optimism Monorepo:", err)
		}
	} else {
		log.Println("Optimism Monorepo is already cloned. Skipping cloning step.")
	}


	log.Println("Entering the Optimism Monorepo...")
	os.Chdir("optimism")

	log.Println("Installing required modules...")
	installCmd := exec.Command("pnpm", "install")
	err := runCommandWithOutput(installCmd)
	if err != nil {
		log.Fatal("Error installing required modules:", err)
	}

	log.Println("Building op-node, op-batcher, and op-proposer...")
	makeCmd := exec.Command("make", "op-node", "op-batcher", "op-proposer")
	err = runCommandWithOutput(makeCmd)
	if err != nil {
		log.Fatal("Error making:", err)
	}

	log.Println("Building the Optimism Monorepo...")
	buildCmd := exec.Command("pnpm", "build")
	err = runCommandWithOutput(buildCmd)
	if err != nil {
		log.Fatal("Error building the Optimism Monorepo:", err)
	}

	fmt.Println("Optimism Monorepo and packages built successfully!")
}
