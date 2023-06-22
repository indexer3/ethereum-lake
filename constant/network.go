package constant

//go:generate go run -mod=mod github.com/dmarkham/enumer -type=Network -transform=snake -trimprefix=Network -json
type Network int64

const (
	NetworkEthereum        Network = 1
	NetworkOptimism        Network = 10
	NetworkBSC             Network = 56
	NetworkEthereumClassic Network = 61
	NetworkPolygon         Network = 137
	NetworkArbitrumOne     Network = 42161
	NetworkArbitrumNova    Network = 42170
)

func ChainID(network Network) uint64 {
	return uint64(network)
}
