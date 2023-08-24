package model

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

func (b *RPCBlock) ToCommonTransactionList() ([]RawTransaction, error) {

	return nil, nil
}

func (b *RPCBlock) ToOptimismTransactionList() []RawOptimismTransaction {
	return nil
}

type Block struct {
}
