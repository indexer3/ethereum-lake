package constant

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	EthereumDecimals = 18
)

const (
	USDTDecimalsInEthereum = 6
	USDTDecimalsInArbitrum = 6
	USDTDecimalsInOptimism = 6
	USDTDecimalsInPolygon  = 6
	BUSDDecimalsInBSC      = 18
)

const (
	USDTAddressInEthereum = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	USDTAddressInPolygon  = "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"
	USDTAddressInArbitrum = "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"
	USDTAddressInOptimism = "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58"
	BUSDAddressInBSC      = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
)

func USDContractAddress(network Network) common.Address {
	switch network {
	case NetworkEthereum:
		return common.HexToAddress(USDTAddressInEthereum)
	case NetworkPolygon:
		return common.HexToAddress(USDTAddressInPolygon)
	case NetworkArbitrumOne:
		return common.HexToAddress(USDTAddressInArbitrum)
	case NetworkOptimism:
		return common.HexToAddress(USDTAddressInOptimism)
	case NetworkBSC:
		return common.HexToAddress(BUSDAddressInBSC)
	default:
		return common.HexToAddress(USDTAddressInEthereum)
	}
}

func USDDecimals(network Network) uint8 {
	switch network {
	case NetworkEthereum:
		return USDTDecimalsInEthereum
	case NetworkPolygon:
		return USDTDecimalsInPolygon
	case NetworkArbitrumOne:
		return USDTDecimalsInArbitrum
	case NetworkOptimism:
		return USDTDecimalsInOptimism
	case NetworkBSC:
		return BUSDDecimalsInBSC
	default:
		return USDTDecimalsInEthereum
	}
}
