package event

import "github.com/ethereum/go-ethereum/crypto"

var (
	TransferEvent = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)
