package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestOnGMX(t *testing.T) {
	arbClient, err := NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/arbitrum"})
	require.Nil(t, err)
	accountAddress1 := common.HexToAddress("0xa329ac2efffea563159897d7828866cfaed42167")
	ctx := context.Background()

	t.Run("test on account address1", func(t *testing.T) {
		rewards, err := arbClient.StakeEthReward(ctx, accountAddress1, nil)
		require.Nil(t, err)

		glpRewards, err := arbClient.StakeGlpRewardEth(ctx, accountAddress1, nil)
		require.Nil(t, err)

		fmt.Println(rewards)
		fmt.Println(glpRewards)
	})
}
