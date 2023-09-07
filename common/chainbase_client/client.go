package chainbase_client

import (
	"github.com/go-resty/resty/v2"
	"github.com/indexer3/ethereum-lake/common/config"
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
			SetHeader("x-api-key", config.RelayerConf.ChainbaseAPIKey),
	}
}
