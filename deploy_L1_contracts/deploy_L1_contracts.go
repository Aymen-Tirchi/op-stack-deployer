package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type GoerliConfig struct {
	NumDeployConfirmations         int    `json:"numDeployConfirmations"`
	FinalSystemOwner               string `json:"finalSystemOwner"`
	PortalGuardian                 string `json:"portalGuardian"`
	Controller                     string `json:"controller"`
	L1StartingBlockTag             string `json:"l1StartingBlockTag"`
	L1ChainID                      int    `json:"l1ChainID"`
	L2ChainID                      int    `json:"l2ChainID"`
	L2BlockTime                    int    `json:"l2BlockTime"`
	MaxSequencerDrift              int    `json:"maxSequencerDrift"`
	SequencerWindowSize            int    `json:"sequencerWindowSize"`
	ChannelTimeout                 int    `json:"channelTimeout"`
	P2PSequencerAddress            string `json:"p2pSequencerAddress"`
	BatchInboxAddress              string `json:"batchInboxAddress"`
	BatchSenderAddress             string `json:"batchSenderAddress"`
	L2OutputOracleSubmissionInterval int `json:"l2OutputOracleSubmissionInterval"`
	L2OutputOracleStartingBlockNumber int `json:"l2OutputOracleStartingBlockNumber"`
	L2OutputOracleStartingTimestamp int `json:"l2OutputOracleStartingTimestamp"`
	L2OutputOracleProposer          string `json:"l2OutputOracleProposer"`
	L2OutputOracleChallenger        string `json:"l2OutputOracleChallenger"`
	FinalizationPeriodSeconds       int    `json:"finalizationPeriodSeconds"`
	ProxyAdminOwner                string `json:"proxyAdminOwner"`
	BaseFeeVaultRecipient           string `json:"baseFeeVaultRecipient"`
	L1FeeVaultRecipient             string `json:"l1FeeVaultRecipient"`
	SequencerFeeVaultRecipient      string `json:"sequencerFeeVaultRecipient"`
	BaseFeeVaultMinimumWithdrawalAmount string `json:"baseFeeVaultMinimumWithdrawalAmount"`
	L1FeeVaultMinimumWithdrawalAmount string `json:"l1FeeVaultMinimumWithdrawalAmount"`
	SequencerFeeVaultMinimumWithdrawalAmount string `json:"sequencerFeeVaultMinimumWithdrawalAmount"`
	BaseFeeVaultWithdrawalNetwork   int    `json:"baseFeeVaultWithdrawalNetwork"`
	L1FeeVaultWithdrawalNetwork     int    `json:"l1FeeVaultWithdrawalNetwork"`
	SequencerFeeVaultWithdrawalNetwork int `json:"sequencerFeeVaultWithdrawalNetwork"`
	GasPriceOracleOverhead         int    `json:"gasPriceOracleOverhead"`
	GasPriceOracleScalar           int    `json:"gasPriceOracleScalar"`
	EnableGovernance               bool   `json:"enableGovernance"`
	GovernanceTokenSymbol          string `json:"governanceTokenSymbol"`
	GovernanceTokenName            string `json:"governanceTokenName"`
	GovernanceTokenOwner           string `json:"governanceTokenOwner"`
	L2GenesisBlockGasLimit         string `json:"l2GenesisBlockGasLimit"`
	L2GenesisBlockBaseFeePerGas    string `json:"l2GenesisBlockBaseFeePerGas"`
	L2GenesisRegolithTimeOffset    string `json:"l2GenesisRegolithTimeOffset"`
	Eip1559Denominator             int    `json:"eip1559Denominator"`
	Eip1559Elasticity              int    `json:"eip1559Elasticity"`
}

func updateGoerliConfig(configFilePath string) error {
	// Read the existing JSON data from the file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a struct
	var config GoerliConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	// Update the l1StartingBlockTag field with the desired value
	config.L1StartingBlockTag = "0x6ffc1bf3754c01f6bb9fe057c1578b87a8571ce2e9be5ca14bace6eccfd336c7"

	// Marshal the updated struct back to JSON
	updatedData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(configFilePath, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
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

func main() {
	log.Println("Entering the contracts-bedrock package...")
	os.Chdir("optimism/packages/contracts-bedrock")

	if _, err := os.Stat("deployments/getting-started"); os.IsNotExist(err) {
		log.Println("Creating a getting-started deployment directory...")
		mkdirCmd := exec.Command("mkdir", "deployments/getting-started")
		err := mkdirCmd.Run()
		if err != nil {
			log.Fatal("Error creating the deployments/getting-started directory:", err)
		}
	} else {
		log.Println("The deployments/getting-started directory already exists")
	}

	configFilePath := "deploy-config/goerli.json"
	err := updateGoerliConfig(configFilePath)
	if err != nil {
		log.Fatal("Error updating goerli.json: ", err)
	}

	fmt.Println("goerli.json updated successfully!")

	var privateKeyAdmin, rpcURL string

	fmt.Println("Enter the private key of the Admin: ")
	fmt.Scan(&privateKeyAdmin)
	fmt.Println("Enter the L1 node URL (ETH_RPC_URL): ")
	fmt.Scan(&rpcURL)
	if privateKeyAdmin == "" || rpcURL == "" {
		log.Fatal("The private key and/or RPC URL are not set")
	}

	log.Println("Deploying the L1 smart contracts...")

	deployCmd1 := exec.Command("forge", "script", "scripts/Deploy.s.sol:Deploy", "--private-key="+privateKeyAdmin, "--broadcast", "--rpc-url="+rpcURL)
	err = runCommandWithOutput(deployCmd1)
	if err != nil {
		log.Fatal("Error deploying the L1 contracts: ", err)
	}

	deployCmd2 := exec.Command("forge", "script", "scripts/Deploy.s.sol:Deploy", "--sig", "sync()", "--private-key="+privateKeyAdmin, "--broadcast", "--rpc-url="+rpcURL)
	err = runCommandWithOutput(deployCmd2)
	if err != nil {
		log.Fatal("Error deploying the L1 contracts: ", err)
	}

	log.Println("L1 smart contracts deployed successfully!")
}
