package common

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Handle[T any](ctx context.Context, handleItem T, fns []func(ctx context.Context, item T) error) error {
	var eg errgroup.Group

	for _, f := range fns {
		fn := f

		eg.Go(func() error {
			return fn(ctx, handleItem)
		})
	}

	return eg.Wait()
}
