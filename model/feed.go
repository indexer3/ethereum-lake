package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

type FeedCursor struct {
	// Network          constant.Network `json:"network"`
	// AccountAddress   common.Address   `json:"account_address"`
	BlockNumber      uint64 `json:"block_number"`
	TransactionIndex uint64 `json:"transaction_index"`
}

type Feed struct {
	TxHash    common.Hash                 `json:"tx_hash"`
	From      common.Address              `json:"from"`
	To        common.Address              `json:"to"`
	GasUsed   uint64                      `json:"gas_used"`
	GasPrice  uint64                      `json:"gas_price"`
	MethodID  string                      `json:"method_id"`
	Network   constant.Network            `json:"network"`
	TokenLoss []ERC20TransferWithMetadata `json:"token_loss"`
	TokenGain []ERC20TransferWithMetadata `json:"token_gain"`
}
