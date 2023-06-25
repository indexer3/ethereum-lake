package engine

import (
	"context"
	"sync"

	"github.com/indexer3/ethereum-lake/common/database"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/common/utils"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type IEngine interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context)
}

type Engine struct {
	dbs       []database.IDatabase
	done      chan struct{}
	startOnce sync.Once
	taskMu    sync.RWMutex
	Tasks     []ITask
	TaskIndex chan DispatchTaskIndex
}

func NewLakeEngine(dbs []database.IDatabase) IEngine {
	return &Engine{
		dbs:       dbs,
		done:      make(chan struct{}, 1),
		startOnce: sync.Once{},
		taskMu:    sync.RWMutex{},
		TaskIndex: make(chan DispatchTaskIndex, viper.GetInt(config.TaskChannelSize)),
	}
}

func (e *Engine) AddTask(task ITask) {
	e.taskMu.Lock()
	e.Tasks = append(e.Tasks, task)
	e.taskMu.Unlock()
}

func (e *Engine) Start(ctx context.Context) error {
	e.startOnce.Do(func() {
		go func() {
			for {
				select {
				case taskIndex := <-e.TaskIndex:
					for _, task := range e.Tasks {

						ts := task
						go utils.Recoverable(func() {
							// handle from cache
							err := ts.Handle(ctx, taskIndex)
							if err != nil {
								log.Error("failed to handle task", zap.String("task", task.Name()), zap.Any("taskIndex", taskIndex), zap.Error(err))
							}
						})
					}
				case <-e.done:
					log.Info("engine stopped")
					return
				case <-ctx.Done():
					log.Info("engine stopped by context")
				}
			}
		}()
	})

	return nil
}

func (e *Engine) Stop(ctx context.Context) {
	e.done <- struct{}{}
	log.Info("engine stop signal sent")
}
