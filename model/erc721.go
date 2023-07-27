package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ERC721Transfer struct {
	BlockHash       common.Hash    `gorm:"column:block_hash" ch:"block_hash" json:"block_hash,omitempty"`
	BlockNumber     uint64         `gorm:"column:block_number" ch:"block_number" json:"block_number,omitempty"`
	TxHash          common.Hash    `gorm:"column:tx_hash" ch:"tx_hash" json:"tx_hash,omitempty"`
	ContractAddress common.Address `gorm:"column:contract_address" ch:"contract_address" json:"contract_address,omitempty"`
	From            common.Address `gorm:"column:from" ch:"from" json:"from,omitempty"`
	To              common.Address `gorm:"column:to" ch:"to" json:"to,omitempty"`
	TokenID         *big.Int       `gorm:"column:token_id" ch:"token_id" json:"token_id,omitempty"`
}

type ERC721Token struct {
	ContractAddress common.Address `gorm:"column:contract_address" ch:"contract_address" json:"contract_address,omitempty"`
	Name            string         `gorm:"column:name" ch:"name" json:"name,omitempty"`
	Symbol          string         `gorm:"column:symbol" ch:"symbol" json:"symbol,omitempty"`
	URI             string         `gorm:"column:uri" ch:"uri" json:"uri,omitempty"`
}
