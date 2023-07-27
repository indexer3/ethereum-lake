package engine

import (
	"context"
	"math/big"

	"github.com/indexer3/ethereum-lake/constant"
)

type DispatchTaskIndex struct {
	BlockNumber *big.Int `json:"block_number,omitempty"`

	Network constant.Network `json:"network,omitempty"`
	TxHash  string           `json:"tx_hash,omitempty"`
}

type ITask interface {
	Name() string
	Handle(ctx context.Context, taskIndex DispatchTaskIndex) error
}
