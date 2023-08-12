package chainbase_client

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (c *ChainbaseCli) GetEventLogsByTxHashes(ctx context.Context, network constant.Network, txHashes []common.Hash) ([]types.Log, error) {
	txHashesStrings := make([]string, 0)

	lo.ForEach[common.Hash](txHashes, func(txhash common.Hash, _ int) {
		txHashesStrings = append(txHashesStrings, "'"+txhash.String()+"'")
	})

	joinCommaQuery := strings.Join(txHashesStrings, ",")

	var resp model.SQLQueryEventLog

	query := fmt.Sprintf(`select * from transaction_logs where %s.transaction_hash in (%s)`, network.String(), joinCommaQuery)

	_, err := c.cli.R().SetContext(ctx).SetBody(model.ChainbaseQueryRequest{
		Query: query,
	}).SetResult(&resp).Post("/v1/dw/query")
	if err != nil {
		log.Error("failed to get response from chainbase", zap.Error(err))
		return nil, err
	}

	if resp.Code != OkCode || resp.Data.ErrMsg != "" {
		log.Error("failed to get response from chainbase", zap.Any("response", resp))
		return nil, fmt.Errorf("failed to execute eventlogs query")
	}

	logs := make([]types.Log, 0)
	lo.ForEach(resp.Data.Result, func(log model.EventLog, _ int) {
		logs = append(logs, log.ToEthereumLog())
	})

	return logs, nil
}
