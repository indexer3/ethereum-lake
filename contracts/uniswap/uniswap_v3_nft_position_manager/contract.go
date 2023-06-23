package uniswap_v3_nft_position_manager

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_nft_position_manager.abi --pkg uniswap_v3_nft_position_manager --type UniswapV3NFTPositionManager --out ./uniswap_v3_nft_position_manager.go
var (
	UniswapV3NFTPositionManagerAddress = common.HexToAddress("")
)
