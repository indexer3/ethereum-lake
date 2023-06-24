package one_inch_oracle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./one_inch_oracle.abi --pkg one_inch_oracle --type OneInchOracle --out ./one_inch_oracle.go
var (
	oneInchPriceOralceAddressEthereum = common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
	oneInchPriceOralceAddressPolygon  = common.HexToAddress("0x7F069df72b7A39bCE9806e3AfaF579E54D8CF2b9")
	oneInchPriceOracleAddressArbitrum = common.HexToAddress("0x735247fb0a604c0adC6cab38ACE16D0DbA31295F")
	oneInchPriceOracleAddressOptimism = common.HexToAddress("0x11DEE30E710B8d4a8630392781Cc3c0046365d4c")
	oneInchPriceOracleAddressBSC      = common.HexToAddress("0xfbD61B037C325b959c0F6A7e69D8f37770C2c550")

	usdtAddressInEthereum = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	usdtAddressInPolygon  = common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	usdtAddressInArbitrum = common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9")
	usdtAddressInOptimism = common.HexToAddress("0x94b008aa00579c1307b0ef2c499ad98a8ce58e58")
	busdAddressInBSC      = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
)

func OneInchPriceOracle(network constant.Network) common.Address {
	switch network {
	case constant.NetworkEthereum:
		return oneInchPriceOralceAddressEthereum
	case constant.NetworkPolygon:
		return oneInchPriceOralceAddressPolygon
	case constant.NetworkArbitrumOne:
		return oneInchPriceOracleAddressArbitrum
	case constant.NetworkBSC:
		return oneInchPriceOracleAddressBSC
	case constant.NetworkOptimism:
		return oneInchPriceOracleAddressOptimism
	default:
		return oneInchPriceOralceAddressEthereum
	}
}

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
