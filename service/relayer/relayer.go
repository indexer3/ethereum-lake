package relayer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/constant"
)

type RelayerService[T any] interface {
	AccountPortfolio(ctx context.Context, account common.Address, network constant.Network) (T, error)
}
