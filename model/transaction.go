package model




type Transaction struct {
	BlockHash            string `json:"block_hash"`
	BlockNumber          string `json:"block_number"`
	BlockTimestamp       string `json:"block_timestamp"`
	ContractAddress      string `json:"contract_address"`
	CumulativeGasUsed    string `json:"cumulative_gas_used"`
	EffectiveGasPrice    string `json:"effective_gas_price"`
	FromAddress          string `json:"from_address"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gas_price"`
	GasUsed              string `json:"gas_used"`
	Input                string `json:"input"`
	LogsCount            int    `json:"logs_count"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	Nonce                string `json:"nonce"`
	R                    string `json:"r"`
	S                    string `json:"s"`
	Status               string `json:"status"`
	ToAddress            string `json:"to_address"`
	TransactionHash      string `json:"transaction_hash"`
	TransactionIndex     int    `json:"transaction_index"`
	Type                 int    `json:"type"`
	V                    string `json:"v"`
	Value                string `json:"value"`
}
