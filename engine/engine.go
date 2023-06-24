package engine

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/database"
)

type IEngine interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context)
}

type Engine struct {
	dbs []database.IDatabase
}

func NewLakeEngine(dbs []database.IDatabase) IEngine {
	return &Engine{
		dbs: dbs,
	}
}

func (e *Engine) AddTask(task ITask) {

}

func (e *Engine) Start(ctx context.Context) error {
	return nil
}

func (e *Engine) Stop(ctx context.Context) {

}
