# op-stack-deployer

OP Stack Deployer is a tool that simplifies the process of setting up and deploying an OP Stack chain<!-- on the Ethereum Goerli testnet -->. This guide will walk you through the steps required to spin up your own OP Stack chain and perform tests or customize it for your specific needs.

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
- You also need to the to configure the `.envrc` file in the `optimism/packages/contracts-bedrock` directory, the `ETH_RPC_URL` should be the L1 RPC URL that you have used in the previous script, replace the `PRIVATE_KEY` with the actual private key of the `Admin` which is in the `keys.txt`, and the `DEPLOYMENT_CONTEXT` stays the same which is `getting-started`.

## Contributing

Contributions to Op-Stack Deployer are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
