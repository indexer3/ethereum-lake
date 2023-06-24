package engine

import (
	"context"
	"sync"

	"github.com/indexer3/ethereum-lake/common/database"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/common/utils"
)

type IEngine interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context)
}

type Engine struct {
	dbs       []database.IDatabase
	Tasks     []ITask
	TaskIndex chan DispatchTaskIndex
}

var taskMu sync.RWMutex
var engineStart sync.Once

func NewLakeEngine(dbs []database.IDatabase) IEngine {
	return &Engine{
		dbs: dbs,
	}
}

func (e *Engine) AddTask(task ITask) {
	taskMu.Lock()
	e.Tasks = append(e.Tasks, task)
	taskMu.Unlock()
}

func (e *Engine) Start(ctx context.Context) error {
	engineStart.Do(func() {
		for {
			select {
			case taskIndex := <-e.TaskIndex:
				for _, task := range e.Tasks {
					go utils.Recoverable(func() {
						task.Handle(ctx, taskIndex)
					})
				}
			case <-ctx.Done():
				log.Info("engine stopped by context")
			}
		}
	})

	return nil
}

func (e *Engine) Stop(ctx context.Context) {

}
