package erc20

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexer3/ethereum-lake/constant/event"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./erc20.abi --pkg erc20 --type ERC20 --out ./erc20.go

func FilterERC20Transfer(l types.Log) bool {
	return len(l.Topics) == 3 && l.Topics[0] == event.TransferEvent
}

func FilterERC20Approval(l types.Log) bool {
	return len(l.Topics) == 3 && l.Topics[0] == event.ApprovalEvent
}
