package model

import "github.com/ethereum/go-ethereum/core/types"

// Optimism RPC Receipt Compatibility
type RPCReceipt struct {
	TransactionHash   string `json:"transactionHash,omitempty"`
	TransactionIndex  string `json:"transactionIndex,omitempty"`
	BlockHash         string `json:"blockHash,omitempty"`
	BlockNumber       string `json:"blockNumber,omitempty"`
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	CumulativeGasUsed string `json:"cumulativeGasUsed,omitempty"`
	GasUsed           string `json:"gasUsed,omitempty"`

	Status          string      `json:"status,omitempty"`
	GasPrice        string      `json:"gasPrice,omitempty"`
	ContractAddress string      `json:"contractAddress,omitempty"`
	Logs            []types.Log `json:"logs,omitempty"`

	// no such field in optimism
	EffectiveGasPrice string `json:"effectiveGasPrice,omitempty"`

	// optimism fields
	L1Fee       string `json:"l1Fee,omitempty"`
	L1FeeScalar string `json:"l1FeeScalar,omitempty"`
	L1GasPrice  string `json:"l1GasPrice,omitempty"`
	L1GasUsed   string `json:"l1GasUsed,omitempty"`
}

type Receipt struct {
	TransactionHash   string  `json:"transactionHash,omitempty"`
	TransactionIndex  uint    `json:"transactionIndex,omitempty"`
	BlockHash         string  `json:"blockHash,omitempty"`
	BlockNumber       uint64  `json:"blockNumber,omitempty"`
	From              string  `json:"from,omitempty"`
	To                string  `json:"to,omitempty"`
	GasPrice          uint64  `json:"gasPrice"`
	GasUsed           uint64  `json:"gasUsed,omitempty"`
	EffectiveGasPrice uint64  `json:"effectiveGasPrice,omitempty"`
	ContractAddress   string  `json:"contractAddress,omitempty"`
	L1Fee             uint64  `json:"l1Fee,omitempty"`
	L1FeeScalar       float64 `json:"l1FeeScalar,omitempty"`
	L1GasPrice        uint64  `json:"l1GasPrice,omitempty"`
	L1GasUsed         uint64  `json:"l1GasUsed,omitempty"`
	Status            uint64  `json:"status,omitempty"`
}

func (r *RPCReceipt) ToReceipt() (*Receipt, error) {
	var (
		err error
	)

	return nil, err
}
