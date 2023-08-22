package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts"
	one_inch "github.com/indexer3/ethereum-lake/contracts/1inch"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ OneInchRPC = (*NodeClient)(nil)

type OneInchRPC interface {
	TokenPrice(ctx context.Context, network constant.Network, tokenAddress common.Address, blockNumber *big.Int) (*decimal.Decimal, error)
}

func (n *NodeClient) TokenPrice(ctx context.Context, network constant.Network, tokenAddress common.Address, blockNumber *big.Int) (*decimal.Decimal, error) {
	getRateCall, err := contracts.ABIs[contracts.ContractTypeOneInchOracle].Pack("getRate", tokenAddress, constant.USDContractAddress(network), true)
	if err != nil {
		log.Error("failed to pack one inch getRate data", zap.Error(err))
		return nil, err
	}

	// decimalCall, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("decimals")
	// if err != nil {
	// 	log.Error("failed to pack erc20 decimals data", zap.Error(err))
	// 	return nil, err
	// }

	// packCall, err := contracts.ABIs[contracts.ContractTypeMulticall].Pack("aggregate3", []multicall.Multicall3Call3{
	// 	{
	// 		Target:       one_inch.OneInchPriceOracle(network),
	// 		CallData:     getRateCall,
	// 		AllowFailure: false,
	// 	},
	// 	{
	// 		Target:       tokenAddress,
	// 		CallData:     decimalCall,
	// 		AllowFailure: false,
	// 	},
	// })
	// if err != nil {
	// 	log.Error("failed to pack multicall aggregate data", zap.Error(err))
	// 	return nil, err
	// }

	resultBytes, err := n.ETHClient().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr(one_inch.OneInchPriceOracle(network)),
		Data: getRateCall,
	}, blockNumber)
	if err != nil {
		log.Error("failed to call multicall aggregate data", zap.Error(err))
		return nil, err
	}

	// var results []multicall.Multicall3Result
	// err = contracts.ABIs[contracts.ContractTypeMulticall].UnpackIntoInterface(&results, "aggregate3", resultBytes)
	// if err != nil {
	// 	log.Error("failed to unpack multicall aggregate3 result", zap.String("tokenAddress", tokenAddress.String()), zap.Error(err))
	// 	return nil, err
	// }

	// if len(results) != 2 {
	// 	log.Error("failed to parse multicall aggregate data", zap.Error(err))
	// 	return nil, err
	// }

	var rate = new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeOneInchOracle].UnpackIntoInterface(&rate, "getRate", resultBytes)
	if err != nil {
		log.Error("failed to unpack one inch getRate result", zap.String("tokenAddress", tokenAddress.String()), zap.Error(err))
		return nil, err
	}

	// var decimals uint8
	// err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&decimals, "decimals", results[1].ReturnData)
	// if err != nil {
	// 	log.Error("failed to unpack erc20 decimals result", zap.String("tokenAddress", tokenAddress.String()), zap.Error(err))
	// 	return nil, err
	// }

	return lo.ToPtr[decimal.Decimal](decimal.NewFromBigInt(rate, -int32(constant.USDDecimals(network)))), nil
}
