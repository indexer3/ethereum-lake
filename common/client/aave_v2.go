package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/aave"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ AaveV2RPC = (*NodeClient)(nil)

type AaveV2RPC interface {
	AaveV2UnClaimedRewards(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AbptStakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AbptStakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AccountTVL(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountHealthRate(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountActiveAssets(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (deposits []model.ERC20Token, borrows []model.ERC20Token, err error)

	AaveV2AccountATokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountDebt(ctx context.Context, network constant.Network, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2RewardToken(ctx context.Context, network constant.Network, blockNumber *big.Int) (common.Address, error)
}

func (n *NodeClient) SupportNetwork(network constant.Network) bool {
	return network == constant.NetworkEthereum || network == constant.NetworkPolygon
}

func (n *NodeClient) SupportNetworks() []constant.Network {
	return []constant.Network{constant.NetworkEthereum, constant.NetworkPolygon}
}

func (n *NodeClient) AaveV2UnClaimedRewards(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2Incentives].Pack("getUserUnclaimedRewards", account)
	if err != nil {
		log.Error("failed to pack aave v2 getUserUnclaimedRewards data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2Incentives(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 getUserUnclaimedRewards", zap.Error(err))
		return nil, err
	}

	var rewards = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeAaveV2Incentives].UnpackIntoInterface(&rewards, "getUserUnclaimedRewards", result)
	if err != nil {
		log.Error("failed to unpack aave v2 getUserUnclaimedRewards", zap.Error(err))
		return nil, err
	}

	res := decimal.NewFromBigInt(rewards, 0)

	return lo.ToPtr(res), nil
}

func (n *NodeClient) AaveV2StakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2StakedToken].Pack("getTotalRewardsBalance", account)
	if err != nil {
		log.Error("failed to pack aave v2 getTotalRewardsBalance data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2StakedToken(network)),
		Data: calldata,
	}, nil)
	if err != nil {
		log.Error("failed to call aave v2 getTotalRewardsBalance", zap.Error(err))
		return nil, err
	}

	var rewards = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeAaveV2StakedToken].UnpackIntoInterface(&rewards, "getTotalRewardsBalance", result)
	if err != nil {
		log.Error("failed to unpack aave v2 getTotalRewardsBalance", zap.Error(err))
		return nil, err
	}

	res := decimal.NewFromBigInt(rewards, 0)

	return lo.ToPtr(res), nil
}

func (n *NodeClient) AaveV2StakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AbptStakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AbptStakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AccountTVL(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AccountHealthRate(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AccountActiveAssets(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (deposits []model.ERC20Token, borrows []model.ERC20Token, err error) {
	return nil, nil, nil
}

func (n *NodeClient) AaveV2AccountATokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2AccountDebt(ctx context.Context, network constant.Network, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	return nil, nil
}

func (n *NodeClient) AaveV2RewardToken(ctx context.Context, network constant.Network, blockNumber *big.Int) (common.Address, error) {
	return common.HexToAddress(""), nil
}
