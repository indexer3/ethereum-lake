package model

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SQLQueryEventLog = ChainbaseQueryResponse[SQLQueryResp[[]EventLog]]

type EthereumLog types.Log

// How EventLog looks like:
//
//	{
//		"address": "0x5564886ca2c518d1964e5fcea4f423b41db9f561",
//		"block_number": "141710",
//		"block_timestamp": "2015-08-25 16:19:32",
//		"data": "0x",
//		"log_index": 0,
//		"topic0": "0xa6697e974e6a320f454390be03f74955e8978f1a6971ea6730542e37b66179bc",
//		"topic1": "0x6d616b6572000000000000000000000000000000000000000000000000000000",
//		"topic2": "",
//		"topic3": "",
//		"topics_count": 2,
//		"transaction_hash": "0x00efda1c9b71bb98139fc59a77b88b63f34d9f87b716f5f340ccf74535833dc4",
//		"transaction_index": 0
//	  }
func (l *EthereumLog) ToEventLog(blockHeader types.Header) EventLog {
	var (
		topic0, topic1, topic2, topic3 string
	)

	if len(l.Topics) >= 1 {
		topic0 = l.Topics[0].Hex()
	}

	if len(l.Topics) >= 2 {
		topic1 = l.Topics[1].Hex()
	}

	if len(l.Topics) >= 3 {
		topic2 = l.Topics[2].Hex()
	}

	if len(l.Topics) == 4 {
		topic3 = l.Topics[3].Hex()
	}

	return EventLog{
		Address:          l.Address.Hex(),
		BlockNumber:      blockHeader.Number.String(),
		BlockTimestamp:   time.Unix(int64(blockHeader.Time), 0).String(),
		Data:             common.Bytes2Hex(l.Data),
		LogIndex:         l.Index,
		Topic0:           topic0,
		Topic1:           topic1,
		Topic2:           topic2,
		Topic3:           topic3,
		TopicsCount:      len(l.Topics),
		TransactionHash:  l.TxHash.Hex(),
		TransactionIndex: l.TxIndex,
	}
}

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

	bn, _ := strconv.ParseUint(e.BlockNumber, 10, 64)

	return types.Log{
		Address:     common.HexToAddress(e.Address),
		Topics:      topics,
		Data:        common.Hex2Bytes(e.Data),
		BlockNumber: bn,
		TxHash:      common.HexToHash(e.TransactionHash),
		TxIndex:     e.TransactionIndex,
		Index:       e.LogIndex,
	}
}
