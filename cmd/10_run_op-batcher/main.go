package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func runCommand(cmdStr string){
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("command execution error: ", err)
	}
}

func getBatcherPrivateKey(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var batcherPrivateKey string
	isBatcher := false

	for _, line := range lines {
		if strings.Contains(line, "Batcher:") {
			isBatcher = true
		} else if isBatcher && strings.HasPrefix(line, "Private Key:") {
			batcherPrivateKey = strings.TrimSpace(strings.TrimPrefix(line, "Private Key:"))
			break
		}
	}

	if batcherPrivateKey == "" {
		return "", errors.New("batcher private key not found in keys.txt")
	}

	return strings.TrimPrefix(batcherPrivateKey, "0x"), nil
}

func loadEnvironment() {
	if err := godotenv.Load("optimism/packages/contracts-bedrock/.envrc"); err != nil {
		log.Fatal("Error loading environment variables from .envrc: ", err)
	}
}

func setEnvironmentVariables(batcherPrivateKey string) {
	os.Setenv("BATCHER_KEY", batcherPrivateKey)
	os.Setenv("L1_RPC", os.Getenv("ETH_RPC_URL"))
}

// func printEnvironmentVariables() {
// 	log.Println("BATCHER_KEY:", os.Getenv("BATCHER_KEY"))
// 	log.Println("L1_RPC:", os.Getenv("L1_RPC"))
// }

func main() {
	log.Println("setting system variables... ")
	keysPath := "keys.txt"
	batcherPrivateKey, err := getBatcherPrivateKey(keysPath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	loadEnvironment()
	setEnvironmentVariables(batcherPrivateKey)
	// printEnvironmentVariables()

	log.Println("entering the op-batcher directory... ")
	err = os.Chdir("optimism/op-batcher")
	if err != nil {
		log.Fatal("error entering the op-batcher directory: ", err)
	}

	log.Println("running op-batcher")
	runCmd := `./bin/op-batcher \
    --l2-eth-rpc=http://localhost:8545 \
    --rollup-rpc=http://localhost:8547 \
    --poll-interval=1s \
    --sub-safety-margin=6 \
    --num-confirmations=1 \
    --safe-abort-nonce-too-low-count=3 \
    --resubmission-timeout=30s \
    --rpc.addr=0.0.0.0 \
    --rpc.port=8548 \
    --rpc.enable-admin \
    --max-channel-duration=1 \
    --l1-eth-rpc=$L1_RPC \
    --private-key=$BATCHER_KEY`

	runCommand(runCmd)

	log.Println("op-batcher running successfully! ")

}