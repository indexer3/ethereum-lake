package seaport

import "github.com/ethereum/go-ethereum/common"

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./seaport.abi --pkg seaport --type Seaport --out ./seaport.go
var (
	SeaportContractAddress = common.HexToAddress("0x00000000006c3852cbef3e08e8df289169ede581")
)
