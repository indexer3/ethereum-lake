package uniswap

import (
	"context"

	"github.com/indexer3/ethereum-lake/engine"
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

func (t *UniswapTask) Run(ctx context.Context) error {
	return nil
}

func (t *UniswapTask) Handle(ctx context.Context, taskIndex engine.DispatchTaskIndex) error {
	return nil
}
