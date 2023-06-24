package client

import (
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/indexer3/ethereum-lake/contracts"
)

type NodeClient struct {
	ethClients   []*wrapClient
	limitChan    chan struct{}
	mu           sync.RWMutex
	contractAbis map[contracts.ContractType]*abi.ABI
}

type wrapClient struct {
	client    *rpc.Client
	ethclient *ethclient.Client
}

func (w *wrapClient) RPCClient() *rpc.Client {
	return w.client
}

func (w *wrapClient) ETHClient() *ethclient.Client {
	return w.ethclient
}

func NewNodeClientsWithEndpoints(endpoints []string) (*NodeClient, error) {
	ethClients := make([]*wrapClient, 0)
	for _, endpoint := range endpoints {
		callClient, err := rpc.Dial(endpoint)
		if err != nil {
			return nil, err
		}

		ethClient, err := ethclient.Dial(endpoint)
		if err != nil {
			return nil, err
		}
		ethClients = append(ethClients, &wrapClient{
			client:    callClient,
			ethclient: ethClient,
		})
	}
	return &NodeClient{
		ethClients:   ethClients,
		limitChan:    make(chan struct{}, 10),
		mu:           sync.RWMutex{},
		contractAbis: make(map[contracts.ContractType]*abi.ABI),
	}, nil
}

func (n *NodeClient) RegisterContractABI(contractType contracts.ContractType, abi *abi.ABI) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.contractAbis[contractType] = abi
}

func (c *NodeClient) Client() *wrapClient {
	if len(c.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return c.ethClients[rand.Intn(len(c.ethClients))]
}
