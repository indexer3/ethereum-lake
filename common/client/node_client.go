package client

import (
	"context"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samber/lo"
)

type NodeClient struct {
	ethClients []*ethclient.Client
	limitChan  chan struct{}
}

func NewNodeClientsWithEndpoints(endpoints []string) (*NodeClient, error) {
	ethClients := make([]*ethclient.Client, 0)
	for _, endpoint := range endpoints {
		ethClient, err := ethclient.Dial(endpoint)
		if err != nil {
			return nil, err
		}
		ethClients = append(ethClients, ethClient)
	}

	return &NodeClient{
		ethClients: ethClients,
		limitChan:  make(chan struct{}, 10),
	}, nil
}

func (n *NodeClient) Client() *ethclient.Client {
	if len(n.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return n.ethClients[rand.Intn(len(n.ethClients))]
}

// Latest Client returns the client with the latest block number.
// It only works on specific node, like 'http://192.xxx.xxx.xx:8045'.
// It will NOT work when connected to a gateway like 'https://ankr.com/eth',
// because the request itself will be sent to one of the nodes randomly.
func (n *NodeClient) LatestClient() *ethclient.Client {
	if len(n.ethClients) == 0 {
		return nil
	}
	var (
		latestBlockNumber uint64
		latestClient      *ethclient.Client
	)

	lo.ForEach[*ethclient.Client](n.ethClients, func(cli *ethclient.Client, _ int) {
		blockNumber, err := cli.BlockNumber(context.Background())
		if err != nil {
			return
		}

		if blockNumber > latestBlockNumber {
			latestBlockNumber = blockNumber
			latestClient = cli
		}
	})

	return latestClient
}
