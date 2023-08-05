package chainbase_client

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/spf13/viper"
)

type ChainbaseCli struct {
	cli *resty.Client
}

func NewChainbaseClient() *ChainbaseCli {
	return &ChainbaseCli{
		cli: resty.New().SetBaseURL("https://api.chainbase.io/v1").
			SetHeader("x-api-key", viper.GetString(config.ChainbaseSQLAPIKey)),
	}
}

func (c *ChainbaseCli) GetTransactionFeeds(ctx context.Context) (*model.ChainbaseSQLTransaction, error) {
	var resp model.ChainbaseSQLTransaction
	_, err := c.cli.R().SetContext(ctx).SetResult(&resp).Post("/v1/dw/query")
	if err != nil {
		log.Error("failed to ")
	}
	return &resp, nil
}
