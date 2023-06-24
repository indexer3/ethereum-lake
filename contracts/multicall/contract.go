package multicall

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./multicall.abi --pkg multicall --type Multicall --out ./multicall.go
var (
	Multicall3Address = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
)
