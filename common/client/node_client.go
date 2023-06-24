package client

import (
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type NodeClient struct {
	ethClients   []*wrapClient
	limitChan    chan struct{}
	mu           sync.RWMutex
	contractAbis map[common.Address]*abi.ABI
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
		contractAbis: make(map[common.Address]*abi.ABI),
	}, nil
}

func (n *NodeClient) RegisterContractABI(contractAddress common.Address, abi *abi.ABI) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.contractAbis[contractAddress] = abi
}

func (c *NodeClient) Client() *wrapClient {
	if len(c.ethClients) == 0 {
		return nil
	}
	rand.NewSource(time.Now().Unix())
	return c.ethClients[rand.Intn(len(c.ethClients))]
}
