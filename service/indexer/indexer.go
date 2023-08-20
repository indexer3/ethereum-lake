package indexer

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/client"
)

type CommonEthereumProcess struct {
	nodeClient *client.NodeClient
}

func NewCommonProcess(nodeClient *client.NodeClient) *CommonEthereumProcess {
	return &CommonEthereumProcess{
		nodeClient: nodeClient,
	}
}

func (c *CommonEthereumProcess) GetLatestBlock(ctx context.Context) (uint64, error) {
	return c.nodeClient.Client().BlockNumber(ctx)
}
