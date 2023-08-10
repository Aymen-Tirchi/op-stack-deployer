package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/joho/godotenv"
)

type Addresses struct {
	Admin     string
	Proposer  string
	Batcher   string
	Sequencer string
}

func readAddresses(filename string) (Addresses, error) {
	var addresses Addresses
	data, err := os.ReadFile(filename)
	if err != nil {
		return addresses, err
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		// Skip empty lines or lines with insufficient data
		if strings.TrimSpace(line) == "" || i+1 >= len(lines) {
			continue
		}

		// Extract the address line right after each role
		switch {
		case strings.HasPrefix(line, "Admin:"):
			addresses.Admin = extractAddress(lines[i+1])
		case strings.HasPrefix(line, "Proposer:"):
			addresses.Proposer = extractAddress(lines[i+1])
		case strings.HasPrefix(line, "Batcher:"):
			addresses.Batcher = extractAddress(lines[i+1])
		case strings.HasPrefix(line, "Sequencer:"):
			addresses.Sequencer = extractAddress(lines[i+1])
		}
	}

	return addresses, nil
}

func extractAddress(line string) string {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return ""
}


type BlockInfo struct {
	Hash      string
	Number    string
	Timestamp string
}

func getBlockInfo(rpcURL string) (BlockInfo, error) {
	var blockInfo BlockInfo

	cmd := exec.Command("cast", "block", "finalized", "--rpc-url", rpcURL)
	output, err := cmd.CombinedOutput()  // This captures both standard output and error
	if err != nil {
		return blockInfo, fmt.Errorf("failed to execute the cast command: %v, output: %s", err, output)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "hash") {
			blockInfo.Hash = strings.TrimSpace(strings.TrimPrefix(line, "hash"))
		} else if strings.HasPrefix(line, "number") {
			blockInfo.Number = strings.TrimSpace(strings.TrimPrefix(line, "number"))
		} else if strings.HasPrefix(line, "timestamp") {
			blockInfo.Timestamp = strings.TrimSpace(strings.TrimPrefix(line, "timestamp"))
		}
	}

	return blockInfo, nil
}

func main() {

	log.Println("Entering the contracts-bedrock package...")
	os.Chdir("optimism/packages/contracts-bedrock")

  if _, err:= os.Stat(".envrc"); os.IsNotExist(err) {
    log.Println("Copying the environment file...")
    cpCmd := exec.Command("cp", ".envrc.example", ".envrc")
    err := cpCmd.Run()
    if err != nil {
      log.Fatal("Error copying the environment file: ", err)
    }
  }

	if err := godotenv.Load(".envrc"); err != nil {
		log.Fatal("Error loading environment variables from .envrc: ", err)
	}

	log.Println("Reading keys.txt file...")
	addresses, err := readAddresses("../../../keys.txt")
	if err != nil {
		log.Fatal("Error reading addresses from keys.txt: ", err)
	}

	// Get block information using the cast command
	rpcURL := os.Getenv("ETH_RPC_URL")

	// Get block information using the cast command
	blockInfo, err := getBlockInfo(rpcURL)
	if err != nil {
		log.Fatalf("Error getting block information: %v", err)
	}

	// Generate the configuration data
	configData := fmt.Sprintf(`{
		"numDeployConfirmations": 1,

		"finalSystemOwner": "%s",
		"portalGuardian": "%s",
		"controller": "%s",

		"l1StartingBlockTag": "%s",

		"l1ChainID": 5,
		"l2ChainID": 42069,
		"l2BlockTime": 2,

		"maxSequencerDrift": 600,
		"sequencerWindowSize": 3600,
		"channelTimeout": 300,

		"p2pSequencerAddress": "%s",
		"batchInboxAddress": "0xff00000000000000000000000000000000042069",
		"batchSenderAddress": "%s",

		"l2OutputOracleSubmissionInterval": 120,
		"l2OutputOracleStartingBlockNumber": 0,
		"l2OutputOracleStartingTimestamp": %s,

		"l2OutputOracleProposer": "%s",
		"l2OutputOracleChallenger": "%s",

		"finalizationPeriodSeconds": 12,

		"proxyAdminOwner": "%s",
		"baseFeeVaultRecipient": "%s",
		"l1FeeVaultRecipient": "%s",
		"sequencerFeeVaultRecipient": "%s",

		"baseFeeVaultMinimumWithdrawalAmount": "0x8ac7230489e80000",
		"l1FeeVaultMinimumWithdrawalAmount": "0x8ac7230489e80000",
		"sequencerFeeVaultMinimumWithdrawalAmount": "0x8ac7230489e80000",
		"baseFeeVaultWithdrawalNetwork": 0,
		"l1FeeVaultWithdrawalNetwork": 0,
		"sequencerFeeVaultWithdrawalNetwork": 0,

		"gasPriceOracleOverhead": 2100,
		"gasPriceOracleScalar": 1000000,

		"enableGovernance": true,
		"governanceTokenSymbol": "OP",
		"governanceTokenName": "Optimism",
		"governanceTokenOwner": "%s",

		"l2GenesisBlockGasLimit": "0x1c9c380",
		"l2GenesisBlockBaseFeePerGas": "0x3b9aca00",
		"l2GenesisRegolithTimeOffset": "0x0",

		"eip1559Denominator": 50,
		"eip1559Elasticity": 10
	}`, addresses.Admin, addresses.Admin, addresses.Admin, blockInfo.Hash, addresses.Sequencer, addresses.Batcher, blockInfo.Timestamp, addresses.Proposer, addresses.Admin, addresses.Admin, addresses.Admin, addresses.Admin, addresses.Admin, addresses.Admin)

	// Save the configuration data to a file
	outputFilePath := "deploy-config/getting-started.json"
	err = os.WriteFile(outputFilePath, []byte(configData), 0644)
	if err != nil {
		log.Fatal("Error writing configuration data to file: ", err)
	}

	log.Printf("Configuration file has been updated and saved to %s\n", outputFilePath)
}