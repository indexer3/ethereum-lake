package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/indexer3/ethereum-lake/contracts"
	"github.com/indexer3/ethereum-lake/contracts/multicall"

	"github.com/samber/lo"
)

type CallParam struct {
	To          common.Address `json:"to"`
	CallData    string         `json:"data"`
	BlockNumber *big.Int       `json:"blockNumber"`
}

func (n *NodeClient) AggregatedCalls(ctx context.Context, calls []multicall.Multicall3Call3) ([]multicall.Multicall3Result, error) {
	cli := n.Client()
	if cli == nil {
		return nil, fmt.Errorf("no client available")
	}

	n.mu.RLock()
	defer n.mu.Unlock()
	calldata, err := n.contractAbis[contracts.ContractTypeMulticall].Pack("aggregate3", calls)
	if err != nil {
		return nil, err
	}

	callMsg := ethereum.CallMsg{
		To:   lo.ToPtr(multicall.Multicall3Address),
		Data: calldata,
	}

	res, err := cli.CallContract(ctx, callMsg, nil)
	if err != nil {
		return nil, err
	}

	var results []multicall.Multicall3Result
	if err := n.contractAbis[contracts.ContractTypeMulticall].UnpackIntoInterface(&results, "aggregate3", res); err != nil {
		return nil, err
	}

	return results, nil
}

func (n *NodeClient) Call(ctx context.Context, callParam CallParam) ([]byte, error) {
	cli := n.Client()
	if cli == nil {
		return nil, fmt.Errorf("no available client")
	}

	var (
		_res        string
		blockNumber string
	)

	blockNumber = lo.If[string](callParam.BlockNumber == nil, "latest").Else(hexutil.EncodeBig(callParam.BlockNumber))

	err := cli.Client().CallContext(ctx, &_res, "eth_call", callParam, blockNumber)
	if err != nil {
		return nil, err
	}

	decodeRes, err := hexutil.Decode(_res)
	if err != nil {
		return nil, err
	}

	return decodeRes, nil
}
