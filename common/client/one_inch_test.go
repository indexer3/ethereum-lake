package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	lakecommon "github.com/indexer3/ethereum-lake/common"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/stretchr/testify/require"
)

func TestOnPriceOracle(t *testing.T) {
	ethCli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	require.NoError(t, err)

	wethInETHContractAddress := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	uniswapInETHTokenContractAddress := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")

	wethInPolygonContractAddress := common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619")
	uniswapInPolygonTokenContractAddress := common.HexToAddress("0xb33eaad8d922b1083446dc23f610c2567fb5180f")

	wethInBSCContractAddress := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	uniswapInBSCTokenContractAddress := common.HexToAddress("0xBf5140A22578168FD562DCcF235E5D43A02ce9B1")

	ctx := context.Background()

	polygonCli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/polygon"})
	require.NoError(t, err)

	bscCli, err := NewNodeClientsWithEndpoints([]string{"https://bsc-dataseed1.binance.org/"})
	require.NoError(t, err)

	wethPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, wethInETHContractAddress, nil)
	require.NoError(t, err)

	fmt.Println("ethereum network: WETH price:", wethPrice)

	uniPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, uniswapInETHTokenContractAddress, nil)
	require.NoError(t, err)

	fmt.Println("ethereum network: UNI price:", uniPrice)

	t.Run("test on polygon token price", func(t *testing.T) {
		polygonWethPrice, err := polygonCli.TokenPrice(ctx, constant.NetworkPolygon, wethInPolygonContractAddress, nil)
		require.NoError(t, err)

		fmt.Println("polygon network: WETH price:", polygonWethPrice, " rate ", lakecommon.GetDiffPercent(wethPrice, polygonWethPrice))

		polygonUniPrice, err := polygonCli.TokenPrice(ctx, constant.NetworkPolygon, uniswapInPolygonTokenContractAddress, nil)
		require.NoError(t, err)

		fmt.Println("polygon network: UNI price:", polygonUniPrice, " rate ", lakecommon.GetDiffPercent(uniPrice, polygonUniPrice))
	})

	t.Run("test on bsc token price", func(t *testing.T) {
		wbnbPrice, err := bscCli.TokenPrice(ctx, constant.NetworkBSC, wethInBSCContractAddress, nil)
		require.NoError(t, err)

		fmt.Println("bsc network: WBNB price:", wbnbPrice)

		bnbUniPrice, err := bscCli.TokenPrice(ctx, constant.NetworkBSC, uniswapInBSCTokenContractAddress, nil)
		require.NoError(t, err)

		fmt.Println("bsc network: UNI price:", bnbUniPrice, " rate ", lakecommon.GetDiffPercent(uniPrice, bnbUniPrice))
	})
}
