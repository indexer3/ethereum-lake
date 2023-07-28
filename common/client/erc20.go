package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/multicall"
	"github.com/indexer3/ethereum-lake/model"
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
