package client

import (
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
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

func (c *NodeClient) Client() *ethclient.Client {
	if len(c.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return c.ethClients[rand.Intn(len(c.ethClients))]
}
