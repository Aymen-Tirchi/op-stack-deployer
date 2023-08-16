package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func runCommandWithOutput(cmd *exec.Cmd) error {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	multiReader := io.MultiReader(stdoutPipe, stderrPipe)

	go func() {
		io.Copy(os.Stdout, multiReader)
	}()

	return cmd.Run()
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

	log.Println("getting the rollup address runs successfully")
}