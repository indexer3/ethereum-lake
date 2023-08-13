package constant

import "github.com/ethereum/go-ethereum/common"

const (
	ethereumAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	zeroAddress     = "0x0000000000000000000000000000000000000000"
)

func EthereumAddress() common.Address {
	return common.HexToAddress(ethereumAddress)
}

func ZeroAddress() common.Address {
	return common.HexToAddress(zeroAddress)
}
