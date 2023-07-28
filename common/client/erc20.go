package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	indexerCommon "github.com/indexer3/ethereum-lake/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (n *NodeClient) AggregatedERC20Token(ctx context.Context, address common.Address) (*model.ERC20Token, error) {
	calls := make([]multicall.Multicall3Call3, 0)

	var err error
	symbolCall := multicall.Multicall3Call3{
		Target:       address,
		AllowFailure: false,
	}

	symbolCall.CallData, err = contracts.ABIs[contracts.ContractTypeERC20].Pack("symbol")
	if err != nil {
		return nil, err
	}

	calls = append(calls, symbolCall)

	totalSupplyCall := multicall.Multicall3Call3{
		Target:       address,
		AllowFailure: false,
	}
	totalSupplyCall.CallData, err = contracts.ABIs[contracts.ContractTypeERC20].Pack("totalSupply")
	if err != nil {
		return nil, err
	}
	calls = append(calls, totalSupplyCall)

	decimalsCall := multicall.Multicall3Call3{
		Target:       address,
		AllowFailure: false,
	}
	decimalsCall.CallData, err = contracts.ABIs[contracts.ContractTypeERC20].Pack("decimals")
	if err != nil {
		return nil, err
	}

	calls = append(calls, decimalsCall)

	results, err := n.AggregatedCalls(ctx, calls)
	if err != nil {
		return nil, err
	}

	for i, res := range results {
		if !res.Success {
			return nil, fmt.Errorf("failed to call erc20  %d th method, contract address: %s", i, address.String())
		}
	}

	if len(results) != 3 {
		return nil, fmt.Errorf("failed to match the result, len of result: %d", len(results))
	}

	erc20Token := new(model.ERC20Token)
	erc20Token.ContractAddress = address

	var symbol string
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&symbol, "symbol", results[0].ReturnData)
	if err != nil {
		return nil, err
	}
	var totalSupply *big.Int
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&totalSupply, "totalSupply", results[1].ReturnData)
	if err != nil {
		return nil, err
	}
	var decimals uint8
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&decimals, "decimals", results[2].ReturnData)
	if err != nil {
		return nil, err
	}

	return &model.ERC20Token{
		ContractAddress: address,
		Decimals:        decimals,
		Symbol:          symbol,
	}, nil
}

func (n *NodeClient) ERC20Symbol(ctx context.Context, address common.Address) (string, error) {
	_calldata, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("symbol")
	if err != nil {
		return "", err
	}

	data, err := n.Call(ctx, CallParam{
		To:       address,
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return "", err
	}

	var symbol string
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&symbol, "symbol", data)
	if err != nil {
		return "", err
	}

	return symbol, nil
}

func (n *NodeClient) ERC20TotalSupply(ctx context.Context, address common.Address) (*big.Int, error) {
	_calldata, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("totalSupply")
	if err != nil {
		return nil, err
	}

	data, err := n.Call(ctx, CallParam{
		To:       address,
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return nil, err
	}

	totalSupply := new(big.Int)
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&totalSupply, "totalSupply", data)
	if err != nil {
		return nil, err
	}

	return totalSupply, nil
}

func (n *NodeClient) ERC20Decimals(ctx context.Context, address common.Address) (uint8, error) {
	_calldata, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("decimals")
	if err != nil {
		return 0, err
	}

	data, err := n.Call(ctx, CallParam{
		To:       address,
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return 0, err
	}

	var decimals uint8
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&decimals, "decimals", data)
	if err != nil {
		return 0, err
	}

	return decimals, nil
}

func (n *NodeClient) ERC20Balance(ctx context.Context, tokenAddress, accountAddress common.Address) (*big.Int, error) {
	_calldata, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("balanceOf", accountAddress)
	if err != nil {
		return nil, err
	}

	data, err := n.Call(ctx, CallParam{
		To:       tokenAddress,
		CallData: hexutil.Encode(_calldata),
	})
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&balance, "balanceOf", data)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (n *NodeClient) IsERC20(ctx context.Context, address common.Address) (bool, error) {
	c := n.Client()
	if c == nil {
		return false, fmt.Errorf("no client available")
	}

	code, err := c.CodeAt(ctx, address, nil)
	if err != nil {
		log.Error("failed to get code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	methodIDs, err := indexerCommon.ParseMethodIDFromCode(code)
	if err != nil {
		log.Error("failed to parse method id from code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	eventHashes, err := indexerCommon.ParseEventHashFromCode(code)
	if err != nil {
		log.Error("failed to parse event hash from code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	erc20MustMethods := constant.ERC20MustMethodIDs()
	erc20MustEvents := constant.ERC20MustEventIDs()

	for methodId := range erc20MustMethods {
		// method ID not in code
		if !lo.Contains(methodIDs, strings.ToLower(methodId)) {
			return false, nil
		}
	}

	for event := range erc20MustEvents {
		if !lo.Contains(eventHashes, strings.ToLower(event)) {
			return false, nil
		}
	}

	return true, nil
}

func (n *NodeClient) ERC20Meta(ctx context.Context, address common.Address) (*model.ERC20Token, error) {
	decimalsCallData, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("decimals")
	if err != nil {
		log.Error("failed to pack erc20 decimals calldata", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	symbolCallData, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("symbol")
	if err != nil {
		log.Error("failed to pack erc20 symbol calldata", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	nameCallData, err := contracts.ABIs[contracts.ContractTypeERC20].Pack("name")
	if err != nil {
		log.Error("failed to pack erc20 name calldata", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	calls := []multicall.Multicall3Call3{
		{
			Target:       address,
			AllowFailure: false,
			CallData:     decimalsCallData,
		},
		{
			Target:       address,
			AllowFailure: false,
			CallData:     symbolCallData,
		},
		{
			Target:       address,
			AllowFailure: false,
			CallData:     nameCallData,
		},
	}

	calldata, err := contracts.ABIs[contracts.ContractTypeMulticall].Pack("aggregate3", calls)
	if err != nil {
		log.Error("failed to pack multicall aggregate3 calldata", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	resultBytes, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   lo.ToPtr(multicall.Multicall3Address),
		Data: calldata,
	}, nil)
	if err != nil {
		log.Error("failed to call multicall aggregate3", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	var results []multicall.Multicall3Result
	err = contracts.ABIs[contracts.ContractTypeMulticall].UnpackIntoInterface(&results, "aggregate3", resultBytes)
	if err != nil {
		log.Error("failed to unpack multicall aggregate3 result", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	if len(results) != 3 {
		log.Error("failed to unpack multicall aggregate3 result", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	for i, res := range results {
		if !res.Success {
			err = fmt.Errorf("failed to get %d th call", i)
			log.Error("failed to get ith call", zap.Int("index", i), zap.Error(err))
			return nil, err
		}
	}

	var res = new(model.ERC20Token)

	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&res.Decimals, "decimals", results[0].ReturnData)
	if err != nil {
		log.Error("failed to unpack erc20 decimals result", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&res.Symbol, "symbol", results[1].ReturnData)
	if err != nil {
		log.Error("failed to unpack erc20 symbol result", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	err = contracts.ABIs[contracts.ContractTypeERC20].UnpackIntoInterface(&res.Name, "name", results[2].ReturnData)
	if err != nil {
		log.Error("failed to unpack erc20 name result", zap.String("address", address.String()), zap.Error(err))
		return nil, err
	}

	return res, nil
}
