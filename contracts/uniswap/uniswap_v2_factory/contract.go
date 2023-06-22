package uniswap_v2_factory

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v2_factory.abi --pkg uniswap_v2_factory --type UniswapV2Factory --out ./uniswap_v2_factory.go

var (
	UniswapV2FactoryAddress = common.HexToAddress("0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f")
)
