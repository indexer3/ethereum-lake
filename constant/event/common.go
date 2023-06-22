package event

import "github.com/ethereum/go-ethereum/crypto"

var (
	TranferEvent = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)
