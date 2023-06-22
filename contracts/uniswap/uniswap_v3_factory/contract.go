package uniswap_v3_factory

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_factory.abi --pkg uniswap_v3_factory --type UniswapV3Factory --out ./uniswap_v3_factory.go

var (
	UniswapV3FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
)
