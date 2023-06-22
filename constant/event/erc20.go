package event

import "github.com/ethereum/go-ethereum/crypto"

var (
	ApprovalEvent = crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))
)
