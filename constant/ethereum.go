package constant

import "github.com/ethereum/go-ethereum/common"

const (
	ethereumAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

func EthereumAddress() common.Address {
	return common.HexToAddress(ethereumAddress)
}
