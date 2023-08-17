package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/shopspring/decimal"
)

type AaveV3RPC interface {
	NetworkSupport
	AaveV3AllReservesTokens(ctx context.Context, blockNumber *big.Int) ([]model.ERC20Token, error)
	AaveV3AssetBalance(ctx context.Context, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV3LiquidationThreshold(ctx context.Context, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV3UserRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
}
