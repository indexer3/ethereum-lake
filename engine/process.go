package engine

import (
	"context"
	"time"

	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var _ Worker = (*WorkflowProcess)(nil)

type Worker interface {
	Run(ctx context.Context) error
}

// Need to be implemented in different process
type Process interface {
	Handle(ctx context.Context, curBlock uint64) error
	GetLatestBlock(ctx context.Context) (uint64, error)
}

type WorkflowProcess struct {
	Process
}

func (w *WorkflowProcess) Run(ctx context.Context) error {
	endCursor := viper.GetUint64(config.EndCursor)

	if endCursor == 0 {
		return w.Stream(ctx)
	}

	return w.Sync(ctx)
}

func (w *WorkflowProcess) Stream(ctx context.Context) error {
	startCursor := viper.GetUint64(config.StartCursor)
	targetBlock, err := w.GetLatestBlock(ctx)
	if err != nil {
		log.Error("failed to get latest block", zap.Error(err))
		return err
	}

	log.Info("start streaming", zap.Uint64("startCursor", startCursor), zap.Uint64("targetBlock", targetBlock))

	cursor := startCursor

	for {
		if cursor <= targetBlock {
			if err := w.Handle(ctx, cursor); err != nil {
				log.Error("failed to handle block", zap.Uint64("cursor", cursor), zap.Error(err))
				continue
			}

			cursor++
			continue
		}

		if cursor > targetBlock {
			// Wait for new block
			time.Sleep(time.Duration(viper.GetUint64(config.CatchUpSleepInterval)) * time.Second)

			// update target block
			subTargetBlock, err := w.GetLatestBlock(ctx)
			if err != nil {
				log.Error("failed to get latest block", zap.Error(err))
				continue
			}

			if subTargetBlock > targetBlock {
				targetBlock = subTargetBlock
			}
		}
	}
}

func (w *WorkflowProcess) Sync(ctx context.Context) error {
	startCursor := viper.GetUint64(config.StartCursor)
	targetBlock := viper.GetUint64(config.EndCursor)

	log.Info("start syncing", zap.Uint64("startCursor", startCursor), zap.Uint64("targetBlock", targetBlock))

	for cursor := startCursor; cursor <= targetBlock; {
		if err := w.Handle(ctx, cursor); err != nil {
			log.Error("failed to handle block", zap.Uint64("cursor", cursor), zap.Error(err))
			continue
		}

		cursor++
	}

	log.Info("syncing finished", zap.Uint64("startCursor", startCursor), zap.Uint64("targetBlock", targetBlock))

	return nil
}
