package relayer

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/stretchr/testify/require"

	"github.com/spf13/viper"
)

func TestOnFetchFeed(t *testing.T) {
	viper.Set(config.ChainbaseSQLAPIKey, "demo")
	feedRelayer := NewFeedRelayer([]string{"https://rpc.ankr.com/eth"})
	ctx := context.Background()

	feeds, err := feedRelayer.Feeds(ctx, constant.NetworkEthereum,
		common.HexToAddress("0xc101c69340feb4d0c474bf8fc34f5266f3de8a15"), nil, 10)
	require.NoError(t, err)

	t.Log(feeds)
}
