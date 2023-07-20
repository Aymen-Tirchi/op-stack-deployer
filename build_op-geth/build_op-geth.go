package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Println("Cloning the op-geth repo...")
	cloneCmd := exec.Command("git", "clone", "https://github.com/ethereum-optimism/op-geth.git")
	err := runCommandWithOutput(cloneCmd)
	if err != nil {
		log.Fatal("error cloning the op-geth repo: ", err)
	}

	log.Println("building the op-geth...")
	err = os.Chdir("../op-geth")
	if err != nil {
		log.Fatal("Error changing working directory: ", err)
	}
	makeCmd := exec.Command("make", "geth")
	err = runCommandWithOutput(makeCmd)
	if err != nil {
		log.Fatal("error building the op-geth: ", err)
	}
	log.Println("op-geth repo and packages built successfully!")
}

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
