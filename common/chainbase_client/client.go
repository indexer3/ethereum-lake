package chainbase_client

import (
	"github.com/go-resty/resty/v2"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
)

type ChainbaseCli struct {
	cli *resty.Client
}

const DefaultTransactionLimit = 20

const (
	OkCode = 0
)

func NewChainbaseClient() *ChainbaseCli {
	return &ChainbaseCli{
		cli: resty.New().SetBaseURL("https://api.chainbase.online").
			SetHeader("x-api-key", viper.GetString(config.ChainbaseSQLAPIKey)),
	}
}
