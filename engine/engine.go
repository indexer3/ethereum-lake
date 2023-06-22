package engine

import "context"

type IEngine interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context)
}

type Engine struct {
}

func (e *Engine) Start(ctx context.Context) error {
	return nil
}

func (e *Engine) Stop(ctx context.Context) {

}
