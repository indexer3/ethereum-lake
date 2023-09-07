package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type SynthetixRPC interface {
	SynthetixGetSusdDebtBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	SynthetixGetSethDebtBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	SynthetixGetCollateral(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	SynthetixRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error)
	SynthetixGetVestingDetails(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, uint64, *decimal.Decimal, error)
	SynthetixLiquidatorRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
}

func (n *NodeClient) SynthetixGetSusdDebtBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) SynthetixGetSethDebtBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) SynthetixGetCollateral(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) SynthetixRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error) {
	return nil, nil, nil
}

func (n *NodeClient) SynthetixGetVestingDetails(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, uint64, *decimal.Decimal, error) {
	return nil, 0, nil, nil
}

func (n *NodeClient) SynthetixLiquidatorRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}
