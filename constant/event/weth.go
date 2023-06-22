package event

import "github.com/ethereum/go-ethereum/crypto"

var (
	WithdrawalEvent = crypto.Keccak256Hash([]byte("Withdrawal(address,uint256)"))
	DepositEvent    = crypto.Keccak256Hash([]byte("Deposit(address,uint256)"))
)
