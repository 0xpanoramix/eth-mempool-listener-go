package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"os"
)

var (
	NodeEndpoint = ""
)

func init() {
	NodeEndpoint = os.Getenv("NODE_ENDPOINT")
}

func main() {
	ctx := context.Background()
	txnsHash := make(chan common.Hash)

	baseClient, err := rpc.Dial(NodeEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	ethClient, err := ethclient.Dial(NodeEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	subscriber := gethclient.New(baseClient)
	_, err = subscriber.SubscribePendingTransactions(ctx, txnsHash)

	if err != nil {
		log.Fatalln(err)
	}

	signer := types.NewLondonSigner(chainID)

	defer func() {
		baseClient.Close()
		ethClient.Close()
	}()

	for txnHash := range txnsHash {
		txn, _, err := ethClient.TransactionByHash(ctx, txnHash)
		if err != nil {
			continue
		}

		message, err := txn.AsMessage(signer, nil)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(txnHash.String())
		fmt.Println(message.To())
	}
}
