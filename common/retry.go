package common

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/indexer3/ethereum-lake/common/log"
	"go.uber.org/zap"
)

func Retry[T any](ctx context.Context, maxRetries uint64, fn func() (T, error)) (T, error) {
	var defaultBackOff backoff.BackOff = backoff.NewExponentialBackOff()
	b := backoff.WithMaxRetries(defaultBackOff, maxRetries)
	b = backoff.WithContext(b, ctx)

	result, err := backoff.RetryWithData[T](func() (T, error) {
		return fn()
	}, b)

	if err != nil {
		log.Error("failed to execute retry function, max retry exceeded", zap.Error(err))
		return result, err
	}

	return result, nil
}
