package indexer

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/client"
)

type CommonProcess struct {
	nodeClient *client.NodeClient
}

func NewCommonProcess(nodeClient *client.NodeClient) *CommonProcess {
	return &CommonProcess{
		nodeClient: nodeClient,
	}
}

func (c *CommonProcess) GetLatestBlock(ctx context.Context) (uint64, error) {
	return c.nodeClient.Client().BlockNumber(ctx)
}
