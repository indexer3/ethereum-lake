package usd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

var (
	usdtAddressInEthereum = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	usdtAddressInPolygon  = common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	usdtAddressInArbitrum = common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9")
	usdtAddressInOptimism = common.HexToAddress("0x94b008aa00579c1307b0ef2c499ad98a8ce58e58")
	busdAddressInBSC      = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
)

func USDAddress(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return usdtAddressInEthereum
	case constant.NetworkPolygon:
		return usdtAddressInPolygon
	case constant.NetworkArbitrumOne:
		return usdtAddressInArbitrum
	case constant.NetworkOptimism:
		return usdtAddressInOptimism
	case constant.NetworkBSC:
		return busdAddressInBSC

	default:
		return usdtAddressInEthereum
	}
}
