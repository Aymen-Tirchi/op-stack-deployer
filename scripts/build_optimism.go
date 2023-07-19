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

func main() {
	log.Println("Cloning the Optimism Monorepo...")
	cloneCmd := exec.Command("git", "clone", "https://github.com/ethereum-optimism/optimism.git")
	err := runCommandWithOutput(cloneCmd)
	if err != nil {
		log.Fatal("Error cloning the Optimism Monorepo:", err)
	}

	log.Println("Entering the Optimism Monorepo...")
	os.Chdir("/optimism")

	log.Println("Installing required modules...")
	installCmd := exec.Command("pnpm", "install")
	installCmd.Dir = os.ExpandEnv("$HOME/optimism")
	err = runCommandWithOutput(installCmd)
	if err != nil {
		log.Fatal("Error installing required modules:", err)
	}

	log.Println("Building op-node, op-batcher, and op-proposer...")
	makeCmd := exec.Command("make", "op-node", "op-batcher", "op-proposer")
	makeCmd.Dir = os.ExpandEnv("$HOME/optimism")
	err = runCommandWithOutput(makeCmd)
	if err != nil {
		log.Fatal("Error making:", err)
	}

	log.Println("Building the Optimism Monorepo...")
	buildCmd := exec.Command("pnpm", "build")
	buildCmd.Dir = os.ExpandEnv("$HOME/optimism")
	err = runCommandWithOutput(buildCmd)
	if err != nil {
		log.Fatal("Error building the Optimism Monorepo:", err)
	}

	log.Println("Optimism Monorepo and packages built successfully!")
}
