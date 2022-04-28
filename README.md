# eth-mempool-listener-go
A simple program able to listen to the pending transactions of the Ethereum mempool.

## How does it work ?

It creates a set of clients to query the mempool from a node using the endpoint you've provided as an environment 
variable.
Then, it converts the transaction to a readable message so that you can use its fields properly.

Take a look at the code and feel free to use it for your own needs !

## Getting started !

### Installation

Make sure you've downloaded the required dependencies using :
```shell
go get ./
```

### Quickstart

Provide the node endpoint in a .envrc file or export it in the shell environment :
```shell
export NODE_ENDPOINT=RUN_YOUR_OWN_NODE_TO_IMPROVE_DECENTRALISATION_;)
```

Then, run :
```shell
go run main.go
```

## Author

Made with ❤️ by 🤖 [Luca Georges François](https://github.com/PtitLuca) 🤖
