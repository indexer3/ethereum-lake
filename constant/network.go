package constant

//go:generate go run -mod=mod github.com/dmarkham/enumer -type=Network -transform=snake -trimprefix=Network -json
type Network int64

const (
	NetworkEthereum        Network = 1
	NetworkGoerli          Network = 5
	NetworkOptimism        Network = 10
	NetworkBSC             Network = 56
	NetworkEthereumClassic Network = 61
	NetworkPolygon         Network = 137
	NetworkBase            Network = 8453
	NetworkArbitrumOne     Network = 42161
	NetworkArbitrumNova    Network = 42170
	NetworkBaseGoerli      Network = 84531
)

func ChainID(network Network) uint64 {
	return uint64(network)
}
