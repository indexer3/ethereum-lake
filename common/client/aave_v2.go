package client

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/aave"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var _ AaveV2RPC = (*NodeClient)(nil)

type AaveV2RPC interface {
	AaveV2UnClaimedRewards(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2StakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AbptStakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AbptStakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)

	AaveV2AccountData(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*model.AaveV2AccountData, error)
	AaveV2AccountHealthRate(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountActiveAssets(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (deposits, borrows []common.Address, err error)

	AaveV2AccountATokenBalance(ctx context.Context, network constant.Network, account, assetAddress common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2AccountDebt(ctx context.Context, network constant.Network, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	AaveV2RewardToken(ctx context.Context, network constant.Network, blockNumber *big.Int) (*common.Address, error)
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
	}, blockNumber)
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
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2StakedToken].Pack("balanceOf", account)
	if err != nil {
		log.Error("failed to pack aave v2 balanceOf data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2StakedToken(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 balanceOf", zap.Error(err))
		return nil, err
	}

	var balance = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeAaveV2StakedToken].UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		log.Error("failed to unpack aave v2 balanceOf", zap.Error(err))
		return nil, err
	}

	res := decimal.NewFromBigInt(balance, 0)

	return lo.ToPtr(res), nil
}

func (n *NodeClient) AaveV2AbptStakedReward(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2StakedAbpt].Pack("getTotalRewardsBalance", account)
	if err != nil {
		log.Error("failed to pack aave v2 getTotalRewardsBalance data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2StakedAbptToken(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 getTotalRewardsBalance", zap.Error(err))
		return nil, err
	}

	var rewards = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeAaveV2StakedAbpt].UnpackIntoInterface(&rewards, "getTotalRewardsBalance", result)
	if err != nil {
		log.Error("failed to unpack aave v2 getTotalRewardsBalance", zap.Error(err))
		return nil, err
	}

	res := decimal.NewFromBigInt(rewards, 0)

	return lo.ToPtr(res), nil
}

func (n *NodeClient) AaveV2AbptStakedTokenBalance(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2StakedAbpt].Pack("balanceOf", account)
	if err != nil {
		log.Error("failed to pack aave v2 balanceOf data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2StakedAbptToken(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 balanceOf", zap.Error(err))
		return nil, err
	}

	var balance = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeAaveV2StakedAbpt].UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		log.Error("failed to unpack aave v2 balanceOf", zap.Error(err))
		return nil, err
	}

	res := decimal.NewFromBigInt(balance, 0)

	return lo.ToPtr(res), nil
}

func (n *NodeClient) AaveV2AccountData(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*model.AaveV2AccountData, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].Pack("getUserAccountData", account)
	if err != nil {
		log.Error("failed to pack aave v2 getUserAccountData data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2LendingPool(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 getUserAccountData", zap.Error(err))
		return nil, err
	}

	var data model.AaveV2AccountData
	err = contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].UnpackIntoInterface(&data, "getUserAccountData", result)
	if err != nil {
		log.Error("failed to unpack aave v2 getUserAccountData", zap.Error(err))
		return nil, err
	}

	return lo.ToPtr(data), nil
}

func (n *NodeClient) AaveV2AccountHealthRate(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	data, err := n.AaveV2AccountData(ctx, network, account, blockNumber)
	if err != nil {
		log.Error("failed to get aave v2 account data", zap.Error(err))
		return nil, err
	}

	var healthRate float64
	if data.Ltv.Int64() != 0 {
		rate := new(big.Float).SetInt(data.HealthFactor)
		rate = rate.Quo(rate, new(big.Float).SetFloat64(math.Pow10(18)))
		healthRate, _ = rate.Float64()
	}

	return lo.ToPtr(decimal.NewFromFloat(healthRate)), nil
}

func (n *NodeClient) AaveV2AccountActiveAssets(ctx context.Context, network constant.Network, account common.Address, blockNumber *big.Int) (deposits, borrows []common.Address, err error) {
	calls := make([]multicall.Multicall3Call3, 0)

	var (
		userConfigurationCall = multicall.Multicall3Call3{
			Target:       aave.AaveV2LendingPool(network),
			AllowFailure: false,
		}
		reservesListCall = multicall.Multicall3Call3{
			Target:       aave.AaveV2LendingPool(network),
			AllowFailure: false,
		}
	)

	userConfigurationCall.CallData, err = contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].Pack("getUserConfiguration", account)
	if err != nil {
		log.Error("failed to pack aave v2 getUserConfiguration data", zap.Error(err))
		return nil, nil, err
	}

	reservesListCall.CallData, err = contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].Pack("getReservesList")
	if err != nil {
		log.Error("failed to pack aave v2 getReservesList data", zap.Error(err))
		return nil, nil, err
	}

	calls = append(calls, userConfigurationCall, reservesListCall)

	result, err := n.AggregatedCalls(ctx, calls, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 getUserConfiguration", zap.Error(err))
		return nil, nil, err
	}

	if len(result) != 2 {
		err = fmt.Errorf("failed to call aave v2 getUserConfiguration, result length is not 2, %+v", result)
		return nil, nil, err
	}

	for i, res := range result {
		if !res.Success {
			err = fmt.Errorf("failed to call erc20 %d th method, contract address: %s", i, aave.AaveV2LendingPool(network).String())
			return nil, nil, err
		}
	}

	var (
		userConfiguration = new(aave.DataTypesUserConfigurationMap)
		reservesList      = make([]common.Address, 0)
	)

	err = contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].UnpackIntoInterface(&userConfiguration, "getUserConfiguration", result[0].ReturnData)
	if err != nil {
		log.Error("failed to unpack aave v2 getUserConfiguration", zap.Error(err))
		return nil, nil, err
	}

	err = contracts.ABIs[contracts.ContractTypeAaveV2LendingPool].UnpackIntoInterface(&reservesList, "getReservesList", result[1].ReturnData)
	if err != nil {
		log.Error("failed to unpack aave v2 getReservesList", zap.Error(err))
		return nil, nil, err
	}

	bitmap := userConfiguration.Data
	for i := 0; i < bitmap.BitLen(); i++ {
		bit := bitmap.Bit(i)
		if bit == 1 {
			if i%2 == 0 {
				borrows = append(borrows, reservesList[i/2])
			} else {
				deposits = append(deposits, reservesList[i/2])
			}
		}
	}

	return deposits, borrows, nil
}

func (n *NodeClient) AaveV2AccountATokenBalance(ctx context.Context, network constant.Network, account common.Address, assetAddress common.Address,
	blockNumber *big.Int) (*decimal.Decimal, error) {

	var (
		eg   errgroup.Group
		data = new(model.AaveV2UserReserveData)
	)

	eg.Go(func() error {
		calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2DataProvider].Pack("getUserReserveData", account)
		if err != nil {
			log.Error("failed to pack aave v2 getUserReserveData data", zap.Error(err))
			return err
		}

		result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
			To:   lo.ToPtr[common.Address](aave.AaveV2DataProvider(network)),
			Data: calldata,
		}, blockNumber)
		if err != nil {
			log.Error("failed to call aave v2 getUserReserveData", zap.Error(err))
			return err
		}

		err = contracts.ABIs[contracts.ContractTypeAaveV2DataProvider].UnpackIntoInterface(&data, "getUserReserveData", result)
		if err != nil {
			log.Error("failed to unpack aave v2 getUserReserveData", zap.Error(err))
			return err
		}

		return nil
	})

	var decimals uint8
	eg.Go(func() error {
		subDecimals, err := n.ERC20Decimals(ctx, assetAddress)
		if err != nil {
			log.Error("failed to get erc20 decimals", zap.Error(err))
			return err
		}

		decimals = subDecimals
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Error("failed to get aave v2 account aToken balance", zap.Error(err))
		return nil, err
	}

	balance := new(big.Float).SetInt(data.CurrentATokenBalance)

	balance = balance.Quo(balance, new(big.Float).SetFloat64(math.Pow10(int(decimals))))

	res, _ := balance.Float64()
	return lo.ToPtr(decimal.NewFromFloat(res)), nil
}

func (n *NodeClient) AaveV2AccountDebt(ctx context.Context, network constant.Network, account common.Address, asset common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	var (
		eg       errgroup.Group
		decimals uint8
		data     model.AaveV2UserReserveData
	)

	eg.Go(func() error {
		calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2DataProvider].Pack("getUserReserveData", asset, account)
		if err != nil {
			log.Error("failed to pack aave v2 getUserReserveData data", zap.Error(err))
			return err
		}

		result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
			To:   lo.ToPtr[common.Address](aave.AaveV2DataProvider(network)),
			Data: calldata,
		}, blockNumber)
		if err != nil {
			log.Error("failed to call aave v2 getUserReserveData", zap.Error(err))
			return err
		}

		err = contracts.ABIs[contracts.ContractTypeAaveV2DataProvider].UnpackIntoInterface(&data, "getUserReserveData", result)
		if err != nil {
			log.Error("failed to unpack aave v2 getUserReserveData", zap.Error(err))
			return err
		}

		return nil
	})

	eg.Go(func() error {
		subDecimals, err := n.ERC20Decimals(ctx, asset)
		if err != nil {
			log.Error("failed to get erc20 decimals", zap.Error(err))
			return err
		}

		decimals = subDecimals
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Error("failed to get aave v2 account debt", zap.Error(err))
		return nil, err
	}

	debt := new(big.Float).SetInt(data.CurrentStableDebt)
	debt = debt.Add(debt, new(big.Float).SetInt(data.CurrentVariableDebt))
	debt = debt.Quo(debt, new(big.Float).SetFloat64(math.Pow10(int(decimals))))

	debtFloat, _ := debt.Float64()

	return lo.ToPtr(decimal.NewFromFloat(debtFloat)), nil
}

func (n *NodeClient) AaveV2RewardToken(ctx context.Context, network constant.Network, blockNumber *big.Int) (*common.Address, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeAaveV2Incentives].Pack("REWARD_TOKEN")
	if err != nil {
		log.Error("failed to pack aave v2 REWARD_TOKEN data", zap.Error(err))
		return nil, err
	}

	result, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](aave.AaveV2Incentives(network)),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call aave v2 REWARD_TOKEN", zap.Error(err))
		return nil, err
	}

	var rewardToken common.Address
	err = contracts.ABIs[contracts.ContractTypeAaveV2Incentives].UnpackIntoInterface(&rewardToken, "REWARD_TOKEN", result)
	if err != nil {
		log.Error("failed to unpack aave v2 REWARD_TOKEN", zap.Error(err))
		return nil, err
	}

	return lo.ToPtr(rewardToken), nil
}
