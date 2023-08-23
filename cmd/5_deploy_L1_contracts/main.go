package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

type sepoliaConfig struct {
  FinalSystemOwner               string `json:"finalSystemOwner"`
  PortalGuardian                 string `json:"portalGuardian"`
  L1StartingBlockTag             string `json:"l1StartingBlockTag"`
  L1ChainID                      int    `json:"l1ChainID"`
  L1BlockTime                    int    `json:"l1BlockTime"`
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
  SystemConfigStartBlock         int `json:"systemConfigStartBlock"`
}

func updatesepoliaConfig(configFilePath string) error {
  // Read the existing JSON data from the file
  data, err := os.ReadFile(configFilePath)
  if err != nil {
	return err
  }
  // Unmarshal the JSON data into a struct
  var config sepoliaConfig
  err = json.Unmarshal(data, &config)
  if err != nil {
	return err
  }
  config.SystemConfigStartBlock = 0 
  config.L1BlockTime = 12

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
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()
  if err != nil {
      log.Fatalf("Command execution error: %v\n", err)
  }
  return nil
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

  configFilePath := "deploy-config/getting-started.json"
  err := updatesepoliaConfig(configFilePath)
  if err != nil {
    log.Fatal("Error updating getting-started.json: ", err)
  }

  configFilePath = "deploy-config/sepolia.json"
  err = updatesepoliaConfig(configFilePath)
  if err != nil {
    log.Fatal("Error updating sepolia.json: ", err)
  }

  log.Println("sepolia.json updated successfully!")

  if err := godotenv.Load(".envrc"); err != nil {
    log.Fatal("Error loading environment variables from .envrc: ", err)
  }

  rpcURL := os.Getenv("ETH_RPC_URL")
  privateKeyAdmin := os.Getenv("PRIVATE_KEY")

  log.Println("Deploying the L1 smart contracts...")

  deployCmd1 := exec.Command("forge", "script", "scripts/Deploy.s.sol:Deploy", "--private-key="+privateKeyAdmin, "--broadcast", "--rpc-url="+rpcURL)
  err = runCommandWithOutput(deployCmd1)
  if err != nil {
    log.Println("Error deploying the L1 contracts with deployCmd1:", err)
  }

  deployCmd2 := exec.Command("forge", "script", "scripts/Deploy.s.sol:Deploy", "--sig", "sync()", "--private-key="+privateKeyAdmin, "--broadcast", "--rpc-url="+rpcURL)
  err = runCommandWithOutput(deployCmd2)
  if err != nil {
    log.Println("Error deploying the L1 contracts with deployCmd2:", err)
  }

  fmt.Println("L1 smart contracts deployed successfully!")
}
