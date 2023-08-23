package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

type L2OutputOracleProxyInfo struct {
	Address string `json:"address"`
}

func runCommand(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal("command execution error: ", err)
	}
}

func getProposerPrivateKey(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	var proposerPrivateKey string
	isProposer := false

	for _, line := range lines {
		if strings.Contains(line, "Proposer:") {
			isProposer = true
		} else if isProposer && strings.HasPrefix(line, "Private Key:") {
			proposerPrivateKey = strings.TrimSpace(strings.TrimPrefix(line, "Private Key:"))
			break
		}
	}

	if proposerPrivateKey == "" {
		return "", errors.New("proposer private key not found in keys.txt")
	}

	return strings.TrimPrefix(proposerPrivateKey, "0x"), nil
}

func getL2OutputOracleAddress(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var oracleInfo L2OutputOracleProxyInfo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&oracleInfo)
	if err != nil {
		return "", err
	}

	return oracleInfo.Address, nil
}

func loadEnvironment() {
	if err := godotenv.Load("optimism/packages/contracts-bedrock/.envrc"); err != nil {
		log.Fatal("Error loading environment variables from .envrc: ", err)
	}
}

func setEnvironmentVariables(proposerPrivateKey, l2ooAddress string) {
	os.Setenv("PROPOSER_KEY", proposerPrivateKey)
	os.Setenv("L1_RPC", os.Getenv("ETH_RPC_URL"))
	os.Setenv("L2OO_ADDR", l2ooAddress)
}

// func printEnvironmentVariables() {
// 	log.Println("PROPOSER_KEY:", os.Getenv("PROPOSER_KEY"))
// 	log.Println("L1_RPC:", os.Getenv("L1_RPC"))
// 	log.Println("L2OO_ADDR:", os.Getenv("L2OO_ADDR"))
// }

func main() {
	log.Println("setting system variables... ")
	keysPath := "keys.txt"
	proposerPrivateKey, err := getProposerPrivateKey(keysPath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	l2ooAddress, err := getL2OutputOracleAddress("optimism/packages/contracts-bedrock/deployments/getting-started/L2OutputOracleProxy.json")
	if err != nil {
		log.Fatal("Error:", err)
	}

	loadEnvironment()
	setEnvironmentVariables(proposerPrivateKey, l2ooAddress)
	// printEnvironmentVariables()

	log.Println("entering the op-proposer directory... ")
	err = os.Chdir("optimism/op-proposer")
	if err != nil {
		log.Fatal("error entering the op-proposer directory: ", err)
	}

	runCmd := `./bin/op-proposer \
		--poll-interval=12s \
		--rpc.port=8560 \
		--rollup-rpc=http://localhost:8547 \
		--l2oo-address=$L2OO_ADDR \
		--private-key=$PROPOSER_KEY \
		--l1-eth-rpc=$L1_RPC`

	runCommand(runCmd)

}
