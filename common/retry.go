package common

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Retry[T any](ctx context.Context, fn func() (T, error)) (T, error) {
	var defaultBackOff backoff.BackOff = backoff.NewExponentialBackOff()
	b := backoff.WithMaxRetries(defaultBackOff, viper.GetUint64(config.RPCMaxRetry))
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
