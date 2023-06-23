package weth

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./weth.abi --pkg weth --type WETH --out ./weth.go
var (
	WETHAddress = common.HexToAddress("")
)
