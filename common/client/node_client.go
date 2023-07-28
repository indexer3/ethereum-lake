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

func (n *NodeClient) Client() *ethclient.Client {
	if len(n.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return n.ethClients[rand.Intn(len(n.ethClients))]
}
