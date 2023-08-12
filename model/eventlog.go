package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type SQLQueryEventLog = ChainbaseQueryResponse[SQLQueryResp[[]EventLog]]

type EventLog struct {
	Address          string `json:"address,omitempty"`
	BlockNumber      string `json:"block_number,omitempty"`
	BlockTimestamp   string `json:"block_timestamp,omitempty"`
	Data             string `json:"data,omitempty"`
	LogIndex         uint   `json:"log_index,omitempty"`
	Topic0           string `json:"topic0,omitempty"`
	Topic1           string `json:"topic1,omitempty"`
	Topic2           string `json:"topic2,omitempty"`
	Topic3           string `json:"topic3,omitempty"`
	TopicsCount      int    `json:"topics_count,omitempty"`
	TransactionHash  string `json:"transaction_hash,omitempty"`
	TransactionIndex uint   `json:"transaction_index,omitempty"`
}

func (e EventLog) ToEthereumLog() types.Log {
	topics := make([]common.Hash, e.TopicsCount)
	if e.Topic0 != "" {
		topics[0] = common.HexToHash(e.Topic0)
	}

	if e.Topic1 != "" {
		topics[1] = common.HexToHash(e.Topic1)
	}

	if e.Topic2 != "" {
		topics[2] = common.HexToHash(e.Topic2)
	}

	if e.Topic3 != "" {
		topics[3] = common.HexToHash(e.Topic3)
	}

	bn, _ := hexutil.DecodeBig(e.BlockNumber)

	return types.Log{
		Address:     common.HexToAddress(e.Address),
		Topics:      topics,
		Data:        common.Hex2Bytes(e.Data),
		BlockNumber: bn.Uint64(),
		TxHash:      common.HexToHash(e.TransactionHash),
		TxIndex:     e.TransactionIndex,
		Index:       e.LogIndex,
	}
}
