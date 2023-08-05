package chainbase_client

import (
	"github.com/go-resty/resty/v2"
	"github.com/indexer3/ethereum-lake/model"
)

type ChainbaseCli struct {
	c *resty.Client
}

func NewChainbaseClient() *ChainbaseCli {
	return &ChainbaseCli{
		c: resty.New().SetBaseURL("https://api.chainbase.io/v1"),
	}
}

func (c *ChainbaseCli) GetTransactionFeeds() (*model.ChainbaseSQLTransaction, error) {
	
	return nil, nil
}
