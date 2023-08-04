package opensea

import (
	"context"

	"github.com/indexer3/ethereum-lake/index_engine"
)

var _ index_engine.Task = (*OpenSeaTask)(nil)

type OpenSeaTask struct {
}

func NewOpenSeaTask() *OpenSeaTask {
	return &OpenSeaTask{}
}

func (t *OpenSeaTask) Name() string {
	return "opensea"
}

func (t *OpenSeaTask) Handle(ctx context.Context, taskIndex index_engine.DispatchTaskIndex) error {
	return nil
}
