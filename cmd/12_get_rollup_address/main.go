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

func main () {
	log.Println("entering the contracts-bedrock directory... ")
	err := os.Chdir("optimism/packages/contracts-bedrock")
	if err != nil {
		log.Fatal("error entering the contracts-bedrock directory: ", err)
	}

	log.Println("getting the rollup address... ")
	getAddress := exec.Command("bash", "-c", "cat deployments/getting-started/L1StandardBridgeProxy.json | jq -r .address")
	err = runCommandWithOutput(getAddress)
	if err != nil {
		log.Fatal("error getting the rollup address: ", err)
	}

	fmt.Println("getting the rollup address runs successfully")
}