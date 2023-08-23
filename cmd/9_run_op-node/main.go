package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func runCommand(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("command execution error: ", err)
	}
}

func getSequencerPrivateKey(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var sequencerPrivateKey string
	isSequencer := false

	for _, line := range lines {
		if strings.Contains(line, "Sequencer:") {
			isSequencer = true
		} else if isSequencer && strings.HasPrefix(line, "Private Key:") {
			sequencerPrivateKey = strings.TrimSpace(strings.TrimPrefix(line, "Private Key:"))
			break
		}
	}

	if sequencerPrivateKey == "" {
		return "", errors.New("sequencer private key not found in keys.txt")
	}

	return strings.TrimPrefix(sequencerPrivateKey, "0x"), nil
}


func loadEnvironment() {
	if err := godotenv.Load("optimism/packages/contracts-bedrock/.envrc"); err != nil {
		log.Fatal("Error loading environment variables from .envrc: ", err)
	}
}

func setEnvironmentVariables(sequencerPrivateKey string) {
	os.Setenv("SEQ_KEY", sequencerPrivateKey)
	os.Setenv("L1_RPC", os.Getenv("ETH_RPC_URL"))
	os.Setenv("RPC_KIND", "any")
}

// func printEnvironmentVariables() {
// 	log.Println("SEQ_KEY:", os.Getenv("SEQ_KEY"))
// 	log.Println("L1_RPC:", os.Getenv("L1_RPC"))
// 	log.Println("RPC_KIND:", os.Getenv("RPC_KIND"))
// }

func main() {
	log.Println("setting system variables... ")
	keysPath := "keys.txt"
	sequencerPrivateKey, err := getSequencerPrivateKey(keysPath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	loadEnvironment()
	setEnvironmentVariables(sequencerPrivateKey)
	log.Println("Environment variables set successfully!")

	log.Println("entering the op-node directory... ")
	err = os.Chdir("optimism/op-node")
	if err != nil {
		log.Fatal("error entering the op-node directory: ", err)
	}

	log.Println("running op-node... ")
	runCmd := `./bin/op-node \
	--l2=http://localhost:8551 \
	--l2.jwt-secret=./jwt.txt \
	--sequencer.enabled \
	--sequencer.l1-confs=3 \
	--verifier.l1-confs=3 \
	--rollup.config=./rollup.json \
	--rpc.addr=0.0.0.0 \
	--rpc.port=8547 \
	--p2p.disable \
	--rpc.enable-admin \
	--p2p.sequencer.key=$SEQ_KEY \
	--l1=$L1_RPC \
	--l1.rpckind=$RPC_KIND`

	runCommand(runCmd)
}