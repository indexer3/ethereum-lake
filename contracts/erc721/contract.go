package erc721

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexer3/ethereum-lake/constant/event"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./erc721.abi --pkg erc721 --type ERC721 --out ./erc721.go

func FilterERC721Transfer(l types.Log) bool {
	return len(l.Topics) == 4 && l.Topics[0] == event.TransferEvent
}
