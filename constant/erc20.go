package constant

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ERC-20 Standard
// https://eips.ethereum.org/EIPS/eip-20

// Methods
const (
	erc20TransferMethod = "transfer(address,uint256)"
	erc20ApproveMethod  = "approve(address,uint256)"

	erc20BalanceOfMethod    = "balanceOf(address)"
	erc20AllowanceMethod    = "allowance(address,address)"
	erc20TotalSupplyMethod  = "totalSupply()"
	erc20TransderFromMethod = "transferFrom(address,address,uint256)"
)

// Events

const (
	erc20TransferEvent = "Transfer(address,address,uint256)"
	erc20ApprovalEvent = "Approval(address,address,uint256)"
)

var erc20MethodMust = []string{
	crypto.Keccak256Hash([]byte(erc20TransferMethod)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc20ApproveMethod)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc20BalanceOfMethod)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc20AllowanceMethod)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc20TotalSupplyMethod)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc20TransderFromMethod)).String()[2:10],
}

var erc20EventMust = []string{
	crypto.Keccak256Hash([]byte(erc20ApprovalEvent)).String()[2:],
	crypto.Keccak256Hash([]byte(erc20TransferEvent)).String()[2:],
}

func ERC20MustEventIDs() map[string]struct{} {
	seen := make(map[string]struct{})
	for i := 0; i < len(erc20EventMust); i++ {
		seen[strings.ToLower(erc20EventMust[i])] = struct{}{}
	}
	return seen
}

func ERC20MustMethodIDs() map[string]struct{} {
	seen := make(map[string]struct{})
	for i := 0; i < len(erc20MethodMust); i++ {
		seen[strings.ToLower(erc20MethodMust[i])] = struct{}{}
	}
	return seen
}

func ERC20TransferMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20TransferMethod))
}

func ERC20ApproveMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20ApproveMethod))
}

func ERC20BalanceOfMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20BalanceOfMethod))
}

func ERC20AllowanceMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20AllowanceMethod))
}

func ERC20TotalSupplyMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20TotalSupplyMethod))
}

func ERC20TransderFromMethod() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20TransderFromMethod))
}

func ERC20ApprovalEvent() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20ApprovalEvent))
}

func ERC20TransferEvent() common.Hash {
	return crypto.Keccak256Hash([]byte(erc20TransferEvent))
}
