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
	ethereumCli, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	require.NoError(t, err)

	curveAddress := common.HexToAddress("0x445fe580ef8d70ff569ab36e80c647af338db351")
	morphoAddress := common.HexToAddress("0x777777c9898d384f785ee44acfe945efdff5f3e0")
	_ = morphoAddress

	someoneAddress := common.HexToAddress("0x2728FB60405dF11B57407a9e78E48FA1cC98607F")

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

}
