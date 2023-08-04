package chainbase_client

import (
	"github.com/go-resty/resty/v2"
)

type ChainbaseCli struct {
	c *resty.Client
}

func NewChainbaseClient() *ChainbaseCli {
	return &ChainbaseCli{
		c: resty.New().SetBaseURL("https://api.chainbase.io/v1"),
	}
}
