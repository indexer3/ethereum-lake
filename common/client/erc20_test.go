package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestOnERC20(t *testing.T) {
	cli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	require.NoError(t, err)

	var (
		ctx           = context.Background()
		usdtContract  = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
		punksContract = common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB")
	)

	t.Run("test on IsERC20", func(t *testing.T) {
		res, err := cli.IsERC20(ctx, usdtContract)
		require.NoError(t, err)

		require.True(t, res)

		res, err = cli.IsERC20(ctx, punksContract)
		require.NoError(t, err)

		require.False(t, res)
	})

	t.Run("test on erc20 metadata", func(t *testing.T) {
		symbol, err := cli.ERC20Symbol(ctx, usdtContract)
		require.Nil(t, err)

		require.Equal(t, "USDT", symbol)

		decimals, err := cli.ERC20Decimals(ctx, usdtContract)
		require.Nil(t, err)

		require.Equal(t, uint8(6), decimals)

		totalSupply, err := cli.ERC20TotalSupply(ctx, usdtContract)
		require.Nil(t, err)

		fmt.Println(totalSupply)

		meta, err := cli.ERC20Meta(ctx, usdtContract)
		require.Nil(t, err)

		require.Equal(t, "USDT", meta.Symbol)
		require.Equal(t, uint8(6), meta.Decimals)
		require.Equal(t, "Tether USD", meta.Name)
	})

	t.Run("test on erc20 balance", func(t *testing.T) {
	})

}
