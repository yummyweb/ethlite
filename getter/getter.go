package getter

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlockHeader(cl *ethclient.Client) *types.Header {
	h, err := cl.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic("Could not retrieve latest block header!")
	}

	hbytes, err := h.MarshalJSON()
	if err != nil {
		panic("Could not Marshal header!")
	}

	err = os.WriteFile("/tmp/ethlite-cache/" + h.Hash().String(), hbytes, 0644)
	// Most probably, the directory `ethlite-cache` doesn't exist, so we try and create that
	// directory, and then write the block header cache.
	if err != nil {
		err = os.MkdirAll("/tmp/ethlite-cache", 0755)
		if err != nil {
			panic("Could not create cache for block header!")
		}

		err = os.WriteFile("/tmp/ethlite-cache/" + h.Hash().String(), hbytes, 0644)
		if err != nil {
			panic("Could not create cache for block header!")
		}
	}

	return h
}

// Recursively get all blocks
func GetBlocksAll(cl *ethclient.Client, hash common.Hash) {
	if hash == common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000") {
		fmt.Println("[INFO] Reached coinbase transaction")
		fmt.Println("[INFO] Read all blocks")
		return
	}

	header, err := cl.HeaderByHash(context.Background(), hash)

	if err != nil {
		panic("Could not retrieve block header by hash!")
	}

	fmt.Println(header)

	GetBlocksAll(cl, header.ParentHash)
}