package model

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type RPCBlock struct {
	Timestamp     string                `json:"timestamp"`
	BaseFeePerGas string                `json:"baseFeePerGas"`
	GasLimit      string                `json:"gasLimit"`
	Number        string                `json:"number"`
	Difficulty    string                `json:"difficulty"`
	Hash          string                `json:"hash"`
	Miner         string                `json:"miner"`
	Nonce         string                `json:"nonce"`
	Transactions  []RPCBlockTransaction `json:"transactions"`
}

type RPCBlockTransaction struct {
	BlockHash            string `json:"blockHash,omitempty"`
	BlockNumber          string `json:"blockNumber,omitempty"`
	From                 string `json:"from,omitempty"`
	GasLimit             string `json:"gas,omitempty"`
	GasPrice             string `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	Hash                 string `json:"hash,omitempty"`
	Input                string `json:"input,omitempty"`
	Nonce                string `json:"nonce,omitempty"`
	To                   string `json:"to,omitempty"`
	TransactionIndex     string `json:"transactionIndex,omitempty"`
	Value                string `json:"value,omitempty"`
	Type                 string `json:"type,omitempty"`
	ChainId              string `json:"chainId,omitempty"`

	// Optimism specific fields
	L1TxOrigin    string `json:"l1TxOrigin,omitempty"`
	L1BlockNumber string `json:"l1BlockNumber,omitempty"`
	L1Timestamp   string `json:"l1Timestamp,omitempty"`
	QueueOrigin   string `json:"queueOrigin,omitempty"`
}

func (b *RPCBlock) GetCommonTransactionList() ([]RawTransaction, error) {

	return nil, nil
}

func (b *RPCBlock) GetOptimismTransactionList() []RawOptimismTransaction {
	return nil
}

type Block struct {
	Number          *big.Int       `json:"number"`
	Hash            common.Hash    `json:"hash"`
	Coinbase        common.Address `json:"coinbase"`
	ParentHash      common.Hash    `json:"parent_hash"`
	ReceiptHash     common.Hash    `json:"receipt_hash"`
	UncleHash       common.Hash    `json:"uncle_hash"`
	MixDigest       common.Hash    `json:"mix_digest"`
	Root            common.Hash    `json:"root"`
	Bloom           types.Bloom    `json:"bloom"`
	Nonce           [8]byte        `json:"nonce"`
	Extra           []byte         `json:"extra"`
	BaseFee         *big.Int       `json:"base_fee"`
	GasLimit        *big.Int       `json:"gas_limit"`
	GasUsed         uint64         `json:"gas_used"`
	Size            uint64         `json:"size"`
	Difficulty      *big.Int       `json:"difficulty"`
	TotalDifficulty *big.Int       `json:"total_difficulty"`
	Timestamp       time.Time      `json:"timestamp"`
}
