package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/stretchr/testify/require"
)

func TestOnAaveV2(t *testing.T) {
	ctx := context.Background()

	polygonCli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/polygon"})
	require.NoError(t, err)

	ethereumCli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	require.NoError(t, err)

	curveAddress := common.HexToAddress("0x445fe580ef8d70ff569ab36e80c647af338db351")
	morphoAddress := common.HexToAddress("0x777777c9898d384f785ee44acfe945efdff5f3e0")
	_ = morphoAddress

	someoneAddress := common.HexToAddress("0x2728FB60405dF11B57407a9e78E48FA1cC98607F")
	someAddress2 := common.HexToAddress("0xa976ea51b9ba3232706af125a92e32788dc08ddc")

	t.Run("test on unclaimed rewards aave v2", func(t *testing.T) {
		unclaimed, err := polygonCli.AaveV2UnClaimedRewards(ctx, constant.NetworkPolygon, curveAddress, nil)
		require.NoError(t, err)
		fmt.Println(unclaimed)
	})

	t.Run("test on staked rewards aave v2", func(t *testing.T) {
		staked, err := ethereumCli.AaveV2StakedReward(ctx, constant.NetworkEthereum, someoneAddress, nil)
		require.NoError(t, err)
		fmt.Println(staked)
	})

	t.Run("test on health factor aave v2", func(t *testing.T) {
		healthFactor, err := ethereumCli.AaveV2AccountHealthRate(ctx, constant.NetworkEthereum, someAddress2, nil)
		require.NoError(t, err)
		fmt.Println(healthFactor)
	})

	t.Run("test on user account data aave v2", func(t *testing.T) {
		accountData, err := ethereumCli.AaveV2AccountData(ctx, constant.NetworkEthereum, someAddress2, nil)
		require.NoError(t, err)
		fmt.Printf("%+v", accountData)
	})

	t.Run("test on user account debt aave v2 usdt", func(t *testing.T) {
		accountData, err := ethereumCli.AaveV2AccountDebt(ctx, constant.NetworkEthereum, someAddress2,
			constant.USDContractAddress(constant.NetworkEthereum), nil)
		require.NoError(t, err)
		fmt.Printf("%+v", accountData)
	})
}
