package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/joho/godotenv"
)

const (
	opNodeDir        = "cmd/main.go"
	deployConfigDir  = "../packages/contracts-bedrock/deploy-config/getting-started.json"
	gettingStartedDir = "../packages/contracts-bedrock/deployments/getting-started/"
)

func runCommandWithOutput(cmd *exec.Cmd) error {
	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &outputBuf

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution error: %v\nOutput:\n%s\n", err, outputBuf.String())
	}
	return nil
}

func changeWorkingDirectory(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("failed to change the working directory: %v", err)
	}
	log.Println("Changed the current working directory to", dir)
	return nil
}

func main() {

	if err := godotenv.Load("optimism/packages/contracts-bedrock/.envrc"); err != nil {
		log.Fatal("Error loading environment variables from .envrc: ", err)
	}

	scriptDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get the current working directory: ", err)
	}
	log.Println("Current working directory:", scriptDir)

	opNodePath := filepath.Join(scriptDir, "optimism/op-node")
	err = changeWorkingDirectory(opNodePath)
	if err != nil {
		log.Fatal(err)
	}

	rpcURL := os.Getenv("ETH_RPC_URL")

	if _, err := os.Stat("genesis.json"); os.IsNotExist(err) {
		log.Println("Creating genesis.json and rollup.json... ")
		runCmd := exec.Command("go", "run", opNodeDir, "genesis", "l2",
			"--deploy-config", deployConfigDir,
			"--deployment-dir", gettingStartedDir,
			"--outfile.l2", "genesis.json",
			"--outfile.rollup", "rollup.json",
			"--l1-rpc", rpcURL)
		err = runCommandWithOutput(runCmd)
		if err != nil {
			log.Fatal("error creating genesis.json and rollup.json: ", err)
		}
	} else {
		log.Fatal("genesis.json already exists")
	}
	
	if _, err := os.Stat("jwt.txt"); os.IsNotExist(err) {
		log.Println("generating jwt.txt file... ")
		jwtCmd := exec.Command("sh", "-c", "openssl rand -hex 32 > jwt.txt")
		err = runCommandWithOutput(jwtCmd)
		if err != nil {
			log.Fatal("error generating jwt.txt: ", err)
		}
	} else {
		log.Fatal("jwt.txt already exists")
	}

	log.Println("copying genesis.json and rollup.json into op-geth... ")
	copyCmd1 := exec.Command("cp", "genesis.json", "../../op-geth")
	err = runCommandWithOutput(copyCmd1)
	if err != nil {
		log.Fatal("error copying genesis.json: ", err)
	}
	copyCmd2 := exec.Command("cp", "jwt.txt", "../../op-geth")
	err = runCommandWithOutput(copyCmd2)
	if err != nil {
		log.Fatal("error copying jwt.txt: ", err)
	}

	fmt.Println("Generated L2 config files successfully")
}
