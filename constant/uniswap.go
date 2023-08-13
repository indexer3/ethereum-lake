package constant

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func UniswapPostitionTransferEvent() common.Hash {
	return crypto.Keccak256Hash([]byte(erc721TransferEventSig))
}
