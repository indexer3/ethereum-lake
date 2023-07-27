package uniswap

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexer3/ethereum-lake/common/cache"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/engine"
	"go.uber.org/zap"
)

var _ engine.ITask = (*UniswapTask)(nil)

type UniswapTask struct {
}

func NewUniswapTask() *UniswapTask {
	return &UniswapTask{}
}

func (t *UniswapTask) Name() string {
	return "uniswap"
}

func (t *UniswapTask) Handle(ctx context.Context, taskIndex engine.DispatchTaskIndex) error {
	var (
		blockInfo *types.Block
		txInfo    *types.Transaction
	)

	if taskIndex.BlockNumber != nil {
		blockBytes, err := cache.GlobalCache().Get(taskIndex.BlockNumber.String())
		if err != nil {
			log.Error("failed to block info from cache", zap.String("task", t.Name()), zap.Any("taskIndex", taskIndex), zap.Error(err))
			return err
		}

		err = json.Unmarshal(blockBytes, &blockInfo)
		if err != nil {
			log.Error("failed to unmarshal block info", zap.String("task", t.Name()), zap.Any("taskIndex", taskIndex), zap.Error(err))
			return err
		}
	}

	if taskIndex.TxHash != "" {
		txBytes, err := cache.GlobalCache().Get(strings.ToLower(taskIndex.TxHash))
		if err != nil {
			log.Error("failed to tx info from cache", zap.String("task", t.Name()), zap.Any("taskIndex", taskIndex), zap.Error(err))
			return err
		}

		err = json.Unmarshal(txBytes, &txInfo)
		if err != nil {
			log.Error("failed to unmarshal tx info", zap.String("task", t.Name()), zap.Any("taskIndex", taskIndex), zap.Error(err))
			return err
		}
	}

	return nil 
}
