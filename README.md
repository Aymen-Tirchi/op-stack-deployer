# op-stack-deployer

OP Stack Deployer is a tool that simplifies setting up and deploying an OP Stack chain on the Ethereum Sepolia testnet. This guide will walk you through the steps required to spin up your OP Stack chain and perform tests or customize it for your specific needs.

This project is based on the [Creating your own L2 rollup testnet]([https://stack.optimism.io/docs/build/getting-started/#](https://docs.optimism.io/operators/chain-operators/tutorials/create-l2-rollup)).

## Prerequisites

Before getting started, ensure that you have the following software installed: `Git`, `Go`, `Node`, `Pnpm`, `Foundry`, `Make`, `jq`, and `direnv`.

## Getting Started

First of all clone the op-stack-deployer repository
```bash
git clone https://github.com/Aymen-Tirchi/op-stack-deployer.git && cd op-stack-deployer
```
And then follow these steps: 

1. Build the Optimism Monorepo

```bash
go run cmd/1_build_optimism/main.go
```
This script will automatically clone the Optimism Monorepo, install the required modules, build the necessary packages, and generate the Optimism Monorepo and packages successfully.

2. Build op-geth

```bash
go run cmd/2_build_op-geth/main.go
```
This script will automatically clone the op-geth repo, build the necessary packages, and generate the op-geth repo and packages successfully.

3. Generate some keys

```bash
go run cmd/3_generate_keys/main.go
```
This script will generate the keys of each role and store them in a text file named `keys.txt` in the root directory of the project. The `keys.txt` file will contain the addresses and private keys for the `Admin`, `Proposer`, `Batcher`, and `Sequencer` accounts.

4. Configure your network

- Before you run the script go to this path `optimism/packages/contracts-bedrock` you will find `.envrc.`, set the `ETH_RPC_URL` that you are using, and replace the `PRIVATE_KEY` with the actual private key of the `Admin` which is in the `keys.txt`, and the `DEPLOYMENT_CONTEXT` stays the same which is `getting-started`.

```bash
go run cmd/4_configure_network/main.go
```
- This script will automatically configure your network based on the generated keys and the provided L1 node RPC URL. It will configure `getting-started.json` in the `optimism/packages/contracts-bedrock/deploy-config` directory, which contains all the required parameters for your network setup.

5. Deploy the L1 contracts
- Before running the `deploy_L1_contracts.go` script, ensure that you have funded your `Admin` address with some Sepolia test ETH (at least 0.5 ETH). Having sufficient test ETH will cover the gas costs and ensure the successful deployment of the L1 contracts.

- Now you can run the script :
```bash
go run cmd/5_deploy_L1_contracts/main.go
```
The script will start deploying all the L1 smart contracts. During the deployment process, you may see various transaction logs and updates. Once the deployment is successful, you will receive a confirmation message.

6. Generate the L2 config files
```bash
sudo go run cmd/6_L2_config/main.go
```
This script will automatically create the necessary L2 configuration files `genesis.json`, `rollup.json`, and `jwt.txt`. These files are crucial for the configuration and secure communication between the op-node and op-geth.

7. Initialize op-geth
```bash
go run cmd/7_Initialize_op-geth/main.go
```
This script will create a data directory and initialize the `op-geth` with the `genesis.json` we generated in the previous script.

8. Run the node software

before running anything make sure you fund your `batcher` and `proposer` addresses with at least 0.5 Sepolia test ETH, to ensure that it can continue operating without running out of ETH for gas.
- Run op-geth 
```bash
go run cmd/8_run_op-geth/main.go
```
This script will run the op-geth node.
- Run op-node

```bash
go run cmd/9_run_op-node/main.go 
```
This script will set up system variables and run `op-node`.

- Run op-batcher
 
```bash
go run cmd/10_run_op-batcher/main.go 
```
This script will set up system variables and run `op-batcher`.

might have warning messages similar to: 
```bash
WARN [03-21|14:13:55.248] Error calculating L2 block range         err="failed to get sync status: Post \"http://localhost:8547\": context deadline exceeded"
```
This means that `op-node` is not yet synchronized up to the present time. Just wait until it is.

- Run op-proposer

```bash
go run cmd/11_run_op-proposer/main.go 
```
This script will set up system variables and run the `op-proposer`.

check out this [Rollup Operations](https://stack.optimism.io/docs/build/operations/#) 

9. Get some ETH on your Rollup

To get the address of your Rollup run the following command
```bash
go run cmd/12_get_rollup_address/main.go 
```
and now you can fund your rollup address with some ETH. It may take up to 5 minutes for that ETH to appear in your wallet on L2.

Congratulations, You have a complete OP Stack-based EVM Rollup.

10. Use your Rollup

To see your rollup in action, you can use the [optimism-tutorial](https://github.com/ethereum-optimism/optimism-tutorial/tree/main).

## Contributing

Contributions to Op-Stack Deployer are welcome! If you have any issues or ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
