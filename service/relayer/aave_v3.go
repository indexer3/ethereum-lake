package relayer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/shopspring/decimal"
)

type IAaveV3Relayer interface {
	GetAllReservesTokens(ctx context.Context, blockNumber *big.Int) ([]model.ERC20Token, error)
	GetAssetBalance(ctx context.Context, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetLiquidationThreshold(ctx context.Context, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetUserRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
}
