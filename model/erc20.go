package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ERC20Token struct {
	ContractAddress common.Address `gorm:"column:contract_address" ch:"contract_address" json:"contract_address,omitempty"`
	Name            string         `gorm:"column:name" ch:"name" json:"name,omitempty"`
	Symbol          string         `gorm:"column:symbol" ch:"symbol" json:"symbol,omitempty"`
	Decimals        uint8          `gorm:"column:decimals" ch:"decimals" json:"decimals,omitempty"`
}

type ERC20Transfer struct {
	BlockHash       common.Hash    `gorm:"column:block_hash" ch:"block_hash" json:"block_hash,omitempty"`
	BlockNumber     uint64         `gorm:"column:block_number" ch:"block_number" json:"block_number,omitempty"`
	TxHash          common.Hash    `gorm:"column:tx_hash" ch:"tx_hash" json:"tx_hash,omitempty"`
	ContractAddress common.Address `gorm:"column:contract_address" ch:"contract_address" json:"contract_address,omitempty"`
	From            common.Address `gorm:"column:from" ch:"from" json:"from,omitempty"`
	To              common.Address `gorm:"column:to" ch:"to" json:"to,omitempty"`
	Value           *big.Int       `gorm:"column:value" ch:"value" json:"value,omitempty"`
}

type ERC20TransferWithMetadata struct {
	ERC20Token
	BlockHash       common.Hash    `gorm:"column:block_hash" ch:"block_hash" json:"block_hash,omitempty"`
	BlockNumber     uint64         `gorm:"column:block_number" ch:"block_number" json:"block_number,omitempty"`
	TxHash          common.Hash    `gorm:"column:tx_hash" ch:"tx_hash" json:"tx_hash,omitempty"`
	ContractAddress common.Address `gorm:"column:contract_address" ch:"contract_address" json:"contract_address,omitempty"`
	From            common.Address `gorm:"column:from" ch:"from" json:"from,omitempty"`
	To              common.Address `gorm:"column:to" ch:"to" json:"to,omitempty"`
	Value           *big.Int       `gorm:"column:value" ch:"value" json:"value,omitempty"`
}

func (e *ERC20TransferWithMetadata) SetMeta(meta ERC20Token) *ERC20TransferWithMetadata {
	e.ERC20Token = meta
	return e
}

type TokenChanges struct {
	TokenLoss []ERC20TransferWithMetadata `json:"token_loss"`
	TokenGain []ERC20TransferWithMetadata `json:"token_gain"`
}
