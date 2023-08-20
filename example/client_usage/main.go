package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/client"
	"github.com/indexer3/ethereum-lake/constant"
)

func main() {
	ethCli, err := client.NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	if err != nil {
		panic(err)
	}

	wethInETHContractAddress := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	uniswapInETHTokenContractAddress := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")

	ctx := context.Background()

	wethPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, wethInETHContractAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("ethereum network: WETH price:", wethPrice)

	uniPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, uniswapInETHTokenContractAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("ethereum network: UNI price:", uniPrice)
}
