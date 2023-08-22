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

type RawTransaction struct {
	Hash              string `json:"hash,omitempty"`
	Status            uint64 `json:"status,omitempty"`
	BlockHash         string `json:"blockHash,omitempty"`
	BlockNumber       uint64 `json:"blockNumber,omitempty"`
	Timestamp         int64  `json:"timestamp"`
	From              string `json:"from,omitempty"`
	To                string `json:"to,omitempty"`
	Value             string `json:"value,omitempty"`
	TransactionFee    uint64 `json:"transactionFee"`
	GasPrice          uint64 `json:"gasPrice,omitempty"`
	GasLimit          uint64 `json:"gasLimit"`
	GasUsed           uint64 `json:"gas,omitempty"`
	BaseFee           uint64 `json:"baseFee"`
	MaxFeePerGas      uint64 `json:"maxFeePerGas,omitempty"`
	PriorityFeePerGas uint64 `json:"priorityFeePerGas,omitempty"`
	TxType            uint8  `json:"txType,omitempty"`
	MethodID          string `json:"methodID,omitempty"`
	Nonce             uint64 `json:"nonce,omitempty"`
	TransactionIndex  uint   `json:"transactionIndex,omitempty"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	//Category          uint64 `json:"category,omitempty"` // 0: eth transfer， 1: contract create， 2: contract call
	//IsEIP1559         bool   `json:"isEIP1559"`
	//InputData         string `json:"input,omitempty"`
}

type RawOptimismTransaction struct {
	Hash                   string  `json:"hash,omitempty"`
	BlockHash              string  `json:"blockHash,omitempty"`
	BlockNumber            uint64  `json:"blockNumber,omitempty"`
	Timestamp              int64   `json:"timestamp"`
	From                   string  `json:"from,omitempty"`
	To                     string  `json:"to,omitempty"`
	Value                  string  `json:"value,omitempty"`
	TransactionFee         uint64  `json:"transactionFee"`
	L2GasPrice             uint64  `json:"l2GasPrice"`
	L2GasUsedByTransaction uint64  `json:"l2GasUsedByTransaction"`
	L1GasUsedByTransaction uint64  `json:"l1GasUsedByTransaction"`
	L1GasPrice             uint64  `json:"l1GasPrice"`
	L1FeeScalar            float64 `json:"l1FeeScalar"`
	Nonce                  uint64  `json:"nonce,omitempty"`
	TransactionIndex       uint    `json:"transactionIndex,omitempty"`
	ContractAddress        string  `json:"contractAddress,omitempty"`
	MethodID               string  `json:"methodID,omitempty"`
	Status                 uint64  `json:"status,omitempty"`
}
