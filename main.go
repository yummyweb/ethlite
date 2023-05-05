package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yummyweb/ethlite/getter"
)

func main() {
	// Check if some command is provided. If not, error.
	if len(os.Args) <= 1 {
		log.Fatal("NO COMMAND PROVIDED! EXITING...")
	}
    commands := os.Args[1:]
    fmt.Println(commands)

	if commands[0] == "start" {
		fmt.Println("[INFO] Starting light client...")
		fmt.Println("[INFO] Retrieving latest block header...")
		cl := InstantiateClient()
		h := getter.GetLatestBlockHeader(cl)
		fmt.Printf("[INFO] Latest header: %v   Parent hash: %v\n", h.Root, h.ParentHash)
		getter.GetBlocksAll(cl, h.ParentHash)
	} else {
		log.Fatal("[ERROR] Invalid command: ", commands[0])
	}
}

func InstantiateClient() *ethclient.Client {
	client, err := ethclient.Dial("/Users/antarikshverma/dev-chain/geth.ipc")

	if err != nil {
		panic("Could not open IPC file!")
	}

	return client
}