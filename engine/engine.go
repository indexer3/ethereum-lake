package engine

import (
	"context"
	"encoding/json"
	"math"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/cache"
	"github.com/indexer3/ethereum-lake/common/client"
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
	cache     *cache.CommonCache
	taskMu    sync.RWMutex
	ethClient *client.NodeClient
	Tasks     []ITask
	TaskIndex chan DispatchTaskIndex
}

func NewLakeEngine(dbs []database.IDatabase) IEngine {
	cli, err := client.NewNodeClientsWithEndpoints(viper.GetStringSlice(config.RPCEndpoints))
	if err != nil {
		log.Fatal("failed to create node client", zap.Error(err))
	}

	return &Engine{
		dbs:       dbs,
		done:      make(chan struct{}, 1),
		startOnce: sync.Once{},
		taskMu:    sync.RWMutex{},
		ethClient: cli,
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
						// prepare the task data once to reduce rpc calls
						// TODO: replace with redis
						if taskIndex.BlockNumber != nil {
							blockInfo, err := e.ethClient.Client().BlockByNumber(ctx, taskIndex.BlockNumber)
							if err != nil {
								log.Error("failed to get block info", zap.Any("taskIndex", taskIndex), zap.Error(err))
								continue
							}

							blockBytes, err := json.Marshal(blockInfo)
							if err != nil {
								log.Error("failed to marshal block info", zap.Any("taskIndex", taskIndex), zap.Error(err))
								continue
							}

							e.cache.Set(taskIndex.BlockNumber.String(), blockBytes)
						}

						if taskIndex.TxHash != "" {
							hash := common.HexToHash(taskIndex.TxHash)

							txInfo, _, err := e.ethClient.Client().TransactionByHash(ctx, hash)
							if err != nil {
								log.Error("failed to get tx info", zap.Any("taskIndex", taskIndex), zap.Error(err))
								continue
							}

							txBytes, err := json.Marshal(txInfo)
							if err != nil {
								log.Error("failed to marshal tx info", zap.Any("taskIndex", taskIndex), zap.Error(err))
								continue
							}

							e.cache.Set(taskIndex.TxHash, txBytes)
						}

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

		go func() {
			endCursor := viper.GetInt(config.EndCursor)

			// has no end
			if endCursor <= 0 {
				endCursor = math.MaxInt64
			}

			for i := viper.GetInt(config.StartCursor); i < endCursor; i++ {

			}
		}()
	})

	return nil
}

func (e *Engine) Stop(ctx context.Context) {
	e.done <- struct{}{}
	log.Info("engine stop signal sent")
}
