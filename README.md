# op-stack-deployer

OP Stack Deployer is a tool that simplifies the process of setting up and deploying an OP Stack chain<!-- on the Ethereum Goerli testnet -->. This guide will walk you through the steps required to spin up your own OP Stack chain and perform tests or customize it for your specific needs.

This project is based on the [OP Stack Getting Started doc](https://stack.optimism.io/docs/build/getting-started/#).

## Prerequisites

Before getting started, ensure that you have the following software installed: `Git`, `Go`, `Node`, `Pnpm`, `Foundry`, `Make`, `jq`, and `direnv`.

1. Build the Source Code:

```bash
go run build_optimism/build_optimism.go
```
This will automatically clone the Optimism Monorepo, install required modules, build the necessary packages, and generate the Optimism Monorepo and packages successfully.

2. Build op-geth

```bash
go run build_op-geth/build_op-geth.go
```
This will automatically clone the op-geth repo, build the necessary packages, and generate the op-geth repo and packages successfully.

## Contributing

Contributions to Op-Stack Deployer are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).