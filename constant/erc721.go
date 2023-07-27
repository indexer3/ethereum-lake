package constant

import (
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// ERC-721 Standard
// https://eips.ethereum.org/EIPS/eip-721

// Methods
const (
	erc721BalanceOfMethodSig                 = "balanceOf(address)"
	erc721OwnerOfMethodSig                   = "ownerOf(uint256)"
	erc721SafeTransferFromMethodSig          = "safeTransferFrom(address,address,uint256)"
	erc721SafeTransferFromWithBytesMethodSig = "safeTransferFrom(address,address,uint256,bytes)"
	erc721TransferFromMethodSig              = "transferFrom(address,address,uint256)"
	erc721ApproveMethodSig                   = "approve(address,uint256)"
	erc721GetApprovedMethodSig               = "getApproved(uint256)"
	erc721SetApprovalMethodSig               = "setApprovalForAll(address,bool)"
	erc721IsApprovedMethodSig                = "isApprovedForAll(address,address)"
)

// Events
const (
	erc721TransferEventSig       = "Transfer(address,address,uint256)"
	erc721ApprovalEventSig       = "Approval(address,address,uint256)"
	erc721ApprovalForAllEventSig = "ApprovalForAll(address,address,bool)"
)

var erc721MethodSigMust = [][]byte{
	crypto.Keccak256([]byte(erc721BalanceOfMethodSig)),
	crypto.Keccak256([]byte(erc721OwnerOfMethodSig)),
	crypto.Keccak256([]byte(erc721SafeTransferFromMethodSig)),
	crypto.Keccak256([]byte(erc721SafeTransferFromWithBytesMethodSig)),
	crypto.Keccak256([]byte(erc721TransferFromMethodSig)),
	crypto.Keccak256([]byte(erc721ApproveMethodSig)),
	crypto.Keccak256([]byte(erc721GetApprovedMethodSig)),
	crypto.Keccak256([]byte(erc721SetApprovalMethodSig)),
	crypto.Keccak256([]byte(erc721IsApprovedMethodSig)),
}

var erc721MethodSigMustString = []string{
	crypto.Keccak256Hash([]byte(erc721BalanceOfMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721OwnerOfMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721SafeTransferFromMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721SafeTransferFromWithBytesMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721TransferFromMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721ApproveMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721GetApprovedMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721SetApprovalMethodSig)).String()[2:10],
	crypto.Keccak256Hash([]byte(erc721IsApprovedMethodSig)).String()[2:10],
}

var erc721EventSigMustString = []string{
	crypto.Keccak256Hash([]byte(erc721TransferEventSig)).String()[2:],
	crypto.Keccak256Hash([]byte(erc721ApprovalEventSig)).String()[2:],
	crypto.Keccak256Hash([]byte(erc721ApprovalForAllEventSig)).String()[2:],
}

func ERC721MustEventIDs() map[string]struct{} {
	seen := make(map[string]struct{})
	for i := 0; i < len(erc721EventSigMustString); i++ {
		seen[strings.ToLower(erc721EventSigMustString[i])] = struct{}{}
	}
	return seen
}

func ERC721MustMethodIDs() map[string]struct{} {
	seen := make(map[string]struct{})
	for i := 0; i < len(erc721MethodSigMustString); i++ {
		seen[strings.ToLower(erc721MethodSigMustString[i])] = struct{}{}
	}
	return seen
}
