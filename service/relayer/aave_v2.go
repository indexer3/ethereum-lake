package relayer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/shopspring/decimal"
)

type IAaveV2Relayer interface {
	GetUnClaimedRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetStakedReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetStakedTokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	GetAbptStakedReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetAbptStakedTokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	GetAccountTVL(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetAccountHealthRate(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetAccountActiveAssets(ctx context.Context, account common.Address, blockNumber *big.Int) (deposits []model.ERC20Token, borrows []model.ERC20Token, err error)

	GetAccountATokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetAccountDebt(ctx context.Context, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GetRewardToken(ctx context.Context, blockNumber *big.Int) (common.Address, error)
}
