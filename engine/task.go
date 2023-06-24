package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

type DispatchTaskIndex struct {
	BlockNumber uint64 `json:"block_number,omitempty"`

	Network constant.Network `json:"network,omitempty"`
	TxHash  common.Hash      `json:"tx_hash,omitempty"`
}

type ITask interface {
	Name() string
	Run(ctx context.Context) error
	Handle(ctx context.Context, taskIndex DispatchTaskIndex) error
}
