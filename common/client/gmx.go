package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/gmx"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ GMXRPC = (*NodeClient)(nil)

type GMXRPC interface {
	StakeEthReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	StakeEsGmxReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	StakeGlp(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	StakeGlpRewardEth(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	StakeGmx(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	StakeEsGmx(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
	GmxVesterYield(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error)
	GlpVesterYield(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error)
}

func (n *NodeClient) StakeEthReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXRewardTracker].Pack("claimable", account)
	if err != nil {
		log.Error("failed to pack claimable for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.RewardTrackerSbfGmxAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx claimable", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXRewardTracker].UnpackIntoInterface(&res, "claimable", result)
	if err != nil {
		log.Error("failed to unpack gmx claimable result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) StakeEsGmxReward(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXRewardTracker].Pack("claimable", account)
	if err != nil {
		log.Error("failed to pack claimable for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.RewardTrackerSGMXAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx claimable", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXRewardTracker].UnpackIntoInterface(&res, "claimable", result)
	if err != nil {
		log.Error("failed to unpack gmx claimable result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) StakeGlp(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXStakedGLP].Pack("balanceOf", account)
	if err != nil {
		log.Error("failed to pack balanceOf for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.StakedGlpAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx balanceOf", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXStakedGLP].UnpackIntoInterface(&res, "balanceOf", result)
	if err != nil {
		log.Error("failed to unpack gmx balanceOf result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) StakeGlpRewardEth(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXRewardTracker].Pack("claimable", account)
	if err != nil {
		log.Error("failed to pack claimable for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.RewardTrackerGlpAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx claimable", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXRewardTracker].UnpackIntoInterface(&res, "claimable", result)
	if err != nil {
		log.Error("failed to unpack gmx claimable result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) StakeGmx(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXRewardTracker].Pack("depositBalances", account, gmx.GMXTokenAddress())
	if err != nil {
		log.Error("failed to pack depositBalances for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.RewardTrackerSGMXAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx depositBalances", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXRewardTracker].UnpackIntoInterface(&res, "depositBalances", result)
	if err != nil {
		log.Error("failed to unpack gmx depositBalances result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) StakeEsGmx(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeGMXRewardTracker].Pack("depositBalances", account, gmx.EsTokenAddress())
	if err != nil {
		log.Error("failed to pack depositBalances for gmx", zap.Error(err))
		return nil, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr[common.Address](gmx.RewardTrackerSGMXAddress()),
		Data: calldata,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call gmx depositBalances", zap.Error(err))
		return nil, err
	}

	var res = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXRewardTracker].UnpackIntoInterface(&res, "depositBalances", result)
	if err != nil {
		log.Error("failed to unpack gmx depositBalances result", zap.Error(err))
		return nil, err
	}

	reward := decimal.NewFromBigInt(res, 0)
	return lo.ToPtr(reward), nil
}

func (n *NodeClient) GmxVesterYield(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error) {
	balanceOfCall, err := contracts.ABIs[contracts.ContractTypeGMXVester].Pack("balanceOf", account)
	if err != nil {
		log.Error("failed to pack balanceOf for gmx", zap.Error(err))
		return nil, nil, err
	}

	claimableCall, err := contracts.ABIs[contracts.ContractTypeGMXVester].Pack("claimable", account)
	if err != nil {
		log.Error("failed to pack claimable for gmx", zap.Error(err))
		return nil, nil, err
	}

	multicallData := []multicall.Multicall3Call3{
		{
			Target:       gmx.VesterGMXAddress(),
			CallData:     balanceOfCall,
			AllowFailure: false,
		},
		{
			Target:       gmx.VesterGMXAddress(),
			CallData:     claimableCall,
			AllowFailure: false,
		},
	}

	results, err := n.AggregatedCalls(ctx, multicallData)
	if err != nil {
		log.Error("failed to aggregate call", zap.Error(err))
		return nil, nil, err
	}

	for i, res := range results {
		if !res.Success {
			return nil, nil, fmt.Errorf("failed to call erc20  %d th method, contract address: %s", i, account.String())
		}
	}

	if len(results) != 2 {
		return nil, nil, fmt.Errorf("failed to match the result, len of result: %d", len(results))
	}

	var balance = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXVester].UnpackIntoInterface(&balance, "balanceOf", results[0].ReturnData)
	if err != nil {
		log.Error("failed to unpack gmx balanceOf result", zap.Error(err))
		return nil, nil, err
	}

	var claimable = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXVester].UnpackIntoInterface(&balance, "claimable", results[1].ReturnData)
	if err != nil {
		log.Error("failed to unpack gmx balanceOf result", zap.Error(err))
		return nil, nil, err
	}

	yeildGmxBalance := decimal.NewFromBigInt(balance, 0)
	claimableNum := decimal.NewFromBigInt(claimable, 0)

	return lo.ToPtr(yeildGmxBalance), lo.ToPtr(claimableNum), nil
}

func (n *NodeClient) GlpVesterYield(ctx context.Context, account common.Address, blockNumber *big.Int) (*decimal.Decimal, *decimal.Decimal, error) {
	balanceOfCall, err := contracts.ABIs[contracts.ContractTypeGMXVester].Pack("balanceOf", account)
	if err != nil {
		log.Error("failed to pack balanceOf for gmx", zap.Error(err))
		return nil, nil, err
	}

	claimableCall, err := contracts.ABIs[contracts.ContractTypeGMXVester].Pack("claimable", account)
	if err != nil {
		log.Error("failed to pack claimable for gmx", zap.Error(err))
		return nil, nil, err
	}

	multicallData := []multicall.Multicall3Call3{
		{
			Target:       gmx.VesterGlpAddress(),
			CallData:     balanceOfCall,
			AllowFailure: false,
		},
		{
			Target:       gmx.VesterGlpAddress(),
			CallData:     claimableCall,
			AllowFailure: false,
		},
	}

	results, err := n.AggregatedCalls(ctx, multicallData)
	if err != nil {
		log.Error("failed to aggregate call", zap.Error(err))
		return nil, nil, err
	}

	for i, res := range results {
		if !res.Success {
			return nil, nil, fmt.Errorf("failed to call erc20  %d th method, contract address: %s", i, account.String())
		}
	}

	if len(results) != 2 {
		return nil, nil, fmt.Errorf("failed to match the result, len of result: %d", len(results))
	}

	var balance = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXVester].UnpackIntoInterface(&balance, "balanceOf", results[0].ReturnData)
	if err != nil {
		log.Error("failed to unpack gmx balanceOf result", zap.Error(err))
		return nil, nil, err
	}

	var claimable = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeGMXVester].UnpackIntoInterface(&balance, "claimable", results[1].ReturnData)
	if err != nil {
		log.Error("failed to unpack gmx balanceOf result", zap.Error(err))
		return nil, nil, err
	}

	yeildGlpBalance := decimal.NewFromBigInt(balance, 0)
	claimableNum := decimal.NewFromBigInt(claimable, 0)

	return lo.ToPtr(yeildGlpBalance), lo.ToPtr(claimableNum), nil
}
