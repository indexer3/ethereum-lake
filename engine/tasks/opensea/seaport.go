package opensea

import (
	"context"

	"github.com/indexer3/ethereum-lake/engine"
)

var _ engine.ITask = (*OpenSeaTask)(nil)

type OpenSeaTask struct {
}

func NewOpenSeaTask() *OpenSeaTask {
	return &OpenSeaTask{}
}

func (t *OpenSeaTask) Name() string {
	return "opensea"
}

func (t *OpenSeaTask) Run(ctx context.Context) error {
	return nil
}

func (t *OpenSeaTask) Handle(ctx context.Context, taskIndex engine.DispatchTaskIndex) error {
	return nil
}
