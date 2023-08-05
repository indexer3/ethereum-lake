package model

type Transaction struct {
	BlockHash            string `json:"block_hash,omitempty"`
	BlockNumber          string `json:"block_number,omitempty"`
	BlockTimestamp       string `json:"block_timestamp,omitempty"`
	ContractAddress      string `json:"contract_address,omitempty"`
	CumulativeGasUsed    string `json:"cumulative_gas_used,omitempty"`
	EffectiveGasPrice    string `json:"effective_gas_price,omitempty"`
	FromAddress          string `json:"from_address,omitempty"`
	Gas                  string `json:"gas,omitempty"`
	GasPrice             string `json:"gas_price,omitempty"`
	GasUsed              string `json:"gas_used,omitempty"`
	Input                string `json:"input,omitempty"`
	LogsCount            int    `json:"logs_count,omitempty"`
	MaxFeePerGas         string `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas,omitempty"`
	Nonce                string `json:"nonce,omitempty"`
	R                    string `json:"r,omitempty"`
	S                    string `json:"s,omitempty"`
	Status               string `json:"status,omitempty"`
	ToAddress            string `json:"to_address,omitempty"`
	TransactionHash      string `json:"transaction_hash,omitempty"`
	TransactionIndex     int    `json:"transaction_index,omitempty"`
	Type                 int    `json:"type,omitempty"`
	V                    string `json:"v,omitempty"`
	Value                string `json:"value,omitempty"`
}
