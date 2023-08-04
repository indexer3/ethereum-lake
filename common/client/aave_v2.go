package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/shopspring/decimal"
)

type IAaveV2RPC interface {
	NetworkSupport
	AaveV2UnClaimedRewards(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedTokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AbptStakedReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AbptStakedTokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AccountTVL(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountHealthRate(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountActiveAssets(ctx context.Context, account common.Address, blockNumber *big.Int) (deposits []model.ERC20Token, borrows []model.ERC20Token, err error)

	AaveV2AccountATokenBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountDebt(ctx context.Context, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2RewardToken(ctx context.Context, blockNumber *big.Int) (common.Address, error)
}

func (n *NodeClient) SupportNetwork(network constant.Network) bool {
	return network == constant.NetworkEthereum || network == constant.NetworkPolygon
}

func (n *NodeClient) SupportNetworks() []constant.Network {
	return []constant.Network{constant.NetworkEthereum, constant.NetworkPolygon}
}
