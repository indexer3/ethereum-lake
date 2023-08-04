package client

import "github.com/indexer3/ethereum-lake/constant"

type NetworkSupport interface {
	SupportNetwork(network constant.Network) bool
	SupportNetworks() []constant.Network
}
