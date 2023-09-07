package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type AuraRPC interface {
	AaveV3AuraLockedPosition(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV3AuraLockedReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV3AuraFarmingPool(ctx context.Context, contractAddress common.Address, blockNumber *big.Int) (common.Address, error)
	AaveV3AuraFarmingPoolID(ctx context.Context, contractAddress common.Address, blockNumber *big.Int) (*big.Int, error)
	AaveV3AuraFarmingBalanceEarned(ctx context.Context, account common.Address, contractAddress common.Address, blockNumber *big.Int) (farm, earned *decimal.Decimal, err error)
}

type AuraTokenRPC interface {
	ReductionPerCliff(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	TotalSupply(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	EmissionsMaxSupply(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
	TotalCliffs(ctx context.Context, blockNumber *big.Int) (*big.Int, error)
}
