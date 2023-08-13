package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v2_factory.abi --pkg uniswap --type UniswapV2Factory --out ./uniswap_v2_factory.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v2_pair.abi --pkg uniswap --type UniswapV2Pair --out ./uniswap_v2_pair.go

var (
	UniswapV2FactoryAddress = common.HexToAddress("0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f")
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_factory.abi --pkg uniswap --type UniswapV3Factory --out ./uniswap_v3_factory.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_pool.abi --pkg uniswap --type UniswapV3Pool --out ./uniswap_v3_pool.go
var (
	Uniswapv3FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_nft_position_manager.abi --pkg uniswap --type UniswapV3NFTPositionManager --out ./uniswap_v3_nft_position_manager.go
var (
	UniswapV3NFTPositionManagerAddress = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	Uniswapv3NFTDesciptorAddress       = common.HexToAddress("0x42B24A95702b9986e82d421cC3568932790A48Ec")
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./uniswap_v3_router.abi --pkg uniswap --type UniswapV3Router --out ./uniswap_v3_router.go
var (
	Uniswapv3SwapRouterAddress   = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	Uniswapv3SwapRouter02Address = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
)
