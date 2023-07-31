# op-stack-deployer

OP Stack Deployer is a tool that simplifies the process of setting up and deploying an OP Stack chainon the Ethereum Goerli testnet. This guide will walk you through the steps required to spin up your own OP Stack chain and perform tests or customize it for your specific needs.

This project is based on the [OP Stack Getting Started doc](https://stack.optimism.io/docs/build/getting-started/#).

## Prerequisites

Before getting started, ensure that you have the following software installed: `Git`, `Go`, `Node`, `Pnpm`, `Foundry`, `Make`, `jq`, and `direnv`.

## Getting Started

1. Build the Optimism Monorepo

```bash
go run build_optimism/build_optimism.go
```
This script will automatically clone the Optimism Monorepo, install the required modules, build the necessary packages, and generate the Optimism Monorepo and packages successfully.

2. Build op-geth

```bash
go run build_op-geth/build_op-geth.go
```
This script will automatically clone the op-geth repo, build the necessary packages, and generate the op-geth repo and packages successfully.

3. Generate some keys

```bash
go run generate_keys/generate_keys.go
```
This script will generate the keys of each role and store them in a text file named `keys.txt` in the root directory of the project. The `keys.txt` file will contain the addresses and private keys for the Admin, Proposer, Batcher, and Sequencer accounts.

4. Configure your network

```bash
go run configure_network/configure_network.go
```
- This script will automatically configure your network based on the generated keys and the provided L1 node RPC URL. It will configure `getting-started.json` in the `optimism/packages/contracts-bedrock/deploy-config` directory, which contains all the required parameters for your network setup. Make sure to fill in the correct values for the `ETH_RPC_URL` to ensure a successful deployment.
- You also need to configure the `.envrc` file in the `optimism/packages/contracts-bedrock` directory, the `ETH_RPC_URL` should be the L1 RPC URL that you have used in the previous script, replace the `PRIVATE_KEY` with the actual private key of the `Admin` which is in the `keys.txt`, and the `DEPLOYMENT_CONTEXT` stays the same which is `getting-started`.

5. Deploy the L1 contracts
- Before running the `deploy_L1_contracts.go` script, ensure that you have funded your Admin address with some Goerli test ETH (at least 0.5 ETH). Having sufficient test ETH will cover the gas costs and ensure the successful deployment of the L1 contracts.

- Now you can run the script :
```bash
go run deploy_L1_contracts/deploy_L1_contracts.go
```
The script will start deploying all the L1 smart contracts. During the deployment process, you may see various transaction logs and updates. Once the deployment is successful, you will receive a confirmation message.

6. Generate the L2 config files
```bash
sudo go run L2_config/L2_config.go
```
This script will automatically create the necessary L2 configuration files `genesis.json`, `rollup.json`, and `jwt.txt`. These files are crucial for the configuration and secure communication between the op-node and op-geth.

7. Initialize op-geth
```bash
go run Initialize_op-geth/initialize_op-geth.go
```
This script will create a data directory and initialize the `op-geth` with the `genesis.json` we generated in the previous script.

8. Run the node software
- Run op-geth 
```bash
go run run_op-geth/run_op-geth.go
```
This script will run the op-geth node.
- Run op-node

Before running op-node, ensure you've exported the following environment variables:
```bash
export SEQ_KEY=<Sequencer PrivateKey>
```
Replace `<Sequencer PrivateKey>` with the actual Sequencer Private key from `keys.txt`.
```bash
export L1_RPC=<ETH_RPC_URL>
```
Replace `<ETH_RPC_URL>` with the URL for the L1 (such as Goerli) you're using.
```bash
export RPC_KIND=<L1 server>
```
Replace `<L1 server>` with The type of L1 server to which you connecting (e.g., `alchemy`, `quicknode`).

and then run this command: 
```bash
go run run_op-node/run_op-node.go
```
This script will run the op-node.

- Run op-batcher

Before running op-batcher, ensure you've exported the following environment variables:
```bash
export BATCHER_KEY=<Batcher PrivateKey>
```
Replace `<Batcher PrivateKey>` with the actual Batcher Private key from `keys.txt`.
```bash
export L1_RPC=<ETH_RPC_URL>
```
Replace `<ETH_RPC_URL>` with the URL for the L1 (such as Goerli) you're using.

And make sure you fund your batcher address with at least 1 Goerli test ETH, to ensure that it can continue operating without running out of ETH for gas.

and then run this command: 
```bash
go run run_op-batcher/run_op-batcher.go
```
This script will run the op-batcher.

- Run op-proposer

Before running op-proposer, ensure you've exported the following environment variables:
```bash
export PROPOSER_KEY=<Proposer PrivateKey>
```
Replace `<Proposer PrivateKey>` with the actual Proposer Private key from `keys.txt`.
```bash
export L1_RPC=<ETH_RPC_URL>
```
Replace `<ETH_RPC_URL>` with the URL for the L1 (such as Goerli) you're using.
```bash
export L2OO_ADDR=<L2OutputOracleProxy address>
```
Replace `<L2OutputOracleProxy address>` with the actual address of `L2OutputOracleProxy` found in `optimism/packages/contracts-bedrock/deployments/goerli/L2OutputOracleProxy.json`.

and then run this command: 
```bash
go run run_op-proposer/run_op-proposer.go
```
This script will run the op-proposer.

9. Get some ETH on your Rollup

To get the address of your rollup run the following command
```bash
go run get_rollup_address/get_rollup_address.go
```
and know you can fund your rollup address with some ETH. It may take up to 5 minutes for that ETH to appear in your wallet on L2.

Congratulations, You have a complete OP Stack based EVM Rollup.

10. Use your Rollup

- Before starting set Set RPC URL and the contract address.
```bash
export ETH_RPC_URL=<ETH_RPC_URL>
```
```bash
export GREETER=0x575E9B4f2c3945d7CF07cb76628d29DF471692B8
```

- And then run this command: 
```bash
go run rollup/rollup.go
```

## Contributing

Contributions to Op-Stack Deployer are welcome! If you have any issues or ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
