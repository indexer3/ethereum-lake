package constant

import "github.com/ethereum/go-ethereum/crypto"

// ERC-1155 Standard
// https://eips.ethereum.org/EIPS/eip-1155
var (
	erc1155SafeTransferFromSig      = crypto.Keccak256([]byte("safeTransferFrom(address,address,uint256,uint256,bytes)"))
	erc1155SafeBatchTransferFromSig = crypto.Keccak256([]byte("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)"))
	erc1155BalanceOfSig             = crypto.Keccak256([]byte("balanceOf(address,uint256)"))
	erc1155BalanceOfBatchSig        = crypto.Keccak256([]byte("balanceOfBatch(address[],uint256[])"))
	erc1155SetApprovalForAllSig     = crypto.Keccak256([]byte("setApprovalForAll(address,bool)"))
	erc1155IsApprovedForAllSig      = crypto.Keccak256([]byte("isApprovedForAll(address,address)"))
)

var erc1155SigMust = [][]byte{
	erc1155SafeTransferFromSig,
	erc1155SafeBatchTransferFromSig,
	erc1155BalanceOfSig,
	erc1155BalanceOfBatchSig,
	erc1155SetApprovalForAllSig,
	erc1155IsApprovedForAllSig,
}
