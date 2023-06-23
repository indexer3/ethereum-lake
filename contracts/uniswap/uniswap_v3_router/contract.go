package uniswap_v3_router

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_router.abi --pkg uniswap_v3_router --type UniswapV3Router --out ./uniswap_v3_router.go
var (
	UniswapV3RouterAddress = common.HexToAddress("")
)
