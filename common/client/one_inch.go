package client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/shopspring/decimal"
)

var _ OneInchRPC = (*NodeClient)(nil)

type OneInchRPC interface {
	TokenPrice(ctx context.Context, network constant.Network, tokenAddress common.Address) (*decimal.Decimal, error)
}

func (n *NodeClient) TokenPrice(ctx context.Context, network constant.Network, tokenAddress common.Address) (*decimal.Decimal, error) {
	return nil, nil
}
