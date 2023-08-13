package chainbase_client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts/uniswap"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (c *ChainbaseCli) GetUniswapHoldingPositionNFT(ctx context.Context, network constant.Network, account common.Address) ([]types.Log, error) {
	query := fmt.Sprintf(`select block_number, block_timestamp, transaction_hash, transaction_index, log_index, address, data, topics_count, topic0, topic1, topic2, topic3
		 from %s.transaction_logs where address = '%s' and topic0 = '%s' and topic1 = '%s' and topic2 = '%s'`,
		network.String(),
		uniswap.UniswapV3NFTPositionManagerAddress.String(),
		constant.UniswapPostitionTransferEvent().String(),
		common.HexToHash(constant.ZeroBigInt().String()).String(),
		account.String())

	var resp model.SQLQueryEventLog
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
